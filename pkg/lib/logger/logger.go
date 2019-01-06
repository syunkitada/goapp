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

	"github.com/rs/xid"
	"golang.org/x/net/context"
	"google.golang.org/grpc/peer"

	"github.com/syunkitada/goapp/pkg/config"
)

var (
	conf          *config.Config
	name          string
	Logger        *log.Logger
	stdoutLogger  *log.Logger
	logTimeFormat string
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

func NewTraceContext(host string, app string) *TraceContext {
	return &TraceContext{
		TraceId:  xid.New().String(),
		Host:     host,
		App:      app,
		Metadata: map[string]string{},
	}
}

func NewCtlTraceContext(app string) *TraceContext {
	return &TraceContext{
		TraceId: xid.New().String(),
		Host:    conf.Default.Host,
		App:     app,
		Metadata: map[string]string{
			"Username": conf.Ctl.Username,
			"Project":  conf.Ctl.Project,
			"ApiUrl":   conf.Ctl.ApiUrl,
		},
	}
}

func NewAuthproxyActionTraceContext(host string, app string, traceId string, user string) *TraceContext {
	return &TraceContext{
		TraceId: xid.New().String(),
		Host:    host,
		App:     app,
		Metadata: map[string]string{
			"AuthUser": user,
		},
	}
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

func Init() {
	conf = &config.Conf

	stdoutLogger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	logTimeFormat = conf.Default.LogTimeFormat

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
	return "Time=\"" + time.Now().Format(logTimeFormat) + "\""
}

func convertTags(tctx *TraceContext) string {
	tags := " TraceId=\"" + tctx.TraceId + " Host=\"" + tctx.Host + "\" App=\"" + tctx.App + "\" Func=\"" + tctx.Func + "\""
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
	stdoutLogger.Print(fatalLog + " " + fmt.Sprint(args...))
	os.Exit(1)
}

func StdoutFatalf(format string, args ...interface{}) {
	stdoutLogger.Print(fatalLog + " " + fmt.Sprintf(format, args...))
	os.Exit(1)
}

func Info(ctx *TraceContext, args ...interface{}) {
	Logger.Print(timePrefix() + " Level=\"" + infoLog +
		"\" Msg=\"" + fmt.Sprint(args...) + "\"" + convertTags(ctx))
}

func Infof(ctx *TraceContext, format string, args ...interface{}) {
	Logger.Print(timePrefix() + " Level=\"" + infoLog +
		"\" Msg=\"" + fmt.Sprintf(format, args...) + "\"" + convertTags(ctx))
}

func Warning(ctx *TraceContext, err error, args ...interface{}) {
	Logger.Print(timePrefix() + " Level=\"" + warningLog +
		"\" Err=\"" + err.Error() + "\" Msg=\"" + fmt.Sprint(args...) + "\"" + convertTags(ctx))
}

func Warningf(ctx *TraceContext, err error, format string, args ...interface{}) {
	Logger.Print(timePrefix() + " Level=\"" + warningLog +
		"\" Err=\"" + err.Error() + "\" Msg=\"" + fmt.Sprintf(format, args...) + "\"" + convertTags(ctx))
}

func Error(ctx *TraceContext, err error, args ...interface{}) {
	Logger.Print(timePrefix() + " Level=\"" + errorLog +
		"\" Err=\"" + err.Error() + "\" Msg=\"" + fmt.Sprint(args...) + "\"" + convertTags(ctx))
}

func Errorf(ctx *TraceContext, err error, format string, args ...interface{}) {
	Logger.Print(timePrefix() + " Level=\"" + errorLog +
		"\" Err=\"" + err.Error() + "\" Msg=\"" + fmt.Sprintf(format, args...) + "\"" + convertTags(ctx))
}

func StartTrace(tctx *TraceContext) time.Time {
	startTime := time.Now()
	tctx.Func = getFunc(0)
	Info(tctx, "StartTrace")
	return startTime
}

func EndTrace(tctx *TraceContext, startTime time.Time, err error) {
	tctx.Func = getFunc(0)
	tctx.Metadata["Latency"] = strconv.FormatInt(time.Now().Sub(startTime).Nanoseconds()/1000000, 10)
	if err != nil {
		Error(tctx, err, "EndTrace")
	} else {
		Info(tctx, "EndTrace")
	}
}

func EndGrpcTrace(tctx *TraceContext, startTime time.Time, statusCode int64, err string) {
	tctx.Func = getFunc(0)
	tctx.Metadata["Latency"] = strconv.FormatInt(time.Now().Sub(startTime).Nanoseconds()/1000000, 10)
	tctx.Metadata["StatusCode"] = strconv.FormatInt(statusCode, 10)
	if err != "" {
		Error(tctx, fmt.Errorf(err), "EndTrace")
	} else {
		Info(tctx, "EndTrace")
	}
}
