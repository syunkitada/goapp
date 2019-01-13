package monitor_api

import (
	"fmt"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/syunkitada/goapp/pkg/base"
	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/monitor/monitor_api/monitor_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/monitor/monitor_indexer"
	"github.com/syunkitada/goapp/pkg/monitor/monitor_model/monitor_model_api"
)

type MonitorApiServer struct {
	base.BaseApp
	conf            *config.Config
	monitorModelApi *monitor_model_api.MonitorModelApi
	indexersMap     map[string][]monitor_indexer.Indexer
}

func NewMonitorApiServer(conf *config.Config) *MonitorApiServer {
	indexersMap := map[string][]monitor_indexer.Indexer{}
	for _, indexer := range conf.Monitor.Indexers {
		for _, index := range indexer.Indexes {
			newIndexer, err := monitor_indexer.NewIndexer(&indexer)
			if err != nil {
				logger.StdoutFatal(err)
			}
			if indexers, ok := indexersMap[index]; ok {
				indexersMap[index] = append(indexers, newIndexer)
			} else {
				indexersMap[index] = []monitor_indexer.Indexer{newIndexer}
			}
		}
	}

	conf.Monitor.ApiApp.Name = "monitor.api"
	server := MonitorApiServer{
		BaseApp:         base.NewBaseApp(conf, &conf.Monitor.ApiApp.AppConfig),
		conf:            conf,
		monitorModelApi: monitor_model_api.NewMonitorModelApi(conf),
		indexersMap:     indexersMap,
	}

	server.RegisterDriver(&server)

	return &server
}

func (srv *MonitorApiServer) RegisterGrpcServer(grpcServer *grpc.Server) error {
	monitor_api_grpc_pb.RegisterMonitorApiServer(grpcServer, srv)
	return nil
}

func (srv *MonitorApiServer) Status(ctx context.Context, statusRequest *monitor_api_grpc_pb.StatusRequest) (*monitor_api_grpc_pb.StatusReply, error) {
	return &monitor_api_grpc_pb.StatusReply{Msg: "Status"}, nil
}

//
// Node
//
func (srv *MonitorApiServer) GetNode(ctx context.Context, req *monitor_api_grpc_pb.GetNodeRequest) (*monitor_api_grpc_pb.GetNodeReply, error) {
	tctx := logger.NewGrpcTraceContext(srv.Host, srv.Name, ctx)
	startTime := logger.StartTrace(tctx)
	rep := srv.monitorModelApi.GetNode(tctx, req)
	logger.EndGrpcTrace(tctx, startTime, rep.StatusCode, rep.Err)
	return rep, nil
}

func (srv *MonitorApiServer) UpdateNode(ctx context.Context, req *monitor_api_grpc_pb.UpdateNodeRequest) (*monitor_api_grpc_pb.UpdateNodeReply, error) {
	tctx := logger.NewGrpcTraceContext(srv.Host, srv.Name, ctx)
	startTime := logger.StartTrace(tctx)
	rep := srv.monitorModelApi.UpdateNode(tctx, req)
	logger.EndGrpcTrace(tctx, startTime, rep.StatusCode, rep.Err)
	return rep, nil
}

//
// Report
//
func (srv *MonitorApiServer) Report(ctx context.Context, req *monitor_api_grpc_pb.ReportRequest) (*monitor_api_grpc_pb.ReportReply, error) {
	tctx := logger.NewGrpcTraceContext(srv.Host, srv.Name, ctx)
	startTime := logger.StartTrace(tctx)

	if indexers, ok := srv.indexersMap[req.Index]; ok {
		for _, indexer := range indexers {
			indexer.Report(tctx, req)
		}
	} else {
		logger.Warningf(tctx, fmt.Errorf("InvalidIndex"), "index=%v", req.Index)
	}

	rep := &monitor_api_grpc_pb.ReportReply{}
	logger.EndGrpcTrace(tctx, startTime, rep.StatusCode, rep.Err)
	return rep, nil
}

//
// Get
//
func (srv *MonitorApiServer) GetHost(ctx context.Context, req *monitor_api_grpc_pb.GetHostRequest) (*monitor_api_grpc_pb.GetHostReply, error) {
	tctx := logger.NewGrpcTraceContext(srv.Host, srv.Name, ctx)
	startTime := logger.StartTrace(tctx)

	hostMap := map[string]*monitor_api_grpc_pb.Host{}
	if indexers, ok := srv.indexersMap[req.Index]; ok {
		for _, indexer := range indexers {
			err := indexer.GetHost(tctx, req, hostMap)
			if err != nil {
				logger.Warningf(tctx, err, "Failed GetHost: index=%v", req.Index)
			}
		}
	} else {
		logger.Warningf(tctx, fmt.Errorf("InvalidIndex"), "index=%v", req.Index)
	}

	rep := &monitor_api_grpc_pb.GetHostReply{
		HostMap: hostMap,
	}
	logger.EndGrpcTrace(tctx, startTime, rep.StatusCode, rep.Err)
	return rep, nil
}

// GetUserState
func (srv *MonitorApiServer) GetUserState(ctx context.Context, req *monitor_api_grpc_pb.GetUserStateRequest) (*monitor_api_grpc_pb.GetUserStateReply, error) {
	tctx := logger.NewGrpcTraceContext(srv.Host, srv.Name, ctx)
	startTime := logger.StartTrace(tctx)

	indexMap := map[string]*monitor_api_grpc_pb.IndexState{}
	for name, _ := range srv.indexersMap {
		indexMap[name] = &monitor_api_grpc_pb.IndexState{
			Name: name,
		}
	}

	rep := &monitor_api_grpc_pb.GetUserStateReply{
		IndexMap: indexMap,
	}
	logger.EndGrpcTrace(tctx, startTime, rep.StatusCode, rep.Err)
	return rep, nil
}

func (srv *MonitorApiServer) GetLog(ctx context.Context, req *monitor_api_grpc_pb.GetLogRequest) (*monitor_api_grpc_pb.GetLogReply, error) {
	tctx := logger.NewGrpcTraceContext(srv.Host, srv.Name, ctx)
	startTime := logger.StartTrace(tctx)

	// if indexers, ok := srv.indexersMap[req.Index]; ok {
	// 	for _, indexer := range indexers {
	// 		indexer.Report(tctx, req.Logs)
	// 	}
	// } else {
	// 	logger.Warningf(tctx, fmt.Errorf("InvalidIndex"), "index=%v", req.Index)
	// }

	rep := &monitor_api_grpc_pb.GetLogReply{}
	logger.EndGrpcTrace(tctx, startTime, rep.StatusCode, rep.Err)
	return rep, nil
}
