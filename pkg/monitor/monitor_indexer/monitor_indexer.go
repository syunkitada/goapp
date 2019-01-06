package monitor_indexer

import (
	"fmt"

	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/monitor/monitor_api/monitor_api_grpc_pb"
)

type Indexer interface {
	Report(logs []*monitor_api_grpc_pb.Log) error
}

func NewIndexer(indexerConfig *config.MonitorIndexerConfig) (Indexer, error) {
	switch indexerConfig.Driver {
	case "influxdb":
		return NewInfluxdbIndexer(indexerConfig)
	}

	return nil, fmt.Errorf("InvalidDriver: %v", indexerConfig.Driver)
}
