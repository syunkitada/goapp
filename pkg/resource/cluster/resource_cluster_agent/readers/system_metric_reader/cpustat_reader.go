package system_metric_reader

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/lib/str_utils"
	"github.com/syunkitada/goapp/pkg/resource/config"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

type CpuProcessorStat struct {
	ReportStatus int // 0, 1(GetReport), 2(Reported)
	Timestamp    time.Time
	Processor    int64
	Mhz          int64
	User         int64
	Nice         int64
	System       int64
	Idle         int64
	Iowait       int64
	Irq          int64
	Softirq      int64
	Steal        int64
	Guest        int64
	GuestNice    int64
}

type CpuStat struct {
	ReportStatus  int // 0, 1(GetReport), 2(Reported)
	timestamp     time.Time
	intr          int64
	ctx           int64
	btime         int64
	processes     int64
	procs_running int64
	procs_blocked int64
	softirq       int64
}

type CpuStatReader struct {
	conf        *config.ResourceMetricSystemConfig
	cpus        []spec.NumaNodeCpuSpec
	cacheLength int
	cpuStats    []CpuStat
}

func NewCpuStatReader(conf *config.ResourceMetricSystemConfig, cpus []spec.NumaNodeCpuSpec) SubMetricReader {
	return &CpuStatReader{
		conf:        conf,
		cpus:        cpus,
		cacheLength: conf.CacheLength,
		cpuStats:    make([]CpuStat, 0, conf.CacheLength),
	}
}

func (reader *CpuStatReader) Read(tctx *logger.TraceContext) {
	timestamp := time.Now()
	var tmpReader *bufio.Reader

	// Read /proc/cpuinfo
	cpuinfo, _ := os.Open("/proc/cpuinfo")
	defer cpuinfo.Close()
	tmpReader = bufio.NewReader(cpuinfo)

	cpuProcessorStats := make([]CpuProcessorStat, len(reader.cpus))

	var tmpProcessor int
	var tmpCpuMhz int64
	var tmpBytes []byte
	var tmpErr error
	for {
		tmpBytes, _, tmpErr = tmpReader.ReadLine()
		if tmpErr != nil {
			break
		}

		splited := str_utils.SplitSpaceColon(string(tmpBytes))
		if len(splited) < 1 {
			continue
		}
		switch splited[0] {
		case "processor":
			tmpProcessor, _ = strconv.Atoi(splited[1])
		case "cpu MHz":
			tmpCpuMhz, _ = strconv.ParseInt(splited[1], 10, 64)
			for i := 0; i < 20; i++ {
				if _, _, tmpErr = tmpReader.ReadLine(); tmpErr != nil {
					break
				}
			}
		}
		cpuProcessorStats[tmpProcessor] = CpuProcessorStat{
			ReportStatus: 0,
			Timestamp:    timestamp,
			Mhz:          tmpCpuMhz,
			Processor:    int64(tmpProcessor),
		}
	}

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

	f, _ := os.Open("/proc/stat")
	defer f.Close()
	tmpReader = bufio.NewReader(f)

	tmpBytes, _, _ = tmpReader.ReadLine()

	for i := 0; i < len(reader.cpus); i++ {
		tmpBytes, _, _ = tmpReader.ReadLine()
		cpu := strings.Split(string(tmpBytes), " ")
		user, _ := strconv.ParseInt(cpu[1], 10, 64)
		nice, _ := strconv.ParseInt(cpu[2], 10, 64)
		system, _ := strconv.ParseInt(cpu[3], 10, 64)
		idle, _ := strconv.ParseInt(cpu[4], 10, 64)
		iowait, _ := strconv.ParseInt(cpu[5], 10, 64)
		irq, _ := strconv.ParseInt(cpu[6], 10, 64)
		softirq, _ := strconv.ParseInt(cpu[7], 10, 64)
		steal, _ := strconv.ParseInt(cpu[8], 10, 64)
		guest, _ := strconv.ParseInt(cpu[9], 10, 64)
		guestNice, _ := strconv.ParseInt(cpu[10], 10, 64)
		cpuProcessorStats[i].User = user
		cpuProcessorStats[i].Nice = nice
		cpuProcessorStats[i].System = system
		cpuProcessorStats[i].Idle = idle
		cpuProcessorStats[i].Iowait = iowait
		cpuProcessorStats[i].Irq = irq
		cpuProcessorStats[i].Softirq = softirq
		cpuProcessorStats[i].Steal = steal
		cpuProcessorStats[i].Guest = guest
		cpuProcessorStats[i].GuestNice = guestNice
	}

	tmpBytes, _, _ = tmpReader.ReadLine()
	intr, _ := strconv.ParseInt(strings.Split(string(tmpBytes), " ")[1], 10, 64)
	tmpBytes, _, _ = tmpReader.ReadLine()
	ctx, _ := strconv.ParseInt(strings.Split(string(tmpBytes), " ")[1], 10, 64)
	tmpBytes, _, _ = tmpReader.ReadLine()
	btime, _ := strconv.ParseInt(strings.Split(string(tmpBytes), " ")[1], 10, 64)
	tmpBytes, _, _ = tmpReader.ReadLine()
	processes, _ := strconv.ParseInt(strings.Split(string(tmpBytes), " ")[1], 10, 64)
	tmpBytes, _, _ = tmpReader.ReadLine()
	procs_running, _ := strconv.ParseInt(strings.Split(string(tmpBytes), " ")[1], 10, 64)
	tmpBytes, _, _ = tmpReader.ReadLine()
	procs_blocked, _ := strconv.ParseInt(strings.Split(string(tmpBytes), " ")[1], 10, 64)
	tmpBytes, _, _ = tmpReader.ReadLine()
	softirq, _ := strconv.ParseInt(strings.Split(string(tmpBytes), " ")[1], 10, 64)
	stat := CpuStat{
		ReportStatus:  0,
		timestamp:     timestamp,
		intr:          intr,
		ctx:           ctx,
		btime:         btime,
		processes:     processes,
		procs_running: procs_running,
		procs_blocked: procs_blocked,
		softirq:       softirq,
	}

	if len(reader.cpuStats) > reader.cacheLength {
		reader.cpuStats = reader.cpuStats[1:]
	}
	reader.cpuStats = append(reader.cpuStats, stat)

}

func (reader *CpuStatReader) ReportMetrics() (metrics []spec.ResourceMetric) {
	metrics = make([]spec.ResourceMetric, len(reader.cpuStats))

	for _, stat := range reader.cpuStats {
		// for cpuName, cpu := range stat.cpuMap {
		// 	metrics = append(metrics, spec.ResourceMetric{
		// 		Name: "system_cpu",
		// 		Time: timestamp,
		// 		Tag: map[string]string{
		// 			"cpu": cpuName,
		// 		},
		// 		Metric: map[string]interface{}{
		// 			"user":       cpu[0],
		// 			"nice":       cpu[1],
		// 			"system":     cpu[2],
		// 			"idle":       cpu[3],
		// 			"iowait":     cpu[4],
		// 			"irq":        cpu[5],
		// 			"softirq":    cpu[6],
		// 			"steal":      cpu[7],
		// 			"guest":      cpu[8],
		// 			"guest_nice": cpu[9],
		// 		},
		// 	})
		// }

		if stat.ReportStatus == ReportStatusReported {
			continue
		}

		metrics = append(metrics, spec.ResourceMetric{
			Name: "system_cpu",
			Time: stat.timestamp,
			Metric: map[string]interface{}{
				"intr":          stat.intr,
				"ctx":           stat.ctx,
				"btime":         stat.btime,
				"processes":     stat.processes,
				"procs_running": stat.procs_running,
				"procs_blocked": stat.procs_blocked,
				"softirq":       stat.softirq,
			},
		})

		stat.ReportStatus = 1
	}

	return
}

func (reader *CpuStatReader) ReportEvents() (events []spec.ResourceEvent) {
	return
}

func (reader *CpuStatReader) Reported() {
	for i := range reader.cpuStats {
		reader.cpuStats[i].ReportStatus = ReportStatusReported
	}
	return
}
