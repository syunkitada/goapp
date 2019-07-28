package qemu_driver

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"

	"github.com/syunkitada/goapp/pkg/lib/exec_utils"
	"github.com/syunkitada/goapp/pkg/lib/json_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/lib/os_utils"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_agent/compute_models"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

func (driver *QemuDriver) syncActivatingAssignmentMap(tctx *logger.TraceContext,
	assignmentMap map[uint]resource_model.ComputeAssignmentEx,
	computeNetnsPortsMap map[uint][]compute_models.NetnsPort) error {
	var err error
	for _, assignment := range assignmentMap {
		if err = driver.syncActivatingAssignment(tctx, assignment, computeNetnsPortsMap[assignment.ID]); err != nil {
			return err
		}
	}
	return nil
}

func (driver *QemuDriver) syncActivatingAssignment(tctx *logger.TraceContext,
	assignment resource_model.ComputeAssignmentEx, netnsPorts []compute_models.NetnsPort) error {
	var err error
	compute := assignment.Spec.Compute

	vmDir := filepath.Join(driver.vmsDir, compute.Name)
	vmImagePath := filepath.Join(vmDir, "img")
	vmConfigImagePath := filepath.Join(vmDir, "config.img")
	configDir := filepath.Join(vmDir, "config")
	vmMetaDataConfigFilePath := filepath.Join(configDir, "meta-data")
	vmUserDataConfigFilePath := filepath.Join(configDir, "user-data")

	if err = os_utils.Mkdir(vmDir, 0755); err != nil {
		return err
	}

	// Initialize Image
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

	// Initialize Network
	defaultGateway := ""
	for _, port := range compute.Ports {
		if port.Gateway != "" {
			defaultGateway = port.Gateway
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
	userdataFile, err := os.OpenFile(vmUserDataConfigFilePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer func() {
		if err := userdataFile.Close(); err != nil {
			logger.Warningf(tctx, "Failed close %s", vmUserDataConfigFilePath)
		}
	}()

	t := template.Must(template.ParseFiles(driver.userdataTmpl))
	t.Execute(userdataFile, map[string]interface{}{
		"DefaultGateway": defaultGateway,
		"Ports":          netnsPorts,
	})

	if _, err = exec_utils.Cmdf(tctx, "genisoimage -o %s -V config-2 -r -J %s",
		vmConfigImagePath, configDir); err != nil {
		return err
	}

	// TODO create systemd service and start vm

	return nil
}
