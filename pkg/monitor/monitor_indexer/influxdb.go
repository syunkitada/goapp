package monitor_indexer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func (indexer *InfluxdbIndexer) Report(tctx *logger.TraceContext, req *monitor_api_grpc_pb.ReportRequest) error {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	data := ""
	for _, log := range req.Logs {
		tags := ",Project=" + req.Project
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

func (indexer *InfluxdbIndexer) GetHost(tctx *logger.TraceContext, req *monitor_api_grpc_pb.GetHostRequest, hostMap map[string]*monitor_api_grpc_pb.Host) error {
	// TODO FIX query
	query := "show tag values from \"monitor-api\" with key = \"Host\""
	for _, client := range indexer.logClients {
		result, err := client.Query(query)
		for _, v := range result.Results[0].Series[0].Values {
			host := v[1]
			hostMap[host] = &monitor_api_grpc_pb.Host{
				Name: host,
			}
		}
		if err != nil {
			logger.Warning(tctx, err, "Failed Query")
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
		queryUrl:   userPassUrlDb[1] + "/query?db=" + userPassUrlDb[2],
		writeUrl:   userPassUrlDb[1] + "/write?db=" + userPassUrlDb[2],
		username:   userPass[0],
		password:   userPass[1],
		httpClient: &http.Client{},
	}, nil
}

type QueryResult struct {
	Results []Result
}

type Result struct {
	Series []Series
}

type Series struct {
	Name    string
	Columns []string
	Values  [][]string
}

func (c *Client) Query(data string) (*QueryResult, error) {
	var err error

	req, err := http.NewRequest("GET", c.queryUrl, nil)
	if err != nil {
		return nil, err
	}

	query := req.URL.Query()
	query.Add("q", data)
	req.URL.RawQuery = query.Encode()

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result *QueryResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	if resp.StatusCode == 200 {
		return result, nil
	}

	return nil, fmt.Errorf("InvalidStatusCode: %v", resp.StatusCode)
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
