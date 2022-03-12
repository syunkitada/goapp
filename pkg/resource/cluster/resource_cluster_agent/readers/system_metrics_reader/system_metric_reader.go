package system_metrics_reader

import (
	"bufio"
	"os"
	"strconv"

	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/lib/str_utils"
	"github.com/syunkitada/goapp/pkg/resource/config"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

type SystemMetricsReader struct {
	conf      *config.ResourceMetricsSystemConfig
	name      string
	NumaNodes []spec.NumaNodeSpec
	Cpus      []spec.NumaNodeCpuSpec

	NetDevStatMap map[string]NetDevStat

	subReaders []SubMetricsReader
}

func New(conf *config.ResourceMetricsSystemConfig) *SystemMetricsReader {
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

	reader := &SystemMetricsReader{
		conf:          conf,
		name:          "system",
		NumaNodes:     numaNodes,
		Cpus:          cpus,
		NetDevStatMap: map[string]NetDevStat{},
	}

	reader.subReaders = []SubMetricsReader{
		NewCpuReader(conf, cpus),
		NewProcReader(conf, reader),
		NewDiskReader(conf),
		NewDiskFsReader(conf),
		NewMemReader(conf, reader),
		NewMemBuddyinfoReader(conf),
		NewNetReader(conf),
		NewNetDevReader(conf, reader),
		NewUptimeReader(conf),
		NewLoginReader(conf),
	}

	return reader
}

const (
	ReportStatusReported = 2
)

type SubMetricsReader interface {
	Read(tctx *logger.TraceContext)
	ReportMetrics() []spec.ResourceMetric
	ReportEvents() []spec.ResourceEvent
	Reported()
}

func (reader *SystemMetricsReader) GetNumaNodes(tctx *logger.TraceContext) []spec.NumaNodeSpec {
	return reader.NumaNodes
}

func (reader *SystemMetricsReader) Read(tctx *logger.TraceContext) (err error) {
	for _, r := range reader.subReaders {
		r.Read(tctx)
	}
	return
}

func (reader *SystemMetricsReader) GetName() string {
	return reader.name
}

func (reader *SystemMetricsReader) Report() ([]spec.ResourceMetric, []spec.ResourceEvent) {
	metrics := make([]spec.ResourceMetric, 0, 1000)
	events := make([]spec.ResourceEvent, 0, 1000)

	for _, r := range reader.subReaders {
		metrics = append(metrics, r.ReportMetrics()...)
	}

	for _, r := range reader.subReaders {
		events = append(events, r.ReportEvents()...)
	}

	return metrics, events
}

func (reader *SystemMetricsReader) Reported() {
	for _, r := range reader.subReaders {
		r.Reported()
	}
}
