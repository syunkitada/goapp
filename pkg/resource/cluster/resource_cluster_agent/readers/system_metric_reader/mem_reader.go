package system_metric_reader

import (
	"bufio"
	"fmt"
	"io/ioutil"
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

type MemStat struct {
	ReportStatus int // 0, 1(GetReport), 2(Reported)
	Timestamp    time.Time
	NodeId       int
	MemTotal     int64
	MemFree      int64
	MemUsed      int64
	MemAvailable int64
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

type MemReader struct {
	conf               *config.ResourceMetricSystemConfig
	cacheLength        int
	lenNodes           int
	memStats           []MemStat
	systemMetricReader *SystemMetricReader
	tmpVmStat          *TmpVmStat
	vmStats            []VmStat

	checkAvailableOccurences         int
	checkAvailableReissueDuration    int
	checkAvailableWarnAvailableRatio float64
	checkAvailableWarnNodeCounters   []int

	checkPgscanOccurences              int
	checkPgscanReissueDuration         int
	checkPgscanWarnPgscanDirect        int64
	checkPgscanWarnPgscanDirectCounter int
}

func NewMemReader(conf *config.ResourceMetricSystemConfig, systemMetricReader *SystemMetricReader) SubMetricReader {
	lenNodes := len(systemMetricReader.NumaNodes)
	return &MemReader{
		conf:               conf,
		cacheLength:        conf.CacheLength,
		memStats:           make([]MemStat, 0, conf.CacheLength),
		systemMetricReader: systemMetricReader,
		lenNodes:           lenNodes,

		checkAvailableOccurences:         conf.Mem.CheckAvailable.Occurences,
		checkAvailableReissueDuration:    conf.Mem.CheckAvailable.ReissueDuration,
		checkAvailableWarnAvailableRatio: conf.Mem.CheckAvailable.WarnAvailableRatio,
		checkAvailableWarnNodeCounters:   make([]int, 0, len(systemMetricReader.NumaNodes)),

		checkPgscanOccurences:              conf.Mem.CheckPgscan.Occurences,
		checkPgscanReissueDuration:         conf.Mem.CheckPgscan.ReissueDuration,
		checkPgscanWarnPgscanDirect:        conf.Mem.CheckPgscan.WarnPgscanDirect,
		checkPgscanWarnPgscanDirectCounter: 0,
	}
}

func (reader *MemReader) Read(tctx *logger.TraceContext) {
	timestamp := time.Now()

	// Read /proc/vmstat
	if reader.tmpVmStat == nil {
		reader.tmpVmStat = reader.readTmpVmStat(tctx)
	} else {
		if len(reader.vmStats) > reader.cacheLength {
			reader.vmStats = reader.vmStats[1:]
		}

		tmpVmStat := reader.readTmpVmStat(tctx)

		diffPgscanDirect := tmpVmStat.PgscanDirect - reader.tmpVmStat.PgscanDirect
		if diffPgscanDirect > reader.checkPgscanWarnPgscanDirect {
			reader.checkPgscanWarnPgscanDirectCounter += 1
		} else {
			reader.checkPgscanWarnPgscanDirectCounter = 0
		}

		reader.vmStats = append(reader.vmStats, VmStat{
			ReportStatus:     0,
			Timestamp:        timestamp,
			DiffPgscanKswapd: tmpVmStat.PgscanKswapd - reader.tmpVmStat.PgscanKswapd,
			DiffPgscanDirect: diffPgscanDirect,
			DiffPgfault:      tmpVmStat.Pgfault - reader.tmpVmStat.Pgfault,
			DiffPswapin:      tmpVmStat.Pswapin - reader.tmpVmStat.Pswapin,
			DiffPswapout:     tmpVmStat.Pswapout - reader.tmpVmStat.Pswapout,
		})
		reader.tmpVmStat = tmpVmStat
	}

	// Read /sys/devices/system/node/node.*/hugepages
	// Read /sys/devices/system/node/node.*/meminfo
	var tmpReader *bufio.Reader
	var tmpBytes []byte
	var tmpErr error
	var tmpFile *os.File
	for id, node := range reader.systemMetricReader.NumaNodes {
		tmpBytes, _ = ioutil.ReadFile("/sys/devices/system/node/node" + strconv.Itoa(id) + "/hugepages/hugepages-1048576kB/nr_hugepages")
		nr1GHugepages, _ := strconv.Atoi(string(tmpBytes))

		tmpBytes, _ = ioutil.ReadFile("/sys/devices/system/node/node" + strconv.Itoa(id) + "/hugepages/hugepages-1048576kB/free_hugepages")
		free1GHugepages, _ := strconv.Atoi(string(tmpBytes))

		node.Total1GMemory = nr1GHugepages
		node.Used1GMemory = nr1GHugepages - free1GHugepages

		if tmpFile, tmpErr = os.Open("/sys/devices/system/node/node" + strconv.Itoa(id) + "/meminfo"); tmpErr != nil {
			continue
		}
		tmpReader = bufio.NewReader(tmpFile)

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

		if len(reader.memStats) > reader.cacheLength {
			reader.memStats = reader.memStats[1:]
		}

		memAvailable := memFree + inactive + kReclaimable + sReclaimable

		if len(reader.checkAvailableWarnNodeCounters) < id+1 {
			reader.checkAvailableWarnNodeCounters = append(
				reader.checkAvailableWarnNodeCounters, 0)
		}

		if memAvailable < int64(float64(node.TotalMemory-(nr1GHugepages*1000000))*reader.checkAvailableWarnAvailableRatio) {
			reader.checkAvailableWarnNodeCounters[id] += 1
		} else {
			reader.checkAvailableWarnNodeCounters[id] = 0
		}

		reader.memStats = append(reader.memStats, MemStat{
			ReportStatus: 0,
			Timestamp:    timestamp,
			NodeId:       id,
			MemTotal:     memTotal,
			MemFree:      memFree,
			MemUsed:      memUsed,
			MemAvailable: memAvailable,
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
}

func (reader *MemReader) ReportMetrics() (metrics []spec.ResourceMetric) {
	metrics = make([]spec.ResourceMetric, 0, len(reader.vmStats)+len(reader.memStats))
	for _, stat := range reader.vmStats {
		if stat.ReportStatus == ReportStatusReported {
			continue
		}
		metrics = append(metrics, spec.ResourceMetric{
			Name: "system_vmstat",
			Time: stat.Timestamp,
			Metric: map[string]interface{}{
				"pgscan_kswapd": stat.DiffPgscanKswapd,
				"pgscan_direct": stat.DiffPgscanDirect,
				"pgfault":       stat.DiffPgfault,
				"pswapin":       stat.DiffPswapin,
				"pswapout":      stat.DiffPswapout,
			},
		})
	}

	for _, stat := range reader.memStats {
		if stat.ReportStatus == ReportStatusReported {
			continue
		}

		reclaimable := (stat.Inactive + stat.KReclaimable + stat.SReclaimable) * 1000
		metrics = append(metrics, spec.ResourceMetric{
			Name: "system_mem",
			Time: stat.Timestamp,
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

	return
}

func (reader *MemReader) ReportEvents() (events []spec.ResourceEvent) {
	if len(reader.vmStats) == 0 {
		return
	}

	stats := reader.memStats[len(reader.memStats)-reader.lenNodes:]
	msgs := []string{}
	eventCheckAvailableLevel := consts.EventLevelSuccess
	for _, stat := range stats {
		if reader.checkAvailableWarnNodeCounters[stat.NodeId] > reader.checkAvailableOccurences {
			eventCheckAvailableLevel = consts.EventLevelWarning
		}
		msgs = append(msgs,
			fmt.Sprintf("node:%d,total=%d,free=%d,available=%d",
				stat.NodeId,
				stat.MemTotal,
				stat.MemFree,
				stat.MemAvailable))
	}

	events = append(events, spec.ResourceEvent{
		Name:            "CheckMemAvailable",
		Time:            stats[0].Timestamp,
		Level:           eventCheckAvailableLevel,
		Msg:             strings.Join(msgs, ", "),
		ReissueDuration: reader.checkAvailableReissueDuration,
	})

	stat := reader.vmStats[len(reader.vmStats)-1]
	eventCheckPgscanLevel := consts.EventLevelSuccess
	if reader.checkPgscanWarnPgscanDirectCounter > reader.checkPgscanOccurences {
		eventCheckPgscanLevel = consts.EventLevelWarning
	}
	events = append(events, spec.ResourceEvent{
		Name:  "CheckMemPgscan",
		Time:  stat.Timestamp,
		Level: eventCheckPgscanLevel,
		Msg: fmt.Sprintf("Pgscan kswapd=%d, direct=%d",
			stat.DiffPgscanKswapd,
			stat.DiffPgscanDirect,
		),
		ReissueDuration: reader.checkPgscanReissueDuration,
	})

	return
}

func (reader *MemReader) Reported() {
	for i := range reader.vmStats {
		reader.vmStats[i].ReportStatus = ReportStatusReported
	}

	for i := range reader.memStats {
		reader.memStats[i].ReportStatus = ReportStatusReported
	}
	return
}

func (reader *MemReader) readTmpVmStat(tctx *logger.TraceContext) (tmpVmStat *TmpVmStat) {
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
