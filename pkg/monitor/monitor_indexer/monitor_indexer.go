package monitor_indexer

import (
	"fmt"

	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/monitor/monitor_api/monitor_api_grpc_pb"
)

type Indexer interface {
	Report(tctx *logger.TraceContext, req *monitor_api_grpc_pb.ReportRequest) error
	GetHost(tctx *logger.TraceContext, projectName string, hostMap map[string]*monitor_api_grpc_pb.Host) error
}

func NewIndexer(index string, indexerConfig *config.MonitorIndexerConfig) (Indexer, error) {
	switch indexerConfig.Driver {
	case "influxdb":
		return NewInfluxdbIndexer(index, indexerConfig)
	}

	return nil, fmt.Errorf("InvalidDriver: %v", indexerConfig.Driver)
}
