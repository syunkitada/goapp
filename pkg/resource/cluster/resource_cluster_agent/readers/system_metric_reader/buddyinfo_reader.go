package system_metric_reader

import (
	"bufio"
	"os"
	"strconv"
	"time"

	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/lib/str_utils"
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

// ReadDiskStat read /proc/buddyinfo
//
// Output example is below.
// $ /proc/buddyinfo
//                           4K     8k    16k    32k    64k   128k   256k   512k     1M     2M     4M
// Node 0, zone      DMA      0      0      0      1      2      1      1      0      1      1      3
// Node 0, zone    DMA32      3      3      3      3      3      2      5      6      5      2    874
// Node 0, zone   Normal  24727  53842  18419  15120  10448   4451   1761    804    382    105    229
func (reader *SystemMetricReader) ReadBuddyinfoStat(tctx *logger.TraceContext) {
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