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
	ReportStatus int // 0, 1(GetReport), 2(Reported)
	timestamp    time.Time
	uptime       int64
}

type UptimeMetricReader struct {
	conf        *config.ResourceMetricSystemConfig
	uptimeStats []UptimeStat
	cacheLength int
}

func NewUptimeMetricReader(conf *config.ResourceMetricSystemConfig) SubMetricReader {
	return &UptimeMetricReader{
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
func (reader *UptimeMetricReader) Read(tctx *logger.TraceContext) {
	timestamp := time.Now()

	procUptime, _ := os.Open("/proc/uptime")
	defer procUptime.Close()
	tmpReader := bufio.NewReader(procUptime)
	tmpBytes, _, _ := tmpReader.ReadLine()
	uptimeWords := strings.Split(string(tmpBytes), " ")
	uptime, _ := strconv.ParseInt(uptimeWords[0], 10, 64)
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

func (reader *UptimeMetricReader) ReportMetrics() (metrics []spec.ResourceMetric) {
	metrics = make([]spec.ResourceMetric, len(reader.uptimeStats))
	for _, stat := range reader.uptimeStats {
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

func (reader *UptimeMetricReader) ReportEvents() (events []spec.ResourceEvent) {
	return
}

func (reader *UptimeMetricReader) Reported() {
	for _, stat := range reader.uptimeStats {
		stat.ReportStatus = 2
	}
	return
}
