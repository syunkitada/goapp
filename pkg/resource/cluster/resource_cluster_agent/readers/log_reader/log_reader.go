package log_reader

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/config"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

const (
	LogFormatDefault = 0
	LogFormatTcpDump = 1
)

type LogEvent struct {
	name            string
	handler         string
	level           string
	pattern         *regexp.Regexp
	reissueDuration int
}

type LogReader struct {
	name               string
	path               string
	logFormat          int
	maxInitialReadSize int64
	offset             int64
	eventsMap          map[string][]LogEvent
	logs               []spec.ResourceLog
	events             []spec.ResourceEvent
}

func New(baseConf *base_config.Config, name string, logConf *config.ResourceLogConfig) (*LogReader, error) {
	var logFormat int
	switch logConf.LogFormat {
	case "":
		logFormat = LogFormatDefault
	default:
		return nil, fmt.Errorf("Invalid LogFormat: %v", logConf.LogFormat)
	}

	eventsMap := map[string][]LogEvent{}
	for checkName, check := range logConf.CheckMap {
		logEvent := LogEvent{
			name:            logConf.CheckPrefix + checkName,
			handler:         check.Handler,
			level:           check.Level,
			pattern:         regexp.MustCompile(check.Pattern),
			reissueDuration: check.ReissueDuration,
		}
		if events, ok := eventsMap[check.Key]; ok {
			events = append(events, logEvent)
			eventsMap[check.Key] = events
		} else {
			events = []LogEvent{logEvent}
			eventsMap[check.Key] = events
		}
	}

	reader := &LogReader{
		name:               name,
		path:               path.Join(baseConf.LogDir, logConf.Path),
		logFormat:          logFormat,
		maxInitialReadSize: logConf.MaxInitialReadSize,
		eventsMap:          eventsMap,
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

func (reader *LogReader) Report() ([]spec.ResourceLog, []spec.ResourceEvent) {
	return reader.logs, reader.events
}

func (reader *LogReader) Reported() {
	reader.ClearLogs()
}

func (reader *LogReader) ClearLogs() {
	reader.logs = []spec.ResourceLog{}
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

	fmt.Println("DEBUG ReadUntil", reader.path)

	// Open file each time, because of the file may be deleted or replaced
	file, err = os.Open(reader.path)
	defer file.Close()
	if err != nil {
		return err
	}

	ioreader = bufio.NewReaderSize(file, 10240)

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
							if events, ok := reader.eventsMap[keyStr]; ok {
								for _, event := range events {
									if event.pattern.MatchString(valueStr) {
										if timeStr, ok := logMap["Time"]; ok {
											t, err := time.Parse(time.RFC3339, timeStr)
											if err != nil {
												t = time.Now()
											}
											tstr := strconv.FormatInt(t.UnixNano(), 10)
											reader.events = append(reader.events, spec.ResourceEvent{
												Name:            event.name,
												Time:            tstr,
												Level:           event.level,
												Handler:         event.handler,
												Msg:             line,
												ReissueDuration: event.reissueDuration,
											})
										}
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
					reader.logs = append(reader.logs, spec.ResourceLog{
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
