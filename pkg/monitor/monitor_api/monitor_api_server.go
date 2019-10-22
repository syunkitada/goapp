package monitor_api

import (
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
			newIndexer, err := monitor_indexer.NewIndexer(index, &indexer)
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
// IgnoreAlert
//
func (srv *MonitorApiServer) GetIgnoreAlert(ctx context.Context, req *monitor_api_grpc_pb.GetIgnoreAlertRequest) (*monitor_api_grpc_pb.GetIgnoreAlertReply, error) {
	tctx := logger.NewGrpcTraceContext(srv.Host, srv.Name, ctx)
	startTime := logger.StartTrace(tctx)
	rep := srv.monitorModelApi.GetIgnoreAlert(tctx, req)
	logger.EndGrpcTrace(tctx, startTime, rep.StatusCode, rep.Err)
	return rep, nil
}

func (srv *MonitorApiServer) CreateIgnoreAlert(ctx context.Context, req *monitor_api_grpc_pb.CreateIgnoreAlertRequest) (*monitor_api_grpc_pb.CreateIgnoreAlertReply, error) {
	tctx := logger.NewGrpcTraceContext(srv.Host, srv.Name, ctx)
	startTime := logger.StartTrace(tctx)
	rep := srv.monitorModelApi.CreateIgnoreAlert(tctx, req)
	logger.EndGrpcTrace(tctx, startTime, rep.StatusCode, rep.Err)
	return rep, nil
}

func (srv *MonitorApiServer) UpdateIgnoreAlert(ctx context.Context, req *monitor_api_grpc_pb.UpdateIgnoreAlertRequest) (*monitor_api_grpc_pb.UpdateIgnoreAlertReply, error) {
	tctx := logger.NewGrpcTraceContext(srv.Host, srv.Name, ctx)
	startTime := logger.StartTrace(tctx)
	rep := srv.monitorModelApi.UpdateIgnoreAlert(tctx, req)
	logger.EndGrpcTrace(tctx, startTime, rep.StatusCode, rep.Err)
	return rep, nil
}

func (srv *MonitorApiServer) DeleteIgnoreAlert(ctx context.Context, req *monitor_api_grpc_pb.DeleteIgnoreAlertRequest) (*monitor_api_grpc_pb.DeleteIgnoreAlertReply, error) {
	tctx := logger.NewGrpcTraceContext(srv.Host, srv.Name, ctx)
	startTime := logger.StartTrace(tctx)
	rep := srv.monitorModelApi.DeleteIgnoreAlert(tctx, req)
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
		logger.Warningf(tctx, "index=%v", req.Index)
	}

	rep := &monitor_api_grpc_pb.ReportReply{}
	logger.EndGrpcTrace(tctx, startTime, rep.StatusCode, rep.Err)
	return rep, nil
}

//
// Get
//
func (srv *MonitorApiServer) GetIndex(ctx context.Context, req *monitor_api_grpc_pb.GetIndexRequest) (*monitor_api_grpc_pb.GetIndexReply, error) {
	tctx := logger.NewGrpcTraceContext(srv.Host, srv.Name, ctx)
	startTime := logger.StartTrace(tctx)

	indexMap := map[string]*monitor_api_grpc_pb.Index{}
	for index, indexers := range srv.indexersMap {
		for _, indexer := range indexers {
			err := indexer.GetIndex(tctx, req.ProjectName, indexMap)
			if err != nil {
				logger.Warningf(tctx, "Failed GetIndex: index=%v, err=%v", index, err)
			}
		}
	}

	rep := &monitor_api_grpc_pb.GetIndexReply{
		IndexMap: indexMap,
	}
	logger.EndGrpcTrace(tctx, startTime, rep.StatusCode, rep.Err)
	return rep, nil
}

func (srv *MonitorApiServer) GetHost(ctx context.Context, req *monitor_api_grpc_pb.GetHostRequest) (*monitor_api_grpc_pb.GetHostReply, error) {
	tctx := logger.NewGrpcTraceContext(srv.Host, srv.Name, ctx)
	startTime := logger.StartTrace(tctx)

	hostMap := map[string]*monitor_api_grpc_pb.Host{}
	if req.Index == "all" {
		for index, indexers := range srv.indexersMap {
			for _, indexer := range indexers {
				err := indexer.GetHost(tctx, req.ProjectName, hostMap)
				if err != nil {
					logger.Warningf(tctx, "Failed GetHost: index=%v, err=%v", index, err)
				}
			}
		}
	} else {
		if indexers, ok := srv.indexersMap[req.Index]; ok {
			for _, indexer := range indexers {
				err := indexer.GetHost(tctx, req.ProjectName, hostMap)
				if err != nil {
					logger.Warningf(tctx, "Failed GetHost: index=%v, err=%v", req.Index, err)
				}
			}
		} else {
			logger.Warningf(tctx, "InvalidIndex: index=%v", req.Index)
		}
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

	indexMap := map[string]*monitor_api_grpc_pb.Index{}
	for index, indexers := range srv.indexersMap {
		for _, indexer := range indexers {
			err := indexer.GetIndex(tctx, req.ProjectName, indexMap)
			if err != nil {
				logger.Warningf(tctx, "Failed GetIndex: index=%v, err=%v", index, err)
			}
		}
	}

	rep := &monitor_api_grpc_pb.GetUserStateReply{
		IndexMap: indexMap,
	}
	logger.EndGrpcTrace(tctx, startTime, rep.StatusCode, rep.Err)
	return rep, nil
}

// GetIndexState
func (srv *MonitorApiServer) GetIndexState(ctx context.Context, req *monitor_api_grpc_pb.GetIndexStateRequest) (*monitor_api_grpc_pb.GetIndexStateReply, error) {
	tctx := logger.NewGrpcTraceContext(srv.Host, srv.Name, ctx)
	startTime := logger.StartTrace(tctx)

	hostMap := map[string]*monitor_api_grpc_pb.Host{}
	if indexers, ok := srv.indexersMap[req.Index]; ok {
		for _, indexer := range indexers {
			err := indexer.GetHost(tctx, req.ProjectName, hostMap)
			if err != nil {
				logger.Warningf(tctx, "Failed GetHost: index=%v, err=%v", req.Index, err)
			}
		}
	} else {
		logger.Warningf(tctx, "InvalidIndex: index=%v", req.Index)
	}

	rep := &monitor_api_grpc_pb.GetIndexStateReply{
		HostMap: hostMap,
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
