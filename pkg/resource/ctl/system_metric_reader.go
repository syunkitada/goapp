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
	isProcsStat := false
	if systemMetricReaderCmdTarget == "procs" {
		isProcsStat = true
	}
	isProcStat := false
	if systemMetricReaderCmdTarget == "proc" {
		isProcStat = true
	}
	isNetDevStat := false
	if systemMetricReaderCmdTarget == "netdev" {
		isNetDevStat = true
	}
	isTcpNetStat := false
	if systemMetricReaderCmdTarget == "tcpnetstat" {
		isTcpNetStat = true
	}
	isIpNetStat := false
	if systemMetricReaderCmdTarget == "ipnetstat" {
		isIpNetStat = true
	}

	var kilo int64 = 1000
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
			case "system_netdevstat":
				if !isNetDevStat {
					continue
				}
				fmt.Printf("%s netdevstat inter=%s: rbps=%d rpps=%d rerrs=%d rdrops=%d tbps=%d tpps=%d terrs=%d tdrops=%d\n",
					metric.Time.Format(timeFormat), metric.Tag["interface"],
					metric.Metric["receive_bytes_per_sec"], metric.Metric["receive_packets_per_sec"],
					metric.Metric["receive_errors"], metric.Metric["receive_drops"],
					metric.Metric["transmit_bytes_per_sec"], metric.Metric["transmit_packets_per_sec"],
					metric.Metric["transmit_errors"], metric.Metric["transmit_drops"],
				)
				continue
			case "system_tcp_netstat":
				if !isTcpNetStat {
					continue
				}
				fmt.Printf("%s tcpnetstat: tw=%d twr=%d abort=%d abortf=%d ret=%d retf=%d drops=%d ldrops=%d lovers=%d dacks=%d\n",
					metric.Time.Format(timeFormat),
					metric.Metric["tw"], metric.Metric["tw_recycled"],
					metric.Metric["abort"], metric.Metric["abort_failed"],
					metric.Metric["retrans"], metric.Metric["retrans_failed"],
					metric.Metric["drops"],
					metric.Metric["listen_drops"], metric.Metric["listen_overflows"],
					metric.Metric["delayed_acks"],
				)
				continue
			case "system_ip_netstat":
				if !isIpNetStat {
					continue
				}
				fmt.Printf("%s ipnetstat: noroutes=%d truncatedpkts=%d csumerrors=%d\n",
					metric.Time.Format(timeFormat),
					metric.Metric["in_no_routes"], metric.Metric["in_truncated_pkts"], metric.Metric["in_csum_errors"],
				)
				continue
			case "system_procs":
				if !isProcsStat {
					continue
				}
				fmt.Printf("%s procsstat: procs=%d runs=%d sleeps=%d dsleeps=%d zonbies=%d others=%d\n",
					metric.Time.Format(timeFormat),
					metric.Metric["procs"], metric.Metric["runs"], metric.Metric["sleeps"],
					metric.Metric["disk_sleeps"], metric.Metric["zonbies"], metric.Metric["others"],
				)
				continue
			case "system_proc":
				if !isProcStat {
					continue
				}
				state := ""
				switch metric.Metric["state"].(int64) {
				case 3:
					state = "R"
				case 2:
					state = "D"
				case 1:
					state = "S"
				case 0:
					state = "N"
				case -1:
					state = "Z"
				}

				fmt.Printf("%s procsstat cmd=%s pid=%s: state=%s th=%d vmM=%d rssM=%d hpM=%d ctxs=%d nctxs=%d uutil=%d sutil=%d gutil=%d cgutil=%d\n",
					metric.Time.Format(timeFormat), metric.Tag["cmd"], metric.Tag["pid"], state, metric.Metric["threads"],
					metric.Metric["vm_size_kb"].(int64)/kilo, metric.Metric["vm_rss_kb"].(int64)/kilo, metric.Metric["hugetlb_pages"].(int64)/kilo,
					metric.Metric["voluntary_ctxt_switches"], metric.Metric["nonvoluntary_ctxt_switches"],
					metric.Metric["user_util"], metric.Metric["system_util"],
					metric.Metric["guest_util"], metric.Metric["cguest_util"],
				)
				continue
			}
		}
		reader.Reported()
	}

	return
}
