package system_metric_reader

import (
	"bufio"
	"os"
	"strings"
	"syscall"
	"time"

	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/config"
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

type FsStatReader struct {
	conf        *config.ResourceMetricSystemConfig
	cacheLength int
	fsStats     []FsStat
	fsStatTypes []string
}

func NewFsStatReader(conf *config.ResourceMetricSystemConfig) SubMetricReader {
	return &FsStatReader{
		conf:        conf,
		cacheLength: conf.CacheLength,
		fsStatTypes: []string{"ext4"},
	}
}

// Read filesystem stat
func (reader *FsStatReader) Read(tctx *logger.TraceContext) {
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

func (reader *FsStatReader) ReportMetrics() (metrics []spec.ResourceMetric) {
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

func (reader *FsStatReader) ReportEvents() (events []spec.ResourceEvent) {
	return
}

func (reader *FsStatReader) Reported() {
	for i := range reader.fsStats {
		reader.fsStats[i].ReportStatus = ReportStatusReported
	}
	return
}
