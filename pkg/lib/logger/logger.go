package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/rs/xid"
	"golang.org/x/net/context"
	"google.golang.org/grpc/peer"

	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/config"
)

var (
	conf          *config.Config
	name          string
	Logger        *log.Logger
	stdoutLogger  *log.Logger
	LogTimeFormat string
)

const (
	debugLog   = "DEBUG"
	infoLog    = "INFO"
	warningLog = "WARNING"
	errorLog   = "ERROR"
	fatalLog   = "FATAL"
	benchLog   = "BENCH"
	traceLog   = "TRACE"
)

type TraceContext struct {
	mtx      *sync.Mutex
	code     uint8
	err      error
	Host     string
	App      string
	Func     string
	TraceId  string
	timeout  int
	Metadata map[string]string
}

func (tctx *TraceContext) SetTimeout(timeout int) {
	tctx.timeout = timeout
}

func (tctx *TraceContext) GetTimeout() int {
	return tctx.timeout
}

func (tctx *TraceContext) GetTraceId() string {
	return tctx.TraceId
}

func NewTraceContext(host string, app string) *TraceContext {
	return &TraceContext{
		mtx:      new(sync.Mutex),
		TraceId:  xid.New().String(),
		Host:     host,
		App:      app,
		Metadata: map[string]string{},
		timeout:  3,
		code:     0,
		err:      nil,
	}
}

func NewTraceContextWithTraceId(traceId string, host string, app string) *TraceContext {
	return &TraceContext{
		TraceId:  traceId,
		Host:     host,
		App:      app,
		Metadata: map[string]string{},
		timeout:  3,
		code:     0,
		err:      nil,
	}
}

func NewCtlTraceContext(app string) *TraceContext {
	return &TraceContext{
		TraceId: xid.New().String(),
		Host:    config.Conf.Default.Host,
		App:     app,
		timeout: 3,
		Metadata: map[string]string{
			"Username": config.Conf.Ctl.Username,
			"Project":  config.Conf.Ctl.Project,
			"ApiUrl":   config.Conf.Ctl.ApiUrl,
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

func InitLogger(baseConf *base_config.Config) {
	stdoutLogger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	LogTimeFormat = baseConf.LogTimeFormat

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

	if baseConf.LogDir == "stdout" {
		Logger = log.New(os.Stdout, "", 0)
	} else {
		logfilePath := filepath.Join(baseConf.LogDir, name+".log")
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

func Fatal(tctx *TraceContext, args ...interface{}) {
	tctx.Func = getFunc(0)
	Logger.Print(timePrefix() + " Level=\"" + fatalLog +
		"\" Msg=\"" + fmt.Sprint(args...) + "\"" + convertTags(tctx))
	os.Exit(1)
}

func Fatalf(tctx *TraceContext, format string, args ...interface{}) {
	tctx.Func = getFunc(0)
	Logger.Print(timePrefix() + " Level=\"" + fatalLog +
		"\" Msg=\"" + fmt.Sprintf(format, args...) + "\"" + convertTags(tctx))
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

func Warning(tctx *TraceContext, args ...interface{}) {
	tctx.Func = getFunc(0)
	Logger.Print(timePrefix() + " Level=\"" + warningLog +
		"\" Msg=\"" + fmt.Sprint(args...) + "\"" + convertTags(tctx))
}

func Warningf(tctx *TraceContext, format string, args ...interface{}) {
	tctx.Func = getFunc(0)
	Logger.Print(timePrefix() + " Level=\"" + warningLog +
		"\" Msg=\"" + fmt.Sprintf(format, args...) + "\"" + convertTags(tctx))
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
