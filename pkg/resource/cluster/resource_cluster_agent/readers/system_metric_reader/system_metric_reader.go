package system_metric_reader

import (
	"bufio"
	"os"
	"strconv"

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
	loginStats        []LoginStat
	cpuStats          []CpuStat
	memStats          []MemStat
	diskStats         []DiskStat
	fsStats           []FsStat
	buddyinfoStats    []BuddyinfoStat

	tmpVmStat *TmpVmStat
	vmStats   []VmStat

	// network
	tmpNetDevStatMap map[string]TmpNetDevStat
	netDevStats      []NetDevStat
	tmpTcpExtStat    *TmpTcpExtStat
	tmpIpExtStat     *TmpIpExtStat
	tcpExtStats      []TcpExtStat
	ipExtStats       []IpExtStat

	procsStats []ProcsStat
	procStats  []ProcStat

	subReaders []SubMetricReader
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

	reader := &SystemMetricReader{
		conf:              conf,
		name:              "system",
		enableLogin:       conf.EnableLogin,
		enableCpu:         conf.EnableCpu,
		enableMemory:      conf.EnableMemory,
		enableProc:        conf.EnableProc,
		cacheLength:       conf.CacheLength,
		numaNodes:         numaNodes,
		cpus:              cpus,
		loginStats:        make([]LoginStat, 0, conf.CacheLength),
		cpuStats:          make([]CpuStat, 0, conf.CacheLength),
		memStats:          make([]MemStat, 0, conf.CacheLength),
		diskStats:         make([]DiskStat, 0, conf.CacheLength),
		buddyinfoStats:    make([]BuddyinfoStat, 0, conf.CacheLength),
		procsStats:        make([]ProcsStat, 0, conf.CacheLength),
		procStats:         make([]ProcStat, 0, conf.CacheLength),
		tcpExtStats:       make([]TcpExtStat, 0, conf.CacheLength),
		ipExtStats:        make([]IpExtStat, 0, conf.CacheLength),
		netDevStatFilters: []string{"lo"},
		fsStatTypes:       []string{"ext4"},
	}

	reader.subReaders = []SubMetricReader{
		NewUptimeMetricReader(conf),
		NewDiskMetricReader(conf),
	}

	return reader
}

type SubMetricReader interface {
	Read(tctx *logger.TraceContext)
	ReportMetrics() []spec.ResourceMetric
	ReportEvents() []spec.ResourceEvent
	Reported()
}

func (reader *SystemMetricReader) GetNumaNodes(tctx *logger.TraceContext) []spec.NumaNodeSpec {
	return reader.numaNodes
}

func (reader *SystemMetricReader) Read(tctx *logger.TraceContext) (err error) {
	reader.ReadLoginStat(tctx)

	reader.ReadCpuStat(tctx)

	if reader.enableProc {
		// Read /proc/
		reader.ReadProc(tctx)
	}

	reader.ReadMemStat(tctx)
	reader.ReadVmStat(tctx)
	reader.ReadBuddyinfoStat(tctx)

	reader.ReadFsStat(tctx)

	reader.ReadNetDevStat(tctx)
	reader.ReadNetStat(tctx)

	for _, r := range reader.subReaders {
		r.Read(tctx)
	}

	return
}

func (reader *SystemMetricReader) GetName() string {
	return reader.name
}

func (reader *SystemMetricReader) Report() ([]spec.ResourceMetric, []spec.ResourceEvent) {
	metrics := make([]spec.ResourceMetric, 0, 100)
	events := make([]spec.ResourceEvent, 0, 100)

	metrics = append(metrics, reader.GetLoginStatMetrics()...)
	metrics = append(metrics, reader.GetCpuStatMetrics()...)
	metrics = append(metrics, reader.GetProcStatMetrics()...)
	metrics = append(metrics, reader.GetMemStatMetrics()...)
	metrics = append(metrics, reader.GetVmStatMetrics()...)
	metrics = append(metrics, reader.GetFsStatMetrics()...)
	metrics = append(metrics, reader.GetNetDevStatMetrics()...)
	metrics = append(metrics, reader.GetNetStatMetrics()...)

	for _, r := range reader.subReaders {
		metrics = append(metrics, r.ReportMetrics()...)
	}

	// TODO check metrics and issue events

	return metrics, events
}

func (reader *SystemMetricReader) Reported() {
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

	for _, r := range reader.subReaders {
		r.Reported()
	}
}
