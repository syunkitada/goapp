package monitor_indexer

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/monitor/monitor_api/monitor_api_grpc_pb"
)

type InfluxdbIndexer struct {
	logClients    []*Client
	metricClients []*Client
}

func NewInfluxdbIndexer(indexerConfig *config.MonitorIndexerConfig) (*InfluxdbIndexer, error) {
	logClients := []*Client{}
	for _, databaseInfo := range indexerConfig.LogDatabases {
		client, err := NewClient(databaseInfo)
		if err != nil {
			return nil, err
		}
		logClients = append(logClients, client)
	}

	metricClients := []*Client{}
	for _, databaseInfo := range indexerConfig.MetricDatabases {
		client, err := NewClient(databaseInfo)
		if err != nil {
			return nil, err
		}
		metricClients = append(metricClients, client)
	}

	return &InfluxdbIndexer{
		logClients:    logClients,
		metricClients: metricClients,
	}, nil
}

func (indexer *InfluxdbIndexer) Report(tctx *logger.TraceContext, logs []*monitor_api_grpc_pb.Log) error {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err) }()

	data := ""
	for _, log := range logs {
		tags := ""
		logstr := ""
		values := ""
		for key, value := range log.Log {
			logstr += " " + key + "=\\\"" + value + "\\\""
			switch key {
			case "Host":
				tags += ",Host=" + value
			case "App":
				tags += ",App=" + value
			case "Func":
				tags += ",Func=" + value
			case "Level":
				tags += ",Level=" + value
			case "TraceId":
				tags += ",TraceId=" + value
			case "Latency":
				values += ",Latency=" + value
			default:
				values += "," + key + "=\"" + value + "\""
			}
		}

		timestamp, err := time.Parse(logger.LogTimeFormat, log.Time)
		if err != nil {
			continue
		}
		data += log.Name + tags + " Log=\"" + logstr[1:] + "\"" + values + " " + strconv.FormatInt(timestamp.UnixNano(), 10) + "\n"
	}

	for _, client := range indexer.logClients {
		err := client.Write(data)
		if err != nil {
			logger.Warning(tctx, err, "Failed Write")
		}
	}

	return nil
}

type Client struct {
	queryUrl   string
	writeUrl   string
	username   string
	password   string
	httpClient *http.Client
}

func NewClient(databaseInfo string) (*Client, error) {
	userPassUrlDb := strings.Split(databaseInfo, "@")
	if len(userPassUrlDb) != 3 {
		return nil, fmt.Errorf("Invalid influxdb connection")
	}

	userPass := strings.Split(userPassUrlDb[0], ":")
	if len(userPass) != 2 {
		return nil, fmt.Errorf("Invalid influxdb connection")
	}

	return &Client{
		queryUrl:   userPassUrlDb[1] + "/query",
		writeUrl:   userPassUrlDb[1] + "/write?db=" + userPassUrlDb[2],
		username:   userPass[0],
		password:   userPass[1],
		httpClient: &http.Client{},
	}, nil
}

func (c *Client) Write(data string) error {
	var err error
	req, err := http.NewRequest("POST", c.writeUrl, bytes.NewBuffer([]byte(data)))
	if err != nil {
		return err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode == 204 {
		return nil
	}

	return fmt.Errorf("InvalidStatusCode: %v", resp.StatusCode)
}
