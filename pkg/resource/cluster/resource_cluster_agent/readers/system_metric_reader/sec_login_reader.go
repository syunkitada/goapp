package system_metric_reader

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/config"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

type LoginStat struct {
	ReportStatus int // 0, 1(GetReport), 2(Reported)
	Users        []UserStat
	Timestamp    time.Time
}

type UserStat struct {
	ReportStatus int // 0, 1(GetReport), 2(Reported)
	User         string
	Tty          string
	From         string
	Login        string
	Idle         string
	Jcpu         string
	Pcpu         string
	What         string
}

type SecLoginReader struct {
	conf        *config.ResourceMetricSystemConfig
	cacheLength int
	loginStats  []LoginStat
}

func NewSecLoginReader(conf *config.ResourceMetricSystemConfig) SubMetricReader {
	return &SecLoginReader{
		conf:        conf,
		cacheLength: conf.CacheLength,
		loginStats:  make([]LoginStat, 0, conf.CacheLength),
	}
}

func (reader *SecLoginReader) Read(tctx *logger.TraceContext) {
	timestamp := time.Now()

	// Don't read /var/run/utmp, because of this is binary
	// Read w -h
	// USER    TTY      FROM          LOGIN@   IDLE   JCPU   PCPU  WHAT
	// hoge    pts/8    192.168.1.1   09:34    2.00s  0.10s  0.00s tmux a
	out, err := exec.Command("w", "-h").Output()
	users := []UserStat{}
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range strings.Split(string(out), "\n") {
		l := strings.Split(line, " ")
		if len(l) < 8 {
			continue
		}
		users = append(users, UserStat{
			User:  l[0],
			Tty:   l[1],
			From:  l[2],
			Login: l[3],
			Idle:  l[4],
			Jcpu:  l[5],
			Pcpu:  l[6],
			What:  l[7],
		})
	}
	if len(reader.loginStats) > reader.cacheLength {
		reader.loginStats = reader.loginStats[1:]
	}
	reader.loginStats = append(reader.loginStats, LoginStat{
		Timestamp: timestamp,
		Users:     users,
	})
}

func (reader *SecLoginReader) ReportMetrics() (metrics []spec.ResourceMetric) {
	metrics = make([]spec.ResourceMetric, 0, len(reader.loginStats))
	for _, stat := range reader.loginStats {
		if stat.ReportStatus == ReportStatusReported {
			continue
		}
		userSet := map[string]bool{}
		for _, user := range stat.Users {
			userSet[user.User] = true
		}
		metrics = append(metrics, spec.ResourceMetric{
			Name: "system_login",
			Time: stat.Timestamp,
			Tag:  map[string]string{},
			Metric: map[string]interface{}{
				"users":      len(stat.Users),
				"uniq_users": len(userSet),
			},
		})
	}

	return
}

func (reader *SecLoginReader) ReportEvents() (events []spec.ResourceEvent) {
	return
}

func (reader *SecLoginReader) Reported() {
	for i := range reader.loginStats {
		reader.loginStats[i].ReportStatus = ReportStatusReported
	}
	return
}
