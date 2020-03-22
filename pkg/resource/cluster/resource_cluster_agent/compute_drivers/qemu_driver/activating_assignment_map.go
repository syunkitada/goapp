package qemu_driver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/syunkitada/goapp/pkg/lib/exec_utils"
	"github.com/syunkitada/goapp/pkg/lib/ip_utils"
	"github.com/syunkitada/goapp/pkg/lib/json_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/lib/os_utils"
	"github.com/syunkitada/goapp/pkg/lib/template_utils"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_agent/compute_models"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

type VmMetadata struct {
	netnsPorts []compute_models.NetnsPort
}

func (driver *QemuDriver) syncActivatingAssignmentMap(tctx *logger.TraceContext,
	assignmentMap map[uint]spec.ComputeAssignmentEx,
	dummyComputeNetnsPortsMap map[uint][]compute_models.NetnsPort) (err error) {

	// 既存のvmMetadataを集める
	files, err := ioutil.ReadDir(driver.conf.VmsDir)
	if err != nil {
		return
	}
	assignedNetnsIds := make([]bool, 4096)
	vmMetadataMap := map[string]VmMetadata{}
	for _, file := range files {
		if file.IsDir() {
			metadataFilePath := filepath.Join(driver.conf.VmsDir, file.Name(), "metadata")
			if _, tmpErr := os.Stat(metadataFilePath); tmpErr != nil {
				logger.Warningf(tctx, "Invalid vm directory: %s", tmpErr.Error())
				continue
			}
			var vmMetadata VmMetadata
			os_utils.LoadDataFile(tctx, metadataFilePath, &vmMetadata)
			vmMetadataMap[file.Name()] = vmMetadata
			for i := range vmMetadata.netnsPorts {
				port := vmMetadata.netnsPorts[i]
				assignedNetnsIds[port.Id] = true
			}
		}
	}

	for _, assignment := range assignmentMap {
		vmMetadata, ok := vmMetadataMap[assignment.Spec.Name]
		if !ok {
			vmMetadata = VmMetadata{}
		}

		for j, port := range assignment.Spec.Ports {
			existPort := false
			for _, netnsPort := range vmMetadata.netnsPorts {
				if netnsPort.VmIp == port.Ip {
					existPort = true
					break
				}
			}
			if !existPort {
				// インターフェイスの最大文字数が15なので、ベース文字数は12とする
				var netnsId int
				for id, assigned := range assignedNetnsIds {
					if !assigned {
						netnsId = id
						assignedNetnsIds[id] = true
						break
					}
				}

				netnsName := fmt.Sprintf("com-%d", netnsId)
				netnsGateway := ip_utils.AddIntToIp(driver.conf.VmNetnsGatewayStartIp, j)
				netnsIp := ip_utils.AddIntToIp(driver.conf.VmNetnsStartIp, netnsId)
				netnsPort := compute_models.NetnsPort{
					Id:           netnsId,
					Name:         netnsName,
					NetnsGateway: netnsGateway.String(),
					NetnsIp:      netnsIp.String(),
					VmGateway:    port.Gateway,
					VmIp:         port.Ip,
					VmMac:        port.Mac,
					VmSubnet:     port.Subnet,
					Kind:         port.Kind,
				}

				switch port.Kind {
				case "Local":
					var netSpec spec.NetworkV4LocalSpec
					if tmpErr := json_utils.Unmarshal(port.Spec, &netSpec); tmpErr != nil {
						logger.Warningf(tctx, "Invalid port.Spec: %s", tmpErr.Error())
						continue
					}
					netnsPort.NetworkV4LocalSpec = netSpec
				}

				vmMetadata.netnsPorts = append(vmMetadata.netnsPorts, netnsPort)
			}
		}

		if err = driver.syncActivatingAssignment(tctx, assignment, &vmMetadata); err != nil {
			return err
		}
	}
	return nil
}

func (driver *QemuDriver) syncActivatingAssignment(tctx *logger.TraceContext,
	assignment spec.ComputeAssignmentEx, vmMetadata *VmMetadata) error {
	var err error
	compute := assignment.Spec

	vmDir := filepath.Join(driver.conf.VmsDir, compute.Name)
	vmImagePath := filepath.Join(vmDir, "img")
	vmMonitorSocketPath := filepath.Join(vmDir, "monitor.sock")
	vmSerialSocketPath := filepath.Join(vmDir, "serial.sock")
	vmConfigImagePath := filepath.Join(vmDir, "config.img")
	configDir := filepath.Join(vmDir, "config")
	vmMetadataFilePath := filepath.Join(vmDir, "metadata")
	vmServiceShFilePath := filepath.Join(vmDir, "service.sh")
	vmServiceFilePath := filepath.Join(driver.conf.SystemdDir, compute.Name+".service")
	vmMetaDataConfigFilePath := filepath.Join(configDir, "meta-data")
	vmUserDataConfigFilePath := filepath.Join(configDir, "user-data")

	if err = os_utils.Mkdir(vmDir, 0755); err != nil {
		return err
	}

	if err = os_utils.SaveDataFileIfNotExist(tctx, vmMetadataFilePath, &vmMetadata); err != nil {
		return err
	}

	// Initialize Image
	srcImagePath := filepath.Join(driver.conf.ImagesDir, compute.ImageSpec.Name)
	if !os_utils.PathExists(srcImagePath) {
		tctx.SetTimeout(3600)
		var specBytes []byte
		switch compute.ImageSpec.Kind {
		case "Url":
			specBytes, err = json.Marshal(&compute.ImageSpec.Spec)
			if err != nil {
				return err
			}
			var imageUrlSpec spec.ImageUrlSpec
			if err = json.Unmarshal(specBytes, &imageUrlSpec); err != nil {
				return err
			}
			splitedUrl := strings.Split(imageUrlSpec.Url, "/")
			tmpSrcImagePath := filepath.Join(driver.conf.ImagesDir, splitedUrl[len(splitedUrl)-1])
			if !os_utils.PathExists(tmpSrcImagePath) {
				if _, err = exec_utils.Cmdf(tctx, "wget -O %s %s", tmpSrcImagePath, imageUrlSpec.Url); err != nil {
					return err
				}
			}
			var outputPath string
			if outputPath, err = os_utils.UnArchiveFile(tctx, tmpSrcImagePath); err != nil {
				return err
			}
			if _, err = exec_utils.Cmdf(tctx, "mv %s %s", outputPath, srcImagePath); err != nil {
				return err
			}
		}
	}

	if !os_utils.PathExists(vmImagePath) {
		if _, err = exec_utils.Cmdf(tctx, "cp %s %s", srcImagePath, vmImagePath); err != nil {
			return err
		}
	}

	// Create ConfigDrive
	if err = os_utils.Mkdir(configDir, 0755); err != nil {
		return err
	}

	metaData := map[string]interface{}{
		"hostname": compute.Name,
	}
	metaDataBytes, err := json_utils.YamlMarshal(metaData)
	if err != nil {
		return err
	}
	if err = ioutil.WriteFile(vmMetaDataConfigFilePath, []byte(metaDataBytes), 0644); err != nil {
		return err
	}

	var resolvers []spec.Resolver
	for _, port := range vmMetadata.netnsPorts {
		switch port.Kind {
		case "Local":
			for _, resolver := range port.NetworkV4LocalSpec.Resolvers {
				exists := false
				for _, r := range resolvers {
					if r.Resolver == resolver.Resolver {
						exists = true
					}
				}
				if !exists {
					resolvers = append(resolvers, resolver)
				}
			}
		}
	}

	defaultGateway := vmMetadata.netnsPorts[0].VmGateway
	if err = template_utils.ExecTemplate(tctx, driver.userDataTmpl, vmUserDataConfigFilePath, 0644,
		map[string]interface{}{
			"DefaultGateway": defaultGateway,
			"Ports":          vmMetadata.netnsPorts,
			"Resolvers":      resolvers,
		}); err != nil {
		return err
	}

	if _, err = exec_utils.Cmdf(tctx, "genisoimage -o %s -V cidata -r -J %s %s",
		vmConfigImagePath, vmMetaDataConfigFilePath, vmUserDataConfigFilePath); err != nil {
		return err
	}

	if err = template_utils.ExecTemplate(tctx, driver.vmServiceShTmpl, vmServiceShFilePath, 0755,
		map[string]interface{}{
			"Compute":           compute,
			"Ports":             vmMetadata.netnsPorts,
			"VmImagePath":       vmImagePath,
			"VmConfigImagePath": vmConfigImagePath,
			"MonitorSocketPath": vmMonitorSocketPath,
			"SerialSocketPath":  vmSerialSocketPath,
		}); err != nil {
		return err
	}

	if err = template_utils.ExecTemplate(tctx, driver.vmServiceTmpl, vmServiceFilePath, 0755,
		map[string]interface{}{
			"Compute":             compute,
			"VmServiceShFilePath": vmServiceShFilePath,
		}); err != nil {
		return err
	}

	tctx.SetTimeout(5)
	if _, err = exec_utils.Cmdf(tctx, "systemctl daemon-reload"); err != nil {
		return err
	}

	if _, err = exec_utils.Cmdf(tctx, "systemctl start %s", compute.Name); err != nil {
		return err
	}

	return nil
}
