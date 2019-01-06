package monitor_indexer

import (
	"fmt"

	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/monitor/monitor_api/monitor_api_grpc_pb"
)

type InfluxdbIndexer struct {
}

func NewInfluxdbIndexer(indexerConfig *config.MonitorIndexerConfig) (*InfluxdbIndexer, error) {
	return &InfluxdbIndexer{}, nil
}

func (indexer *InfluxdbIndexer) Report(logs []*monitor_api_grpc_pb.Log) error {
	fmt.Println(logs)
	return nil
}
