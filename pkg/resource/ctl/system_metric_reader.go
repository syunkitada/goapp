package ctl

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_agent/readers/system_metric_reader"
	"github.com/syunkitada/goapp/pkg/resource/config"
)

var (
	systemMetricReaderCmdInterval int
	systemMetricReaderCmdTarget   string
)

var systemMetricReaderCmd = &cobra.Command{
	Use:   "system-metric-reader",
	Short: "system-metric-reader",
	Run: func(cmd *cobra.Command, args []string) {
		ctl := NewCtl(&config.BaseConf, &config.MainConf)
		if tmpErr := ctl.SystemMetricReader(); tmpErr != nil {
			logger.StdoutFatalf("Failed SystemMetricReader: %s\n", tmpErr.Error())
			os.Exit(1)
		}
	},
}

func init() {
	systemMetricReaderCmd.PersistentFlags().IntVarP(&systemMetricReaderCmdInterval, "interval", "i", 1, "interval")
	systemMetricReaderCmd.PersistentFlags().StringVarP(&systemMetricReaderCmdTarget, "target", "t", "all", "metrics target")
	RootCmd.AddCommand(systemMetricReaderCmd)
}

func (ctl *Ctl) SystemMetricReader() (err error) {
	tctx := logger.NewTraceContext(ctl.baseConf.Host, "system-metric-reader")
	clusterConf, ok := ctl.mainConf.Resource.ClusterMap[ctl.mainConf.Resource.ClusterName]
	if !ok {
		err = fmt.Errorf("Invalid conf: cluster is not found: cluster=%s", ctl.mainConf.Resource.ClusterName)
	}

	interval := time.Duration(systemMetricReaderCmdInterval) * time.Second

	reader := system_metric_reader.New(&clusterConf.Agent.Metric.System)
	if err = reader.Read(tctx); err != nil {
		return
	}
	time.Sleep(interval)
	_, _ = reader.Report()

	for {
		time.Sleep(interval)
		if err = reader.Read(tctx); err != nil {
			return
		}

		metrics, _ := reader.Report()
		for _, metric := range metrics {
			switch metric.Name {
			case "system_diskstat":
				timestampInt, _ := strconv.ParseInt(metric.Time, 10, 64)
				timestampInt /= 1000000000
				timestamp := time.Unix(timestampInt, 0)
				fmt.Printf("%s diskstat dev=%s: rps=%d, rbps=%d, rmsps=%d, wps=%d, wbps=%d, wmsps=%d, pios=%d\n",
					timestamp.Format("15:04:05"),
					metric.Tag["dev"],
					metric.Metric["reads_per_sec"], metric.Metric["read_bytes_per_sec"], metric.Metric["read_ms_per_sec"],
					metric.Metric["writes_per_sec"], metric.Metric["write_bytes_per_sec"], metric.Metric["write_ms_per_sec"],
					metric.Metric["progress_ios"],
				)
			}
		}
		reader.Reported()
	}

	return
}
