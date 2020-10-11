package system_metric_reader

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/config"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

type UptimeStat struct {
	ReportStatus int
	timestamp    time.Time
	uptime       int64
}

type UptimeReader struct {
	conf        *config.ResourceMetricSystemConfig
	uptimeStats []UptimeStat
	cacheLength int
}

func NewUptimeReader(conf *config.ResourceMetricSystemConfig) SubMetricReader {
	return &UptimeReader{
		conf:        conf,
		cacheLength: conf.CacheLength,
		uptimeStats: make([]UptimeStat, 0, conf.CacheLength),
	}
}

// Read read /proc/uptime.
//
// Output example is below.
// uptime(s)  idle(s)
// 2906.26 5507.43
func (reader *UptimeReader) Read(tctx *logger.TraceContext) {
	timestamp := time.Now()

	procUptime, _ := os.Open("/proc/uptime")
	defer procUptime.Close()
	tmpReader := bufio.NewReader(procUptime)
	tmpBytes, _, _ := tmpReader.ReadLine()
	uptimeWords := strings.Split(string(tmpBytes), " ")
	uptimeF, _ := strconv.ParseFloat(uptimeWords[0], 64)
	uptime := int64(uptimeF)
	uptimeStat := UptimeStat{
		ReportStatus: 0,
		timestamp:    timestamp,
		uptime:       uptime,
	}

	if len(reader.uptimeStats) > reader.cacheLength {
		reader.uptimeStats = reader.uptimeStats[1:]
	}

	reader.uptimeStats = append(reader.uptimeStats, uptimeStat)
}

func (reader *UptimeReader) ReportMetrics() (metrics []spec.ResourceMetric) {
	metrics = make([]spec.ResourceMetric, 0, len(reader.uptimeStats))
	for _, stat := range reader.uptimeStats {
		if stat.ReportStatus == ReportStatusReported {
			continue
		}

		metrics = append(metrics, spec.ResourceMetric{
			Name: "system_uptime",
			Time: stat.timestamp,
			Tag:  map[string]string{},
			Metric: map[string]interface{}{
				"uptime": stat.uptime,
			},
		})
	}
	return
}

func (reader *UptimeReader) ReportEvents() (events []spec.ResourceEvent) {
	return
}

func (reader *UptimeReader) Reported() {
	for i := range reader.uptimeStats {
		reader.uptimeStats[i].ReportStatus = ReportStatusReported
	}
	return
}
