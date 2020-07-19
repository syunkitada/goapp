package influxdb_driver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/config"
	"github.com/syunkitada/goapp/pkg/resource/db_model"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
	api_spec "github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

type FilterEventRule struct {
	ReNode  *regexp.Regexp
	ReMsg   *regexp.Regexp
	ReCheck *regexp.Regexp
	ReLevel *regexp.Regexp
	Until   *time.Time
}

type InfluxdbDriver struct {
	clusterConf      *config.ResourceClusterConfig
	eventClients     []*Client
	logClients       []*Client
	metricClients    []*Client
	mtx              *sync.Mutex
	filterEventRules []FilterEventRule
}

func New(clusterConf *config.ResourceClusterConfig) *InfluxdbDriver {
	tsdbConf := clusterConf.TimeSeriesDatabase

	eventClients := []*Client{}
	for _, connection := range tsdbConf.EventDatabases {
		client, err := NewClient(connection)
		if err != nil {
			logger.StdoutFatalf("Failed init client: %v", err)
		}
		eventClients = append(eventClients, client)
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
		eventClients:  eventClients,
		logClients:    logClients,
		metricClients: metricClients,
		mtx:           new(sync.Mutex),
	}
}

func (driver *InfluxdbDriver) SetFilterEventRules(tctx *logger.TraceContext, eventRules []db_model.EventRule) {
	// Project string     `gorm:"not null;size:50;"`
	// Node    string     `gorm:"not null;size:255;"`
	// Msg     string     `gorm:"not null;size:255;"`
	// Check   string     `gorm:"not null;size:255;"`
	// Level   string     `gorm:"not null;size:50;"`
	// Until   *time.Time `gorm:""`

	var filterEventRules []FilterEventRule
	for _, rule := range eventRules {
		filter := FilterEventRule{}
		if rule.Node != "" {
			re, tmpErr := regexp.Compile(rule.Node)
			if tmpErr != nil {
				logger.Errorf(tctx, tmpErr, "Failed compile rule: %s", rule.Node)
			}
			filter.ReNode = re
		}
		if rule.Check != "" {
			re, tmpErr := regexp.Compile(rule.Check)
			if tmpErr != nil {
				logger.Errorf(tctx, tmpErr, "Failed compile rule: %s", rule.Check)
			}
			filter.ReCheck = re
		}
		if rule.Msg != "" {
			re, tmpErr := regexp.Compile(rule.Msg)
			if tmpErr != nil {
				logger.Errorf(tctx, tmpErr, "Failed compile rule: %s", rule.Msg)
			}
			filter.ReMsg = re
		}
		if rule.Level != "" {
			re, tmpErr := regexp.Compile(rule.Level)
			if tmpErr != nil {
				logger.Errorf(tctx, tmpErr, "Failed compile rule: %s", rule.Level)
			}
			filter.ReLevel = re
		}
		filter.Until = rule.Until
		filterEventRules = append(filterEventRules, filter)
	}
	driver.mtx.Lock()
	driver.filterEventRules = filterEventRules
	driver.mtx.Unlock()
}

func (driver *InfluxdbDriver) Report(tctx *logger.TraceContext, input *api_spec.ReportNode) error {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	filterEventRules := driver.filterEventRules
	eventsData := ""
	var filterEvent bool
	for _, event := range input.Events {
		filterEvent = false
		for _, filter := range filterEventRules {
			if filter.ReCheck != nil {
				if filter.ReCheck.MatchString(event.Name) {
					filterEvent = true
				}
			}
			if filter.ReNode != nil {
				if filter.ReNode.MatchString(input.Name) {
					filterEvent = true
				} else {
					filterEvent = false
				}
			}
			if filter.ReMsg != nil {
				if msg, ok := event.Tag["Msg"]; ok {
					if filter.ReMsg.MatchString(msg) {
						filterEvent = true
					} else {
						filterEvent = false
					}
				}
			}
			if filterEvent {
				break
			}
		}
		if filterEvent {
			continue
		}

		tags := ",Check=" + event.Name + ",Level=" + event.Level + ",Project=" + input.Project + ",Node=" + input.Name
		for key, value := range event.Tag {
			switch key {
			case "Project":
				continue
			case "Node":
				continue
			default:
				tags += "," + key + "=" + value
			}
		}

		eventsData += "events" + tags + " ReissueDuration=" + strconv.Itoa(event.ReissueDuration) + ",Msg=\"" + event.Msg + "\" " + event.Time + "\n"
	}

	for _, client := range driver.eventClients {
		err := client.Write(eventsData)
		if err != nil {
			logger.Warning(tctx, err, "Failed Write events")
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

	for _, client := range driver.metricClients {
		err := client.Write(metricsData)
		if err != nil {
			logger.Warning(tctx, err, "Failed Write")
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

func (driver *InfluxdbDriver) GetNode(tctx *logger.TraceContext, input *api_spec.GetNodeMetrics) (data []api_spec.MetricsGroup, err error) {
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	until := "now()"
	if input.UntilTime != nil {
		until = fmt.Sprintf("'%s'", input.UntilTime.Format(time.RFC3339))
	}
	from := "-6h"
	if input.FromTime != "" {
		from = input.FromTime
	}

	whereStr := fmt.Sprintf("WHERE Node = '%s' AND time < %s AND time > %s %s",
		input.Name, until, until, from)
	suffixQuery := fmt.Sprintf("%s GROUP BY time(1m)", whereStr)

	var systemMetrics []api_spec.Metric
	driver.GetMetrics(tctx,
		&systemMetrics,
		"ProcsRunning",
		fmt.Sprintf("SELECT max(procs_running), mean(procs_blocked) FROM system_cpu %s fill(null)", suffixQuery),
		[]string{"procs_running", "procs_blocked"})

	// Memo: derivativeを使う場合は、mean(processes)も一緒に取得しないと、開始時間からのメトリックスが存在しない期間分が取得できない(nullで埋められない)
	driver.GetMetrics(tctx,
		&systemMetrics,
		"NewProcesses/Min",
		fmt.Sprintf("SELECT non_negative_derivative(max(processes), 1m), max(processes) FROM system_cpu %s fill(null)", suffixQuery),
		[]string{"new_processes"})

	driver.GetMetrics(tctx,
		&systemMetrics,
		"Procs",
		fmt.Sprintf("SELECT max(procs) FROM system_procs %s fill(null)", suffixQuery),
		[]string{"procs"})

	data = append(data, api_spec.MetricsGroup{
		Name:    "system",
		Metrics: systemMetrics,
	})

	var systemMemMetrics []api_spec.Metric
	driver.GetMetrics(tctx,
		&systemMemMetrics,
		"Mem",
		fmt.Sprintf("SELECT max(mem_total), max(mem_free), max(reclaimable) FROM system_mem %s fill(null)", suffixQuery),
		[]string{"total", "free", "reclaimable"})

	driver.GetMetrics(tctx,
		&systemMemMetrics,
		"Slab",
		fmt.Sprintf("SELECT max(slab), max(s_reclaimable), max(s_unreclaim) FROM system_mem %s fill(null)", suffixQuery),
		[]string{"slab", "slab_reclaimable"})

	driver.GetMetrics(tctx,
		&systemMemMetrics,
		"PgScan",
		fmt.Sprintf("SELECT max(pgscan_kswapd), max(pgscan_direct) FROM system_vmstat %s fill(null)", suffixQuery),
		[]string{"pgscan_kswapd", "pgscan_direct"})

	driver.GetMetrics(tctx,
		&systemMemMetrics,
		"Pgfault",
		fmt.Sprintf("SELECT max(pgfault) FROM system_vmstat %s fill(null)", suffixQuery),
		[]string{"pgfault"})

	driver.GetMetrics(tctx,
		&systemMemMetrics,
		"Pswap",
		fmt.Sprintf("SELECT max(pswapin), max(pswapout) FROM system_vmstat %s fill(null)", suffixQuery),
		[]string{"pswapin", "pswapout"})

	data = append(data, api_spec.MetricsGroup{
		Name:    "system_mem",
		Metrics: systemMemMetrics,
	})

	var systemDiskMetrics []api_spec.Metric
	driver.GetMetrics(tctx,
		&systemDiskMetrics,
		"Disk",
		fmt.Sprintf("SELECT max(total_size), max(free_size) FROM system_fsstat %s fill(null)", suffixQuery),
		[]string{"total_size", "free_size"})

	driver.GetMetrics(tctx,
		&systemDiskMetrics,
		"Block Reads/Writes",
		fmt.Sprintf("SELECT max(reads_per_sec), max(writes_per_sec) FROM system_diskstat %s fill(null)", suffixQuery),
		[]string{"reads_per_sec", "writes_per_sec"})

	driver.GetMetrics(tctx,
		&systemDiskMetrics,
		"Block ReadBytes/WriteBytes",
		fmt.Sprintf("SELECT max(read_bytes_per_sec), max(write_bytes_per_sec) FROM system_diskstat %s fill(null)", suffixQuery),
		[]string{"read_bytes_per_sec", "write_bytes_per_sec"})

	driver.GetMetrics(tctx,
		&systemDiskMetrics,
		"Block ReadMsPerSec/WriteMsPerSec",
		fmt.Sprintf("SELECT max(read_ms_per_sec), max(write_ms_per_sec) FROM system_diskstat %s fill(null)", suffixQuery),
		[]string{"read_ms_per_sec", "write_ms_per_sec"})

	driver.GetMetrics(tctx,
		&systemDiskMetrics,
		"Block ProgressIos",
		fmt.Sprintf("SELECT max(progress_ios) FROM system_diskstat %s fill(null)", suffixQuery),
		[]string{"progress_ios"})

	data = append(data, api_spec.MetricsGroup{
		Name:    "system disk",
		Metrics: systemDiskMetrics,
	})

	var systemNetMetrics []api_spec.Metric
	driver.GetMetrics(tctx,
		&systemNetMetrics,
		"NetDev Bytes/Sec",
		fmt.Sprintf("SELECT max(receive_bytes_per_sec), max(transmit_bytes_per_sec) FROM system_netdevstat %s fill(null)", suffixQuery),
		[]string{"receive_bytes_per_sec", "transmit_bytes_per_sec"})

	driver.GetMetrics(tctx,
		&systemNetMetrics,
		"NetDev Packets/Sec",
		fmt.Sprintf("SELECT max(receive_packets_per_sec), max(transmit_packets_per_sec) FROM system_netdevstat %s fill(null)", suffixQuery),
		[]string{"receive_packets_per_sec", "transmit_packets_per_sec"})

	driver.GetMetrics(tctx,
		&systemNetMetrics,
		"NetDev Errors",
		fmt.Sprintf("SELECT max(receive_errors), max(transmit_errors) FROM system_netdevstat %s fill(null)", suffixQuery),
		[]string{"receive_errors", "transmit_errors"})

	driver.GetMetrics(tctx,
		&systemNetMetrics,
		"NetDev Drops",
		fmt.Sprintf("SELECT max(receive_drops), max(transmit_drops) FROM system_netdevstat %s fill(null)", suffixQuery),
		[]string{"receive_drops", "transmit_drops"})

	data = append(data, api_spec.MetricsGroup{
		Name:    "system netdev",
		Metrics: systemNetMetrics,
	})

	var systemProcMetrics []api_spec.Metric
	driver.GetMetrics(tctx,
		&systemProcMetrics,
		"ProcSched/Min",
		fmt.Sprintf("SELECT non_negative_derivative(max(sched_cpu_time), 1m), non_negative_derivative(max(sched_wait_time), 1m), max(sched_time_slices) FROM system_proc %s, cmd, pid fill(null)", suffixQuery),
		[]string{"sched_cpu_time", "sched_wait_time", "sched_time_slices"})

	data = append(data, api_spec.MetricsGroup{
		Name:    "system proc",
		Metrics: systemProcMetrics,
	})

	return
}

func (driver *InfluxdbDriver) GetMetrics(tctx *logger.TraceContext, metrics *[]api_spec.Metric, name string, query string, keys []string) {
	fmt.Println(query)
	for _, client := range driver.metricClients {
		queryResult, tmpErr := client.Query(query)
		if tmpErr != nil {
			fmt.Println("DEBUG FaledQuery", tmpErr)
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
				values = values[0 : len(values)-1]
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
		if tmpErr != nil {
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
		if tmpErr != nil {
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

func (driver *InfluxdbDriver) GetLogs(tctx *logger.TraceContext, input *api_spec.GetLogs) (data *api_spec.GetLogsData, err error) {
	logs := []map[string]interface{}{}

	query := "SELECT App, Func, Level, Node, Project, TraceId, Msg FROM \"goapp-resource-cluster-agent\" WHERE"
	whereApps := []string{}
	if len(input.Apps) > 0 {
		for _, app := range input.Apps {
			whereApps = append(whereApps, fmt.Sprintf("App = '%s'", app))
		}
		whereAppsStr := strings.Join(whereApps, " OR ")
		query += " (" + whereAppsStr + ")"
	}

	if len(input.Nodes) > 0 {
		whereNodes := []string{}
		for _, node := range input.Nodes {
			whereNodes = append(whereNodes, fmt.Sprintf("Node = '%s'", node))
		}
		whereNodesStr := strings.Join(whereNodes, " OR ")
		if len(whereApps) > 0 {
			query += " AND"
		}
		query += " (" + whereNodesStr + ")"
	}

	query += " AND time > now() - 20d"
	query += " limit 10000"
	keys := []string{"App", "Func", "Level", "Node", "Project", "TraceId", "Msg"}
	for _, client := range driver.logClients {
		queryResult, tmpErr := client.Query(query)
		if tmpErr != nil {
			logger.Warningf(tctx, "Failed Query: %s", tmpErr.Error())
			continue
		}
		for _, result := range queryResult.Results {
			for _, series := range result.Series {
				for _, value := range series.Values {
					v := map[string]interface{}{
						"Time": value[0],
					}
					for i, key := range keys {
						v[key] = value[i+1]
					}
					logs = append(logs, v)
				}
			}
		}
	}

	data = &api_spec.GetLogsData{Logs: logs}
	return
}

func (driver *InfluxdbDriver) GetEvents(tctx *logger.TraceContext, input *api_spec.GetEvents) (data *api_spec.GetEventsData, err error) {
	events := []spec.Event{}

	query := "SELECT Check, Level, ReissueDuration, Node, Project, last(Msg) FROM \"events\" WHERE"
	query += " time > now() - 1d"
	query += " group by Node,Check"
	for _, client := range driver.eventClients {
		queryResult, tmpErr := client.Query(query)
		if tmpErr != nil {
			logger.Warningf(tctx, "Failed Query: %s", tmpErr.Error())
			continue
		}
		for _, result := range queryResult.Results {
			for _, series := range result.Series {
				for _, value := range series.Values {
					sec := int64(value[0].(float64) / 1000)
					v := spec.Event{
						Time:            time.Unix(sec, 0),
						Check:           value[1].(string),
						Level:           value[2].(string),
						ReissueDuration: int(value[3].(float64)),
						Node:            value[4].(string),
						Project:         value[5].(string),
						Msg:             value[6].(string),
					}
					events = append(events, v)
				}
			}
		}
	}

	data = &api_spec.GetEventsData{Events: events}
	return
}

func (driver *InfluxdbDriver) IssueEvent(tctx *logger.TraceContext, input *api_spec.IssueEvent) (err error) {
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	eventsData := ""

	event := input.Event
	tags := ",Check=" + event.Check + ",Level=" + event.Level + ",Project=" + event.Project + ",Node=" + event.Node
	eventsData += "issued_events" + tags + " Msg=\"" + event.Msg + "\" " + strconv.FormatInt(event.Time.UnixNano(), 10) + "\n"
	fmt.Println("DEBUG IssueEvent", eventsData)

	for _, client := range driver.eventClients {
		tmpErr := client.Write(eventsData)
		if tmpErr != nil {
			logger.Warning(tctx, tmpErr, "Failed Write events")
		}
	}
	return
}

func (driver *InfluxdbDriver) GetIssuedEvents(tctx *logger.TraceContext, input *api_spec.GetIssuedEvents) (data *api_spec.GetIssuedEventsData, err error) {
	events := []spec.Event{}

	query := "SELECT Check, Level, Node, Project, last(Msg) FROM \"issued_events\" WHERE"
	if input.Node != "" {
		query += " Node = \"" + input.Node + "\""
	}
	query += " time > now() - 1d"
	query += " group by Node,Check"
	for _, client := range driver.eventClients {
		queryResult, tmpErr := client.Query(query)
		if tmpErr != nil {
			logger.Warningf(tctx, "Failed Query: %s", tmpErr.Error())
			continue
		}
		for _, result := range queryResult.Results {
			for _, series := range result.Series {
				for _, value := range series.Values {
					sec := int64(value[0].(float64) / 1000)
					v := spec.Event{
						Time:    time.Unix(sec, 0),
						Check:   value[1].(string),
						Level:   value[2].(string),
						Node:    value[3].(string),
						Project: value[4].(string),
						Msg:     value[5].(string),
					}
					events = append(events, v)
				}
			}
		}
	}

	data = &api_spec.GetIssuedEventsData{Events: events}
	return
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
