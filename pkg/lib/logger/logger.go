package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"golang.org/x/net/context"
	"google.golang.org/grpc/peer"

	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_grpc_pb"
	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_model"
	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/lib/codes"
	"github.com/syunkitada/goapp/pkg/lib/error_utils"
)

var (
	conf          *config.Config
	name          string
	Logger        *log.Logger
	stdoutLogger  *log.Logger
	LogTimeFormat string
)

const (
	debugLog   = "D"
	infoLog    = "I"
	warningLog = "W"
	errorLog   = "E"
	fatalLog   = "F"
	benchLog   = "B"
	traceLog   = "T"
)

type TraceContext struct {
	Host     string
	App      string
	Func     string
	TraceId  string
	Metadata map[string]string
}

type ActionTraceContext struct {
	TraceContext
	UserName        string
	RoleName        string
	ProjectName     string
	ProjectRoleName string
	ActionName      string
	ActionData      string
}

func NewTraceContext(host string, app string) *TraceContext {
	return &TraceContext{
		TraceId:  xid.New().String(),
		Host:     host,
		App:      app,
		Metadata: map[string]string{},
	}
}

func NewTraceContextWithTraceId(traceId string, host string, app string) *TraceContext {
	return &TraceContext{
		TraceId:  traceId,
		Host:     host,
		App:      app,
		Metadata: map[string]string{},
	}
}

func NewCtlTraceContext(app string) *TraceContext {
	return &TraceContext{
		TraceId: xid.New().String(),
		Host:    config.Conf.Default.Host,
		App:     app,
		Metadata: map[string]string{
			"Username": config.Conf.Ctl.Username,
			"Project":  config.Conf.Ctl.Project,
			"ApiUrl":   config.Conf.Ctl.ApiUrl,
		},
	}
}

func NewAuthproxyActionTraceContext(host string, app string, c *gin.Context) (*ActionTraceContext, error) {
	traceId, traceIdOk := c.Get("TraceId")
	username, usernameOk := c.Get("Username")
	userAuthority, userAuthorityOk := c.Get("UserAuthority")
	action, actionOk := c.Get("Action")

	if !traceIdOk || !usernameOk || !userAuthorityOk || !actionOk {
		return nil, error_utils.NewInvalidRequestError(map[string]bool{
			"TraceId":       traceIdOk,
			"Username":      usernameOk,
			"UserAuthority": userAuthorityOk,
			"Action":        actionOk,
		})
	}
	tmpAuthority := userAuthority.(*authproxy_model.UserAuthority)
	tmpAction := action.(authproxy_model.ActionRequest)
	return &ActionTraceContext{
		TraceContext: TraceContext{
			TraceId: traceId.(string),
			Host:    host,
			App:     app,
			Metadata: map[string]string{
				"AuthUser": username.(string),
			},
		},
		UserName:        username.(string),
		RoleName:        tmpAuthority.ActionProjectService.RoleName,
		ProjectName:     tmpAuthority.ActionProjectService.ProjectName,
		ProjectRoleName: tmpAuthority.ActionProjectService.ProjectRoleName,
		ActionName:      tmpAction.Name,
		ActionData:      tmpAction.Data,
	}, nil
}

func NewGrpcTraceContext(host string, app string, ctx context.Context) *TraceContext {
	var client string
	if pr, ok := peer.FromContext(ctx); ok {
		client = pr.Addr.String()
	} else {
		client = ""
	}

	return &TraceContext{
		TraceId: xid.New().String(),
		Host:    host,
		App:     app,
		Metadata: map[string]string{
			"Client": client,
		},
	}
}

func NewAuthproxyTraceContext(tctx *TraceContext, atctx *ActionTraceContext) *authproxy_grpc_pb.TraceContext {
	if tctx != nil {
		return &authproxy_grpc_pb.TraceContext{
			TraceId: tctx.TraceId,
			App:     tctx.App,
			Host:    tctx.Host,
		}
	}
	if atctx != nil {
		return &authproxy_grpc_pb.TraceContext{
			TraceId:         atctx.TraceId,
			ActionName:      atctx.ActionName,
			UserName:        atctx.UserName,
			RoleName:        atctx.RoleName,
			ProjectName:     atctx.ProjectName,
			ProjectRoleName: atctx.ProjectRoleName,
		}
	}
	return nil
}

func NewGrpcAuthproxyTraceContext(host string, app string, ctx context.Context, atctx *authproxy_grpc_pb.TraceContext) *TraceContext {
	var client string
	if pr, ok := peer.FromContext(ctx); ok {
		client = pr.Addr.String()
	} else {
		client = ""
	}

	tctx := &TraceContext{
		TraceId: atctx.TraceId,
		Host:    host,
		App:     app,
		Metadata: map[string]string{
			"Client":     client,
			"ClientHost": atctx.Host,
			"ClientApp":  atctx.App,
		},
	}

	if atctx.ActionName != "" {
		tctx.Metadata["ActionName"] = atctx.ActionName
	}
	if atctx.UserName != "" {
		tctx.Metadata["UserName"] = atctx.UserName
		tctx.Metadata["RoleName"] = atctx.RoleName
		tctx.Metadata["ProjectName"] = atctx.ProjectName
		tctx.Metadata["ProjectRoleName"] = atctx.ProjectRoleName
	}

	return tctx
}

func SetErrorTraceContext(tctx *authproxy_grpc_pb.TraceContext, statusCode int64, data interface{}) {
	tctx.StatusCode = statusCode
	tctx.Err = codes.GetMsg(statusCode, data)
}

func Init() {
	conf = &config.Conf

	stdoutLogger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	LogTimeFormat = conf.Default.LogTimeFormat

	name = os.Getenv("LOG_FILE")
	if name == "" {
		for _, arg := range os.Args {
			option := strings.Index(arg, "-")
			if option == 0 {
				continue
			}
			slash := strings.LastIndex(arg, "/")
			if slash > 0 {
				arg = arg[slash+1:]
			}
			name += "-" + arg
		}
		name = name[1:]
	}

	if conf.Default.LogDir == "stdout" {
		Logger = log.New(os.Stdout, "", 0)
	} else {
		logfilePath := filepath.Join(conf.Default.LogDir, name+".log")
		logfile, err := os.OpenFile(logfilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			StdoutFatalf("Failed open logfile: %v", logfile)
		} else {
			Logger = log.New(logfile, "", 0)
		}
	}
}

func timePrefix() string {
	return "Time=\"" + time.Now().Format(LogTimeFormat) + "\""
}

func convertTags(tctx *TraceContext) string {
	tags := " TraceId=\"" + tctx.TraceId + "\" Host=\"" + tctx.Host + "\" App=\"" + tctx.App + "\" Func=\"" + tctx.Func + "\""
	for k, v := range tctx.Metadata {
		tags += " " + k + "=\"" + v + "\""
	}
	return tags
}

func getFunc(depth int) string {
	var function string
	pc, file, line, ok := runtime.Caller(2 + depth)
	if !ok {
		file = ""
		line = 1
		function = ""
	} else {
		slash := strings.LastIndex(file, "/")
		if slash > 0 {
			file = file[slash+1:]
		}

		function = runtime.FuncForPC(pc).Name()
		dot := strings.LastIndex(function, ".")
		if dot > 0 {
			function = function[dot+1:]
		}
	}
	return fmt.Sprintf("%s:%d:%s", file, line, function)
}

func StdoutInfo(format string, args ...interface{}) {
	stdoutLogger.Print(infoLog + " " + fmt.Sprint(args...))
}

func StdoutInfof(format string, args ...interface{}) {
	stdoutLogger.Print(infoLog + " " + fmt.Sprintf(format, args...))
}

func StdoutFatal(args ...interface{}) {
	l := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	l.Print(fatalLog + " " + fmt.Sprint(args...))
	os.Exit(1)
}

func StdoutFatalf(format string, args ...interface{}) {
	l := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	l.Print(fatalLog + " " + fmt.Sprintf(format, args...))
	os.Exit(1)
}

func Info(tctx *TraceContext, args ...interface{}) {
	tctx.Func = getFunc(0)
	Logger.Print(timePrefix() + " Level=\"" + infoLog +
		"\" Msg=\"" + fmt.Sprint(args...) + "\"" + convertTags(tctx))
}

func Infof(tctx *TraceContext, format string, args ...interface{}) {
	tctx.Func = getFunc(0)
	Logger.Print(timePrefix() + " Level=\"" + infoLog +
		"\" Msg=\"" + fmt.Sprintf(format, args...) + "\"" + convertTags(tctx))
}

func Warning(tctx *TraceContext, err error, args ...interface{}) {
	tctx.Func = getFunc(0)
	Logger.Print(timePrefix() + " Level=\"" + warningLog +
		"\" Err=\"" + err.Error() + "\" Msg=\"" + fmt.Sprint(args...) + "\"" + convertTags(tctx))
}

func Warningf(tctx *TraceContext, err error, format string, args ...interface{}) {
	tctx.Func = getFunc(0)
	Logger.Print(timePrefix() + " Level=\"" + warningLog +
		"\" Err=\"" + err.Error() + "\" Msg=\"" + fmt.Sprintf(format, args...) + "\"" + convertTags(tctx))
}

func Error(tctx *TraceContext, err error, args ...interface{}) {
	tctx.Func = getFunc(0)
	Logger.Print(timePrefix() + " Level=\"" + errorLog +
		"\" Err=\"" + err.Error() + "\" Msg=\"" + fmt.Sprint(args...) + "\"" + convertTags(tctx))
}

func Errorf(tctx *TraceContext, err error, format string, args ...interface{}) {
	tctx.Func = getFunc(0)
	Logger.Print(timePrefix() + " Level=\"" + errorLog +
		"\" Err=\"" + err.Error() + "\" Msg=\"" + fmt.Sprintf(format, args...) + "\"" + convertTags(tctx))
}

func StartTrace(tctx *TraceContext) time.Time {
	startTime := time.Now()
	tctx.Func = getFunc(0)
	Info(tctx, "StartTrace")
	Logger.Print(timePrefix() + " Level=\"" + infoLog + "\" Msg=\"StartTrace\"" + convertTags(tctx))
	return startTime
}

func EndTrace(tctx *TraceContext, startTime time.Time, err error, depth int) {
	tctx.Func = getFunc(depth)
	tctx.Metadata["Latency"] = strconv.FormatInt(time.Now().Sub(startTime).Nanoseconds()/1000000, 10)
	if err != nil {
		Logger.Print(timePrefix() + " Level=\"" + errorLog + "\" Msg=\"EndTrace\" Err=\"" + err.Error() + "\"" + convertTags(tctx))
	} else {
		Logger.Print(timePrefix() + " Level=\"" + infoLog + "\" Msg=\"EndTrace\"" + convertTags(tctx))
	}
}

func EndGrpcTrace(tctx *TraceContext, startTime time.Time, statusCode int64, err string) {
	tctx.Func = getFunc(0)
	tctx.Metadata["Latency"] = strconv.FormatInt(time.Now().Sub(startTime).Nanoseconds()/1000000, 10)
	tctx.Metadata["StatusCode"] = strconv.FormatInt(statusCode, 10)
	if err != "" {
		Logger.Print(timePrefix() + " Level=\"" + errorLog + "\" Msg=\"EndTrace\"" + convertTags(tctx))
	} else {
		Logger.Print(timePrefix() + " Level=\"" + infoLog + "\" Msg=\"EndTrace\"" + convertTags(tctx))
	}
}
