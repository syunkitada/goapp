package qemu_driver

import (
	"io/ioutil"
	"path/filepath"

	"github.com/syunkitada/goapp/pkg/lib/exec_utils"
	"github.com/syunkitada/goapp/pkg/lib/json_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/lib/os_utils"
	"github.com/syunkitada/goapp/pkg/lib/template_utils"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_agent/compute_models"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

func (driver *QemuDriver) syncActivatingAssignmentMap(tctx *logger.TraceContext,
	assignmentMap map[uint]spec.ComputeAssignmentEx,
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
	assignment spec.ComputeAssignmentEx, netnsPorts []compute_models.NetnsPort) error {
	var err error
	compute := assignment.Spec

	vmDir := filepath.Join(driver.conf.VmsDir, compute.Name)
	vmImagePath := filepath.Join(vmDir, "img")
	vmConfigImagePath := filepath.Join(vmDir, "config.img")
	configDir := filepath.Join(vmDir, "config")
	vmServiceShFilePath := filepath.Join(vmDir, "service.sh")
	vmServiceFilePath := filepath.Join(driver.conf.SystemdDir, compute.Name+".service")
	vmMetaDataConfigFilePath := filepath.Join(configDir, "meta-data")
	vmUserDataConfigFilePath := filepath.Join(configDir, "user-data")

	if err = os_utils.Mkdir(vmDir, 0755); err != nil {
		return err
	}

	// Initialize Image
	srcImagePath := filepath.Join(driver.conf.ImagesDir, compute.ImageSpec.Name)
	if !os_utils.PathExists(srcImagePath) {
		tctx.SetTimeout(3600)

		switch compute.ImageSpec.Kind {
		case "Url":
			imageUrlSpec := compute.ImageSpec.Spec.(spec.ImageUrlSpec)
			if _, err = exec_utils.Cmdf(tctx, "wget -O %s %s", srcImagePath, imageUrlSpec.Url); err != nil {
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

	defaultGateway := netnsPorts[0].NetnsGateway
	if err = template_utils.Template(tctx, vmUserDataConfigFilePath, 0644, driver.conf.UserdataTmpl,
		map[string]interface{}{
			"DefaultGateway": defaultGateway,
			"Ports":          netnsPorts,
		}); err != nil {
		return err
	}

	if _, err = exec_utils.Cmdf(tctx, "genisoimage -o %s -V cidata -r -J %s %s",
		vmConfigImagePath, vmMetaDataConfigFilePath, vmUserDataConfigFilePath); err != nil {
		return err
	}

	if err = template_utils.Template(tctx, vmServiceShFilePath, 0755, driver.conf.VmServiceShTmpl,
		map[string]interface{}{
			"Compute":           compute,
			"Ports":             netnsPorts,
			"VmImagePath":       vmImagePath,
			"VmConfigImagePath": vmConfigImagePath,
		}); err != nil {
		return err
	}

	if err = template_utils.Template(tctx, vmServiceFilePath, 0755, driver.conf.VmServiceTmpl,
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
