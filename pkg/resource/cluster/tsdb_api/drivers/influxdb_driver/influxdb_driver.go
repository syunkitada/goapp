package influxdb_driver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/config"
	api_spec "github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

type InfluxdbDriver struct {
	clusterConf   *config.ResourceClusterConfig
	alertClients  []*Client
	logClients    []*Client
	metricClients []*Client
}

func New(clusterConf *config.ResourceClusterConfig) *InfluxdbDriver {
	tsdbConf := clusterConf.TimeSeriesDatabase

	alertClients := []*Client{}
	for _, connection := range tsdbConf.AlertDatabases {
		client, err := NewClient(connection)
		if err != nil {
			logger.StdoutFatalf("Failed init client: %v", err)
		}
		alertClients = append(alertClients, client)
	}

	logClients := []*Client{}
	for _, connection := range tsdbConf.LogDatabases {
		client, err := NewClient(connection)
		if err != nil {
			logger.StdoutFatalf("Failed init client: %v", err)
		}
		logClients = append(logClients, client)
	}

	metricClients := []*Client{}
	for _, connection := range tsdbConf.MetricDatabases {
		client, err := NewClient(connection)
		if err != nil {
			logger.StdoutFatalf("Failed init client: %v", err)
		}
		metricClients = append(metricClients, client)
	}

	return &InfluxdbDriver{
		clusterConf:   clusterConf,
		alertClients:  alertClients,
		logClients:    logClients,
		metricClients: metricClients,
	}
}

func (driver *InfluxdbDriver) Report(tctx *logger.TraceContext, input *api_spec.ReportNode) error {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	alertsData := ""
	for _, alert := range input.Alerts {
		tags := ",Project=" + input.Project + ",Node=" + input.Name
		for key, value := range alert.Tag {
			switch key {
			case "Project":
				continue
			case "Node":
				continue
			default:
				tags += "," + key + "=" + value
			}
		}

		alertsData += alert.Name + tags + " Msg=\"" + alert.Msg + "\" " + alert.Time + "\n"
	}

	for _, client := range driver.alertClients {
		err := client.Write(alertsData)
		if err != nil {
			logger.Warning(tctx, err, "Failed Write alerts")
		}
	}

	metricsData := ""
	for _, metric := range input.Metrics {
		tags := ",Project=" + input.Project + ",Node=" + input.Name
		values := ""
		for key, value := range metric.Tag {
			switch key {
			case "Project":
				continue
			case "Node":
				continue
			default:
				tags += "," + key + "=" + value
			}
		}

		for key, value := range metric.Metric {
			values += "," + key + "=" + fmt.Sprint(value) + ""
		}
		values = values[1:]
		metricsData += metric.Name + tags + " " + values + " " + metric.Time + "\n"
	}

	fmt.Println("debug metrics", metricsData)

	for _, client := range driver.metricClients {
		err := client.Write(metricsData)
		if err != nil {
			logger.Warning(tctx, err, "Failed Write")
			fmt.Println("DEBUG Failed Write metrics", err)
		}
	}

	logData := ""
	for _, log := range input.Logs {
		tags := ",Project=" + input.Project + ",Node=" + input.Name
		logstr := ""
		values := ""
		for key, value := range log.Log {
			logstr += " " + key + "=\\\"" + value + "\\\""
			switch key {
			case "Project":
				continue
			case "Node":
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

	for _, client := range driver.logClients {
		err := client.Write(logData)
		if err != nil {
			logger.Warning(tctx, err, "Failed Write log")
		}
	}

	return nil
}

func (driver *InfluxdbDriver) GetNode(tctx *logger.TraceContext, input *api_spec.GetNode) (data []api_spec.MetricsGroup, err error) {
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	fmt.Println("DEBUG GetNode")
	// query := "show tag values from \"system_cpu\" with key = \"Host\""
	// query := "select State, Warning, Warnings, Error, Errors, Timestamp from agent where Project = 'admin' group by Host,Kind"
	var systemMetrics []api_spec.Metric

	driver.GetMetrics(tctx,
		&systemMetrics,
		"ProcsRunning",
		fmt.Sprintf("select procs_running, procs_blocked from system_cpu where Node = '%s'", input.Name),
		[]string{"procs_running", "procs_blocked"})

	driver.GetMetrics(tctx,
		&systemMetrics,
		"Processes",
		fmt.Sprintf("select processes from system_cpu where Node = '%s'", input.Name),
		[]string{"processes"})

	data = append(data, api_spec.MetricsGroup{
		Name:    "system",
		Metrics: systemMetrics,
	})

	return
}

func (driver *InfluxdbDriver) GetMetrics(tctx *logger.TraceContext, metrics *[]api_spec.Metric, name string, query string, keys []string) {
	for _, client := range driver.metricClients {
		queryResult, tmpErr := client.Query(query)
		if tmpErr != nil {
			logger.Warningf(tctx, "Failed Query: %s", tmpErr.Error())
			continue
		}
		for _, result := range queryResult.Results {
			for _, series := range result.Series {
				values := []map[string]interface{}{}
				for _, value := range series.Values {
					v := map[string]interface{}{
						"time": value[0],
					}
					for i, key := range keys {
						v[key] = value[i+1]
					}
					values = append(values, v)
				}
				*metrics = append(*metrics, api_spec.Metric{
					Name:   name,
					Values: values,
					Keys:   keys,
				})
			}
		}
	}
}

func (driver *InfluxdbDriver) GetLogParams(tctx *logger.TraceContext, input *api_spec.GetLogParams) (data *api_spec.GetLogParamsData, err error) {
	nodesQuery := "show tag values from \"goapp-resource-cluster-agent\" with key = \"Node\""
	nodes := []string{}
	for _, client := range driver.logClients {
		result, tmpErr := client.Query(nodesQuery)
		if err != nil {
			logger.Warningf(tctx, "Failed Query: %s", tmpErr.Error())
			continue
		}
		for _, s := range result.Results[0].Series {
			for _, v := range s.Values {
				nodes = append(nodes, v[1].(string))
			}
		}
	}

	appsQuery := "show tag values from \"goapp-resource-cluster-agent\" with key = \"App\""
	apps := []string{}
	for _, client := range driver.logClients {
		result, tmpErr := client.Query(appsQuery)
		if err != nil {
			logger.Warningf(tctx, "Failed Query: %s", tmpErr.Error())
			continue
		}
		for _, s := range result.Results[0].Series {
			for _, v := range s.Values {
				apps = append(apps, v[1].(string))
			}
		}
	}

	data = &api_spec.GetLogParamsData{
		LogNodes: nodes,
		LogApps:  apps,
	}
	return
}

func (driver *InfluxdbDriver) GetHost(tctx *logger.TraceContext, projectName string) error {
	// hosts
	// query := "show tag values from \"agent\" with key = \"Host\""
	// query := "select State, Warning, Warnings, Error, Errors, Timestamp from agent where Project = 'admin' group by Host,Kind"
	// for _, client := range indexer.percistentClients {
	// 	result, err := client.Query(query)
	// 	if err != nil {
	// 		logger.Warning(tctx, err, "Failed Query")
	// 		continue
	// 	}

	// 	for _, s := range result.Results[0].Series {
	// 		fmt.Println(s.Values)
	// 		hostMap[s.Tags["Host"]] = &monitor_api_grpc_pb.Host{
	// 			Index:     indexer.index,
	// 			Name:      s.Tags["Host"],
	// 			Kind:      s.Tags["Kind"],
	// 			State:     s.Values[0][1].(float64),
	// 			Warning:   s.Values[0][2].(string),
	// 			Warnings:  s.Values[0][3].(float64),
	// 			Error:     s.Values[0][4].(string),
	// 			Errors:    s.Values[0][5].(float64),
	// 			Timestamp: s.Values[0][6].(float64),
	// 		}
	// 	}
	// }

	// fmt.Println(hostMap)

	return nil
}

type Client struct {
	queryUrl   string
	writeUrl   string
	username   string
	password   string
	httpClient *http.Client
}

func NewClient(connection string) (*Client, error) {
	userPassUrlDb := strings.Split(connection, "@")
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
	query.Add("epoch", "ms")
	query.Add("q", data)
	req.URL.RawQuery = query.Encode()

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() { err = resp.Body.Close() }()
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

	fmt.Println("DEBUG data", data)

	return fmt.Errorf("InvalidStatusCode: %v", resp.StatusCode)
}
