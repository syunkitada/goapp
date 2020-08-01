package system_metric_reader

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/lib/str_utils"
	"github.com/syunkitada/goapp/pkg/resource/config"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

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

type VmStatReader struct {
	conf            *config.ResourceMetricSystemConfig
	tmpDiskStatMap  map[string]TmpDiskStat
	diskStats       []DiskStat
	diskStatFilters []string
}

func (reader *SystemMetricReader) ReadVmStat(tctx *logger.TraceContext) {
	timestamp := time.Now()

	// Read /proc/vmstat
	if reader.tmpVmStat == nil {
		reader.tmpVmStat = reader.ReadTmpVmStat(tctx)
	} else {
		if len(reader.vmStats) > reader.cacheLength {
			reader.vmStats = reader.vmStats[1:]
		}

		tmpVmStat := reader.ReadTmpVmStat(tctx)

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
}

func (reader *SystemMetricReader) ReadTmpVmStat(tctx *logger.TraceContext) (tmpVmStat *TmpVmStat) {
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

func (reader *SystemMetricReader) GetVmStatMetrics() (metrics []spec.ResourceMetric) {
	metrics = make([]spec.ResourceMetric, len(reader.vmStats))
	for _, stat := range reader.vmStats {
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
	return
}
