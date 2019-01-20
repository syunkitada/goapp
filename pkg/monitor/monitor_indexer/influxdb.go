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
	index             string
	logClients        []*Client
	metricClients     []*Client
	percistentClients []*Client
}

func NewInfluxdbIndexer(index string, indexerConfig *config.MonitorIndexerConfig) (*InfluxdbIndexer, error) {
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

	percistentClients := []*Client{}
	for _, databaseInfo := range indexerConfig.PercistentDatabases {
		client, err := NewClient(databaseInfo)
		if err != nil {
			return nil, err
		}
		percistentClients = append(percistentClients, client)
	}

	return &InfluxdbIndexer{
		index:             index,
		logClients:        logClients,
		metricClients:     metricClients,
		percistentClients: percistentClients,
	}, nil
}

func (indexer *InfluxdbIndexer) Report(tctx *logger.TraceContext, req *monitor_api_grpc_pb.ReportRequest) error {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	percistentData := "agent,Project=" + req.Project + ",Kind=" + req.Kind + ",Host=" + req.Host +
		" State=" + strconv.FormatInt(req.State, 10) +
		",Warning=\"" + req.Warning + "\",Warnings=" + strconv.FormatInt(req.Warnings, 10) +
		",Error=\"" + req.Error + "\",Errors=" + strconv.FormatInt(req.Errors, 10) +
		",Timestamp=" + strconv.FormatInt(req.Timestamp, 10) + " 0"

	for _, client := range indexer.percistentClients {
		err := client.Write(percistentData)
		if err != nil {
			logger.Warning(tctx, err, "Failed Write")
		}
	}

	metricsData := ""
	for _, metric := range req.Metrics {
		tags := ",Project=" + req.Project + ",Host=" + req.Host
		values := ""
		for key, value := range metric.Tag {
			switch key {
			case "Project":
				continue
			case "Host":
				continue
			default:
				tags += "," + key + "=" + value
			}
		}

		for key, value := range metric.Metric {
			values += "," + key + "=" + strconv.FormatInt(value, 10) + ""
		}
		values = values[1:]
		metricsData += metric.Name + tags + " " + values + " " + metric.Time + "\n"
	}

	for _, client := range indexer.metricClients {
		err := client.Write(metricsData)
		if err != nil {
			logger.Warning(tctx, err, "Failed Write")
		}
	}

	logData := ""
	for _, log := range req.Logs {
		tags := ",Project=" + req.Project + ",Host=" + req.Host
		logstr := ""
		values := ""
		for key, value := range log.Log {
			logstr += " " + key + "=\\\"" + value + "\\\""
			switch key {
			case "Project":
				continue
			case "Host":
				continue
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
		logData += log.Name + tags + " Log=\"" + logstr[1:] + "\"" + values + " " + strconv.FormatInt(timestamp.UnixNano(), 10) + "\n"
	}

	for _, client := range indexer.logClients {
		err := client.Write(logData)
		if err != nil {
			logger.Warning(tctx, err, "Failed Write")
		}
	}

	return nil
}

func (indexer *InfluxdbIndexer) GetIndex(tctx *logger.TraceContext, projectName string, indexMap map[string]*monitor_api_grpc_pb.Index) error {
	// hosts
	// query := "show tag values from \"agent\" with key = \"Host\""
	query := "select count(State), sum(State) as states, sum(Warnings) as warnings, sum(Errors) as errors from agent where Project = 'admin' group by Host,Kind"
	var count float64 = 0
	var states float64 = 0
	var warnings float64 = 0
	var errors float64 = 0
	for _, client := range indexer.percistentClients {
		result, err := client.Query(query)
		if err != nil {
			logger.Warning(tctx, err, "Failed Query")
			continue
		}

		for _, s := range result.Results[0].Series {
			count = s.Values[0][1].(float64)
			states = s.Values[0][2].(float64)
			warnings = s.Values[0][3].(float64)
			errors = s.Values[0][4].(float64)
		}
	}

	indexMap[indexer.index] = &monitor_api_grpc_pb.Index{
		Name:     indexer.index,
		Count:    count,
		States:   states,
		Warnings: warnings,
		Errors:   errors,
	}

	return nil
}

func (indexer *InfluxdbIndexer) GetHost(tctx *logger.TraceContext, projectName string, hostMap map[string]*monitor_api_grpc_pb.Host) error {
	// hosts
	// query := "show tag values from \"agent\" with key = \"Host\""
	query := "select State, Warning, Warnings, Error, Errors, Timestamp from agent where Project = 'admin' group by Host,Kind"
	for _, client := range indexer.percistentClients {
		result, err := client.Query(query)
		if err != nil {
			logger.Warning(tctx, err, "Failed Query")
			continue
		}

		for _, s := range result.Results[0].Series {
			fmt.Println(s.Values)
			hostMap[s.Tags["Host"]] = &monitor_api_grpc_pb.Host{
				Index:     indexer.index,
				Name:      s.Tags["Host"],
				Kind:      s.Tags["Kind"],
				State:     s.Values[0][1].(float64),
				Warning:   s.Values[0][2].(string),
				Warnings:  s.Values[0][3].(float64),
				Error:     s.Values[0][4].(string),
				Errors:    s.Values[0][5].(float64),
				Timestamp: s.Values[0][6].(float64),
			}
		}
	}

	fmt.Println(hostMap)

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
	Tags    map[string]string
	Columns []string
	Values  [][]interface{}
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
	fmt.Println("DEBUG")
	fmt.Println(string(body))

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
