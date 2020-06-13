package system_metric_reader

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/lib/str_utils"
	"github.com/syunkitada/goapp/pkg/resource/config"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

type SystemMetricReader struct {
	conf         *config.ResourceMetricSystemConfig
	name         string
	enableLogin  bool
	enableCpu    bool
	enableMemory bool
	enableProc   bool
	cacheLength  int
	uptimeStats  []UptimeStat
	loginStats   []LoginStat
	cpuStats     []CpuStat
	procsStats   []ProcsStat
	procStats    []ProcStat
	numaNodeMap  map[string]*spec.NumaNodeSpec
}

func New(conf *config.ResourceMetricSystemConfig) *SystemMetricReader {
	// f, _ := os.Open("/sys/devices/system/node/online")
	// /sys/devices/system/node/node0/meminfo
	nodeOnline, err := os.Open("/sys/devices/system/node/online")
	if err != nil {
		logger.StdoutFatalf("Failed Initialize SystemMetricReader: %v", err)
	}
	defer nodeOnline.Close()

	nodeOnlineBytes, err := ioutil.ReadAll(nodeOnline)
	if err != nil {
		logger.StdoutFatalf("Failed Initialize SystemMetricReader: %v", err)
	}
	splitedNodeServices := strings.Split(strings.TrimRight(string(nodeOnlineBytes), "\n"), ",")

	numaNodeMap := map[string]*spec.NumaNodeSpec{}
	for _, node := range splitedNodeServices {
		id, err := strconv.Atoi(node)
		if err != nil {
			logger.StdoutFatalf("Failed Initialize SystemMetricReader: %v", err)
		}
		numaNodeMap[node] = &spec.NumaNodeSpec{
			Id:     id,
			CpuMap: map[int]spec.NumaNodeCpuSpec{},
		}
	}

	return &SystemMetricReader{
		conf:         conf,
		name:         "system",
		enableLogin:  conf.EnableLogin,
		enableCpu:    conf.EnableCpu,
		enableMemory: conf.EnableMemory,
		enableProc:   conf.EnableProc,
		cacheLength:  conf.CacheLength,
		uptimeStats:  make([]UptimeStat, 0, conf.CacheLength),
		loginStats:   make([]LoginStat, 0, conf.CacheLength),
		cpuStats:     make([]CpuStat, 0, conf.CacheLength),
		procsStats:   make([]ProcsStat, 0, conf.CacheLength),
		procStats:    make([]ProcStat, 0, conf.CacheLength),
		numaNodeMap:  numaNodeMap,
	}
}

type UptimeStat struct {
	reportStatus int // 0, 1(GetReport), 2(Reported)
	timestamp    time.Time
	uptime       int64
}

type LoginStat struct {
	reportStatus int // 0, 1(GetReport), 2(Reported)
	users        []UserStat
	timestamp    time.Time
}

type UserStat struct {
	reportStatus int // 0, 1(GetReport), 2(Reported)
	user         string
	tty          string
	from         string
	login        string
	idle         string
	jcpu         string
	pcpu         string
	what         string
}

type CpuStat struct {
	reportStatus  int // 0, 1(GetReport), 2(Reported)
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

func (reader *SystemMetricReader) GetNumaNodeMap(tctx *logger.TraceContext) map[string]*spec.NumaNodeSpec {
	return reader.numaNodeMap
}

func (reader *SystemMetricReader) Read(tctx *logger.TraceContext) (err error) {
	timestamp := time.Now()

	// Read /proc/uptime
	// uptime(s)  idle(s)
	// 2906.26 5507.43
	// fmt.Println("READ /proc/uptime")
	procUptime, _ := os.Open("/proc/uptime")
	defer procUptime.Close()
	scanner := bufio.NewScanner(procUptime)
	scanner.Scan()
	uptimeText := scanner.Text()
	uptimeWords := strings.Split(uptimeText, " ")
	uptime, _ := strconv.ParseInt(uptimeWords[0], 10, 64)
	uptimeStat := UptimeStat{
		reportStatus: 0,
		timestamp:    timestamp,
		uptime:       uptime,
	}
	if len(reader.uptimeStats) > reader.cacheLength {
		reader.uptimeStats = reader.uptimeStats[1:]
	}
	reader.uptimeStats = append(reader.uptimeStats, uptimeStat)

	if reader.enableLogin {
		// Don't read /var/run/utmp, because of this is binary
		// Read w -h
		// USER    TTY      FROM          LOGIN@   IDLE   JCPU   PCPU  WHAT
		// hoge    pts/8    192.168.1.1   09:34    2.00s  0.10s  0.00s tmux a
		out, err := exec.Command("w", "-h").Output()
		users := []UserStat{}
		if err != nil {
			fmt.Println(err)
		}
		for _, line := range strings.Split(string(out), "\n") {
			l := strings.Split(line, " ")
			if len(l) != 8 {
				continue
			}
			users = append(users, UserStat{
				user:  l[0],
				tty:   l[1],
				from:  l[2],
				login: l[3],
				idle:  l[4],
				jcpu:  l[5],
				pcpu:  l[6],
				what:  l[7],
			})
		}
		if len(reader.loginStats) > reader.cacheLength {
			reader.loginStats = reader.loginStats[1:]
		}
		reader.loginStats = append(reader.loginStats, LoginStat{
			timestamp: timestamp,
			users:     users,
		})
	}

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

		// fmt.Println("READ /proc/stat")
		f, _ := os.Open("/proc/stat")
		defer f.Close()
		scanner = bufio.NewScanner(f)
		lines := make([]string, 0, 20)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}

		var cpu []string
		cpuMap := map[string][]int64{}
		lastIndex := 13 // TODO FIXME len(cpus) + 1
		for i := 0; i < lastIndex; i++ {
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
		// fmt.Println("procs_running", procs_running)
		stat := CpuStat{
			reportStatus:  0,
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

		if len(reader.cpuStats) > reader.cacheLength {
			reader.cpuStats = reader.cpuStats[1:]
		}
		reader.cpuStats = append(reader.cpuStats, stat)
	}

	if reader.enableMemory {
		// Read /proc/vmstat
		// TODO READ Memory
		fmt.Println("READ Memory")
	}

	if reader.enableProc {
		// Read /proc/
		reader.ReadProc(tctx)
	}

	var tmpBytes []byte
	var tmpFile *os.File
	var tmpScanner *bufio.Scanner
	var tmpErr error
	for id, node := range reader.numaNodeMap {
		if tmpBytes, err = ioutil.ReadFile("/sys/devices/system/node/node" + id + "/cpulist"); err != nil {
			return
		}
		cpus := str_utils.ParseRangeFormatStr(string(tmpBytes))
		for _, cpu := range cpus {
			node.CpuMap[cpu] = spec.NumaNodeCpuSpec{}
		}

		nr1GHugepages := 0
		if tmpBytes, err = ioutil.ReadFile("/sys/devices/system/node/node" + id + "/hugepages/hugepages-1048576kB/nr_hugepages"); err == nil {
			nr1GHugepages, _ = strconv.Atoi(string(tmpBytes))
		}

		free1GHugepages := 0
		if tmpBytes, err = ioutil.ReadFile("/sys/devices/system/node/node" + id + "/hugepages/hugepages-1048576kB/free_hugepages"); err == nil {
			free1GHugepages, _ = strconv.Atoi(string(tmpBytes))
		}

		node.Total1GMemory = nr1GHugepages
		node.Used1GMemory = nr1GHugepages - free1GHugepages

		if tmpFile, tmpErr = os.Open("/sys/devices/system/node/node" + id + "/meminfo"); tmpErr != nil {
			continue
		}

		tmpScanner = bufio.NewScanner(tmpFile)
		tmpScanner.Scan()
		memTotal, _ := strconv.ParseInt(str_utils.ParseLastSecondValue(tmpScanner.Text()), 10, 64)
		tmpScanner.Scan()
		memFree, _ := strconv.ParseInt(str_utils.ParseLastSecondValue(tmpScanner.Text()), 10, 64)
		tmpScanner.Scan()
		memUsed, _ := strconv.ParseInt(str_utils.ParseLastSecondValue(tmpScanner.Text()), 10, 64)

		fmt.Println("DEBUG mem", memTotal, memFree, memUsed)
	}

	// TODO /proc/cpuinfo

	// TODO /proc/buddyinfo

	return
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
	var tmpScanner *bufio.Scanner
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

		tmpScanner = bufio.NewScanner(tmpFile)
		tmpScanner.Scan()
		cmd := str_utils.ParseLastSecondValue(tmpScanner.Text())
		tmpScanner.Scan()
		tmpScanner.Scan()
		state := str_utils.ParseLastSecondValue(tmpScanner.Text())
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
			for tmpScanner.Scan() {
				tmpTexts = append(tmpTexts, tmpScanner.Text())
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
				tmpScanner = bufio.NewScanner(tmpFile)
				tmpScanner.Scan()
				tmpText = tmpScanner.Text()
				tmpTexts = strings.Split(tmpText, " ")
				if len(tmpTexts) != 3 {
					logger.Warningf(tctx, "Unexpected Format: path=/proc/[pid]/schedstat, text=%s", tmpText)
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

func (reader *SystemMetricReader) GetName() string {
	return reader.name
}

func (reader *SystemMetricReader) Report() ([]spec.ResourceMetric, []spec.ResourceEvent) {
	metrics := make([]spec.ResourceMetric, 0, 100)
	events := make([]spec.ResourceEvent, 0, 100)

	for _, stat := range reader.uptimeStats {
		timestamp := strconv.FormatInt(stat.timestamp.UnixNano(), 10)
		metrics = append(metrics, spec.ResourceMetric{
			Name: "system_uptime",
			Time: timestamp,
			Tag:  map[string]string{},
			Metric: map[string]interface{}{
				"uptime": stat.uptime,
			},
		})
	}

	for _, stat := range reader.loginStats {
		timestamp := strconv.FormatInt(stat.timestamp.UnixNano(), 10)
		metrics = append(metrics, spec.ResourceMetric{
			Name: "system_login",
			Time: timestamp,
			Tag:  map[string]string{},
			Metric: map[string]interface{}{
				"users": len(stat.users),
			},
		})
	}

	for _, stat := range reader.cpuStats {
		timestamp := strconv.FormatInt(stat.timestamp.UnixNano(), 10)
		for cpuName, cpu := range stat.cpuMap {
			metrics = append(metrics, spec.ResourceMetric{
				Name: "system_cpu",
				Time: timestamp,
				Tag: map[string]string{
					"cpu": cpuName,
				},
				Metric: map[string]interface{}{
					"user":       cpu[0],
					"nice":       cpu[1],
					"system":     cpu[2],
					"idle":       cpu[3],
					"iowait":     cpu[4],
					"irq":        cpu[5],
					"softirq":    cpu[6],
					"steal":      cpu[7],
					"guest":      cpu[8],
					"guest_nice": cpu[9],
				},
			})
		}

		metrics = append(metrics, spec.ResourceMetric{
			Name: "system_cpu",
			Time: timestamp,
			Tag: map[string]string{
				"cpu": "cpu",
			},
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

		stat.reportStatus = 1
	}

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

	fmt.Println("DEBUG procStats55", reader.procStats)
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
		fmt.Println("DEBUG stat", metrics[len(metrics)-1])
	}

	// TODO check metrics and issue events

	return metrics, events
}

func (reader *SystemMetricReader) Reported() {
	for _, stat := range reader.uptimeStats {
		stat.reportStatus = 2
	}
	for _, stat := range reader.loginStats {
		stat.reportStatus = 2
	}
	for _, stat := range reader.cpuStats {
		stat.reportStatus = 2
	}
	for _, stat := range reader.procsStats {
		stat.ReportStatus = 2
	}
	for _, stat := range reader.procStats {
		stat.ReportStatus = 2
	}
}
