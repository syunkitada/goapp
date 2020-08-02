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
	users        []UserStat
	timestamp    time.Time
}

type UserStat struct {
	ReportStatus int // 0, 1(GetReport), 2(Reported)
	user         string
	tty          string
	from         string
	login        string
	idle         string
	jcpu         string
	pcpu         string
	what         string
}

type LoginStatReader struct {
	conf        *config.ResourceMetricSystemConfig
	cacheLength int
	loginStats  []LoginStat
}

func NewLoginStatReader(conf *config.ResourceMetricSystemConfig) SubMetricReader {
	return &LoginStatReader{
		conf:        conf,
		cacheLength: conf.CacheLength,
		loginStats:  make([]LoginStat, 0, conf.CacheLength),
	}
}

func (reader *LoginStatReader) Read(tctx *logger.TraceContext) {
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
		if len(l) != 8 {
			continue
		}
		users = append(users, UserStat{
			user:  l[0],
			tty:   l[1],
			from:  l[2],
			login: l[3],
			idle:  l[4],
			jcpu:  l[5],
			pcpu:  l[6],
			what:  l[7],
		})
	}
	if len(reader.loginStats) > reader.cacheLength {
		reader.loginStats = reader.loginStats[1:]
	}
	reader.loginStats = append(reader.loginStats, LoginStat{
		timestamp: timestamp,
		users:     users,
	})
}

func (reader *LoginStatReader) ReportMetrics() (metrics []spec.ResourceMetric) {
	metrics = make([]spec.ResourceMetric, len(reader.loginStats))
	for _, stat := range reader.loginStats {
		metrics = append(metrics, spec.ResourceMetric{
			Name: "system_login",
			Time: stat.timestamp,
			Tag:  map[string]string{},
			Metric: map[string]interface{}{
				"users": len(stat.users),
			},
		})
	}
	return
}

func (reader *LoginStatReader) ReportEvents() (events []spec.ResourceEvent) {
	return
}

func (reader *LoginStatReader) Reported() {
	for i := range reader.loginStats {
		reader.loginStats[i].ReportStatus = 2
	}
	return
}
