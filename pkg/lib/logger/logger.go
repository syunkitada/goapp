package logger

import (
	"encoding/json"
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
	name   string
	Logger *log.Logger
	Tracer *log.Logger
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

func Init() {
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

	conf := &config.Conf
	if conf.Default.LogDir == "stdout" {
		Logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
		Tracer = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	} else {
		logfilePath := filepath.Join(conf.Default.LogDir, name+".log")
		logfile, err := os.OpenFile(logfilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			FatalStdoutf("Failed open logfile: %v", logfile)
		} else {
			Logger = log.New(logfile, "", log.Ldate|log.Ltime)
			Tracer = log.New(logfile, "", log.Ldate|log.Ltime)
		}
	}
}

func FatalStdoutf(format string, args ...interface{}) {
	stdoutLogger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	stdoutLogger.Printf(fatalLog+" "+format, args)
	os.Exit(1)
}

func Info(source string, metadata map[string]string) {
	msg, err := json.Marshal(metadata)
	if err != nil {
		failedJsonMarshal(source, err.Error())
		return
	}
	Logger.Print(infoLog + " " + source + " " + string(msg))
}

func failedJsonMarshal(source string, msg string) {
	Logger.Print(errorLog + " " + source + " {\"msg\"=\"Failed json.Marshal\"")
}

func Error(source string, metadata map[string]string) {
	msg, err := json.Marshal(metadata)
	if err != nil {
		failedJsonMarshal(source, err.Error())
		return
	}
	Logger.Print(errorLog + " " + source + " " + string(msg))
}

func Warning(source string, metadata map[string]string) {
	msg, err := json.Marshal(metadata)
	if err != nil {
		failedJsonMarshal(source, err.Error())
		return
	}
	Logger.Print(warningLog + " " + source + " " + string(msg))
}

func NewTraceId() string {
	return xid.New().String()
}

func TraceInfo(traceid string, host string, source string, metadata map[string]string) {
	msg, err := json.Marshal(metadata)
	if err != nil {
		failedJsonMarshal(source, err.Error())
		return
	}
	Tracer.Print(traceLog + " " + infoLog + " " + traceid + " " + host + " " + source + " " + string(msg))
}

func TraceError(traceid string, host string, source string, metadata map[string]string) {
	msg, err := json.Marshal(metadata)
	if err != nil {
		failedJsonMarshal(source, err.Error())
		return
	}
	Tracer.Print(traceLog + " " + errorLog + " " + traceid + " " + host + " " + source + " " + string(msg))
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

func StartGrpcTrace(traceId string, host string, name string, ctx context.Context) (time.Time, string) {
	startTime := time.Now()
	var client string

	if pr, ok := peer.FromContext(ctx); ok {
		client = pr.Addr.String()
	} else {
		client = ""
	}

	TraceInfo(traceId, host, name, map[string]string{
		"Msg":    "Start",
		"Client": client,
		"Func":   getFunc(0),
	})

	return startTime, client
}

func EndGrpcTrace(traceId string, host string, name string, startTime time.Time, client string, statusCode int64, err string) {
	TraceInfo(traceId, host, name, map[string]string{
		"Msg":        "End",
		"Client":     client,
		"Func":       getFunc(0),
		"Latency":    strconv.FormatInt(time.Now().Sub(startTime).Nanoseconds()/1000000, 10),
		"StatusCode": strconv.FormatInt(statusCode, 10),
		"Err":        err,
	})
}
