package system_metric_reader

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"
	"time"

	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/config"
	"github.com/syunkitada/goapp/pkg/resource/consts"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

type FsStat struct {
	ReportStatus int
	Timestamp    time.Time
	Path         string
	MountPath    string
	TotalSize    int64
	FreeSize     int64
	UsedSize     int64
	Files        int64
}

type DiskFsReader struct {
	conf        *config.ResourceMetricSystemConfig
	cacheLength int
	fsStats     []FsStat
	fsStatTypes []string

	checkFreeWarnRatio       float64
	checkFreeCritRatio       float64
	checkFreeOccurences      int
	checkFreeReissueDuration int
	checkFreeWarnCounter     int
	checkFreeCritCounter     int
}

func NewDiskFsReader(conf *config.ResourceMetricSystemConfig) SubMetricReader {
	return &DiskFsReader{
		conf:        conf,
		cacheLength: conf.CacheLength,
		fsStatTypes: []string{"ext4"},

		checkFreeWarnRatio:       conf.DiskFs.CheckFree.WarnFreeRatio,
		checkFreeCritRatio:       conf.DiskFs.CheckFree.CritFreeRatio,
		checkFreeOccurences:      conf.DiskFs.CheckFree.Occurences,
		checkFreeReissueDuration: conf.DiskFs.CheckFree.ReissueDuration,
		checkFreeWarnCounter:     0,
		checkFreeCritCounter:     0,
	}
}

// Read filesystem stat
func (reader *DiskFsReader) Read(tctx *logger.TraceContext) {
	timestamp := time.Now()

	// read /proc/self/mounts
	// MEMO: /etc/mtab is symbolic link to /proc/self/mounts
	mountsFile, _ := os.Open("/proc/self/mounts")
	defer mountsFile.Close()
	tmpReader := bufio.NewReader(mountsFile)
	var splitedLine []string
	var tmpBytes []byte
	var tmpErr error
	var isMatch bool
	for {
		tmpBytes, _, tmpErr = tmpReader.ReadLine()
		if tmpErr != nil {
			break
		}
		splitedLine = strings.Split(string(tmpBytes), " ")
		isMatch = false
		for _, fsType := range reader.fsStatTypes {
			if splitedLine[2] == fsType {
				isMatch = true
				break
			}
		}
		if !isMatch {
			continue
		}
		var statfs syscall.Statfs_t
		if tmpErr = syscall.Statfs(splitedLine[1], &statfs); tmpErr != nil {
			continue
		}
		totalSize := int64(statfs.Blocks) * statfs.Bsize
		freeSize := int64(statfs.Bavail) * statfs.Bsize

		if len(reader.fsStats) > reader.cacheLength {
			reader.fsStats = reader.fsStats[1:]
		}

		if freeSize < int64(float64(totalSize)*reader.checkFreeWarnRatio) {

			reader.checkFreeWarnCounter += 1
		} else {
			reader.checkFreeWarnCounter = 0
		}

		if freeSize < int64(float64(totalSize)*reader.checkFreeCritRatio) {

			reader.checkFreeCritCounter += 1
		} else {
			reader.checkFreeCritCounter = 0
		}

		reader.fsStats = append(reader.fsStats, FsStat{
			Timestamp: timestamp,
			Path:      splitedLine[0],
			MountPath: splitedLine[1],
			TotalSize: totalSize,
			FreeSize:  freeSize,
			UsedSize:  totalSize - freeSize,
			Files:     int64(statfs.Files),
		})
	}
}

func (reader *DiskFsReader) ReportMetrics() (metrics []spec.ResourceMetric) {
	metrics = make([]spec.ResourceMetric, 0, len(reader.fsStats))
	for _, stat := range reader.fsStats {
		if stat.ReportStatus == ReportStatusReported {
			continue
		}

		metrics = append(metrics, spec.ResourceMetric{
			Name: "system_fsstat",
			Time: stat.Timestamp,
			Tag: map[string]string{
				"path":       stat.Path,
				"mount_path": stat.MountPath,
			},
			Metric: map[string]interface{}{
				"total_size": stat.TotalSize,
				"free_size":  stat.FreeSize,
				"used_size":  stat.UsedSize,
				"files":      stat.Files,
			},
		})
	}
	return
}

func (reader *DiskFsReader) ReportEvents() (events []spec.ResourceEvent) {
	if len(reader.fsStats) == 0 {
		return
	}

	stat := reader.fsStats[len(reader.fsStats)-1]
	eventCheckFreeLevel := consts.EventLevelSuccess
	if reader.checkFreeCritCounter > reader.checkFreeOccurences {
		eventCheckFreeLevel = consts.EventLevelCritical
	} else if reader.checkFreeWarnCounter > reader.checkFreeOccurences {
		eventCheckFreeLevel = consts.EventLevelWarning
	}
	events = append(events, spec.ResourceEvent{
		Name:  "CheckDiskFsFree",
		Time:  stat.Timestamp,
		Level: eventCheckFreeLevel,
		Msg: fmt.Sprintf("totalGb=%d, freeGb=%d",
			stat.TotalSize/1000000000,
			stat.FreeSize/1000000000,
		),
		ReissueDuration: reader.checkFreeReissueDuration,
	})

	return
}

func (reader *DiskFsReader) Reported() {
	for i := range reader.fsStats {
		reader.fsStats[i].ReportStatus = ReportStatusReported
	}
	return
}
