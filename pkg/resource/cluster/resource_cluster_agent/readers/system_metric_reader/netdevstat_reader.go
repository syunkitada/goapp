package system_metric_reader

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/lib/str_utils"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

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

// ReadNetDevStat read /proc/diskstat
func (reader *SystemMetricReader) ReadNetDevStat(tctx *logger.TraceContext) {
	timestamp := time.Now()

	if reader.tmpNetDevStatMap == nil {
		reader.tmpNetDevStatMap = reader.ReadTmpNetDevStat(tctx)
	} else {
		tmpNetDevStatMap := reader.ReadTmpNetDevStat(tctx)
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
}

func (reader *SystemMetricReader) ReadTmpNetDevStat(tctx *logger.TraceContext) (tmpNetDevStatMap map[string]TmpNetDevStat) {
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

func (reader *SystemMetricReader) GetNetDevStatMetrics() (metrics []spec.ResourceMetric) {
	metrics = make([]spec.ResourceMetric, len(reader.procsStats)+len(reader.memStats))
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
	return
}
