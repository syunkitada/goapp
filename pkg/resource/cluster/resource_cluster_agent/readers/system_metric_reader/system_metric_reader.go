package system_metric_reader

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/lib/str_utils"
	"github.com/syunkitada/goapp/pkg/resource/config"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

type SystemMetricReader struct {
	conf              *config.ResourceMetricSystemConfig
	name              string
	diskStatFilters   []string
	netDevStatFilters []string
	fsStatTypes       []string
	enableLogin       bool
	enableCpu         bool
	enableMemory      bool
	enableProc        bool
	cacheLength       int
	numaNodes         []spec.NumaNodeSpec
	cpus              []spec.NumaNodeCpuSpec
	tmpDiskStatMap    map[string]TmpDiskStat
	tmpVmStat         *TmpVmStat
	uptimeStats       []UptimeStat
	loginStats        []LoginStat
	cpuStats          []CpuStat
	memStats          []MemStat
	diskStats         []DiskStat
	fsStats           []FsStat
	buddyinfoStats    []BuddyinfoStat
	vmStats           []VmStat
	tmpNetDevStatMap  map[string]TmpNetDevStat
	netDevStats       []NetDevStat
	procsStats        []ProcsStat
	procStats         []ProcStat
}

func New(conf *config.ResourceMetricSystemConfig) *SystemMetricReader {
	// f, _ := os.Open("/sys/devices/system/node/online")
	// /sys/devices/system/node/node0/meminfo

	var numaNodes []spec.NumaNodeSpec
	var cpus []spec.NumaNodeCpuSpec

	cpuinfoFile, _ := os.Open("/proc/cpuinfo")
	defer cpuinfoFile.Close()
	tmpReader := bufio.NewReader(cpuinfoFile)

	var tmpProcessor int
	var tmpPhysicalId int
	var tmpCoreId int
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
		case "physical id":
			tmpPhysicalId, _ = strconv.Atoi(splited[1])
		case "core id":
			tmpCoreId, _ = strconv.Atoi(splited[1])
			cpuSpec := spec.NumaNodeCpuSpec{
				PhysicalId: tmpPhysicalId,
				CoreId:     tmpCoreId,
				Processor:  tmpProcessor,
			}
			if len(numaNodes) == tmpPhysicalId {
				numaNodes = append(numaNodes, spec.NumaNodeSpec{
					Id:   tmpPhysicalId,
					Cpus: []spec.NumaNodeCpuSpec{cpuSpec},
				})
			}
			cpus = append(cpus, cpuSpec)

			for i := 0; i < 13; i++ {
				if _, _, tmpErr = tmpReader.ReadLine(); tmpErr != nil {
					break
				}
			}
		}
	}

	return &SystemMetricReader{
		conf:              conf,
		name:              "system",
		enableLogin:       conf.EnableLogin,
		enableCpu:         conf.EnableCpu,
		enableMemory:      conf.EnableMemory,
		enableProc:        conf.EnableProc,
		cacheLength:       conf.CacheLength,
		numaNodes:         numaNodes,
		cpus:              cpus,
		uptimeStats:       make([]UptimeStat, 0, conf.CacheLength),
		loginStats:        make([]LoginStat, 0, conf.CacheLength),
		cpuStats:          make([]CpuStat, 0, conf.CacheLength),
		memStats:          make([]MemStat, 0, conf.CacheLength),
		diskStats:         make([]DiskStat, 0, conf.CacheLength),
		buddyinfoStats:    make([]BuddyinfoStat, 0, conf.CacheLength),
		procsStats:        make([]ProcsStat, 0, conf.CacheLength),
		procStats:         make([]ProcStat, 0, conf.CacheLength),
		diskStatFilters:   []string{"loop"},
		netDevStatFilters: []string{"lo"},
		fsStatTypes:       []string{"ext4"},
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
	reportStatus  int // 0, 1(GetReport), 2(Reported)
	timestamp     time.Time
	intr          int64
	ctx           int64
	btime         int64
	processes     int64
	procs_running int64
	procs_blocked int64
	softirq       int64
}

type MemStat struct {
	ReportStatus int // 0, 1(GetReport), 2(Reported)
	Timestamp    time.Time
	NodeId       int
	MemTotal     int64
	MemFree      int64
	MemUsed      int64
	Active       int64
	Inactive     int64
	ActiveAnon   int64
	InactiveAnon int64
	ActiveFile   int64
	InactiveFile int64
	Unevictable  int64
	Mlocked      int64
	Dirty        int64
	Writeback    int64
	FilePages    int64
	Mapped       int64
	AnonPages    int64
	Shmem        int64
	KernelStack  int64
	PageTables   int64
	NfsUnstable  int64
	Bounce       int64
	WritebackTmp int64
	KReclaimable int64
	Slab         int64
	SReclaimable int64
	SUnreclaim   int64
}

type BuddyinfoStat struct {
	ReportStatus int // 0, 1(GetReport), 2(Reported)
	Timestamp    time.Time
	NodeId       int64
	M4K          int64
	M8K          int64
	M16K         int64
	M32K         int64
	M64K         int64
	M128K        int64
	M256K        int64
	M512K        int64
	M1M          int64
	M2M          int64
	M4M          int64
}

type VmStat struct {
	ReportStatus     int // 0, 1(GetReport), 2(Reported)
	Timestamp        time.Time
	DiffPgscanKswapd int64
	DiffPgscanDirect int64
	DiffPgfault      int64
	DiffPswapin      int64
	DiffPswapout     int64
}

type TmpVmStat struct {
	Timestamp    time.Time
	PgscanKswapd int64
	PgscanDirect int64
	Pgfault      int64
	Pswapin      int64
	Pswapout     int64
}

type TmpDiskStat struct {
	Timestamp         time.Time
	PblockSize        int64
	ReadsCompleted    int64
	ReadsMerges       int64
	ReadSectors       int64
	ReadMs            int64
	WritesCompleted   int64
	WritesMerges      int64
	WriteSectors      int64
	WriteMs           int64
	ProgressIos       int64
	IosMs             int64
	WeightedIosMs     int64
	DiscardsCompleted int64
	DiscardsMerges    int64
	DiscardSectors    int64
	DiscardMs         int64
}

type DiskStat struct {
	ReportStatus        int
	Timestamp           time.Time
	Device              string
	ReadsPerSec         int64
	RmergesPerSec       int64
	ReadBytesPerSec     int64
	ReadMsPerSec        int64
	WritesPerSec        int64
	WmergesPerSec       int64
	WriteBytesPerSec    int64
	WriteMsPerSec       int64
	DiscardsPerSec      int64
	DmergesPerSec       int64
	DiscardBytesPerSec  int64
	DiscardMsPerSec     int64
	ProgressIos         int64
	IosMsPerSec         int64
	WeightedIosMsPerSec int64
}

type FsStat struct {
	ReportStatus int
	Timestamp    time.Time
	TotalSize    int64
	FreeSize     int64
	UsedSize     int64
	Files        int64
}

type TmpNetDevStat struct {
	ReportStatus    int
	Timestamp       time.Time
	ReceiveBytes    int64
	ReceivePackets  int64
	ReceiveErrors   int64
	ReceiveDrops    int64
	TransmitBytes   int64
	TransmitPackets int64
	TransmitErrors  int64
	TransmitDrops   int64
}

type NetDevStat struct {
	ReportStatus          int
	Timestamp             time.Time
	Interface             string
	ReceiveBytesPerSec    int64
	ReceivePacketsPerSec  int64
	ReceiveDiffErrors     int64
	ReceiveDiffDrops      int64
	TransmitBytesPerSec   int64
	TransmitPacketsPerSec int64
	TransmitDiffErrors    int64
	TransmitDiffDrops     int64
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

func (reader *SystemMetricReader) GetNumaNodes(tctx *logger.TraceContext) []spec.NumaNodeSpec {
	return reader.numaNodes
}

func (reader *SystemMetricReader) Read(tctx *logger.TraceContext) (err error) {
	timestamp := time.Now()

	// Read /proc/uptime
	// uptime(s)  idle(s)
	// 2906.26 5507.43
	procUptime, _ := os.Open("/proc/uptime")
	defer procUptime.Close()
	var tmpErr error
	tmpReader := bufio.NewReader(procUptime)
	tmpBytes, _, _ := tmpReader.ReadLine()
	uptimeWords := strings.Split(string(tmpBytes), " ")
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

	// Read /sys/devices/system/node/node.*/hugepages
	// Read /sys/devices/system/node/node.*/meminfo
	var tmpFile *os.File
	for id, node := range reader.numaNodes {
		nr1GHugepages := 0
		if tmpBytes, err = ioutil.ReadFile("/sys/devices/system/node/node" + strconv.Itoa(id) + "/hugepages/hugepages-1048576kB/nr_hugepages"); err == nil {
			nr1GHugepages, _ = strconv.Atoi(string(tmpBytes))
		}

		free1GHugepages := 0
		if tmpBytes, err = ioutil.ReadFile("/sys/devices/system/node/node" + strconv.Itoa(id) + "/hugepages/hugepages-1048576kB/free_hugepages"); err == nil {
			free1GHugepages, _ = strconv.Atoi(string(tmpBytes))
		}

		node.Total1GMemory = nr1GHugepages
		node.Used1GMemory = nr1GHugepages - free1GHugepages

		if tmpFile, tmpErr = os.Open("/sys/devices/system/node/node" + strconv.Itoa(id) + "/meminfo"); tmpErr != nil {
			continue
		}
		tmpReader := bufio.NewReader(tmpFile)

		tmpBytes, _, _ = tmpReader.ReadLine()
		memTotal, _ := strconv.ParseInt(str_utils.ParseLastSecondValue(string(tmpBytes)), 10, 64)
		tmpBytes, _, _ = tmpReader.ReadLine()
		memFree, _ := strconv.ParseInt(str_utils.ParseLastSecondValue(string(tmpBytes)), 10, 64)
		tmpBytes, _, _ = tmpReader.ReadLine()
		memUsed, _ := strconv.ParseInt(str_utils.ParseLastSecondValue(string(tmpBytes)), 10, 64)
		tmpBytes, _, _ = tmpReader.ReadLine()
		active, _ := strconv.ParseInt(str_utils.ParseLastSecondValue(string(tmpBytes)), 10, 64)
		tmpBytes, _, _ = tmpReader.ReadLine()
		inactive, _ := strconv.ParseInt(str_utils.ParseLastSecondValue(string(tmpBytes)), 10, 64)
		tmpBytes, _, _ = tmpReader.ReadLine()
		activeAnon, _ := strconv.ParseInt(str_utils.ParseLastSecondValue(string(tmpBytes)), 10, 64)
		tmpBytes, _, _ = tmpReader.ReadLine()
		inactiveAnon, _ := strconv.ParseInt(str_utils.ParseLastSecondValue(string(tmpBytes)), 10, 64)
		tmpBytes, _, _ = tmpReader.ReadLine()
		activeFile, _ := strconv.ParseInt(str_utils.ParseLastSecondValue(string(tmpBytes)), 10, 64)
		tmpBytes, _, _ = tmpReader.ReadLine()
		inactiveFile, _ := strconv.ParseInt(str_utils.ParseLastSecondValue(string(tmpBytes)), 10, 64)
		tmpBytes, _, _ = tmpReader.ReadLine()
		unevictable, _ := strconv.ParseInt(str_utils.ParseLastSecondValue(string(tmpBytes)), 10, 64)
		tmpBytes, _, _ = tmpReader.ReadLine()
		mlocked, _ := strconv.ParseInt(str_utils.ParseLastSecondValue(string(tmpBytes)), 10, 64)
		tmpBytes, _, _ = tmpReader.ReadLine()
		dirty, _ := strconv.ParseInt(str_utils.ParseLastSecondValue(string(tmpBytes)), 10, 64)
		tmpBytes, _, _ = tmpReader.ReadLine()
		writeback, _ := strconv.ParseInt(str_utils.ParseLastSecondValue(string(tmpBytes)), 10, 64)
		tmpBytes, _, _ = tmpReader.ReadLine()
		filePages, _ := strconv.ParseInt(str_utils.ParseLastSecondValue(string(tmpBytes)), 10, 64)
		tmpBytes, _, _ = tmpReader.ReadLine()
		mapped, _ := strconv.ParseInt(str_utils.ParseLastSecondValue(string(tmpBytes)), 10, 64)
		tmpBytes, _, _ = tmpReader.ReadLine()
		anonPages, _ := strconv.ParseInt(str_utils.ParseLastSecondValue(string(tmpBytes)), 10, 64)
		tmpBytes, _, _ = tmpReader.ReadLine()
		shmem, _ := strconv.ParseInt(str_utils.ParseLastSecondValue(string(tmpBytes)), 10, 64)
		tmpBytes, _, _ = tmpReader.ReadLine()
		kernelStack, _ := strconv.ParseInt(str_utils.ParseLastSecondValue(string(tmpBytes)), 10, 64)
		tmpBytes, _, _ = tmpReader.ReadLine()
		pageTables, _ := strconv.ParseInt(str_utils.ParseLastSecondValue(string(tmpBytes)), 10, 64)
		tmpBytes, _, _ = tmpReader.ReadLine()
		nfsUnstable, _ := strconv.ParseInt(str_utils.ParseLastSecondValue(string(tmpBytes)), 10, 64)
		tmpBytes, _, _ = tmpReader.ReadLine()
		bounce, _ := strconv.ParseInt(str_utils.ParseLastSecondValue(string(tmpBytes)), 10, 64)
		tmpBytes, _, _ = tmpReader.ReadLine()
		writebackTmp, _ := strconv.ParseInt(str_utils.ParseLastSecondValue(string(tmpBytes)), 10, 64)
		tmpBytes, _, _ = tmpReader.ReadLine()
		kReclaimable, _ := strconv.ParseInt(str_utils.ParseLastSecondValue(string(tmpBytes)), 10, 64)
		tmpBytes, _, _ = tmpReader.ReadLine()
		slab, _ := strconv.ParseInt(str_utils.ParseLastSecondValue(string(tmpBytes)), 10, 64)
		tmpBytes, _, _ = tmpReader.ReadLine()
		sReclaimable, _ := strconv.ParseInt(str_utils.ParseLastSecondValue(string(tmpBytes)), 10, 64)
		tmpBytes, _, _ = tmpReader.ReadLine()
		sUnreclaim, _ := strconv.ParseInt(str_utils.ParseLastSecondValue(string(tmpBytes)), 10, 64)

		node.TotalMemory = int(memTotal)
		node.UsedMemory = int(memUsed)

		reader.memStats = append(reader.memStats, MemStat{
			ReportStatus: 0,
			Timestamp:    timestamp,
			NodeId:       id,
			MemTotal:     memTotal,
			MemFree:      memFree,
			MemUsed:      memUsed,
			Active:       active,
			Inactive:     inactive,
			ActiveAnon:   activeAnon,
			InactiveAnon: inactiveAnon,
			ActiveFile:   activeFile,
			InactiveFile: inactiveFile,
			Unevictable:  unevictable,
			Mlocked:      mlocked,
			Dirty:        dirty,
			Writeback:    writeback,
			FilePages:    filePages,
			Mapped:       mapped,
			AnonPages:    anonPages,
			Shmem:        shmem,
			KernelStack:  kernelStack,
			PageTables:   pageTables,
			NfsUnstable:  nfsUnstable,
			Bounce:       bounce,
			WritebackTmp: writebackTmp,
			KReclaimable: kReclaimable,
			Slab:         slab,
			SReclaimable: sReclaimable,
			SUnreclaim:   sUnreclaim,
		})
	}

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
			reportStatus:  0,
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

	if reader.enableProc {
		// Read /proc/
		reader.ReadProc(tctx)
	}

	// Read /proc/vmstat
	if reader.tmpVmStat == nil {
		reader.tmpVmStat = reader.ReadVmStat(tctx)
	} else {
		tmpVmStat := reader.ReadVmStat(tctx)
		reader.vmStats = append(reader.vmStats, VmStat{
			ReportStatus:     0,
			Timestamp:        timestamp,
			DiffPgscanKswapd: tmpVmStat.PgscanKswapd - reader.tmpVmStat.PgscanKswapd,
			DiffPgscanDirect: tmpVmStat.PgscanDirect - reader.tmpVmStat.PgscanDirect,
			DiffPgfault:      tmpVmStat.Pgfault - reader.tmpVmStat.Pgfault,
			DiffPswapin:      tmpVmStat.Pswapin - reader.tmpVmStat.Pswapin,
			DiffPswapout:     tmpVmStat.Pswapout - reader.tmpVmStat.Pswapout,
		})
		reader.tmpVmStat = tmpVmStat
	}

	// Read /proc/buddyinfo
	// $ cat /proc/buddyinfo
	//                           4K     8k    16k    32k    64k   128k   256k   512k     1M     2M     4M
	// Node 0, zone      DMA      0      0      0      1      2      1      1      0      1      1      3
	// Node 0, zone    DMA32      3      3      3      3      3      2      5      6      5      2    874
	// Node 0, zone   Normal  24727  53842  18419  15120  10448   4451   1761    804    382    105    229
	buddyinfoFile, _ := os.Open("/proc/buddyinfo")
	defer buddyinfoFile.Close()
	tmpReader = bufio.NewReader(buddyinfoFile)
	for {
		tmpBytes, _, tmpErr = tmpReader.ReadLine()
		if tmpErr != nil {
			break
		}
		buddyinfo := str_utils.SplitSpace(string(tmpBytes))
		if len(buddyinfo) < 10 {
			continue
		}
		if buddyinfo[3] == "Normal" {
			nodeId, _ := strconv.ParseInt(buddyinfo[1], 10, 64)
			m4K, _ := strconv.ParseInt(buddyinfo[4], 10, 64)
			m8K, _ := strconv.ParseInt(buddyinfo[5], 10, 64)
			m16K, _ := strconv.ParseInt(buddyinfo[6], 10, 64)
			m32K, _ := strconv.ParseInt(buddyinfo[7], 10, 64)
			m64K, _ := strconv.ParseInt(buddyinfo[8], 10, 64)
			m128K, _ := strconv.ParseInt(buddyinfo[9], 10, 64)
			m256K, _ := strconv.ParseInt(buddyinfo[10], 10, 64)
			m512K, _ := strconv.ParseInt(buddyinfo[11], 10, 64)
			m1M, _ := strconv.ParseInt(buddyinfo[12], 10, 64)
			m2M, _ := strconv.ParseInt(buddyinfo[13], 10, 64)
			m4M, _ := strconv.ParseInt(buddyinfo[14], 10, 64)

			reader.buddyinfoStats = append(reader.buddyinfoStats, BuddyinfoStat{
				ReportStatus: 0,
				Timestamp:    timestamp,
				NodeId:       nodeId,
				M4K:          m4K,
				M8K:          m8K,
				M16K:         m16K,
				M32K:         m32K,
				M64K:         m64K,
				M128K:        m128K,
				M256K:        m256K,
				M512K:        m512K,
				M1M:          m1M,
				M2M:          m2M,
				M4M:          m4M,
			})
		}
	}

	// Read /proc/diskstat
	if reader.tmpDiskStatMap == nil {
		reader.tmpDiskStatMap = reader.ReadDiskStat(tctx)
	} else {
		tmpDiskStatMap := reader.ReadDiskStat(tctx)
		for dev, cstat := range tmpDiskStatMap {
			bstat, ok := reader.tmpDiskStatMap[dev]
			if !ok {
				continue
			}
			interval := cstat.Timestamp.Unix() - bstat.Timestamp.Unix()
			readsPerSec := int64((cstat.ReadsCompleted - bstat.ReadsCompleted) / int64(interval))
			rmergesPerSec := int64((cstat.ReadsMerges - bstat.ReadsMerges) / int64(interval))
			readBytesPerSec := int64(((cstat.ReadSectors - bstat.ReadSectors) * cstat.PblockSize) / int64(interval))
			readMsPerSec := int64((cstat.ReadMs - bstat.ReadMs) / int64(interval))

			writesPerSec := int64((cstat.WritesCompleted - bstat.WritesCompleted) / int64(interval))
			wmergesPerSec := int64((cstat.WritesMerges - bstat.WritesMerges) / int64(interval))
			writeBytesPerSec := int64(((cstat.WriteSectors - bstat.WriteSectors) * cstat.PblockSize) / int64(interval))
			writeMsPerSec := int64((cstat.WriteMs - bstat.WriteMs) / int64(interval))

			discardsPerSec := int64((cstat.DiscardsCompleted - bstat.DiscardsCompleted) / int64(interval))
			dmergesPerSec := int64((cstat.DiscardsMerges - bstat.DiscardsMerges) / int64(interval))
			discardBytesPerSec := int64(((cstat.DiscardSectors - bstat.DiscardSectors) * cstat.PblockSize) / int64(interval))
			discardMsPerSec := int64((cstat.DiscardMs - bstat.DiscardMs) / int64(interval))

			iosMsPerSec := int64((cstat.IosMs - bstat.IosMs) / int64(interval))
			weightedIosMsPerSec := int64((cstat.WeightedIosMs - bstat.WeightedIosMs) / int64(interval))

			reader.diskStats = append(reader.diskStats, DiskStat{
				ReportStatus:        0,
				Timestamp:           timestamp,
				Device:              dev,
				ReadsPerSec:         readsPerSec,
				RmergesPerSec:       rmergesPerSec,
				ReadBytesPerSec:     readBytesPerSec,
				ReadMsPerSec:        readMsPerSec,
				WritesPerSec:        writesPerSec,
				WmergesPerSec:       wmergesPerSec,
				WriteBytesPerSec:    writeBytesPerSec,
				WriteMsPerSec:       writeMsPerSec,
				DiscardsPerSec:      discardsPerSec,
				DmergesPerSec:       dmergesPerSec,
				DiscardBytesPerSec:  discardBytesPerSec,
				DiscardMsPerSec:     discardMsPerSec,
				ProgressIos:         cstat.ProgressIos,
				IosMsPerSec:         iosMsPerSec,
				WeightedIosMsPerSec: weightedIosMsPerSec,
			})
		}

		reader.tmpDiskStatMap = tmpDiskStatMap
	}

	// Read mounts
	// /etc/mtab -> ../proc/self/mounts
	mountsFile, _ := os.Open("/proc/self/mounts")
	defer mountsFile.Close()
	tmpReader = bufio.NewReader(mountsFile)
	var splitedLine []string
	var isMatch bool
	for {
		tmpBytes, _, tmpErr = tmpReader.ReadLine()
		if tmpErr != nil {
			break
		}
		splitedLine = strings.Split(string(tmpBytes), " ")
		isMatch = false
		for _, fsType := range reader.fsStatTypes {
			if splitedLine[2] == fsType {
				isMatch = true
				break
			}
		}
		if !isMatch {
			continue
		}
		var statfs syscall.Statfs_t
		if tmpErr = syscall.Statfs(splitedLine[1], &statfs); tmpErr != nil {
			continue
		}
		totalSize := int64(statfs.Blocks) * statfs.Bsize
		freeSize := int64(statfs.Bavail) * statfs.Bsize

		reader.fsStats = append(reader.fsStats, FsStat{
			ReportStatus: 0,
			Timestamp:    timestamp,
			TotalSize:    totalSize,
			FreeSize:     freeSize,
			UsedSize:     totalSize - freeSize,
			Files:        int64(statfs.Files),
		})
	}

	// Read /proc/diskstat
	if reader.tmpNetDevStatMap == nil {
		reader.tmpNetDevStatMap = reader.ReadNetDevStat(tctx)
	} else {
		tmpNetDevStatMap := reader.ReadNetDevStat(tctx)
		for dev, cstat := range tmpNetDevStatMap {
			bstat, ok := reader.tmpNetDevStatMap[dev]
			if !ok {
				continue
			}
			interval := cstat.Timestamp.Unix() - bstat.Timestamp.Unix()
			receiveBytesPerSec := int64((cstat.ReceiveBytes - bstat.ReceiveBytes) / int64(interval))
			receivePacketsPerSec := int64((cstat.ReceivePackets - bstat.ReceivePackets) / int64(interval))
			receiveDiffErrors := int64((cstat.ReceiveErrors - bstat.ReceiveErrors) / int64(interval))
			receiveDiffDrops := int64((cstat.ReceiveDrops - bstat.ReceiveDrops) / int64(interval))
			transmitBytesPerSec := int64((cstat.TransmitBytes - bstat.TransmitBytes) / int64(interval))
			transmitPacketsPerSec := int64((cstat.TransmitPackets - bstat.TransmitPackets) / int64(interval))
			transmitDiffErrors := int64((cstat.TransmitErrors - bstat.TransmitErrors) / int64(interval))
			transmitDiffDrops := int64((cstat.TransmitDrops - bstat.TransmitDrops) / int64(interval))

			reader.netDevStats = append(reader.netDevStats, NetDevStat{
				ReportStatus:          0,
				Timestamp:             timestamp,
				Interface:             dev,
				ReceiveBytesPerSec:    receiveBytesPerSec,
				ReceivePacketsPerSec:  receivePacketsPerSec,
				ReceiveDiffErrors:     receiveDiffErrors,
				ReceiveDiffDrops:      receiveDiffDrops,
				TransmitBytesPerSec:   transmitBytesPerSec,
				TransmitPacketsPerSec: transmitPacketsPerSec,
				TransmitDiffErrors:    transmitDiffErrors,
				TransmitDiffDrops:     transmitDiffDrops,
			})
		}

		reader.tmpNetDevStatMap = tmpNetDevStatMap
	}

	// TODO /proc/net/netstat

	return
}

func (reader *SystemMetricReader) ReadVmStat(tctx *logger.TraceContext) (tmpVmStat *TmpVmStat) {
	// Read /proc/vmstat
	timestamp := time.Now()
	f, _ := os.Open("/proc/vmstat")
	defer f.Close()
	tmpReader := bufio.NewReader(f)
	vmstat := map[string]string{}
	for {
		tmpBytes, _, tmpErr := tmpReader.ReadLine()
		if tmpErr != nil {
			break
		}
		columns := strings.Split(string(tmpBytes), " ")
		vmstat[columns[0]] = columns[1]
	}

	pgscanKswapd, _ := strconv.ParseInt(str_utils.ParseLastValue(vmstat["pgscan_kswapd"]), 10, 64)
	pgscanDirect, _ := strconv.ParseInt(str_utils.ParseLastValue(vmstat["pgscan_direct"]), 10, 64)
	pgfault, _ := strconv.ParseInt(str_utils.ParseLastValue(vmstat["pgfault"]), 10, 64)

	pswapin, _ := strconv.ParseInt(str_utils.ParseLastValue(vmstat["pswapin"]), 10, 64)
	pswapout, _ := strconv.ParseInt(str_utils.ParseLastValue(vmstat["pswapout"]), 10, 64)

	tmpVmStat = &TmpVmStat{
		Timestamp:    timestamp,
		PgscanKswapd: pgscanKswapd,
		PgscanDirect: pgscanDirect,
		Pgfault:      pgfault,
		Pswapin:      pswapin,
		Pswapout:     pswapout,
	}
	return
}

func (reader *SystemMetricReader) ReadDiskStat(tctx *logger.TraceContext) (tmpDiskStatMap map[string]TmpDiskStat) {
	// Read /proc/diskstats

	// 259       0 nvme0n1 94360 70783 6403078 67950 136558 90723 6419592 38105 0 97140 59208 0 0 0 0
	// 259       0 nvme0n1 94360 70783 6403078 67950 136611 90751 6423880 38111 0 97200 59208 0 0 0 0
	// 259       0 nvme0n1 94364 70783 6403230 67951 155638 101247 7087392 41420 0 107356 59208 0 0 0 0

	// Field  1 -- # of reads completed
	// Field  2 -- # of reads merged, field 6 -- # of writes merged
	// Field  3 -- # of sectors read
	// Field  4 -- # of milliseconds spent reading
	// Field  5 -- # of writes completed
	// Field  6 -- # of writes merged
	// Field  7 -- # of sectors written
	// Field  8 -- # of milliseconds spent writing
	// Field  9 -- # of I/Os currently in progress
	// Field 10 -- # of milliseconds spent doing I/Os
	// Field 11 -- weighted # of milliseconds spent doing I/Os
	// Field 12 -- # of discards completed
	// Field 13 -- # of discards merged
	// Field 14 -- # of sectors discarded
	// Field 15 -- # of milliseconds spent discarding

	timestamp := time.Now()
	f, _ := os.Open("/proc/diskstats")
	defer f.Close()
	tmpReader := bufio.NewReader(f)
	tmpDiskStatMap = map[string]TmpDiskStat{}
	var isFiltered bool
	for {
		tmpBytes, _, tmpErr := tmpReader.ReadLine()
		if tmpErr != nil {
			break
		}
		columns := str_utils.SplitSpace(string(tmpBytes))
		isFiltered = false
		for _, filter := range reader.diskStatFilters {
			if strings.Index(columns[2], filter) > -1 {
				isFiltered = true
				break
			}
		}
		if isFiltered {
			continue
		}

		pblockSizeFile, tmpErr := os.Open("/sys/block/" + columns[2] + "/queue/physical_block_size")
		if tmpErr != nil {
			continue
		}
		pblockSizeReader := bufio.NewReader(pblockSizeFile)
		pblockSizeBytes, _, tmpErr := pblockSizeReader.ReadLine()
		pblockSizeFile.Close()
		if tmpErr != nil {
			continue
		}
		pblockSize, _ := strconv.ParseInt(string(pblockSizeBytes), 10, 64)

		readsCompleted, _ := strconv.ParseInt(columns[3], 10, 64)
		readsMerges, _ := strconv.ParseInt(columns[4], 10, 64)
		readSectors, _ := strconv.ParseInt(columns[5], 10, 64)
		readMs, _ := strconv.ParseInt(columns[6], 10, 64)
		writesCompleted, _ := strconv.ParseInt(columns[7], 10, 64)
		writesMerges, _ := strconv.ParseInt(columns[8], 10, 64)
		writeSectors, _ := strconv.ParseInt(columns[9], 10, 64)
		writeMs, _ := strconv.ParseInt(columns[10], 10, 64)

		progressIos, _ := strconv.ParseInt(columns[11], 10, 64)
		iosMs, _ := strconv.ParseInt(columns[12], 10, 64)
		weightedIosMs, _ := strconv.ParseInt(columns[13], 10, 64)

		discardsCompleted, _ := strconv.ParseInt(columns[14], 10, 64)
		discardsMerges, _ := strconv.ParseInt(columns[15], 10, 64)
		discardSectors, _ := strconv.ParseInt(columns[16], 10, 64)
		discardMs, _ := strconv.ParseInt(columns[17], 10, 64)

		tmpDiskStatMap[columns[2]] = TmpDiskStat{
			Timestamp:         timestamp,
			PblockSize:        pblockSize,
			ReadsCompleted:    readsCompleted,
			ReadsMerges:       readsMerges,
			ReadSectors:       readSectors,
			ReadMs:            readMs,
			WritesCompleted:   writesCompleted,
			WritesMerges:      writesMerges,
			WriteSectors:      writeSectors,
			WriteMs:           writeMs,
			ProgressIos:       progressIos,
			IosMs:             iosMs,
			WeightedIosMs:     weightedIosMs,
			DiscardsCompleted: discardsCompleted,
			DiscardsMerges:    discardsMerges,
			DiscardSectors:    discardSectors,
			DiscardMs:         discardMs,
		}
	}

	return
}

func (reader *SystemMetricReader) ReadNetDevStat(tctx *logger.TraceContext) (tmpNetDevStatMap map[string]TmpNetDevStat) {
	// $ cat /proc/net/dev
	// Inter-|   Receive                                                |  Transmit
	//  face |bytes    packets errs drop fifo frame compressed multicast|bytes    packets errs drop fifo colls carrier compressed
	//  com-1-ex:    1426      19    0    0    0     0          0         0     4616      43    0    0    0     0       0          0
	//  enp31s0: 7855580   30554    0    0    0     0          0      1408 19677375   42829    0    0    0     0       0          0
	//      lo: 1442597782 3051437    0    0    0     0          0         0 1442597782 3051437    0    0    0     0       0          0
	// 	 com-0-ex:   29026     447    0    0    0     0          0         0    34621     471    0    0    0     0       0          0
	// 	 com-2-ex:   26578     383    0    0    0     0          0         0    32083     406    0    0    0     0       0          0
	// 	 com-4-ex:   28084     420    0    0    0     0          0         0    33499     442    0    0    0     0       0          0
	// 	 docker0:       0       0    0    0    0     0          0         0        0       0    0    0    0     0       0          0
	timestamp := time.Now()
	netdevFile, _ := os.Open("/proc/net/dev")
	defer netdevFile.Close()
	tmpReader := bufio.NewReader(netdevFile)
	_, _, _ = tmpReader.ReadLine()
	_, _, _ = tmpReader.ReadLine()

	var tmpBytes []byte
	var tmpErr error
	tmpNetDevStatMap = map[string]TmpNetDevStat{}
	var isFiltered bool
	for {
		tmpBytes, _, tmpErr = tmpReader.ReadLine()
		if tmpErr != nil {
			break
		}
		splitedStr := str_utils.SplitColon(string(tmpBytes))
		columns := str_utils.SplitSpace(splitedStr[1])

		isFiltered = false
		for _, filter := range reader.netDevStatFilters {
			if strings.Index(splitedStr[0], filter) > -1 {
				isFiltered = true
				break
			}
		}
		if isFiltered {
			continue
		}

		receiveBytes, _ := strconv.ParseInt(columns[0], 10, 64)
		receivePackets, _ := strconv.ParseInt(columns[1], 10, 64)
		receiveErrors, _ := strconv.ParseInt(columns[2], 10, 64)
		receiveDrops, _ := strconv.ParseInt(columns[3], 10, 64)

		transmitBytes, _ := strconv.ParseInt(columns[8], 10, 64)
		transmitPackets, _ := strconv.ParseInt(columns[9], 10, 64)
		transmitErrors, _ := strconv.ParseInt(columns[10], 10, 64)
		transmitDrops, _ := strconv.ParseInt(columns[11], 10, 64)

		tmpNetDevStatMap[splitedStr[0]] = TmpNetDevStat{
			Timestamp:       timestamp,
			ReceiveBytes:    receiveBytes,
			ReceivePackets:  receivePackets,
			ReceiveErrors:   receiveErrors,
			ReceiveDrops:    receiveDrops,
			TransmitBytes:   transmitBytes,
			TransmitPackets: transmitPackets,
			TransmitErrors:  transmitErrors,
			TransmitDrops:   transmitDrops,
		}
	}

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

	for _, stat := range reader.memStats {
		timestamp := strconv.FormatInt(stat.Timestamp.UnixNano(), 10)

		reclaimable := (stat.Inactive + stat.KReclaimable + stat.SReclaimable) * 1000
		metrics = append(metrics, spec.ResourceMetric{
			Name: "system_mem",
			Time: timestamp,
			Tag: map[string]string{
				"node_id": strconv.Itoa(stat.NodeId),
			},
			Metric: map[string]interface{}{
				"reclaimable":   reclaimable,
				"mem_total":     stat.MemTotal * 1000,
				"mem_free":      stat.MemFree * 1000,
				"mem_used":      stat.MemUsed * 1000,
				"active":        stat.Active * 1000,
				"inactive":      stat.Inactive * 1000,
				"active_anon":   stat.ActiveAnon * 1000,
				"inactive_anon": stat.InactiveAnon * 1000,
				"active_file":   stat.ActiveFile * 1000,
				"inactive_file": stat.InactiveFile * 1000,
				"unevictable":   stat.Unevictable * 1000,
				"mlocked":       stat.Mlocked * 1000,
				"dirty":         stat.Dirty * 1000,
				"writeback":     stat.Writeback * 1000,
				"writeback_tmp": stat.WritebackTmp * 1000,
				"k_reclaimable": stat.KReclaimable * 1000,
				"slab":          stat.Slab * 1000,
				"s_reclaimable": stat.SReclaimable * 1000,
				"s_unreclaim":   stat.SUnreclaim * 1000,
			},
		})
	}

	for _, stat := range reader.vmStats {
		timestamp := strconv.FormatInt(stat.Timestamp.UnixNano(), 10)

		metrics = append(metrics, spec.ResourceMetric{
			Name: "system_vmstat",
			Time: timestamp,
			Metric: map[string]interface{}{
				"pgscan_kswapd": stat.DiffPgscanKswapd,
				"pgscan_direct": stat.DiffPgscanDirect,
				"pgfault":       stat.DiffPgfault,
				"pswapin":       stat.DiffPswapin,
				"pswapout":      stat.DiffPswapout,
			},
		})
	}

	for _, stat := range reader.diskStats {
		timestamp := strconv.FormatInt(stat.Timestamp.UnixNano(), 10)
		metrics = append(metrics, spec.ResourceMetric{
			Name: "system_diskstat",
			Time: timestamp,
			Tag: map[string]string{
				"dev": stat.Device,
			},
			Metric: map[string]interface{}{
				"reads_per_sec":       stat.ReadsPerSec,
				"read_bytes_per_sec":  stat.ReadBytesPerSec,
				"read_ms_per_sec":     stat.ReadMsPerSec,
				"writes_per_sec":      stat.WritesPerSec,
				"write_bytes_per_sec": stat.WriteBytesPerSec,
				"write_ms_per_sec":    stat.WriteMsPerSec,
				"progress_ios":        stat.ProgressIos,
			},
		})
	}

	for _, stat := range reader.fsStats {
		timestamp := strconv.FormatInt(stat.Timestamp.UnixNano(), 10)
		metrics = append(metrics, spec.ResourceMetric{
			Name: "system_fsstat",
			Time: timestamp,
			Metric: map[string]interface{}{
				"total_size": stat.TotalSize,
				"free_size":  stat.FreeSize,
				"used_size":  stat.UsedSize,
				"files":      stat.Files,
			},
		})
	}

	for _, stat := range reader.netDevStats {
		timestamp := strconv.FormatInt(stat.Timestamp.UnixNano(), 10)
		metrics = append(metrics, spec.ResourceMetric{
			Name: "system_netdevstat",
			Time: timestamp,
			Tag: map[string]string{
				"interface": stat.Interface,
			},
			Metric: map[string]interface{}{
				"receive_bytes_per_sec":    stat.ReceiveBytesPerSec,
				"receive_packets_per_sec":  stat.ReceivePacketsPerSec,
				"receive_errors":           stat.ReceiveDiffErrors,
				"receive_drops":            stat.ReceiveDiffDrops,
				"transmit_bytes_per_sec":   stat.TransmitBytesPerSec,
				"transmit_packets_per_sec": stat.TransmitPacketsPerSec,
				"transmit_errors":          stat.TransmitDiffErrors,
				"transmit_drops":           stat.TransmitDiffDrops,
			},
		})
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
	for _, stat := range reader.memStats {
		stat.ReportStatus = 2
	}
	for _, stat := range reader.vmStats {
		stat.ReportStatus = 2
	}
	for _, stat := range reader.diskStats {
		stat.ReportStatus = 2
	}
	for _, stat := range reader.fsStats {
		stat.ReportStatus = 2
	}
	for _, stat := range reader.netDevStats {
		stat.ReportStatus = 2
	}
}
