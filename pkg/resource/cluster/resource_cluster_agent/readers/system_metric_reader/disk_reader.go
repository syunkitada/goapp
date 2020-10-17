package system_metric_reader

import (
	"bufio"
	"fmt"
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

type DiskStat struct {
	ReportStatus        int
	Timestamp           time.Time
	Device              string
	ReadsPerSec         int64
	RmergesPerSec       int64
	ReadBytesPerSec     int64
	ReadMsPerSec        int64
	WritesPerSec        int64
	WmergesPerSec       int64
	WriteBytesPerSec    int64
	WriteMsPerSec       int64
	DiscardsPerSec      int64
	DmergesPerSec       int64
	DiscardBytesPerSec  int64
	DiscardMsPerSec     int64
	ProgressIos         int64
	IosMsPerSec         int64
	WeightedIosMsPerSec int64
}

type TmpDiskStat struct {
	Timestamp         time.Time
	PblockSize        int64
	ReadsCompleted    int64
	ReadsMerges       int64
	ReadSectors       int64
	ReadMs            int64
	WritesCompleted   int64
	WritesMerges      int64
	WriteSectors      int64
	WriteMs           int64
	ProgressIos       int64
	IosMs             int64
	WeightedIosMs     int64
	DiscardsCompleted int64
	DiscardsMerges    int64
	DiscardSectors    int64
	DiscardMs         int64
}

type DiskReader struct {
	conf            *config.ResourceMetricSystemConfig
	cacheLength     int
	tmpDiskStatMap  map[string]TmpDiskStat
	diskStats       []DiskStat
	diskStatFilters []string

	checkIoDelayOccurences           int
	checkIoDelayReissueDuration      int
	checkReadMsPerSecWarnCounterMap  map[string]int
	checkReadMsPerSecCritCounterMap  map[string]int
	checkWriteMsPerSecWarnCounterMap map[string]int
	checkWriteMsPerSecCritCounterMap map[string]int
	checkProgressIosWarnCounterMap   map[string]int
	checkProgressIosCritCounterMap   map[string]int
	checkCritReadMsPerSec            int64
	checkWarnReadMsPerSec            int64
	checkCritWriteMsPerSec           int64
	checkWarnWriteMsPerSec           int64
	checkCritProgressIos             int64
	checkWarnProgressIos             int64
}

func NewDiskReader(conf *config.ResourceMetricSystemConfig) SubMetricReader {
	return &DiskReader{
		conf:            conf,
		cacheLength:     conf.CacheLength,
		diskStatFilters: []string{"loop"},
		tmpDiskStatMap:  map[string]TmpDiskStat{},

		checkIoDelayOccurences:           conf.Disk.CheckIoDelay.Occurences,
		checkIoDelayReissueDuration:      conf.Disk.CheckIoDelay.ReissueDuration,
		checkReadMsPerSecWarnCounterMap:  map[string]int{},
		checkReadMsPerSecCritCounterMap:  map[string]int{},
		checkWriteMsPerSecWarnCounterMap: map[string]int{},
		checkWriteMsPerSecCritCounterMap: map[string]int{},
		checkProgressIosWarnCounterMap:   map[string]int{},
		checkProgressIosCritCounterMap:   map[string]int{},
		checkCritReadMsPerSec:            conf.Disk.CheckIoDelay.CritReadMsPerSec,
		checkWarnReadMsPerSec:            conf.Disk.CheckIoDelay.WarnReadMsPerSec,
		checkCritWriteMsPerSec:           conf.Disk.CheckIoDelay.CritWriteMsPerSec,
		checkWarnWriteMsPerSec:           conf.Disk.CheckIoDelay.WarnWriteMsPerSec,
		checkCritProgressIos:             conf.Disk.CheckIoDelay.CritProgressIos,
		checkWarnProgressIos:             conf.Disk.CheckIoDelay.WarnProgressIos,
	}
}

func (reader *DiskReader) Read(tctx *logger.TraceContext) {
	timestamp := time.Now()

	if reader.tmpDiskStatMap == nil {
		reader.tmpDiskStatMap = reader.readTmpDiskStat(tctx)
	} else {
		tmpDiskStatMap := reader.readTmpDiskStat(tctx)
		for dev, cstat := range tmpDiskStatMap {
			bstat, ok := reader.tmpDiskStatMap[dev]
			if !ok {
				continue
			}
			interval := cstat.Timestamp.Unix() - bstat.Timestamp.Unix()
			readsPerSec := int64((cstat.ReadsCompleted - bstat.ReadsCompleted) / int64(interval))
			rmergesPerSec := int64((cstat.ReadsMerges - bstat.ReadsMerges) / int64(interval))
			readBytesPerSec := int64(((cstat.ReadSectors - bstat.ReadSectors) * cstat.PblockSize) / int64(interval))
			readMsPerSec := int64((cstat.ReadMs - bstat.ReadMs) / int64(interval))

			writesPerSec := int64((cstat.WritesCompleted - bstat.WritesCompleted) / int64(interval))
			wmergesPerSec := int64((cstat.WritesMerges - bstat.WritesMerges) / int64(interval))
			writeBytesPerSec := int64(((cstat.WriteSectors - bstat.WriteSectors) * cstat.PblockSize) / int64(interval))
			writeMsPerSec := int64((cstat.WriteMs - bstat.WriteMs) / int64(interval))

			discardsPerSec := int64((cstat.DiscardsCompleted - bstat.DiscardsCompleted) / int64(interval))
			dmergesPerSec := int64((cstat.DiscardsMerges - bstat.DiscardsMerges) / int64(interval))
			discardBytesPerSec := int64(((cstat.DiscardSectors - bstat.DiscardSectors) * cstat.PblockSize) / int64(interval))
			discardMsPerSec := int64((cstat.DiscardMs - bstat.DiscardMs) / int64(interval))

			iosMsPerSec := int64((cstat.IosMs - bstat.IosMs) / int64(interval))
			weightedIosMsPerSec := int64((cstat.WeightedIosMs - bstat.WeightedIosMs) / int64(interval))

			if len(reader.diskStats) > reader.cacheLength {
				reader.diskStats = reader.diskStats[1:]
			}

			if readMsPerSec > reader.checkCritReadMsPerSec {
				reader.checkReadMsPerSecCritCounterMap[dev] += 1
			} else if readMsPerSec > reader.checkWarnReadMsPerSec {
				reader.checkReadMsPerSecWarnCounterMap[dev] += 1
			} else {
				reader.checkReadMsPerSecCritCounterMap[dev] = 0
				reader.checkReadMsPerSecWarnCounterMap[dev] = 0
			}

			if writeMsPerSec > reader.checkCritWriteMsPerSec {
				reader.checkWriteMsPerSecCritCounterMap[dev] += 1
			} else if readMsPerSec > reader.checkWarnWriteMsPerSec {
				reader.checkWriteMsPerSecWarnCounterMap[dev] += 1
			} else {
				reader.checkWriteMsPerSecCritCounterMap[dev] = 0
				reader.checkWriteMsPerSecWarnCounterMap[dev] = 0
			}

			if cstat.ProgressIos > reader.checkCritProgressIos {
				reader.checkProgressIosCritCounterMap[dev] += 1
			} else if cstat.ProgressIos > reader.checkWarnProgressIos {
				reader.checkProgressIosWarnCounterMap[dev] += 1
			} else {
				reader.checkProgressIosCritCounterMap[dev] = 0
				reader.checkProgressIosWarnCounterMap[dev] = 0
			}

			reader.diskStats = append(reader.diskStats, DiskStat{
				ReportStatus:        0,
				Timestamp:           timestamp,
				Device:              dev,
				ReadsPerSec:         readsPerSec,
				RmergesPerSec:       rmergesPerSec,
				ReadBytesPerSec:     readBytesPerSec,
				ReadMsPerSec:        readMsPerSec,
				WritesPerSec:        writesPerSec,
				WmergesPerSec:       wmergesPerSec,
				WriteBytesPerSec:    writeBytesPerSec,
				WriteMsPerSec:       writeMsPerSec,
				DiscardsPerSec:      discardsPerSec,
				DmergesPerSec:       dmergesPerSec,
				DiscardBytesPerSec:  discardBytesPerSec,
				DiscardMsPerSec:     discardMsPerSec,
				ProgressIos:         cstat.ProgressIos,
				IosMsPerSec:         iosMsPerSec,
				WeightedIosMsPerSec: weightedIosMsPerSec,
			})
		}

		reader.tmpDiskStatMap = tmpDiskStatMap
	}
}

func (reader *DiskReader) ReportMetrics() (metrics []spec.ResourceMetric) {
	metrics = make([]spec.ResourceMetric, 0, len(reader.diskStats))
	for _, stat := range reader.diskStats {
		if stat.ReportStatus == ReportStatusReported {
			continue
		}
		metrics = append(metrics, spec.ResourceMetric{
			Name: "system_diskstat",
			Time: stat.Timestamp,
			Tag: map[string]string{
				"dev": stat.Device,
			},
			Metric: map[string]interface{}{
				"reads_per_sec":       stat.ReadsPerSec,
				"read_bytes_per_sec":  stat.ReadBytesPerSec,
				"read_ms_per_sec":     stat.ReadMsPerSec,
				"writes_per_sec":      stat.WritesPerSec,
				"write_bytes_per_sec": stat.WriteBytesPerSec,
				"write_ms_per_sec":    stat.WriteMsPerSec,
				"progress_ios":        stat.ProgressIos,
			},
		})
	}
	return
}

func (reader *DiskReader) ReportEvents() (events []spec.ResourceEvent) {
	if len(reader.diskStats) == 0 {
		return
	}

	stats := reader.diskStats[len(reader.diskStats)-len(reader.tmpDiskStatMap):]
	checkDiskIoDelayMsgs := []string{}
	eventCheckDiskIoDelayLevel := consts.EventLevelSuccess
	for _, stat := range stats {
		checkReadMsPerSecCritCounter, ok := reader.checkReadMsPerSecCritCounterMap[stat.Device]
		if !ok {
			continue
		}

		checkWriteMsPerSecCritCounter, ok := reader.checkWriteMsPerSecCritCounterMap[stat.Device]
		if !ok {
			continue
		}

		checkProgressIosCritCounter, ok := reader.checkProgressIosCritCounterMap[stat.Device]
		if !ok {
			continue
		}

		if checkReadMsPerSecCritCounter > reader.checkIoDelayOccurences {
			eventCheckDiskIoDelayLevel = consts.EventLevelCritical
		} else if checkWriteMsPerSecCritCounter > reader.checkIoDelayOccurences {
			eventCheckDiskIoDelayLevel = consts.EventLevelCritical
		} else if checkProgressIosCritCounter > reader.checkIoDelayOccurences {
			eventCheckDiskIoDelayLevel = consts.EventLevelCritical
		}

		if eventCheckDiskIoDelayLevel == consts.EventLevelSuccess {
			checkReadMsPerSecWarnCounter, ok := reader.checkReadMsPerSecWarnCounterMap[stat.Device]
			if !ok {
				continue
			}

			checkWriteMsPerSecWarnCounter, ok := reader.checkWriteMsPerSecWarnCounterMap[stat.Device]
			if !ok {
				continue
			}

			checkProgressIosWarnCounter, ok := reader.checkProgressIosWarnCounterMap[stat.Device]
			if !ok {
				continue
			}

			if checkReadMsPerSecWarnCounter > reader.checkIoDelayOccurences {
				eventCheckDiskIoDelayLevel = consts.EventLevelWarning
			} else if checkWriteMsPerSecWarnCounter > reader.checkIoDelayOccurences {
				eventCheckDiskIoDelayLevel = consts.EventLevelWarning
			} else if checkProgressIosWarnCounter > reader.checkIoDelayOccurences {
				eventCheckDiskIoDelayLevel = consts.EventLevelWarning
			}

		}

		checkDiskIoDelayMsgs = append(checkDiskIoDelayMsgs,
			fmt.Sprintf("dev:%s,rbytes=%d,rms=%d,wbytes=%d,wms=%d,ios=%d",
				stat.Device,
				stat.ReadBytesPerSec,
				stat.ReadMsPerSec,
				stat.WriteBytesPerSec,
				stat.WriteMsPerSec,
				stat.ProgressIos,
			))
	}

	events = append(events, spec.ResourceEvent{
		Name:            "CheckDiskIoDelay",
		Time:            stats[0].Timestamp,
		Level:           eventCheckDiskIoDelayLevel,
		Msg:             strings.Join(checkDiskIoDelayMsgs, ", "),
		ReissueDuration: reader.checkIoDelayReissueDuration,
	})

	return
}

func (reader *DiskReader) Reported() {
	for i := range reader.diskStats {
		reader.diskStats[i].ReportStatus = ReportStatusReported
	}
	return
}

func (reader *DiskReader) readTmpDiskStat(tctx *logger.TraceContext) (tmpDiskStatMap map[string]TmpDiskStat) {
	// Read /proc/diskstats

	// 259       0 nvme0n1 94360 70783 6403078 67950 136558 90723 6419592 38105 0 97140 59208 0 0 0 0
	// 259       0 nvme0n1 94360 70783 6403078 67950 136611 90751 6423880 38111 0 97200 59208 0 0 0 0
	// 259       0 nvme0n1 94364 70783 6403230 67951 155638 101247 7087392 41420 0 107356 59208 0 0 0 0

	// Field  1 -- # of reads completed
	// Field  2 -- # of reads merged, field 6 -- # of writes merged
	// Field  3 -- # of sectors read
	// Field  4 -- # of milliseconds spent reading
	// Field  5 -- # of writes completed
	// Field  6 -- # of writes merged
	// Field  7 -- # of sectors written
	// Field  8 -- # of milliseconds spent writing
	// Field  9 -- # of I/Os currently in progress
	// Field 10 -- # of milliseconds spent doing I/Os
	// Field 11 -- weighted # of milliseconds spent doing I/Os
	// Field 12 -- # of discards completed
	// Field 13 -- # of discards merged
	// Field 14 -- # of sectors discarded
	// Field 15 -- # of milliseconds spent discarding

	timestamp := time.Now()
	f, _ := os.Open("/proc/diskstats")
	defer f.Close()
	tmpReader := bufio.NewReader(f)
	tmpDiskStatMap = map[string]TmpDiskStat{}
	var isFiltered bool
	for {
		tmpBytes, _, tmpErr := tmpReader.ReadLine()
		if tmpErr != nil {
			break
		}
		columns := str_utils.SplitSpace(string(tmpBytes))
		isFiltered = false
		for _, filter := range reader.diskStatFilters {
			if strings.Index(columns[2], filter) > -1 {
				isFiltered = true
				break
			}
		}
		if isFiltered {
			continue
		}

		pblockSizeFile, tmpErr := os.Open("/sys/block/" + columns[2] + "/queue/physical_block_size")
		if tmpErr != nil {
			continue
		}
		pblockSizeReader := bufio.NewReader(pblockSizeFile)
		pblockSizeBytes, _, tmpErr := pblockSizeReader.ReadLine()
		pblockSizeFile.Close()
		if tmpErr != nil {
			continue
		}
		pblockSize, _ := strconv.ParseInt(string(pblockSizeBytes), 10, 64)

		readsCompleted, _ := strconv.ParseInt(columns[3], 10, 64)
		readsMerges, _ := strconv.ParseInt(columns[4], 10, 64)
		readSectors, _ := strconv.ParseInt(columns[5], 10, 64)
		readMs, _ := strconv.ParseInt(columns[6], 10, 64)
		writesCompleted, _ := strconv.ParseInt(columns[7], 10, 64)
		writesMerges, _ := strconv.ParseInt(columns[8], 10, 64)
		writeSectors, _ := strconv.ParseInt(columns[9], 10, 64)
		writeMs, _ := strconv.ParseInt(columns[10], 10, 64)

		progressIos, _ := strconv.ParseInt(columns[11], 10, 64)
		iosMs, _ := strconv.ParseInt(columns[12], 10, 64)
		weightedIosMs, _ := strconv.ParseInt(columns[13], 10, 64)

		discardsCompleted, _ := strconv.ParseInt(columns[14], 10, 64)
		discardsMerges, _ := strconv.ParseInt(columns[15], 10, 64)
		discardSectors, _ := strconv.ParseInt(columns[16], 10, 64)
		discardMs, _ := strconv.ParseInt(columns[17], 10, 64)

		tmpDiskStatMap[columns[2]] = TmpDiskStat{
			Timestamp:         timestamp,
			PblockSize:        pblockSize,
			ReadsCompleted:    readsCompleted,
			ReadsMerges:       readsMerges,
			ReadSectors:       readSectors,
			ReadMs:            readMs,
			WritesCompleted:   writesCompleted,
			WritesMerges:      writesMerges,
			WriteSectors:      writeSectors,
			WriteMs:           writeMs,
			ProgressIos:       progressIos,
			IosMs:             iosMs,
			WeightedIosMs:     weightedIosMs,
			DiscardsCompleted: discardsCompleted,
			DiscardsMerges:    discardsMerges,
			DiscardSectors:    discardSectors,
			DiscardMs:         discardMs,
		}
	}

	return
}
