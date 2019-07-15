package libvirt_driver

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/syunkitada/goapp/pkg/lib/exec_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/lib/os_utils"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_agent/compute_drivers/libvirt_driver/libvirt_models"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

func (driver *LibvirtDriver) syncActivatingAssignmentMap(tctx *logger.TraceContext,
	assignmentMap map[uint]resource_model.ComputeAssignmentEx) error {
	fmt.Println("DEBUG driver", assignmentMap)
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
		fmt.Println("DEBUG assignment", assignment)
		compute := assignment.Spec.Compute
		computeId := fmt.Sprintf("compute%d", assignment.ID)
		pciSlot := 1

		vmDir := filepath.Join(driver.vmsDir, computeId)
		if err = os_utils.Mkdir(vmDir, 0755); err != nil {
			return err
		}
		vmImagePath := filepath.Join(vmDir, "img")
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

		emulators := []libvirt_models.DeviceEmulator{}
		emulators = append(emulators, libvirt_models.DeviceEmulator{
			Emulator: "/usr/bin/qemu-system-x86_64",
		})

		disks := []libvirt_models.DeviceDisk{}
		disks = append(disks, libvirt_models.DeviceDisk{
			Type:   "file",
			Device: "disk",
			Driver: libvirt_models.DiskDriver{
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
				Model:  libvirt_models.InterfaceModel{Type: ""},
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

		domain := libvirt_models.Domain{
			ID:   assignment.ID,
			Type: libvirt_models.DomainTypeKvm,
			Name: fmt.Sprintf("compute%d", assignment.ID),
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
					Arch: libvirt_models.OsTypeArchX8664,
					Type: libvirt_models.OsTypeHvm,
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
			},
		}

		var xmlBytes []byte
		if xmlBytes, err = xml.Marshal(&domain); err != nil {
			return err
		}

		if err = ioutil.WriteFile(vmDomainXmlPath, xmlBytes, 0644); err != nil {
			return err
		}
		fmt.Println("DEBUG XML", string(xmlBytes))

		if _, err = exec_utils.Cmdf(tctx, "virsh define %s", vmDomainXmlPath); err != nil {
			return err
		}
	}

	// out, err := exec_utils.Cmdf(1, "brctl addbr %s", "test")
	return nil
}
