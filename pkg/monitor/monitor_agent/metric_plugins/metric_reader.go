package metric_plugins

import (
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/monitor/monitor_api/monitor_api_grpc_pb"
)

type MetricReader interface {
	GetName() string
	Read(tctx *logger.TraceContext) error
	Report() ([]*monitor_api_grpc_pb.Metric, []*monitor_api_grpc_pb.Alert)
	Reported()
}
