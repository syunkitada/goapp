package system_metric_reader

import (
	"fmt"
	"io/ioutil"
	"os/user"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/config"
	"github.com/syunkitada/goapp/pkg/resource/consts"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

type LoginStat struct {
	ReportStatus int // 0, 1(GetReport), 2(Reported)
	Users        []UserStat
	Timestamp    time.Time
}

type UserStat struct {
	ReportStatus  int // 0, 1(GetReport), 2(Reported)
	User          string
	Tty           string
	From          string
	LoginDuration int64
	Idle          string
	Jcpu          string
	Pcpu          string
	What          string
}

type LoginReader struct {
	conf        *config.ResourceMetricSystemConfig
	cacheLength int
	loginStats  []LoginStat

	checkLoginOccurences      int
	checkLoginReissueDuration int
	warnLoginSec              int64
	critLoginSec              int64
}

func NewLoginReader(conf *config.ResourceMetricSystemConfig) SubMetricReader {
	return &LoginReader{
		conf:        conf,
		cacheLength: conf.CacheLength,
		loginStats:  make([]LoginStat, 0, conf.CacheLength),

		checkLoginOccurences:      conf.Login.CheckLogin.Occurences,
		checkLoginReissueDuration: conf.Login.CheckLogin.ReissueDuration,
		warnLoginSec:              conf.Login.CheckLogin.WarnLoginSec,
		critLoginSec:              conf.Login.CheckLogin.CritLoginSec,
	}
}

func (reader *LoginReader) Read(tctx *logger.TraceContext) {
	timestamp := time.Now()

	files, tmpErr := ioutil.ReadDir("/dev/pts")
	if tmpErr != nil {
		logger.Warningf(tctx, "Failed read /dev/pts")
		return
	}

	users := []UserStat{}
	for _, file := range files {
		stat, ok := file.Sys().(*syscall.Stat_t)
		if !ok {
			continue
		}

		var owner string
		uid := strconv.Itoa(int(stat.Uid))
		u, err := user.LookupId(uid)
		if err != nil {
			owner = uid
		} else {
			owner = u.Username
		}

		loginDuration := timestamp.Unix() - file.ModTime().Unix()

		users = append(users, UserStat{
			User:          owner,
			Tty:           file.Name(),
			LoginDuration: loginDuration,
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

func (reader *LoginReader) ReportMetrics() (metrics []spec.ResourceMetric) {
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

func (reader *LoginReader) ReportEvents() (events []spec.ResourceEvent) {
	stat := reader.loginStats[len(reader.loginStats)-1]
	eventCheckLoginLevel := consts.EventLevelSuccess

	var msgs []string
	for _, user := range stat.Users {
		if user.LoginDuration > reader.critLoginSec {
			msgs = append(msgs, fmt.Sprintf("CritUser=%s,Duration=%d", user.User, user.LoginDuration))
			eventCheckLoginLevel = consts.EventLevelCritical
		} else if eventCheckLoginLevel == consts.EventLevelSuccess && user.LoginDuration > reader.warnLoginSec {
			msgs = append(msgs, fmt.Sprintf("WarnUser=%s,Duration=%d", user.User, user.LoginDuration))
			eventCheckLoginLevel = consts.EventLevelWarning
		}
	}

	events = append(events, spec.ResourceEvent{
		Name:            "CheckLogin",
		Time:            stat.Timestamp,
		Level:           eventCheckLoginLevel,
		Msg:             fmt.Sprintf("Users: %d: %s", len(stat.Users), strings.Join(msgs, "\n")),
		ReissueDuration: reader.checkLoginReissueDuration,
	})

	return
}

func (reader *LoginReader) Reported() {
	for i := range reader.loginStats {
		reader.loginStats[i].ReportStatus = ReportStatusReported
	}
	return
}
