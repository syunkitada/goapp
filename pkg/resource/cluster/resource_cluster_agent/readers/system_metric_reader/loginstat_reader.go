package system_metric_reader

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

type LoginStat struct {
	reportStatus int // 0, 1(GetReport), 2(Reported)
	users        []UserStat
	timestamp    time.Time
}

type UserStat struct {
	reportStatus int // 0, 1(GetReport), 2(Reported)
	user         string
	tty          string
	from         string
	login        string
	idle         string
	jcpu         string
	pcpu         string
	what         string
}

func (reader *SystemMetricReader) ReadLoginStat(tctx *logger.TraceContext) {
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

func (reader *SystemMetricReader) GetLoginStatMetrics() (metrics []spec.ResourceMetric) {
	metrics = make([]spec.ResourceMetric, len(reader.loginStats))
	for _, stat := range reader.loginStats {
		timestamp := strconv.FormatInt(stat.timestamp.UnixNano(), 10)
		metrics = append(metrics, spec.ResourceMetric{
			Name: "system_login",
			Time: timestamp,
			Tag:  map[string]string{},
			Metric: map[string]interface{}{
				"users": len(stat.users),
			},
		})
	}
	return
}
