package resource_controller

import (
	"fmt"
	"sync"
	// "time"

	"github.com/golang/glog"
	"golang.org/x/net/context"

	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

func (srv *ResourceControllerServer) MainTask() error {
	glog.Info("Run MainTask")
	if err := srv.UpdateNode(); err != nil {
		return err
	}
	if err := srv.SyncRole(); err != nil {
		return err
	}
	if srv.role == resource_model.RoleMember {
		return nil
	}

	if err := srv.resourceModelApi.CheckNodes(); err != nil {
		return err
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go srv.SyncCompute(&wg)
	wg.Wait()

	// TODO
	// implement with goroutine
	// check compute
	// check container
	// check image
	// check loadbalancer

	return nil
}

func (srv *ResourceControllerServer) UpdateNode() error {
	req := resource_api_grpc_pb.UpdateNodeRequest{
		Name:         srv.conf.Default.Name,
		Kind:         resource_model.KindResourceController,
		Role:         resource_model.RoleMember,
		Status:       resource_model.StatusEnabled,
		StatusReason: "Default",
		State:        resource_model.StateUp,
		StateReason:  "UpdateNode",
	}
	if _, err := srv.resourceApiClient.UpdateNode(&req); err != nil {
		return err
	}

	glog.Info("UpdatedNode")
	return nil
}

func (srv *ResourceControllerServer) SyncRole() error {
	var err error
	nodes, err := srv.resourceModelApi.SyncRole(resource_model.KindResourceController)
	if err != nil {
		return err
	}

	existsSelfNode := false
	existsActiveLeader := false
	for _, node := range nodes {
		if node.Kind != resource_model.KindResourceController {
			continue
		}
		if node.Name == srv.conf.Default.Name && node.Status == resource_model.StatusEnabled && node.State == resource_model.StateUp {
			existsSelfNode = true
			srv.role = node.Role
		}
		if node.Status == resource_model.StatusEnabled && node.State == resource_model.StateUp {
			if node.Role == resource_model.RoleLeader {
				existsActiveLeader = true
			}
		}
	}

	if !existsSelfNode {
		return fmt.Errorf("This node is not activated")
	}

	if !existsActiveLeader {
		return fmt.Errorf("Active Leader is not exists, after ReassignNode")
	}

	glog.Infof("Completed SyncRole: role=%v", srv.role)
	return nil
}

func (srv *ResourceControllerServer) SyncCompute(wg *sync.WaitGroup) {
	defer func() { wg.Done() }()
	var err error

	errChan := make(chan error)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, srv.syncResourceTimeout)
	defer cancel()

	go func() {
		errChan <- srv.resourceModelApi.SyncCompute()
	}()

	select {
	case err = <-errChan:
		if err != nil {
			glog.Errorf("Failed SyncCompute: %v", err)
		} else {
			glog.Info("Complete SyncCompute")
		}
	case <-ctx.Done():
		glog.Errorf("Failed SyncCompute: %v", ctx.Err())
	}
}
