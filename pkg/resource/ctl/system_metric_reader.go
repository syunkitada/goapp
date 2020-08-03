package ctl

import (
	"fmt"
	"os"
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

	isUptimeStat := false
	if systemMetricReaderCmdTarget == "uptime" {
		isUptimeStat = true
	}
	isLoginStat := false
	if systemMetricReaderCmdTarget == "login" {
		isLoginStat = true
	}
	isBuddyinfoStat := false
	if systemMetricReaderCmdTarget == "buddyinfo" {
		isBuddyinfoStat = true
	}
	isFsStat := false
	if systemMetricReaderCmdTarget == "fs" {
		isFsStat = true
	}
	isDiskStat := false
	if systemMetricReaderCmdTarget == "disk" {
		isDiskStat = true
	}
	isVmStat := false
	if systemMetricReaderCmdTarget == "vm" {
		isVmStat = true
	}
	isMemStat := false
	if systemMetricReaderCmdTarget == "mem" {
		isMemStat = true
	}
	isCpuStat := false
	if systemMetricReaderCmdTarget == "cpu" {
		isCpuStat = true
	}
	isProcessorStat := false
	if systemMetricReaderCmdTarget == "processor" {
		isProcessorStat = true
	}

	var mega int64 = 1000000
	var giga int64 = 1000000000
	timeFormat := "15:04:05"
	for {
		time.Sleep(interval)
		if err = reader.Read(tctx); err != nil {
			return
		}

		metrics, _ := reader.Report()
		for _, metric := range metrics {
			switch metric.Name {
			case "system_uptime":
				if !isUptimeStat {
					continue
				}
				fmt.Printf("%s uptimestat: uptime=%d\n",
					metric.Time.Format(timeFormat),
					metric.Metric["uptime"],
				)
			case "system_login":
				if !isLoginStat {
					continue
				}
				fmt.Printf("%s loginstat: users=%d uniq_users=%d\n",
					metric.Time.Format(timeFormat),
					metric.Metric["users"],
					metric.Metric["uniq_users"],
				)
			case "system_buddyinfostat":
				if !isBuddyinfoStat {
					continue
				}
				fmt.Printf("%s buddyinfostat node=%s: 4K=%d, 8K=%d, 16K=%d, 32K=%d, 64K=%d, 128K=%d, 256K=%d, 512K=%d, 1M=%d, 2M=%d, 4M=%d\n",
					metric.Time.Format(timeFormat),
					metric.Tag["node"],
					metric.Metric["4K"], metric.Metric["8K"], metric.Metric["16K"], metric.Metric["32K"],
					metric.Metric["64K"], metric.Metric["128K"], metric.Metric["256K"], metric.Metric["512K"],
					metric.Metric["1M"], metric.Metric["2M"], metric.Metric["4M"],
				)
			case "system_cpu":
				if !isCpuStat {
					continue
				}
				fmt.Printf("%s cpustat: intr=%d ctx=%d btime=%d processes=%d running=%d blocked=%d softirq=%d\n",
					metric.Time.Format(timeFormat),
					metric.Metric["intr"], metric.Metric["ctx"], metric.Metric["btime"], metric.Metric["processes"],
					metric.Metric["procs_running"], metric.Metric["procs_blocked"], metric.Metric["softirq"],
				)
			case "system_processor":
				if !isProcessorStat {
					continue
				}
				fmt.Printf("%s cpustat processor=%s: mhz=%d user=%d nice=%d system=%d idle=%d iowait=%d irq=%d softirq=%d steal=%d guest=%d guestnice=%d\n",
					metric.Time.Format(timeFormat), metric.Tag["processor"],
					metric.Metric["mhz"], metric.Metric["user"], metric.Metric["nice"], metric.Metric["system"], metric.Metric["idle"],
					metric.Metric["iowait"], metric.Metric["irq"], metric.Metric["softirq"],
					metric.Metric["steal"], metric.Metric["guest"], metric.Metric["guestnice"],
				)
			case "system_diskstat":
				if !isDiskStat {
					continue
				}
				fmt.Printf("%s diskstat dev=%s: rps=%d, rbps=%d, rmsps=%d, wps=%d, wbps=%d, wmsps=%d, pios=%d\n",
					metric.Time.Format(timeFormat),
					metric.Tag["dev"],
					metric.Metric["reads_per_sec"], metric.Metric["read_bytes_per_sec"], metric.Metric["read_ms_per_sec"],
					metric.Metric["writes_per_sec"], metric.Metric["write_bytes_per_sec"], metric.Metric["write_ms_per_sec"],
					metric.Metric["progress_ios"],
				)
			case "system_fsstat":
				if !isFsStat {
					continue
				}
				fmt.Printf("%s fsstat path=%s mpath=%s: totalG=%d, freeG=%d, usedG=%d, files=%d\n",
					metric.Time.Format(timeFormat), metric.Tag["path"], metric.Tag["mount_path"],
					metric.Metric["total_size"].(int64)/giga, metric.Metric["free_size"].(int64)/giga, metric.Metric["used_size"].(int64)/giga, metric.Metric["files"],
				)
			case "system_vmstat":
				if !isVmStat {
					continue
				}
				fmt.Printf("%s vmstat: pskswapd=%d, psdirect=%d, pgfault=%d\n",
					metric.Time.Format(timeFormat),
					metric.Metric["pgscan_kswapd"], metric.Metric["pgscan_direct"],
					metric.Metric["pgfault"],
				)
			case "system_mem":
				if !isMemStat {
					continue
				}
				fmt.Printf("%s memstat node=%s: totalM=%d usedM=%d reclM=%d mlockedM=%d dirtyM=%d writebackM=%d slabM=%d\n",
					metric.Time.Format(timeFormat), metric.Tag["node_id"],
					metric.Metric["mem_total"].(int64)/mega, metric.Metric["mem_used"].(int64)/mega,
					metric.Metric["reclaimable"].(int64)/mega, metric.Metric["mlocked"].(int64)/mega,
					metric.Metric["dirty"].(int64)/mega, metric.Metric["writeback"].(int64)/mega, metric.Metric["slab"].(int64)/mega,
				)
			case "system_tcp_netstat":
				// TODO
				continue
			case "system_ip_netstat":
				// TODO
				continue
			case "system_netdevstat":
				// TODO
				continue
			case "system_procs":
				// TODO
				continue
			}
		}
		reader.Reported()
	}

	return
}
