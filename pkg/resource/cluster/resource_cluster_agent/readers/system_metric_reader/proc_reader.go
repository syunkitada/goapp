package system_metric_reader

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/lib/str_utils"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
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
}

const ProcDir = "/proc/"

func (reader *SystemMetricReader) ReadProc(tctx *logger.TraceContext) (err error) {
	timestamp := time.Now()
	var procDirFile *os.File
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
	procMap := map[string]map[string][]string{}
	for _, procCheck := range reader.conf.ProcCheckMap {
		procMap[procCheck.Cmd] = map[string][]string{}
	}
	for _, procFileInfo := range procFileInfos {
		if !procFileInfo.IsDir() {
			continue
		}
		if tmpFile, tmpErr = os.Open(ProcDir + procFileInfo.Name() + "/" + "status"); tmpErr != nil {
			continue
		}
		procs += 1

		tmpReader = bufio.NewReader(tmpFile)
		tmpBytes, _, _ = tmpReader.ReadLine()
		cmd := str_utils.ParseLastSecondValue(string(tmpBytes))
		_, _, _ = tmpReader.ReadLine()
		tmpBytes, _, _ = tmpReader.ReadLine()
		state := str_utils.ParseLastSecondValue(string(tmpBytes))
		switch state {
		case "R":
			procRuns += 1
		case "D":
			procDiskSleeps += 1
		case "S":
			procSleeps += 1
		case "Z":
			procZonbies += 1
		default:
			procOthers += 1
		}

		if cmdMap, ok := procMap[cmd]; ok {
			tmpTexts := make([]string, 0, 53)
			tmpTexts = append(tmpTexts, state)
			for {
				tmpBytes, _, tmpErr = tmpReader.ReadLine()
				if tmpErr != nil {
					break
				}
				tmpTexts = append(tmpTexts, string(tmpBytes))
			}
			cmdMap[procFileInfo.Name()] = tmpTexts
		}
		tmpFile.Close()
	}

	for name, procCheck := range reader.conf.ProcCheckMap {
		if cmdMap, ok := procMap[procCheck.Cmd]; ok {
			for pid, cmdStatus := range cmdMap {
				// Parse /proc/[pid]/schedstat
				// 2554841551 177487694 35200
				// [time spent on the cpu] [time spent waiting on a runqueue] [timeslices run on this cpu]
				if tmpFile, tmpErr = os.Open(ProcDir + pid + "/" + "schedstat"); tmpErr != nil {
					continue
				}
				tmpReader = bufio.NewReader(tmpFile)
				tmpBytes, _, tmpErr = tmpReader.ReadLine()
				if tmpErr != nil {
					continue
				}
				tmpTexts = strings.Split(string(tmpText), " ")
				if len(tmpTexts) != 3 {
					logger.Warningf(tctx, "Unexpected Format: path=/proc/[pid]/schedstat, text=%s", string(tmpText))
					continue
				}
				schedCpuTime, _ := strconv.ParseInt(tmpTexts[0], 10, 64)
				schedWaitTime, _ := strconv.ParseInt(tmpTexts[1], 10, 64)
				schedTimeSlices, _ := strconv.ParseInt(tmpTexts[2], 10, 64)

				// Parse /proc/self/status
				// Name:   qemu-system-x86
				// Umask:  0022
				// State:  S (sleeping)
				var state int64 = 0
				switch cmdStatus[0] {
				case "R":
					state = 3
				case "D":
					state = 2
				case "S":
					state = 1
				case "Z":
					state = -1
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
				vmSizeKb, _ := strconv.ParseInt(str_utils.ParseLastSecondValue(cmdStatus[15]), 10, 64)
				// VmLck:         0 kB
				// VmPin:         0 kB
				// VmHWM:     31584 kB
				// VmRSS:     28784 kB
				vmRssKb, _ := strconv.ParseInt(str_utils.ParseLastSecondValue(cmdStatus[19]), 10, 64)
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
				hugetlbPages, _ := strconv.ParseInt(str_utils.ParseLastSecondValue(cmdStatus[29]), 10, 64)
				// CoreDumping:    0
				// THP_enabled:    1
				// Threads:        4
				threads, _ := strconv.ParseInt(str_utils.ParseLastValue(cmdStatus[32]), 10, 64)
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
				voluntaryCtxtSwitches, _ := strconv.ParseInt(str_utils.ParseLastValue(cmdStatus[51]), 10, 64)
				// nonvoluntary_ctxt_switches:     219
				nonvoluntaryCtxtSwitches, _ := strconv.ParseInt(str_utils.ParseLastValue(cmdStatus[52]), 10, 64)
				//

				stat := ProcStat{
					ReportStatus:             0,
					Timestamp:                timestamp,
					Name:                     name,
					Cmd:                      procCheck.Cmd,
					Pid:                      pid,
					SchedCpuTime:             schedCpuTime,
					SchedWaitTime:            schedWaitTime,
					SchedTimeSlices:          schedTimeSlices,
					State:                    state,
					VmSizeKb:                 vmSizeKb,
					VmRssKb:                  vmRssKb,
					HugetlbPages:             hugetlbPages,
					Threads:                  threads,
					VoluntaryCtxtSwitches:    voluntaryCtxtSwitches,
					NonvoluntaryCtxtSwitches: nonvoluntaryCtxtSwitches,
				}

				if len(reader.procStats) > reader.cacheLength {
					reader.procStats = reader.procStats[1:]
				}
				reader.procStats = append(reader.procStats, stat)
			}
		}
	}

	stat := ProcsStat{
		ReportStatus: 0,
		Timestamp:    timestamp,
		Procs:        procs,
		Runs:         procRuns,
		Sleeps:       procSleeps,
		DiskSleeps:   procDiskSleeps,
		Zonbies:      procZonbies,
		Others:       procOthers,
	}
	if len(reader.procsStats) > reader.cacheLength {
		reader.procsStats = reader.procsStats[1:]
	}
	reader.procsStats = append(reader.procsStats, stat)
	return
}

func (reader *SystemMetricReader) GetProcStatMetrics() (metrics []spec.ResourceMetric) {
	metrics = make([]spec.ResourceMetric, len(reader.procsStats)+len(reader.procStats))

	for _, stat := range reader.procsStats {
		timestamp := strconv.FormatInt(stat.Timestamp.UnixNano(), 10)
		metrics = append(metrics, spec.ResourceMetric{
			Name: "system_procs",
			Time: timestamp,
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

		stat.ReportStatus = 1
	}

	for _, stat := range reader.procStats {
		timestamp := strconv.FormatInt(stat.Timestamp.UnixNano(), 10)
		metrics = append(metrics, spec.ResourceMetric{
			Name: "system_proc",
			Time: timestamp,
			Tag: map[string]string{
				"name": stat.Name,
				"cmd":  stat.Cmd,
				"pid":  stat.Pid,
			},
			Metric: map[string]interface{}{
				"vm_size_kb":                 stat.VmSizeKb,
				"vm_rss_kb":                  stat.VmRssKb,
				"sched_cpu_time":             stat.SchedCpuTime,
				"sched_wait_time":            stat.SchedWaitTime,
				"sched_time_slices":          stat.SchedTimeSlices,
				"hugetlb_pages":              stat.HugetlbPages,
				"threads":                    stat.Threads,
				"voluntary_ctxt_switches":    stat.VoluntaryCtxtSwitches,
				"nonvoluntary_ctxt_switches": stat.NonvoluntaryCtxtSwitches,
			},
		})

		stat.ReportStatus = 1
	}

	return
}
