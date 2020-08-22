package resolver

import (
	"encoding/json"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"

	"github.com/syunkitada/goapp/pkg/resource/consts"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
	api_spec "github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

func (resolver *Resolver) ReportNode(tctx *logger.TraceContext, input *api_spec.ReportNode,
	user *base_spec.UserAuthority) (data *api_spec.ReportNodeData, code uint8, err error) {
	if err = resolver.dbApi.ReportNode(tctx, input); err != nil {
		code = base_const.CodeServerInternalError
		return
	}

	if err = resolver.tsdbApi.ReportNode(tctx, input); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	fmt.Println("DEBUG logs:", len(input.Logs))
	fmt.Println("DEBUG metrics:", len(input.Metrics))
	fmt.Println("DEBUG events:", len(input.Events))
	code = base_const.CodeOk
	data = &api_spec.ReportNodeData{}
	return
}

func (resolver *Resolver) GetNodes(tctx *logger.TraceContext, input *api_spec.GetNodes,
	user *base_spec.UserAuthority) (data *api_spec.GetNodesData, code uint8, err error) {
	var getIssuedEventsData *api_spec.GetIssuedEventsData
	getIssuedEventsData, err = resolver.tsdbApi.GetIssuedEvents(tctx, &api_spec.GetIssuedEvents{})
	if err != nil {
		code = base_const.CodeServerInternalError
		return
	}

	var nodeServices []base_spec.NodeService
	if nodeServices, err = resolver.dbApi.GetNodeServices(tctx, &base_spec.GetNodeServices{}, &base_spec.UserAuthority{}); err != nil {
		code = base_const.CodeServerInternalError
		return
	}

	nodeMap := map[string]api_spec.Node{}
	for _, event := range getIssuedEventsData.Events {
		node, ok := nodeMap[event.Node]
		if !ok {
			node = api_spec.Node{}
		}
		if event.Silenced > 0 {
			node.SilencedEventsData = append(node.SilencedEventsData, event)
		}
		switch event.Level {
		case consts.EventLevelSuccess:
			node.SuccessEventsData = append(node.SuccessEventsData, event)
		case consts.EventLevelCritical:
			node.CriticalEventsData = append(node.CriticalEventsData, event)
		case consts.EventLevelWarning:
			node.WarningEventsData = append(node.WarningEventsData, event)
		}
		nodeMap[event.Node] = node
	}

	for _, service := range nodeServices {
		node, ok := nodeMap[service.Name]
		if !ok {
			node = api_spec.Node{Name: service.Name}
		}
		if service.Status == base_const.StatusDisabled {
			node.DisabledServicesData = append(node.DisabledServicesData, service)
		} else if service.State == base_const.StateUp {
			node.ActiveServicesData = append(node.ActiveServicesData, service)
		} else {
			node.CriticalServicesData = append(node.CriticalServicesData, service)
		}
		nodeMap[service.Name] = node
	}

	var resultNodes []api_spec.Node
	for _, node := range nodeMap {
		node.ActiveServices = len(node.ActiveServicesData)
		node.CriticalServices = len(node.CriticalServicesData)
		node.DisabledServices = len(node.DisabledServicesData)
		node.SuccessEvents = len(node.SuccessEventsData)
		node.CriticalEvents = len(node.CriticalEventsData)
		node.WarningEvents = len(node.WarningEventsData)
		node.SilencedEvents = len(node.SilencedEventsData)
		resultNodes = append(resultNodes, node)
	}

	code = base_const.CodeOk
	data = &spec.GetNodesData{Nodes: resultNodes}
	return
}

func (resolver *Resolver) GetNode(tctx *logger.TraceContext, input *api_spec.GetNode,
	user *base_spec.UserAuthority) (data *api_spec.GetNodeData, code uint8, err error) {

	var getIssuedEventsData *api_spec.GetIssuedEventsData
	getIssuedEventsData, err = resolver.tsdbApi.GetIssuedEvents(
		tctx, &api_spec.GetIssuedEvents{Node: input.Name})
	if err != nil {
		code = base_const.CodeServerInternalError
		return
	}

	node := api_spec.Node{Name: input.Name}
	for _, event := range getIssuedEventsData.Events {
		if event.Silenced > 0 {
			node.SilencedEventsData = append(node.SilencedEventsData, event)
		}
		switch event.Level {
		case consts.EventLevelSuccess:
			node.SuccessEventsData = append(node.SuccessEventsData, event)
		case consts.EventLevelCritical:
			node.CriticalEventsData = append(node.CriticalEventsData, event)
		case consts.EventLevelWarning:
			node.WarningEventsData = append(node.WarningEventsData, event)
		}
	}

	var nodeServices []base_spec.NodeService
	if nodeServices, err = resolver.dbApi.GetNodeServices(tctx, &base_spec.GetNodeServices{
		Name: input.Name,
	}, &base_spec.UserAuthority{}); err != nil {
		code = base_const.CodeServerInternalError
		return
	}

	for _, service := range nodeServices {
		if service.Status == base_const.StatusDisabled {
			node.DisabledServicesData = append(node.DisabledServicesData, service)
		} else if service.State == base_const.StateUp {
			node.ActiveServicesData = append(node.ActiveServicesData, service)
		} else {
			node.CriticalServicesData = append(node.CriticalServicesData, service)
		}
	}

	node.ActiveServices = len(node.ActiveServicesData)
	node.CriticalServices = len(node.CriticalServicesData)
	node.DisabledServices = len(node.DisabledServicesData)
	node.SuccessEvents = len(node.SuccessEventsData)
	node.CriticalEvents = len(node.CriticalEventsData)
	node.WarningEvents = len(node.WarningEventsData)
	node.SilencedEvents = len(node.SilencedEventsData)

	code = base_const.CodeOk
	data = &spec.GetNodeData{Node: node}
	return
}

func (resolver *Resolver) GetNodeMetrics(tctx *logger.TraceContext, input *api_spec.GetNodeMetrics,
	user *base_spec.UserAuthority) (data *api_spec.GetNodeMetricsData, code uint8, err error) {
	node := api_spec.Node{Name: input.Name}
	var metricsGroup []api_spec.MetricsGroup
	if metricsGroup, err = resolver.tsdbApi.GetNode(tctx, input); err != nil {
		return
	}
	node.MetricsGroups = metricsGroup
	data = &spec.GetNodeMetricsData{NodeMetrics: node}
	return
}

func (resolver *Resolver) GetLogs(tctx *logger.TraceContext, input *api_spec.GetLogs, user *base_spec.UserAuthority) (data *api_spec.GetLogsData, code uint8, err error) {
	data, err = resolver.tsdbApi.GetLogs(tctx, input)
	if err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOk
	return
}

func (resolver *Resolver) GetLogParams(tctx *logger.TraceContext, input *api_spec.GetLogParams, user *base_spec.UserAuthority) (data *api_spec.GetLogParamsData, code uint8, err error) {
	data, err = resolver.tsdbApi.GetLogParams(tctx, input)
	if err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOk
	return
}

func (resolver *Resolver) GetEvents(tctx *logger.TraceContext, input *api_spec.GetEvents, user *base_spec.UserAuthority) (data *api_spec.GetEventsData, code uint8, err error) {
	var getIssuedEventsData *api_spec.GetIssuedEventsData
	getIssuedEventsData, err = resolver.tsdbApi.GetIssuedEvents(tctx, &api_spec.GetIssuedEvents{})
	if err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	data = &api_spec.GetEventsData{
		Events: getIssuedEventsData.Events,
	}
	code = base_const.CodeOk
	return
}

func (resolver *Resolver) GetEventRule(tctx *logger.TraceContext, input *api_spec.GetEventRule, user *base_spec.UserAuthority) (data *api_spec.GetEventRuleData, code uint8, err error) {
	var eventRule *spec.EventRule
	if eventRule, err = resolver.dbApi.GetEventRule(tctx, input, user); err != nil {
		if gorm.IsRecordNotFoundError(err) {
			code = base_const.CodeOkNotFound
			return
		}
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOk
	data = &spec.GetEventRuleData{EventRule: *eventRule}
	return
}

func (resolver *Resolver) GetEventRules(tctx *logger.TraceContext, input *api_spec.GetEventRules, user *base_spec.UserAuthority) (data *api_spec.GetEventRulesData, code uint8, err error) {
	var eventRules []spec.EventRule
	if eventRules, err = resolver.dbApi.GetEventRules(tctx, input, user); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOk
	data = &spec.GetEventRulesData{EventRules: eventRules}
	return
}

func (resolver *Resolver) CreateEventRules(tctx *logger.TraceContext, input *api_spec.CreateEventRules, user *base_spec.UserAuthority) (data *api_spec.CreateEventRulesData, code uint8, err error) {
	var specs []spec.EventRule
	if specs, err = resolver.ConvertToEventRuleSpecs(input.Specs); err != nil {
		code = base_const.CodeClientBadRequest
		return
	}
	if err = resolver.dbApi.CreateEventRules(tctx, specs, user); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOkCreated
	data = &spec.CreateEventRulesData{}
	return
}

func (resolver *Resolver) UpdateEventRules(tctx *logger.TraceContext, input *api_spec.UpdateEventRules, user *base_spec.UserAuthority) (data *api_spec.UpdateEventRulesData, code uint8, err error) {
	var specs []spec.EventRule
	if specs, err = resolver.ConvertToEventRuleSpecs(input.Specs); err != nil {
		code = base_const.CodeClientBadRequest
		return
	}
	if err = resolver.dbApi.UpdateEventRules(tctx, specs, user); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOkUpdated
	data = &spec.UpdateEventRulesData{}
	return
}

func (resolver *Resolver) DeleteEventRules(tctx *logger.TraceContext, input *api_spec.DeleteEventRules, user *base_spec.UserAuthority) (data *api_spec.DeleteEventRulesData, code uint8, err error) {
	var specs []spec.EventRule
	if specs, err = resolver.ConvertToEventRuleSpecs(input.Specs); err != nil {
		code = base_const.CodeClientBadRequest
		return
	}
	if err = resolver.dbApi.DeleteEventRules(tctx, specs, user); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOkDeleted
	data = &spec.DeleteEventRulesData{}
	return
}

func (resolver *Resolver) ConvertToEventRuleSpecs(specsStr string) (specs []api_spec.EventRule, err error) {
	var baseSpecs []base_spec.Spec
	if err = json.Unmarshal([]byte(specsStr), &baseSpecs); err != nil {
		return
	}

	for _, base := range baseSpecs {
		if base.Kind != "EventRule" {
			continue
		}
		var specBytes []byte
		if specBytes, err = json.Marshal(base.Spec); err != nil {
			return
		}
		var specData spec.EventRule
		if err = json.Unmarshal(specBytes, &specData); err != nil {
			return
		}
		if err = resolver.Validate.Struct(&specData); err != nil {
			return
		}
		specs = append(specs, specData)
	}
	return
}
