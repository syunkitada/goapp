package system_metric_reader

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/lib/str_utils"
	"github.com/syunkitada/goapp/pkg/resource/config"
	"github.com/syunkitada/goapp/pkg/resource/consts"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

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

type MemBuddyinfoReader struct {
	conf           *config.ResourceMetricSystemConfig
	cacheLength    int
	buddyinfoStats []BuddyinfoStat

	checkPagesOccurences      int
	checkPagesReissueDuration int
	checkPagesWarnMinPages    int64
	checkPagesWarnCounter     int
}

func NewMemBuddyinfoReader(conf *config.ResourceMetricSystemConfig) SubMetricReader {
	return &MemBuddyinfoReader{
		conf:           conf,
		cacheLength:    conf.CacheLength,
		buddyinfoStats: make([]BuddyinfoStat, 0, conf.CacheLength),

		checkPagesOccurences:      conf.MemBuddyinfo.CheckPages.Occurences,
		checkPagesReissueDuration: conf.MemBuddyinfo.CheckPages.ReissueDuration,
		checkPagesWarnMinPages:    conf.MemBuddyinfo.CheckPages.WarnMinPages,
		checkPagesWarnCounter:     0,
	}
}

// ReadDiskStat read /proc/buddyinfo
//
// Output example is below.
// $ /proc/buddyinfo
//                           4K     8k    16k    32k    64k   128k   256k   512k     1M     2M     4M
// Node 0, zone      DMA      0      0      0      1      2      1      1      0      1      1      3
// Node 0, zone    DMA32      3      3      3      3      3      2      5      6      5      2    874
// Node 0, zone   Normal  24727  53842  18419  15120  10448   4451   1761    804    382    105    229
func (reader *MemBuddyinfoReader) Read(tctx *logger.TraceContext) {
	timestamp := time.Now()

	buddyinfoFile, _ := os.Open("/proc/buddyinfo")
	defer buddyinfoFile.Close()
	tmpReader := bufio.NewReader(buddyinfoFile)
	var tmpBytes []byte
	var tmpErr error
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

			if len(reader.buddyinfoStats) > reader.cacheLength {
				reader.buddyinfoStats = reader.buddyinfoStats[1:]
			}

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
}

func (reader *MemBuddyinfoReader) ReportMetrics() (metrics []spec.ResourceMetric) {
	metrics = make([]spec.ResourceMetric, 0, len(reader.buddyinfoStats))
	warnMinPages := reader.checkPagesWarnMinPages
	for _, stat := range reader.buddyinfoStats {
		if stat.ReportStatus == ReportStatusReported {
			continue
		}

		if stat.M4K < warnMinPages {
			reader.checkPagesWarnCounter = 1
		} else if stat.M16K < warnMinPages {
			reader.checkPagesWarnCounter = 1
		} else if stat.M32K < warnMinPages {
			reader.checkPagesWarnCounter = 1
		} else if stat.M64K < warnMinPages {
			reader.checkPagesWarnCounter = 1
		} else if stat.M128K < warnMinPages {
			reader.checkPagesWarnCounter = 1
		} else if stat.M256K < warnMinPages {
			reader.checkPagesWarnCounter = 1
		} else if stat.M512K < warnMinPages {
			reader.checkPagesWarnCounter = 1
		} else if stat.M1M < warnMinPages {
			reader.checkPagesWarnCounter = 1
		} else if stat.M2M < warnMinPages {
			reader.checkPagesWarnCounter = 1
		} else if stat.M4M < warnMinPages {
			reader.checkPagesWarnCounter = 1
		} else {
			reader.checkPagesWarnCounter = 0
		}

		metrics = append(metrics, spec.ResourceMetric{
			Name: "system_buddyinfostat",
			Time: stat.Timestamp,
			Tag: map[string]string{
				"node": strconv.FormatInt(stat.NodeId, 10),
			},
			Metric: map[string]interface{}{
				"4K":   stat.M4K,
				"8K":   stat.M8K,
				"16K":  stat.M16K,
				"32K":  stat.M32K,
				"64K":  stat.M64K,
				"128K": stat.M128K,
				"256K": stat.M256K,
				"512K": stat.M512K,
				"1M":   stat.M1M,
				"2M":   stat.M2M,
				"4M":   stat.M4M,
			},
		})
	}
	return
}

func (reader *MemBuddyinfoReader) ReportEvents() (events []spec.ResourceEvent) {
	if len(reader.buddyinfoStats) == 0 {
		return
	}

	stat := reader.buddyinfoStats[len(reader.buddyinfoStats)-1]
	eventBuddyinfoPagesLevel := consts.EventLevelSuccess
	if reader.checkPagesWarnCounter > reader.checkPagesOccurences {
		eventBuddyinfoPagesLevel = consts.EventLevelWarning
	}

	events = append(events, spec.ResourceEvent{
		Name:  "CheckMemBuddyinfoPages",
		Time:  stat.Timestamp,
		Level: eventBuddyinfoPagesLevel,
		Msg: fmt.Sprintf("Buddyinfo 4K=%d 8K=%d 16K=%d 32K=%d 64K=%d 128K=%d 256K=%d 512K=%d 1M=%d 2M=%d 4M=%d",
			stat.M4K,
			stat.M8K,
			stat.M16K,
			stat.M32K,
			stat.M64K,
			stat.M128K,
			stat.M256K,
			stat.M512K,
			stat.M1M,
			stat.M2M,
			stat.M4M,
		),
		ReissueDuration: reader.checkPagesReissueDuration,
	})

	return
}

func (reader *MemBuddyinfoReader) Reported() {
	for i := range reader.buddyinfoStats {
		reader.buddyinfoStats[i].ReportStatus = ReportStatusReported
	}
	return
}
