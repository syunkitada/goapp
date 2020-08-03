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
	conf      *config.ResourceMetricSystemConfig
	name      string
	NumaNodes []spec.NumaNodeSpec
	Cpus      []spec.NumaNodeCpuSpec

	subReaders []SubMetricReader
}

func New(conf *config.ResourceMetricSystemConfig) *SystemMetricReader {
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
		conf:      conf,
		name:      "system",
		NumaNodes: numaNodes,
		Cpus:      cpus,
	}

	reader.subReaders = []SubMetricReader{
		NewUptimeMetricReader(conf),
		NewLoginStatReader(conf),
		NewCpuStatReader(conf, cpus),
		NewVmStatReader(conf),
		NewMemStatReader(conf, reader),
		NewBuddyinfoStatReader(conf),
		NewFsStatReader(conf),
		NewDiskMetricReader(conf),
		NewNetStatReader(conf),
		NewNetDevStatReader(conf),
	}

	return reader
}

const (
	ReportStatusReported = 2
)

type SubMetricReader interface {
	Read(tctx *logger.TraceContext)
	ReportMetrics() []spec.ResourceMetric
	ReportEvents() []spec.ResourceEvent
	Reported()
}

func (reader *SystemMetricReader) GetNumaNodes(tctx *logger.TraceContext) []spec.NumaNodeSpec {
	return reader.NumaNodes
}

func (reader *SystemMetricReader) Read(tctx *logger.TraceContext) (err error) {
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

	for _, r := range reader.subReaders {
		metrics = append(metrics, r.ReportMetrics()...)
	}

	// TODO check events

	return metrics, events
}

func (reader *SystemMetricReader) Reported() {
	for _, r := range reader.subReaders {
		r.Reported()
	}
}
