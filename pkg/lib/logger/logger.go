package logger

import (
	"encoding/json"
	"log"
	"os"

	"github.com/rs/xid"
)

var (
	Logger *log.Logger
	Tracer *log.Logger
)

const (
	debugLog   = "D"
	infoLog    = "I"
	warningLog = "W"
	errorLog   = "E"
	benchLog   = "B"
	traceLog   = "T"
)

func Init() {
	Logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	Tracer = log.New(os.Stdout, "", log.Ldate|log.Ltime)
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

func TraceInfo(source string, traceid string, metadata map[string]string) {
	msg, err := json.Marshal(metadata)
	if err != nil {
		failedJsonMarshal(source, err.Error())
		return
	}
	Tracer.Print(traceLog + " " + infoLog + " " + traceid + " " + source + " " + string(msg))
}

func TraceError(source string, traceid string, metadata map[string]string) {
	msg, err := json.Marshal(metadata)
	if err != nil {
		failedJsonMarshal(source, err.Error())
		return
	}
	Tracer.Print(traceLog + " " + errorLog + " " + traceid + " " + source + " " + string(msg))
}
