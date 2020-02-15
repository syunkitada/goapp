package resolver

import (
	"fmt"

	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

func (resolver *Resolver) GetEvents(tctx *logger.TraceContext, input *spec.GetEvents, user *base_spec.UserAuthority) (data *spec.GetEventsData, code uint8, err error) {
	if data, err = resolver.dbApi.GetEvents(tctx, input, user); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOk
	return
}

func (resolver *Resolver) GetStatistics(tctx *logger.TraceContext, input *spec.GetStatistics, user *base_spec.UserAuthority) (data *spec.GetStatisticsData, code uint8, err error) {
	code = base_const.CodeOk
	return
}

func (resolver *Resolver) GetTrace(tctx *logger.TraceContext, input *spec.GetTrace, user *base_spec.UserAuthority) (data *spec.GetTraceData, code uint8, err error) {
	code = base_const.CodeOk
	return
}

func (resolver *Resolver) GetLogs(tctx *logger.TraceContext, input *spec.GetLogs, user *base_spec.UserAuthority) (data *spec.GetLogsData, code uint8, err error) {
	if data, err = resolver.dbApi.GetLogs(tctx, input, user); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOk
	return
}

func (resolver *Resolver) GetLogParams(tctx *logger.TraceContext, input *spec.GetLogParams, user *base_spec.UserAuthority) (data *spec.GetLogParamsData, code uint8, err error) {
	if data, err = resolver.dbApi.GetLogParams(tctx, input, user); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOk
	return
}

func (resolver *Resolver) GetEventRule(tctx *logger.TraceContext, input *spec.GetEventRule, user *base_spec.UserAuthority) (data *spec.GetEventRuleData, code uint8, err error) {
	if data, err = resolver.dbApi.GetEventRule(tctx, input, user); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOk
	return
}

func (resolver *Resolver) GetEventRules(tctx *logger.TraceContext, input *spec.GetEventRules, user *base_spec.UserAuthority) (data *spec.GetEventRulesData, code uint8, err error) {
	if data, err = resolver.dbApi.GetEventRules(tctx, input, user); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOk
	return
}

func (resolver *Resolver) CreateEventRules(tctx *logger.TraceContext, input *spec.CreateEventRules, user *base_spec.UserAuthority) (data *spec.CreateEventRulesData, code uint8, err error) {
	fmt.Println("DEBUG CreateEventRules")
	if data, err = resolver.dbApi.CreateEventRules(tctx, input, user); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOk
	return
}

func (resolver *Resolver) UpdateEventRules(tctx *logger.TraceContext, input *spec.UpdateEventRules, user *base_spec.UserAuthority) (data *spec.UpdateEventRulesData, code uint8, err error) {
	if data, err = resolver.dbApi.UpdateEventRules(tctx, input, user); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOk
	return
}

func (resolver *Resolver) DeleteEventRules(tctx *logger.TraceContext, input *spec.DeleteEventRules, user *base_spec.UserAuthority) (data *spec.DeleteEventRulesData, code uint8, err error) {
	if data, err = resolver.dbApi.DeleteEventRules(tctx, input, user); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOk
	return
}
