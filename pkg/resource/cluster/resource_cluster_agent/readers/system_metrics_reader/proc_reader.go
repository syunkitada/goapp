package system_metrics_reader

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

const (
	CmdQemu = "qemu-system-x86"
)

type ProcsStat struct {
	ReportStatus int // 0, 1(GetReport), 2(Reported)
	Timestamp    time.Time
	Procs        int64
	Runs         int64
	Sleeps       int64
	DiskSleeps   int64
	Zonbies      int64
	Others       int64
}

type QemuStat struct {
	NetDevStats []NetDevStat
}

type TmpProcStat struct {
	Timestamp                time.Time
	Name                     string
	Cmd                      string
	Pid                      string
	VmSizeKb                 int64
	VmRssKb                  int64
	State                    int64
	SchedCpuTime             int64
	SchedWaitTime            int64
	SchedTimeSlices          int64
	HugetlbPages             int64
	Threads                  int64
	VoluntaryCtxtSwitches    int64
	NonvoluntaryCtxtSwitches int64

	Utime  int64
	Stime  int64
	Gtime  int64
	Cgtime int64

	Syscr      int64
	Syscw      int64
	ReadBytes  int64
	WriteBytes int64

	Qemu *QemuStat
}

type ProcStat struct {
	ReportStatus             int // 0, 1(GetReport), 2(Reported)
	Timestamp                time.Time
	Name                     string
	Cmd                      string
	Pid                      string
	VmSizeKb                 int64
	VmRssKb                  int64
	State                    int64
	SchedCpuTime             int64
	SchedWaitTime            int64
	SchedTimeSlices          int64
	HugetlbPages             int64
	Threads                  int64
	VoluntaryCtxtSwitches    int64
	NonvoluntaryCtxtSwitches int64

	UserUtil   int64
	SystemUtil int64
	GuestUtil  int64
	CguestUtil int64

	SyscrPerSec      int64
	SyscwPerSec      int64
	ReadBytesPerSec  int64
	WriteBytesPerSec int64

	Qemu *QemuStat

	WarnSchedWaitTimeCounter int
	CritSchedWaitTimeCounter int
	checkProc                config.ResourceMetricsSystemProcCheckProcConfig
}

type ProcReader struct {
	conf                *config.ResourceMetricsSystemConfig
	cacheLength         int
	systemMetricsReader *SystemMetricsReader

	cmdMap         map[string]config.ResourceMetricsSystemProcCheckProcConfig
	pidmax         int
	tmpProcStatMap map[int]TmpProcStat
	procsStats     []ProcsStat
	procStats      []ProcStat

	checkProcStats                  []ProcStat
	checkProcsStatusWarnCounter     int
	checkProcsStatusCritCounter     int
	checkProcsStatusOccurences      int
	checkProcsStatusReissueDuration int

	checkProcStatMap map[int]ProcStat
}

func NewProcReader(conf *config.ResourceMetricsSystemConfig, systemMetricsReader *SystemMetricsReader) SubMetricsReader {
	pidmaxFile, _ := os.Open("/proc/sys/kernel/pid_max")
	defer pidmaxFile.Close()
	tmpReader := bufio.NewReader(pidmaxFile)
	tmpBytes, _, _ := tmpReader.ReadLine()
	pidmax, _ := strconv.Atoi(string(tmpBytes))

	cmdMap := map[string]config.ResourceMetricsSystemProcCheckProcConfig{}
	for name, check := range conf.Proc.CheckProcMap {
		check.Name = name
		cmdMap[check.Cmd] = check
	}

	return &ProcReader{
		conf:                conf,
		cacheLength:         conf.CacheLength,
		systemMetricsReader: systemMetricsReader,
		cmdMap:              cmdMap,
		pidmax:              pidmax,
		procsStats:          make([]ProcsStat, 0, conf.CacheLength),
		procStats:           make([]ProcStat, 0, conf.CacheLength),

		checkProcsStatusWarnCounter:     0,
		checkProcsStatusCritCounter:     0,
		checkProcsStatusOccurences:      conf.Proc.CheckProcsStatus.Occurences,
		checkProcsStatusReissueDuration: conf.Proc.CheckProcsStatus.ReissueDuration,
	}
}

const ProcDir = "/proc/"

func (reader *ProcReader) Read(tctx *logger.TraceContext) {
	timestamp := time.Now()
	var procDirFile *os.File
	var err error
	if procDirFile, err = os.Open(ProcDir); err != nil {
		return
	}
	var procFileInfos []os.FileInfo
	procFileInfos, err = procDirFile.Readdir(-1)
	procDirFile.Close()
	if err != nil {
		return
	}

	var tmpFile *os.File
	var tmpErr error
	var procs int64 = 0
	var procRuns int64 = 0
	var procSleeps int64 = 0
	var procDiskSleeps int64 = 0
	var procZonbies int64 = 0
	var procOthers int64 = 0
	var tmpReader *bufio.Reader
	var tmpBytes []byte
	var tmpText string
	var tmpTexts []string

	// TODO Read self proc stat

	tmpProcStatMap := map[int]TmpProcStat{}
	for _, procFileInfo := range procFileInfos {
		if !procFileInfo.IsDir() {
			continue
		}

		procDir := ProcDir + procFileInfo.Name() + "/"
		if tmpFile, tmpErr = os.Open(procDir + "status"); tmpErr != nil {
			continue
		}
		procs += 1

		tmpReader = bufio.NewReader(tmpFile)
		tmpBytes, _, _ = tmpReader.ReadLine()
		cmd := str_utils.ParseLastValue(string(tmpBytes))

		_, _, _ = tmpReader.ReadLine()
		tmpBytes, _, _ = tmpReader.ReadLine()
		state := str_utils.ParseLastSecondValue(string(tmpBytes))
		var stateInt int64 = 0
		switch state {
		case "R":
			procRuns += 1
			stateInt = 3
		case "D":
			procDiskSleeps += 1
			stateInt = 2
		case "S":
			procSleeps += 1
			stateInt = 1
		case "Z":
			procZonbies += 1
			stateInt = -1
		default:
			procOthers += 1
			stateInt = 0
		}

		if checkProc, ok := reader.cmdMap[cmd]; ok {
			pid, _ := strconv.Atoi(str_utils.ParseLastSecondValue(procFileInfo.Name()))

			statusLines := make([]string, 0, 55)
			for {
				tmpBytes, _, tmpErr := tmpReader.ReadLine()
				if tmpErr != nil {
					break
				}
				statusLines = append(statusLines, string(tmpBytes))
			}
			// Tgid:   23550
			// Ngid:   0
			// Pid:    23550
			// PPid:   23547
			// TracerPid:      0
			// Uid:    0       0       0       0
			// Gid:    0       0       0       0
			// FDSize: 256
			// Groups:
			// NStgid: 23550
			// NSpid:  23550
			// NSpgid: 23550
			// NSsid:  23547
			// VmPeak:  3235840 kB
			// VmSize:  2461756 kB
			vmSizeKb, _ := strconv.ParseInt(str_utils.ParseLastSecondValue(statusLines[14]), 10, 64)
			// VmLck:         0 kB
			// VmPin:         0 kB
			// VmHWM:     31584 kB
			// VmRSS:     28784 kB
			vmRssKb, _ := strconv.ParseInt(str_utils.ParseLastSecondValue(statusLines[18]), 10, 64)
			// RssAnon:           16256 kB
			// RssFile:           12528 kB
			// RssShmem:              0 kB
			// VmData:  2399228 kB
			// VmStk:       132 kB
			// VmExe:     11452 kB
			// VmLib:      7424 kB
			// VmPTE:       572 kB
			// VmSwap:        0 kB
			// HugetlbPages:    2097152 kB
			hugetlbPages, _ := strconv.ParseInt(str_utils.ParseLastSecondValue(statusLines[28]), 10, 64)
			// CoreDumping:    0
			// THP_enabled:    1
			// Threads:        4
			threads, _ := strconv.ParseInt(str_utils.ParseLastValue(statusLines[31]), 10, 64)
			// SigQ:   0/62468
			// SigPnd: 0000000000000000
			// ShdPnd: 0000000000000000
			// SigBlk: 0000000010002240
			// SigIgn: 0000000000001000
			// SigCgt: 0000000180004243
			// CapInh: 0000000000000000
			// CapPrm: 0000003fffffffff
			// CapEff: 0000003fffffffff
			// CapBnd: 0000003fffffffff
			// CapAmb: 0000000000000000
			// NoNewPrivs:     0
			// Seccomp:        0
			// Speculation_Store_Bypass:       thread vulnerable
			// Cpus_allowed:   ffff
			// Cpus_allowed_list:      0-15
			// Mems_allowed:   00000000,00000000,00000000,00000000,00000000,00000000,00000000,00000000,00000000,00000000,00000000,00000000,00000000,00000000,00000000,00000000,00000000,00000000,00000000,00000000,00000000,00000000,00000000,00000000,00000000,00000000,00000000,00000000,00000000,00000000,00000000,00000001
			// Mems_allowed_list:      0
			// voluntary_ctxt_switches:        14415
			voluntaryCtxtSwitches, _ := strconv.ParseInt(str_utils.ParseLastValue(statusLines[50]), 10, 64)
			// nonvoluntary_ctxt_switches:     219
			nonvoluntaryCtxtSwitches, _ := strconv.ParseInt(str_utils.ParseLastValue(statusLines[51]), 10, 64)

			// Parse /proc/[pid]/schedstat
			// 2554841551 177487694 35200
			// [time spent on the cpu] [time spent waiting on a runqueue] [timeslices run on this cpu]
			if tmpFile, tmpErr = os.Open(procDir + "schedstat"); tmpErr != nil {
				continue
			}
			tmpReader = bufio.NewReader(tmpFile)
			tmpBytes, _, tmpErr = tmpReader.ReadLine()
			if tmpErr != nil {
				continue
			}
			tmpTexts = strings.Split(string(tmpBytes), " ")
			if len(tmpTexts) != 3 {
				logger.Warningf(tctx, "Unexpected Format: path=/proc/[pid]/schedstat, text=%s", string(tmpText))
				continue
			}
			// time spent on the cpu
			schedCpuTime, _ := strconv.ParseInt(tmpTexts[0], 10, 64)
			// time spent waiting on a runqueue
			schedWaitTime, _ := strconv.ParseInt(tmpTexts[1], 10, 64)
			// # of timeslices run on this cpu
			schedTimeSlices, _ := strconv.ParseInt(tmpTexts[2], 10, 64)

			// $ cat /proc/24120/stat
			// 24120 (qemu-system-x86) S 24119 24120 24119 0 -1 138412416 23189 0 0 0 2227 753 0 0 20 0 6 0 251962 4969209856 7743 18446744073709551615 1 1 0 0 0 0 268444224 4096 16963 0 0 0 17 9 0 0 0 2041 0 0 0 0 0 0 0 0 0
			if tmpFile, tmpErr = os.Open(procDir + "stat"); tmpErr != nil {
				continue
			}
			tmpReader = bufio.NewReader(tmpFile)
			tmpBytes, _, tmpErr = tmpReader.ReadLine()
			if tmpErr != nil {
				continue
			}
			tmpTexts = strings.Split(string(tmpBytes), " ")
			utime, _ := strconv.ParseInt(tmpTexts[13], 10, 64)
			stime, _ := strconv.ParseInt(tmpTexts[14], 10, 64)
			gtime, _ := strconv.ParseInt(tmpTexts[42], 10, 64)
			cgtime, _ := strconv.ParseInt(tmpTexts[43], 10, 64)
			startTime, _ := strconv.Atoi(tmpTexts[21])

			// $ cat /proc/24120/io
			// rchar: 160323858
			// wchar: 14532026
			// syscr: 48257
			// syscw: 37187
			// read_bytes: 163528704
			// write_bytes: 15466496
			// cancelled_write_bytes: 0
			if tmpFile, tmpErr = os.Open(procDir + "io"); tmpErr != nil {
				continue
			}
			tmpReader = bufio.NewReader(tmpFile)
			_, _, _ = tmpReader.ReadLine()
			_, _, _ = tmpReader.ReadLine()
			tmpBytes, _, _ = tmpReader.ReadLine()
			syscr, _ := strconv.ParseInt(str_utils.ParseLastValue(string(tmpBytes)), 10, 64)
			tmpBytes, _, _ = tmpReader.ReadLine()
			syscw, _ := strconv.ParseInt(str_utils.ParseLastValue(string(tmpBytes)), 10, 64)
			tmpBytes, _, _ = tmpReader.ReadLine()
			readBytes, _ := strconv.ParseInt(str_utils.ParseLastValue(string(tmpBytes)), 10, 64)
			tmpBytes, _, _ = tmpReader.ReadLine()
			writeBytes, _ := strconv.ParseInt(str_utils.ParseLastValue(string(tmpBytes)), 10, 64)

			tmpProcStat := TmpProcStat{
				Timestamp:                timestamp,
				Name:                     checkProc.Name,
				Cmd:                      cmd,
				Pid:                      procFileInfo.Name(),
				SchedCpuTime:             schedCpuTime,
				SchedWaitTime:            schedWaitTime,
				SchedTimeSlices:          schedTimeSlices,
				State:                    stateInt,
				VmSizeKb:                 vmSizeKb,
				VmRssKb:                  vmRssKb,
				HugetlbPages:             hugetlbPages,
				Threads:                  threads,
				VoluntaryCtxtSwitches:    voluntaryCtxtSwitches,
				NonvoluntaryCtxtSwitches: nonvoluntaryCtxtSwitches,

				Utime:  utime,
				Stime:  stime,
				Gtime:  gtime,
				Cgtime: cgtime,

				Syscr:      syscr,
				Syscw:      syscw,
				ReadBytes:  readBytes,
				WriteBytes: writeBytes,
			}

			switch cmd {
			case CmdQemu:
				tmpProcStat.Qemu = reader.GetQemuStat(tctx, &tmpProcStat)
			}

			key := startTime*100000 + pid
			tmpProcStatMap[key] = tmpProcStat
		}

		tmpFile.Close()
	}

	if len(reader.procStats) > reader.cacheLength-len(tmpProcStatMap) {
		reader.procStats = reader.procStats[len(tmpProcStatMap):]
	}

	checkProcStatMap := map[int]ProcStat{}

	if reader.tmpProcStatMap != nil {
		for key, tmpStat := range reader.tmpProcStatMap {
			if stat, ok := tmpProcStatMap[key]; ok {
				interval := stat.Timestamp.Unix() - tmpStat.Timestamp.Unix()

				userUtil := (stat.Utime - tmpStat.Utime) / interval
				systemUtil := (stat.Stime - tmpStat.Stime) / interval
				guestUtil := (stat.Gtime - tmpStat.Gtime) / interval
				cguestUtil := (stat.Cgtime - tmpStat.Cgtime) / interval

				schedTimeSlices := stat.SchedTimeSlices - tmpStat.SchedTimeSlices
				schedCpuTime := ((stat.SchedCpuTime - tmpStat.SchedCpuTime) / interval)
				schedWaitTime := ((stat.SchedWaitTime - tmpStat.SchedWaitTime) / interval)

				checkProc := reader.cmdMap[stat.Cmd]

				procStat := ProcStat{
					Timestamp:                stat.Timestamp,
					Name:                     stat.Name,
					Cmd:                      stat.Cmd,
					Pid:                      stat.Pid,
					VmSizeKb:                 stat.VmSizeKb,
					VmRssKb:                  stat.VmRssKb,
					State:                    stat.State,
					SchedCpuTime:             schedCpuTime,
					SchedWaitTime:            schedWaitTime,
					SchedTimeSlices:          schedTimeSlices,
					HugetlbPages:             stat.HugetlbPages,
					Threads:                  stat.Threads,
					VoluntaryCtxtSwitches:    stat.VoluntaryCtxtSwitches - tmpStat.VoluntaryCtxtSwitches,
					NonvoluntaryCtxtSwitches: stat.NonvoluntaryCtxtSwitches - tmpStat.NonvoluntaryCtxtSwitches,

					UserUtil:   userUtil,
					SystemUtil: systemUtil,
					GuestUtil:  guestUtil,
					CguestUtil: cguestUtil,

					SyscrPerSec:      (stat.Syscr - tmpStat.Syscr) / interval,
					SyscwPerSec:      (stat.Syscw - tmpStat.Syscw) / interval,
					ReadBytesPerSec:  (stat.ReadBytes - tmpStat.ReadBytes) / interval,
					WriteBytesPerSec: (stat.WriteBytes - tmpStat.WriteBytes) / interval,

					Qemu: stat.Qemu,

					checkProc: checkProc,
				}

				reader.procStats = append(reader.procStats, procStat)

				checkProcStat, ok := reader.checkProcStatMap[key]
				if ok {
					procStat.CritSchedWaitTimeCounter = checkProcStat.CritSchedWaitTimeCounter
					procStat.WarnSchedWaitTimeCounter = checkProcStat.WarnSchedWaitTimeCounter
				}
				if schedWaitTime > checkProc.CritSchedWaitTime {
					procStat.CritSchedWaitTimeCounter += 1
				} else if schedWaitTime > checkProc.WarnSchedWaitTime {
					procStat.WarnSchedWaitTimeCounter += 1
				} else {
					procStat.CritSchedWaitTimeCounter = 0
					procStat.WarnSchedWaitTimeCounter = 0
				}

				checkProcStatMap[key] = procStat
			}
		}
	}

	reader.tmpProcStatMap = tmpProcStatMap
	reader.checkProcStatMap = checkProcStatMap

	stat := ProcsStat{
		Timestamp:  timestamp,
		Procs:      procs,
		Runs:       procRuns,
		Sleeps:     procSleeps,
		DiskSleeps: procDiskSleeps,
		Zonbies:    procZonbies,
		Others:     procOthers,
	}
	if len(reader.procsStats) > reader.cacheLength {
		reader.procsStats = reader.procsStats[1:]
	}
	reader.procsStats = append(reader.procsStats, stat)
	return
}

func (reader *ProcReader) ReportMetrics() (metrics []spec.ResourceMetric) {
	metrics = make([]spec.ResourceMetric, 0, len(reader.procsStats)+len(reader.procStats))

	for _, stat := range reader.procsStats {
		if stat.ReportStatus == ReportStatusReported {
			continue
		}
		metrics = append(metrics, spec.ResourceMetric{
			Name: "system_procs",
			Time: stat.Timestamp,
			Tag:  map[string]string{},
			Metric: map[string]interface{}{
				"procs":       stat.Procs,
				"runs":        stat.Runs,
				"sleeps":      stat.Sleeps,
				"disk_sleeps": stat.DiskSleeps,
				"zonbies":     stat.Zonbies,
				"others":      stat.Others,
			},
		})
	}

	for _, stat := range reader.procStats {
		if stat.ReportStatus == ReportStatusReported {
			continue
		}

		metric := map[string]interface{}{
			"vm_size_kb":                 stat.VmSizeKb,
			"vm_rss_kb":                  stat.VmRssKb,
			"state":                      stat.State,
			"sched_cpu_time":             stat.SchedCpuTime,
			"sched_wait_time":            stat.SchedWaitTime,
			"sched_time_slices":          stat.SchedTimeSlices,
			"hugetlb_pages":              stat.HugetlbPages,
			"threads":                    stat.Threads,
			"voluntary_ctxt_switches":    stat.VoluntaryCtxtSwitches,
			"nonvoluntary_ctxt_switches": stat.NonvoluntaryCtxtSwitches,
			"user_util":                  stat.UserUtil,
			"system_util":                stat.SystemUtil,
			"guest_util":                 stat.GuestUtil,
			"cguest_util":                stat.CguestUtil,

			"syscr_per_sec":       stat.SyscrPerSec,
			"syscw_per_sec":       stat.SyscwPerSec,
			"read_bytes_per_sec":  stat.ReadBytesPerSec,
			"write_bytes_per_sec": stat.WriteBytesPerSec,
		}

		if stat.Qemu != nil {
			var receiveBytesPerSec int64
			var receivePacketsPerSec int64
			var receiveDiffErrors int64
			var receiveDiffDrops int64
			var transmitBytesPerSec int64
			var transmitPacketsPerSec int64
			var transmitDiffErrors int64
			var transmitDiffDrops int64
			for _, stat := range stat.Qemu.NetDevStats {
				receiveBytesPerSec += stat.ReceiveBytesPerSec
				receivePacketsPerSec += stat.ReceivePacketsPerSec
				receiveDiffErrors += stat.ReceiveDiffErrors
				receiveDiffDrops += stat.ReceiveDiffDrops
				transmitBytesPerSec += stat.TransmitBytesPerSec
				transmitPacketsPerSec += stat.TransmitPacketsPerSec
				transmitDiffErrors += stat.TransmitDiffErrors
				transmitDiffDrops += stat.TransmitDiffDrops
			}

			metric["receive_bytes_per_sec"] = receiveBytesPerSec
			metric["receive_packets_per_sec"] = receivePacketsPerSec
			metric["receive_errors"] = receiveDiffErrors
			metric["receive_drops"] = receiveDiffDrops
			metric["transmit_bytes_per_sec"] = transmitBytesPerSec
			metric["transmit_packets_per_sec"] = transmitPacketsPerSec
			metric["transmit_errors"] = transmitDiffErrors
			metric["transmit_drops"] = transmitDiffDrops
		}

		metrics = append(metrics, spec.ResourceMetric{
			Name: "system_proc",
			Time: stat.Timestamp,
			Tag: map[string]string{
				"name": stat.Name,
				"cmd":  stat.Cmd,
				"pid":  stat.Pid,
			},
			Metric: metric,
		})

		stat.ReportStatus = 1
	}

	return
}

func (reader *ProcReader) ReportEvents() (events []spec.ResourceEvent) {
	if len(reader.procsStats) == 0 {
		return
	}

	eventCheckProcsStatusLevel := consts.EventLevelSuccess
	stat := reader.procsStats[len(reader.procsStats)-1]

	if reader.checkProcsStatusCritCounter > reader.checkProcsStatusOccurences {
		eventCheckProcsStatusLevel = consts.EventLevelCritical

	} else if reader.checkProcsStatusWarnCounter > reader.checkProcsStatusOccurences {
		eventCheckProcsStatusLevel = consts.EventLevelWarning
	}

	events = append(events, spec.ResourceEvent{
		Name:  "CheckProcsStatus",
		Time:  stat.Timestamp,
		Level: eventCheckProcsStatusLevel,
		Msg: fmt.Sprintf("Procs=%d,Runs=%d,Sleeps=%d,DiskSleeps=%d,Zonbies=%d,Others=%d",
			stat.Procs,
			stat.Runs,
			stat.Sleeps,
			stat.DiskSleeps,
			stat.Zonbies,
			stat.Others,
		),
		ReissueDuration: reader.checkProcsStatusReissueDuration,
	})

	if len(reader.checkProcStatMap) != 0 {
		eventCheckProcLevel := consts.EventLevelSuccess
		var msgs []string
		for _, procStat := range reader.checkProcStatMap {
			if procStat.CritSchedWaitTimeCounter > procStat.checkProc.Occurences {
				eventCheckProcLevel = consts.EventLevelCritical
			} else if eventCheckProcLevel == consts.EventLevelSuccess && procStat.WarnSchedWaitTimeCounter > procStat.checkProc.Occurences {
				eventCheckProcLevel = consts.EventLevelWarning
			}

			msgs = append(msgs, fmt.Sprintf(
				"Pid=%d,Cmd=%d,SchedWaitTime=%d,SystemUtil=%d,UserUtil=%d,GuestUtil=%d",
				procStat.Pid,
				procStat.Cmd,
				procStat.SchedWaitTime,
				procStat.SystemUtil,
				procStat.UserUtil,
				procStat.GuestUtil,
			))
		}

		events = append(events, spec.ResourceEvent{
			Name:            "CheckProc",
			Time:            stat.Timestamp,
			Level:           eventCheckProcLevel,
			Msg:             strings.Join(msgs, "\n"),
			ReissueDuration: reader.checkProcsStatusReissueDuration,
		})
	}

	return
}

func (reader *ProcReader) Reported() {
	for i := range reader.procsStats {
		reader.procsStats[i].ReportStatus = ReportStatusReported
	}
	for i := range reader.procStats {
		reader.procStats[i].ReportStatus = ReportStatusReported
	}
	return
}

func (reader *ProcReader) GetQemuStat(tctx *logger.TraceContext, procStat *TmpProcStat) (qemuStat *QemuStat) {
	var netDevStats []NetDevStat

	cmdlineFile, _ := os.Open("/proc/" + procStat.Pid + "/cmdline")
	defer cmdlineFile.Close()
	tmpReader := bufio.NewReader(cmdlineFile)
	tmpBytes, _, _ := tmpReader.ReadLine()
	cmds := strings.Split(string(tmpBytes), string(byte(0)))
	lenCmds := len(cmds)
	for i := 0; i < lenCmds; i++ {
		switch cmds[i] {
		case "-nic":
			splitedOption := strings.Split(cmds[i+1], ",")
			for _, option := range splitedOption {
				splitedKeyValue := strings.Split(option, "=")
				if splitedKeyValue[0] == "ifname" {
					netDevStat, ok := reader.systemMetricsReader.NetDevStatMap[splitedKeyValue[1]]
					if !ok {
						break
					}
					netDevStats = append(netDevStats, netDevStat)
				}
			}
			i += 1
		}
	}

	qemuStat = &QemuStat{
		NetDevStats: netDevStats,
	}
	return
}
