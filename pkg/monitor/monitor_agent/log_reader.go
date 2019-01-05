package monitor_agent

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/monitor/monitor_api/monitor_api_grpc_pb"
)

const (
	LogFormatDefault = 0
	LogFormatTcpDump = 1
)

type LogAlert struct {
	name    string
	handler string
	pattern *regexp.Regexp
}

type LogReader struct {
	name               string
	path               string
	logFormat          int
	maxInitialReadSize int64
	offset             int64
	alertsMap          map[string][]LogAlert
	logs               []*monitor_api_grpc_pb.Log
}

func NewLogReader(conf *config.Config, name string, logConf *config.MonitorLogConfig) (*LogReader, error) {
	var logFormat int
	switch logConf.LogFormat {
	case "":
		logFormat = LogFormatDefault
	default:
		return nil, fmt.Errorf("Invalid LogFormat: %v", logConf.LogFormat)
	}

	alertsMap := map[string][]LogAlert{}
	for alertName, alert := range logConf.AlertMap {
		logAlert := LogAlert{
			name:    alertName,
			handler: alert.Handler,
			pattern: regexp.MustCompile(alert.Pattern),
		}
		if alerts, ok := alertsMap[alert.Key]; ok {
			alerts = append(alerts, logAlert)
			alertsMap[alert.Key] = alerts
		} else {
			alerts = []LogAlert{logAlert}
			alertsMap[alert.Key] = alerts
		}
	}

	reader := &LogReader{
		name:               name,
		path:               conf.LogPath(logConf.Path),
		logFormat:          logFormat,
		maxInitialReadSize: logConf.MaxInitialReadSize,
		alertsMap:          alertsMap,
	}
	// regexp.MustCompile(`golang`)
	err := reader.Init()
	return reader, err
}

func (reader *LogReader) Init() error {
	var err error
	var file *os.File
	var ioreader *bufio.Reader

	file, err = os.Open(reader.path)
	defer file.Close()
	if err != nil {
		return err
	}

	ioreader = bufio.NewReaderSize(file, 1024)

	if reader.maxInitialReadSize > 0 {
		var endOffset int64
		endOffset, err = file.Seek(0, os.SEEK_END)
		if err != nil {
			return err
		}
		if reader.maxInitialReadSize > endOffset {
			reader.offset = 0
		} else {
			_, err = file.Seek(-1*reader.maxInitialReadSize, os.SEEK_END)
			if err != nil {
				return err
			}
			_, err = ioreader.ReadString('\n')
			if err != nil {
				return err
			}
			reader.offset, err = file.Seek(0, os.SEEK_CUR)
			if err != nil {
				return err
			}
		}
	} else {
		reader.offset = 0
	}

	reader.ClearLogs()

	return nil
}

func (reader *LogReader) GetLogs() []*monitor_api_grpc_pb.Log {
	return reader.logs
}

func (reader *LogReader) ClearLogs() {
	reader.logs = []*monitor_api_grpc_pb.Log{}
}

func (reader *LogReader) ReadUntilEOF(tctx *logger.TraceContext) error {
	var err error
	var file *os.File
	var ioreader *bufio.Reader

	defer func() {
		if err != nil {
			initErr := reader.Init()
			if initErr != nil {
				logger.Error(tctx, initErr, "Failed reader.Init()")
			}
		}
	}()

	// Open file each time, because of the file may be deleted or replaced
	file, err = os.Open(reader.path)
	defer file.Close()
	if err != nil {
		return err
	}

	ioreader = bufio.NewReaderSize(file, 1024)

	endOffset, err := file.Seek(0, os.SEEK_END)
	if err != nil {
		return err
	}

	// Check offset, if reader.offset > endOffset, the file may be replaced
	if reader.offset > endOffset {
		return fmt.Errorf("Invalid offset: current=%v, end=%v", reader.offset, endOffset)
	}

	_, err = file.Seek(reader.offset, 0)
	if err != nil {
		return err
	}

	var line string
	for {
		line, err = ioreader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		line = strings.TrimRight(line, "\n")
		switch reader.logFormat {
		case LogFormatDefault:
			var keyStr string
			var valueStr string
			logMap := map[string]string{}
			strIndex := 0
			foundLog := false
			foundEqual := false
			lastIndex := len(line) - 1
			for i, s := range line {
				if foundEqual {
					if s == '"' {
						if i == strIndex {
							continue
						}
						if i == lastIndex || (line[i-1] != '\\' && line[i+1] == ' ') {
							valueStr = line[strIndex+1 : i]
							logMap[keyStr] = valueStr
							foundLog = true
							foundEqual = false
							strIndex = i + 2
							if alerts, ok := reader.alertsMap[keyStr]; ok {
								for _, alert := range alerts {
									if alert.pattern.MatchString(valueStr) {
										logMap["Alert"+alert.name] = alert.handler
									}
								}
							}
						}
					}
				} else {
					if s == '=' {
						foundEqual = true
						keyStr = line[strIndex:i]
						strIndex = i + 1
					}
				}
			}

			if foundLog {
				if time, ok := logMap["Time"]; ok {
					delete(logMap, "Time")
					reader.logs = append(reader.logs, &monitor_api_grpc_pb.Log{
						Name: reader.name,
						Time: time,
						Log:  logMap,
					})
				}
			}

		default:
			break
		}
	}

	reader.offset, err = file.Seek(0, os.SEEK_CUR)
	return nil
}
