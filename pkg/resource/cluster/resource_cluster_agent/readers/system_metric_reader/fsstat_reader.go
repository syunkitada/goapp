package system_metric_reader

import (
	"bufio"
	"os"
	"strings"
	"syscall"
	"time"

	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

type FsStat struct {
	ReportStatus int
	Timestamp    time.Time
	TotalSize    int64
	FreeSize     int64
	UsedSize     int64
	Files        int64
}

// Read filesystem stat
func (reader *SystemMetricReader) ReadFsStat(tctx *logger.TraceContext) {
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

		reader.fsStats = append(reader.fsStats, FsStat{
			ReportStatus: 0,
			Timestamp:    timestamp,
			TotalSize:    totalSize,
			FreeSize:     freeSize,
			UsedSize:     totalSize - freeSize,
			Files:        int64(statfs.Files),
		})
	}
}

func (reader *SystemMetricReader) GetFsStatMetrics() (metrics []spec.ResourceMetric) {
	metrics = make([]spec.ResourceMetric, len(reader.fsStats))
	for _, stat := range reader.fsStats {
		metrics = append(metrics, spec.ResourceMetric{
			Name: "system_fsstat",
			Time: stat.Timestamp,
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
