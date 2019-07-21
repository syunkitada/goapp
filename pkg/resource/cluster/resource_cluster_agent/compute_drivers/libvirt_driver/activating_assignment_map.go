package libvirt_driver

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/syunkitada/goapp/pkg/lib/exec_utils"
	"github.com/syunkitada/goapp/pkg/lib/json_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/lib/os_utils"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_agent/compute_drivers/libvirt_driver/libvirt_models"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

var reVirshListLine = regexp.MustCompile(` ([-0-9]+) +([a-z0-9]+) +(.*)`)

func (driver *LibvirtDriver) syncActivatingAssignmentMap(tctx *logger.TraceContext,
	assignmentMap map[uint]resource_model.ComputeAssignmentEx) error {

	virshListAll, err := exec_utils.Cmdf(tctx, "virsh list --all")
	if err != nil {
		return err
	}
	domainMap := map[string]libvirt_models.Domain{}
	for i, line := range strings.Split(virshListAll, "\n") {
		if i == 0 {
			continue
		}
		result := reVirshListLine.FindAllStringSubmatch(line, -1)
		if result == nil {
			continue
		}
		fmt.Println("DEBUG state", result[0][3])
		var state string
		switch result[0][3] {
		case "shut off":
			state = resource_model.StateDown
		case "running":
			state = resource_model.StateUp
		default:
			state = resource_model.StateUnknown
		}
		domainMap[result[0][2]] = libvirt_models.Domain{
			Name:  result[0][2],
			State: state,
		}
	}

	out, err := exec_utils.Cmd(tctx, "brctl show")
	if err != nil {
		return err
	}
	bridgeInterfaceMap := map[string]map[string]bool{}
	for i, line := range strings.Split(out, "\n") {
		if i == 0 {
			continue
		}
		columns := strings.Split(line, "\t")
		if len(columns) == 0 {
			break
		}
		bridgeInterfaceMap[columns[0]] = map[string]bool{}
	}

	for _, assignment := range assignmentMap {
		compute := assignment.Spec.Compute
		computeId := fmt.Sprintf("compute%d", assignment.ID)
		pciSlot := 1

		vmDir := filepath.Join(driver.vmsDir, computeId)
		if err = os_utils.Mkdir(vmDir, 0755); err != nil {
			return err
		}
		vmImagePath := filepath.Join(vmDir, "img")
		vmConfigImagePath := filepath.Join(vmDir, "config.img")
		vmDomainXmlPath := filepath.Join(vmDir, "domain.xml")

		srcImagePath := filepath.Join(driver.imagesDir, compute.ImageSpec.Name)
		if !os_utils.PathExists(srcImagePath) {
			tctx.SetTimeout(3600)
			if _, err = exec_utils.Cmdf(tctx, "wget -O %s %s", srcImagePath, compute.ImageSpec.Url); err != nil {
				return err
			}
		}
		if !os_utils.PathExists(vmImagePath) {
			if _, err = exec_utils.Cmdf(tctx, "cp %s %s", srcImagePath, vmImagePath); err != nil {
				return err
			}
		}

		configDir := filepath.Join(vmDir, "config")
		if err = os_utils.Mkdir(configDir, 0755); err != nil {
			return err
		}
		configOpenstackDir := filepath.Join(configDir, "openstack")
		if err = os_utils.Mkdir(configOpenstackDir, 0755); err != nil {
			return err
		}
		configOpenstackLatestDir := filepath.Join(configOpenstackDir, "latest")
		if err = os_utils.Mkdir(configOpenstackLatestDir, 0755); err != nil {
			return err
		}
		vmMetaDataConfigFilePath := filepath.Join(configOpenstackLatestDir, "meta_data.json")
		vmUserDataConfigFilePath := filepath.Join(configOpenstackLatestDir, "user_data")

		metaData := map[string]interface{}{
			"instance-id":    computeId,
			"local-hostname": computeId,
			"network": map[string]interface{}{
				"config": "disabled",
			},
		}
		metaDataBytes, err := json_utils.Marshal(metaData)
		if err != nil {
			return err
		}
		if err = ioutil.WriteFile(vmMetaDataConfigFilePath, []byte(metaDataBytes), 0644); err != nil {
			return err
		}

		userData := fmt.Sprintf("#cloud-config\nnetwork:\n  config: disabled\n")
		if err = ioutil.WriteFile(vmUserDataConfigFilePath, []byte(userData), 0644); err != nil {
			return err
		}

		if _, err = exec_utils.Cmdf(tctx, "genisoimage -o %s -V config-2 -r -J %s",
			vmConfigImagePath, configDir); err != nil {
			return err
		}

		emulators := []libvirt_models.DeviceEmulator{}
		emulators = append(emulators, libvirt_models.DeviceEmulator{
			Emulator: "/usr/bin/qemu-system-x86_64",
		})

		disks := []libvirt_models.DeviceDisk{}
		disks = append(disks, libvirt_models.DeviceDisk{
			Type:   "file",
			Device: "disk",
			Driver: libvirt_models.DiskDriverQcow2{
				Name:  "qemu",
				Type:  "qcow2",
				Cache: "none",
			},
			Source: libvirt_models.DiskSource{File: vmImagePath},
			Target: libvirt_models.DiskTarget{Dev: "hda", Bus: "ide"},
			Alias:  libvirt_models.Alias{Name: "ide0-0-0"},
			Address: libvirt_models.DriveAddress{
				Type:       "drive",
				Controller: 0,
				Bus:        0,
				Target:     0,
				Unit:       0,
			},
		})

		disks = append(disks, libvirt_models.DeviceDisk{
			Type:   "file",
			Device: "cdrom",
			Driver: libvirt_models.DiskDriverRaw{
				Name: "qemu",
				Type: "raw",
			},
			Source: libvirt_models.DiskSource{File: vmConfigImagePath},
			Target: libvirt_models.DiskTarget{Dev: "hdc", Bus: "ide"},
			Alias:  libvirt_models.Alias{Name: "ide0-1-0"},
			Address: libvirt_models.DriveAddress{
				Type:       "drive",
				Controller: 0,
				Bus:        1,
				Target:     0,
				Unit:       0,
			},
		})

		serials := []libvirt_models.DeviceSerial{}
		serials = append(serials, libvirt_models.DeviceSerial{
			Type:   "pty",
			Source: libvirt_models.SerialSource{Path: "/dev/pts/8"},
			Target: libvirt_models.SerialTarget{Port: 0},
			Alias:  libvirt_models.Alias{Name: "serial0"},
		})

		consoles := []libvirt_models.DeviceConsole{}
		consoles = append(consoles, libvirt_models.DeviceConsole{
			Type:   "pty",
			Tty:    "/dev/pts/8",
			Source: libvirt_models.ConsoleSource{Path: "/dev/pts/8"},
			Target: libvirt_models.ConsoleTarget{Type: "serial", Port: 0},
			Alias:  libvirt_models.Alias{Name: "serial0"},
		})

		membaloons := []libvirt_models.DeviceMembaloon{}
		membaloons = append(membaloons, libvirt_models.DeviceMembaloon{
			Model: "virtio",
			Alias: libvirt_models.Alias{Name: "balloon0"},
			Address: libvirt_models.PciAddress{
				Type:     "pci",
				Domain:   "0x0000",
				Bus:      "0x00",
				Slot:     fmt.Sprintf("%#02X", pciSlot),
				Function: "0x0",
			},
		})
		pciSlot += 1

		interfaces := []libvirt_models.DeviceInterface{}
		for i, port := range compute.Ports {
			bridgeName := fmt.Sprintf("br-compute%d", port.NetworkID)
			interfaceMap, ok := bridgeInterfaceMap[bridgeName]
			if !ok {
				if _, err = exec_utils.Cmdf(tctx, "brctl addbr %s", bridgeName); err != nil {
					return err
				}
			}

			if out, err = exec_utils.Cmdf(tctx, "ip addr show dev %s", bridgeName); err != nil {
				return err
			}
			gatewayAddress := fmt.Sprintf("%s/%s", port.Gateway, strings.Split(port.Subnet, "/")[1])
			if strings.Index(out, gatewayAddress) == -1 {
				if _, err = exec_utils.Cmdf(tctx, "ip addr add %s dev %s", gatewayAddress, bridgeName); err != nil {
					return err
				}
			}
			if strings.Index(out, "DOWN") > -1 {
				if _, err = exec_utils.Cmdf(tctx, "ip link set %s up", bridgeName); err != nil {
					return err
				}
			}
			fmt.Println("DEBUG interfaceMap", i, interfaceMap)

			interfaces = append(interfaces, libvirt_models.DeviceInterface{
				Type: "bridge",
				Driver: libvirt_models.InterfaceDriver{
					Name:   "vhost",
					Queues: 1,
				},
				Mac:    libvirt_models.InterfaceMac{Address: port.Mac},
				Source: libvirt_models.InterfaceSource{Bridge: bridgeName},
				Target: libvirt_models.InterfaceTarget{Dev: fmt.Sprintf("tap%d-%d", assignment.ID, i)},
				Model:  libvirt_models.InterfaceModel{Type: "virtio"},
				Alias:  libvirt_models.Alias{Name: fmt.Sprintf("net%d", i)},
				Address: libvirt_models.PciAddress{
					Type:     "pci",
					Domain:   "0x0000",
					Bus:      "0x00",
					Slot:     fmt.Sprintf("%#02X", pciSlot),
					Function: "0x0",
				},
			})
		}

		domainXml := libvirt_models.DomainXML{
			ID: assignment.ID,
			// Type: libvirt_models.DomainTypeKvm,
			Type: libvirt_models.DomainTypeQemu,
			Name: computeId,
			Memory: libvirt_models.Memory{
				Unit:   libvirt_models.UnitMb,
				Memory: compute.Memory,
			},
			CurrentMemory: libvirt_models.Memory{
				Unit:   libvirt_models.UnitMb,
				Memory: compute.Memory,
			},
			Vcpu: libvirt_models.Vcpu{
				Placement: libvirt_models.PlacementStatic,
				Vcpu:      compute.Vcpus,
			},
			Resource: libvirt_models.Resource{
				Partition: "/machine",
			},
			Os: libvirt_models.Os{
				Type: libvirt_models.OsType{
					Arch:    libvirt_models.OsTypeArchX8664,
					Machine: libvirt_models.OsTypeMachinePc,
					Type:    libvirt_models.OsTypeHvm,
				}, Boot: libvirt_models.OsBoot{Dev: "hd"},
			},
			Cpu: libvirt_models.Cpu{
				Mode:     libvirt_models.CpuModeHost,
				Model:    libvirt_models.CpuModel{Fallback: "forbid"},
				Topology: libvirt_models.CpuTopology{Sockets: 1, Cores: 1, Threads: 1},
			},
			Clock: libvirt_models.Clock{
				Offset: "utc",
			},
			OnPoweroff: "destroy",
			OnReboot:   "restart",
			OnCrash:    "restart",
			Devices: libvirt_models.Devices{
				Emulators:  emulators,
				Interfaces: interfaces,
				Disks:      disks,
				Serials:    serials,
				Consoles:   consoles,
			},
		}

		var xmlBytes []byte
		if xmlBytes, err = xml.Marshal(&domainXml); err != nil {
			return err
		}

		if err = ioutil.WriteFile(vmDomainXmlPath, xmlBytes, 0644); err != nil {
			return err
		}

		domain, domainExists := domainMap[computeId]
		if !domainExists {
			if _, err = exec_utils.Cmdf(tctx, "virsh define %s", vmDomainXmlPath); err != nil {
				return err
			}
			if _, err = exec_utils.Cmdf(tctx, "virsh start %s", computeId); err != nil {
				return err
			}
		} else {
			fmt.Println("DEBUG domain.State", domain.State)
			switch domain.State {
			case resource_model.StateDown:
				fmt.Println("DEBUG virsh start")
				if _, err = exec_utils.Cmdf(tctx, "virsh start %s", computeId); err != nil {
					return err
				}
			}
		}
	}

	// out, err := exec_utils.Cmdf(1, "brctl addbr %s", "test")
	return nil
}
