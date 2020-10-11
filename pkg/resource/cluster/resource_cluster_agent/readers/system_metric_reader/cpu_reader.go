package system_metric_reader

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/lib/str_utils"
	"github.com/syunkitada/goapp/pkg/resource/config"
	"github.com/syunkitada/goapp/pkg/resource/consts"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

type CpuProcessorStat struct {
	ReportStatus int
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

type TmpCpuProcessorStat struct {
	Processor int64
	Mhz       int64
	User      int64
	Nice      int64
	System    int64
	Idle      int64
	Iowait    int64
	Irq       int64
	Softirq   int64
	Steal     int64
	Guest     int64
	GuestNice int64
}

type CpuStat struct {
	ReportStatus int
	Timestamp    time.Time
	Intr         int64
	Ctx          int64
	Btime        int64
	Processes    int64
	ProcsRunning int64
	ProcsBlocked int64
	Softirq      int64
}

type TmpCpuStat struct {
	Intr         int64
	Ctx          int64
	Btime        int64
	Processes    int64
	ProcsRunning int64
	ProcsBlocked int64
	Softirq      int64
}

type CpuReader struct {
	conf                 *config.ResourceMetricSystemConfig
	cpus                 []spec.NumaNodeCpuSpec
	cacheLength          int
	cpuStats             []CpuStat
	cpuProcessorStats    []CpuProcessorStat
	tmpCpuStat           *TmpCpuStat
	tmpCpuProcessorStats []TmpCpuProcessorStat

	checkProcsRunningWarnLimit       int64
	checkProcsRunningCritLimit       int64
	checkProcsRunningOccurences      int
	checkProcsRunningReissueDuration int
	checkProcsRunningWarnCounter     int
	checkProcsRunningCritCounter     int

	checkProcsBlockedWarnLimit       int64
	checkProcsBlockedCritLimit       int64
	checkProcsBlockedOccurences      int
	checkProcsBlockedReissueDuration int
	checkProcsBlockedWarnCounter     int
	checkProcsBlockedCritCounter     int
}

func NewCpuReader(conf *config.ResourceMetricSystemConfig, cpus []spec.NumaNodeCpuSpec) SubMetricReader {
	var checkProcsRunningWarnLimit int64 = -1
	var checkProcsRunningCritLimit int64 = -1
	if conf.Cpu.CheckProcsRunning.WarnRateLimit > 0 {
		checkProcsRunningWarnLimit = int64(float64(len(cpus)) * conf.Cpu.CheckProcsRunning.WarnRateLimit)
	}
	if conf.Cpu.CheckProcsRunning.CritRateLimit > 0 {
		checkProcsRunningCritLimit = int64(float64(len(cpus)) * conf.Cpu.CheckProcsRunning.CritRateLimit)
	}

	var checkProcsBlockedWarnLimit int64 = -1
	var checkProcsBlockedCritLimit int64 = -1
	if conf.Cpu.CheckProcsBlocked.WarnRateLimit > 0 {
		checkProcsBlockedWarnLimit = int64(float64(len(cpus)) * conf.Cpu.CheckProcsBlocked.WarnRateLimit)
	}
	if conf.Cpu.CheckProcsBlocked.CritRateLimit > 0 {
		checkProcsBlockedCritLimit = int64(float64(len(cpus)) * conf.Cpu.CheckProcsBlocked.CritRateLimit)
	}

	return &CpuReader{
		conf:        conf,
		cpus:        cpus,
		cacheLength: conf.CacheLength,
		cpuStats:    make([]CpuStat, 0, conf.CacheLength),

		checkProcsRunningWarnLimit:       checkProcsRunningWarnLimit,
		checkProcsRunningCritLimit:       checkProcsRunningCritLimit,
		checkProcsRunningOccurences:      conf.Cpu.CheckProcsRunning.Occurences,
		checkProcsRunningReissueDuration: conf.Cpu.CheckProcsRunning.ReissueDuration,
		checkProcsRunningWarnCounter:     0,
		checkProcsRunningCritCounter:     0,

		checkProcsBlockedWarnLimit:       checkProcsBlockedWarnLimit,
		checkProcsBlockedCritLimit:       checkProcsBlockedCritLimit,
		checkProcsBlockedOccurences:      conf.Cpu.CheckProcsBlocked.Occurences,
		checkProcsBlockedReissueDuration: conf.Cpu.CheckProcsBlocked.ReissueDuration,
		checkProcsBlockedWarnCounter:     0,
		checkProcsBlockedCritCounter:     0,
	}
}

func (reader *CpuReader) Read(tctx *logger.TraceContext) {
	timestamp := time.Now()
	if reader.tmpCpuStat == nil {
		reader.tmpCpuStat, _ = reader.read(tctx)
	} else {
		tmpCpuStat, tmpCpuProcessorStats := reader.read(tctx)
		if len(reader.cpuStats) > reader.cacheLength {
			reader.cpuStats = reader.cpuStats[1:]
		}
		reader.cpuStats = append(reader.cpuStats, CpuStat{
			Timestamp:    timestamp,
			Intr:         tmpCpuStat.Intr - reader.tmpCpuStat.Intr,
			Ctx:          tmpCpuStat.Ctx - reader.tmpCpuStat.Ctx,
			Btime:        tmpCpuStat.Btime - reader.tmpCpuStat.Btime,
			Processes:    tmpCpuStat.Processes - reader.tmpCpuStat.Processes,
			ProcsRunning: tmpCpuStat.ProcsRunning,
			ProcsBlocked: tmpCpuStat.ProcsBlocked,
			Softirq:      tmpCpuStat.Softirq - reader.tmpCpuStat.Softirq,
		})

		if len(reader.cpuProcessorStats) > reader.cacheLength {
			reader.cpuProcessorStats = reader.cpuProcessorStats[len(tmpCpuProcessorStats):]
		}

		for _, stat := range tmpCpuProcessorStats {
			total := stat.User + stat.Nice + stat.Nice + stat.System + stat.Idle + stat.Iowait + stat.Irq + stat.Softirq + stat.Steal + stat.Guest + stat.GuestNice
			reader.cpuProcessorStats = append(reader.cpuProcessorStats, CpuProcessorStat{
				Timestamp: timestamp,
				Processor: stat.Processor,
				Mhz:       stat.Mhz,
				User:      stat.User * 100 / total,
				Nice:      stat.Nice * 100 / total,
				System:    stat.System * 100 / total,
				Idle:      stat.Idle * 100 / total,
				Iowait:    stat.Iowait * 100 / total,
				Irq:       stat.Irq * 100 / total,
				Softirq:   stat.Softirq * 100 / total,
				Steal:     stat.Steal * 100 / total,
				Guest:     stat.Guest * 100 / total,
				GuestNice: stat.GuestNice * 100 / total,
			})
		}

		reader.tmpCpuStat = tmpCpuStat
	}
}

func (reader *CpuReader) read(tctx *logger.TraceContext) (cpuStat *TmpCpuStat, cpuProcessorStats []TmpCpuProcessorStat) {
	var tmpReader *bufio.Reader

	// Read /proc/cpuinfo
	cpuinfo, _ := os.Open("/proc/cpuinfo")
	defer cpuinfo.Close()
	tmpReader = bufio.NewReader(cpuinfo)

	cpuProcessorStats = make([]TmpCpuProcessorStat, len(reader.cpus))

	var tmpProcessor int
	var tmpCpuMhzF float64
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
			tmpCpuMhzF, _ = strconv.ParseFloat(splited[1], 64)
			tmpCpuMhz = int64(tmpCpuMhzF)
			for i := 0; i < 20; i++ {
				if _, _, tmpErr = tmpReader.ReadLine(); tmpErr != nil {
					break
				}
			}
		}
		cpuProcessorStats[tmpProcessor] = TmpCpuProcessorStat{
			Mhz:       tmpCpuMhz,
			Processor: int64(tmpProcessor),
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
	procsRunning, _ := strconv.ParseInt(strings.Split(string(tmpBytes), " ")[1], 10, 64)
	tmpBytes, _, _ = tmpReader.ReadLine()
	procsBlocked, _ := strconv.ParseInt(strings.Split(string(tmpBytes), " ")[1], 10, 64)
	tmpBytes, _, _ = tmpReader.ReadLine()
	softirq, _ := strconv.ParseInt(strings.Split(string(tmpBytes), " ")[1], 10, 64)
	cpuStat = &TmpCpuStat{
		Intr:         intr,
		Ctx:          ctx,
		Btime:        btime,
		Processes:    processes,
		ProcsRunning: procsRunning,
		ProcsBlocked: procsBlocked,
		Softirq:      softirq,
	}

	return
}

func (reader *CpuReader) ReportMetrics() (metrics []spec.ResourceMetric) {
	metrics = make([]spec.ResourceMetric, 0, len(reader.cpuStats))

	for _, stat := range reader.cpuStats {
		if stat.ReportStatus == ReportStatusReported {
			continue
		}

		if reader.checkProcsRunningWarnLimit > 0 && stat.ProcsRunning > reader.checkProcsRunningWarnLimit {
			reader.checkProcsRunningWarnCounter += 1
		} else {
			reader.checkProcsRunningWarnCounter = 0
		}
		if reader.checkProcsRunningCritLimit > 0 && stat.ProcsRunning > reader.checkProcsRunningCritLimit {
			reader.checkProcsRunningCritCounter += 1
		} else {
			reader.checkProcsRunningCritCounter = 0
		}

		if reader.checkProcsBlockedWarnLimit > 0 && stat.ProcsBlocked > reader.checkProcsBlockedWarnLimit {
			reader.checkProcsBlockedWarnCounter += 1
		} else {
			reader.checkProcsBlockedWarnCounter = 0
		}
		if reader.checkProcsBlockedCritLimit > 0 && stat.ProcsBlocked > reader.checkProcsBlockedCritLimit {
			reader.checkProcsBlockedCritCounter += 1
		} else {
			reader.checkProcsBlockedCritCounter = 0
		}

		metrics = append(metrics, spec.ResourceMetric{
			Name: "system_cpu",
			Time: stat.Timestamp,
			Metric: map[string]interface{}{
				"intr":          stat.Intr,
				"ctx":           stat.Ctx,
				"btime":         stat.Btime,
				"processes":     stat.Processes,
				"procs_running": stat.ProcsRunning,
				"procs_blocked": stat.ProcsBlocked,
				"softirq":       stat.Softirq,
			},
		})
	}

	for _, stat := range reader.cpuProcessorStats {
		if stat.ReportStatus == ReportStatusReported {
			continue
		}

		metrics = append(metrics, spec.ResourceMetric{
			Name: "system_processor",
			Time: stat.Timestamp,
			Tag: map[string]string{
				"processor": strconv.FormatInt(stat.Processor, 10),
			},
			Metric: map[string]interface{}{
				"mhz":       stat.Mhz,
				"user":      stat.User,
				"nice":      stat.Nice,
				"system":    stat.System,
				"idle":      stat.Idle,
				"iowait":    stat.Iowait,
				"irq":       stat.Irq,
				"softirq":   stat.Softirq,
				"steal":     stat.Steal,
				"guest":     stat.Guest,
				"guestnice": stat.GuestNice,
			},
		})
	}

	return
}

func (reader *CpuReader) ReportEvents() (events []spec.ResourceEvent) {
	if len(reader.cpuStats) == 0 {
		return
	}

	cpuStat := reader.cpuStats[len(reader.cpuStats)-1]
	eventProcsRunningLevel := consts.EventLevelSuccess
	if reader.checkProcsRunningCritCounter > reader.checkProcsRunningOccurences {
		eventProcsRunningLevel = consts.EventLevelCritical
	} else if reader.checkProcsRunningWarnCounter > reader.checkProcsRunningOccurences {
		eventProcsRunningLevel = consts.EventLevelWarning
	}

	events = append(events, spec.ResourceEvent{
		Name:            "CheckCpuProcsRunning",
		Time:            cpuStat.Timestamp,
		Level:           eventProcsRunningLevel,
		Msg:             fmt.Sprintf("ProcsRunning: %d", cpuStat.ProcsRunning),
		ReissueDuration: reader.checkProcsRunningReissueDuration,
	})

	eventProcsBlockedLevel := consts.EventLevelSuccess
	if reader.checkProcsBlockedCritCounter > reader.checkProcsBlockedOccurences {
		eventProcsBlockedLevel = consts.EventLevelCritical
	} else if reader.checkProcsBlockedWarnCounter > reader.checkProcsBlockedOccurences {
		eventProcsBlockedLevel = consts.EventLevelWarning
	}
	events = append(events, spec.ResourceEvent{
		Name:            "CheckCpuProcsBlocked",
		Time:            cpuStat.Timestamp,
		Level:           eventProcsBlockedLevel,
		Msg:             fmt.Sprintf("ProcsBlocked: %d", cpuStat.ProcsBlocked),
		ReissueDuration: reader.checkProcsBlockedReissueDuration,
	})

	return
}

func (reader *CpuReader) Reported() {
	for i := range reader.cpuStats {
		reader.cpuStats[i].ReportStatus = ReportStatusReported
	}
	for i := range reader.cpuProcessorStats {
		reader.cpuProcessorStats[i].ReportStatus = ReportStatusReported
	}
	return
}
