package resource_controller

import (
	"fmt"
	"sync"

	"golang.org/x/net/context"

	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

func (srv *ResourceControllerServer) MainTask(traceId string) error {
	if err := srv.UpdateNode(traceId); err != nil {
		return err
	}
	if err := srv.SyncRole(traceId); err != nil {
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
	go srv.SyncCompute(traceId, &wg)
	wg.Wait()

	// TODO
	// implement with goroutine
	// check compute
	// check container
	// check image
	// check loadbalancer

	return nil
}

func (srv *ResourceControllerServer) UpdateNode(traceId string) error {
	startTime := logger.StartTaskTrace(traceId, srv.Host, srv.Name)
	req := &resource_api_grpc_pb.UpdateNodeRequest{
		TraceId:      traceId,
		Name:         srv.Host,
		Kind:         resource_model.KindResourceController,
		Role:         resource_model.RoleMember,
		Status:       resource_model.StatusEnabled,
		StatusReason: "Default",
		State:        resource_model.StateUp,
		StateReason:  "UpdateNode",
	}

	rep := srv.resourceModelApi.UpdateNode(req)
	logger.EndTaskTrace(traceId, srv.Host, srv.Name, startTime, rep.Err)
	if rep.Err != "" {
		return fmt.Errorf(rep.Err)
	}

	return nil
}

func (srv *ResourceControllerServer) SyncRole(traceId string) error {
	startTime := logger.StartTaskTrace(traceId, srv.Host, srv.Name)
	nodes, err := srv.resourceModelApi.SyncRole(resource_model.KindResourceController)
	if err != nil {
		logger.EndTaskTrace(traceId, srv.Host, srv.Name, startTime, err)
		return err
	}

	existsSelfNode := false
	existsActiveLeader := false
	for _, node := range nodes {
		if node.Kind != resource_model.KindResourceController {
			continue
		}
		if node.Name == srv.conf.Default.Host && node.Status == resource_model.StatusEnabled && node.State == resource_model.StateUp {
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
		err = fmt.Errorf("This node is not activated")
		logger.EndTaskTrace(traceId, srv.Host, srv.Name, startTime, err)
		return err
	}

	if !existsActiveLeader {
		err = fmt.Errorf("Active Leader is not exists, after ReassignNode")
		logger.EndTaskTrace(traceId, srv.Host, srv.Name, startTime, err)
		return err
	}

	logger.EndTaskTrace(traceId, srv.Host, srv.Name, startTime, "")
	return nil
}

func (srv *ResourceControllerServer) SyncCompute(traceId string, wg *sync.WaitGroup) {
	startTime := logger.StartTaskTrace(traceId, srv.Host, srv.Name)
	defer func() { wg.Done() }()

	errChan := make(chan error)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, srv.syncResourceTimeout)
	defer cancel()

	go func() {
		errChan <- srv.resourceModelApi.SyncCompute(traceId)
	}()

	select {
	case err := <-errChan:
		if err != nil {
			logger.EndTaskTrace(traceId, srv.Host, srv.Name, startTime, "Failed SyncCompute: %v", err)
		} else {
			logger.EndTaskTrace(traceId, srv.Host, srv.Name, startTime, "")
		}
	case <-ctx.Done():
		logger.EndTaskTrace(traceId, srv.Host, srv.Name, startTime, "Failed SyncCompute: %v", ctx.Err())
	}
}
