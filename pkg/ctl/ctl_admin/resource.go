package ctl_admin

import (
// "github.com/golang/glog"
)

func (adminCtl *AdminCtl) ResourceRecreateDatabase() error {
	if err := adminCtl.ResourceModelApi.RecreateDatabase(); err != nil {
		return err
	}

	return nil
}
