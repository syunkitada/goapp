package system

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/monitor/monitor_agent/metric_plugins"
	"github.com/syunkitada/goapp/pkg/monitor/monitor_api/monitor_api_grpc_pb"
)

type SystemMetricReader struct {
	name         string
	enableCpu    bool
	enableMemory bool
	cpuCount     int
	procStats    []ProcStat
}

func NewSystemMetricReader(conf *config.MonitorMetricsSystemConfig) metric_plugins.MetricReader {
	// TODO FIX Calculate cpuCount
	f, _ := os.Open("/sys/devices/system/cpu/online")
	defer f.Close()
	b, _ := ioutil.ReadAll(f)
	cpus := strings.Split(string(b), "-")
	cpus0, err := strconv.Atoi(cpus[0])
	if err != nil {
		logger.StdoutFatalf("Failed Initialize SystemMetricReader: %v", err)
	}
	cpus1, err := strconv.Atoi(strings.TrimRight(cpus[1], "\n"))
	if err != nil {
		logger.StdoutFatalf("Failed Initialize SystemMetricReader: %v", err)
	}
	cpuCount := cpus1 - cpus0 + 1

	return &SystemMetricReader{
		name:         "system",
		enableCpu:    conf.EnableCpu,
		enableMemory: conf.EnableMemory,
		cpuCount:     cpuCount,
		procStats:    make([]ProcStat, 0, 10),
	}
}

type ProcStat struct {
	timestamp     time.Time
	cpuMap        map[string][]int64
	intr          int64
	ctx           int64
	btime         int64
	processes     int64
	procs_running int64
	procs_blocked int64
	softirq       int64
}

func (reader *SystemMetricReader) Read(tctx *logger.TraceContext) error {
	if reader.enableCpu {
		// Read /proc/stat
		//      user   nice system idle    iowait irq softirq steal guest guest_nice
		// cpu  264230 262  60792  8237284 20685  0   2652    0     0     0
		// cpu0 126387 2    30266  4124610 11105  0   1011    0     0     0
		// cpu1 137842 260  30525  4112674 9580   0   1641    0     0     0
		// intr 18316761 ...
		// ctxt 57087643
		// btime 1546819593
		// processes 227393
		// procs_running 1
		// procs_blocked 0
		// softirq 11650881 ...

		fmt.Println("READ /proc/stat")
		timestamp := time.Now()
		f, _ := os.Open("/proc/stat")
		defer f.Close()
		scanner := bufio.NewScanner(f)
		lines := make([]string, 0, reader.cpuCount+20)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}

		var cpu []string
		cpuMap := map[string][]int64{}
		lastIndex := reader.cpuCount + 1
		for i := 1; i < lastIndex; i++ {
			cpu = strings.Split(lines[i], " ")
			v := make([]int64, len(cpu)-1)
			for j, c := range cpu[1:] {
				v[j], _ = strconv.ParseInt(c, 10, 64)
			}
			cpuMap[cpu[0]] = v
		}

		intr, _ := strconv.ParseInt(strings.Split(lines[lastIndex], " ")[1], 10, 64)
		ctx, _ := strconv.ParseInt(strings.Split(lines[lastIndex+1], " ")[1], 10, 64)
		btime, _ := strconv.ParseInt(strings.Split(lines[lastIndex+2], " ")[1], 10, 64)
		processes, _ := strconv.ParseInt(strings.Split(lines[lastIndex+3], " ")[1], 10, 64)
		procs_running, _ := strconv.ParseInt(strings.Split(lines[lastIndex+4], " ")[1], 10, 64)
		procs_blocked, _ := strconv.ParseInt(strings.Split(lines[lastIndex+5], " ")[1], 10, 64)
		softirq, _ := strconv.ParseInt(strings.Split(lines[lastIndex+6], " ")[1], 10, 64)
		stat := ProcStat{
			timestamp:     timestamp,
			cpuMap:        cpuMap,
			intr:          intr,
			ctx:           ctx,
			btime:         btime,
			processes:     processes,
			procs_running: procs_running,
			procs_blocked: procs_blocked,
			softirq:       softirq,
		}
		reader.procStats = append(reader.procStats, stat)
	}

	if reader.enableMemory {
		// Read /proc/vmstat
		fmt.Println("READ Memory")
	}

	fmt.Println(reader.procStats)
	return nil
}

func (reader *SystemMetricReader) GetName() string {
	return reader.name
}

func (reader *SystemMetricReader) GetMetrics() []monitor_api_grpc_pb.Metric {
	metrics := make([]monitor_api_grpc_pb.Metric, 0, 100)

	// TODO convert to metrics
	// TODO check metrics and issue alerts

	metrics = append(metrics, monitor_api_grpc_pb.Metric{
		Name: "system_cpu",
		Time: "",
		Tag: map[string]string{
			"cpu": "cpu0",
		},
		Metric: map[string]int64{
			"user":   0,
			"nice":   0,
			"system": 0,
		},
	})

	return metrics
}
