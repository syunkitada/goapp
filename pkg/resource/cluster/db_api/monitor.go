package db_api

import (
	"encoding/json"
	"fmt"
	"regexp"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/json_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/db_model"
	api_spec "github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

type EventRule struct {
	ReNode         *regexp.Regexp
	ReMsg          *regexp.Regexp
	ReCheck        *regexp.Regexp
	ReLevel        *regexp.Regexp
	Until          *time.Time
	AggregateNode  bool
	AggregateCheck bool
	Priority       int
	Actions        []string
	ContinueNext   bool
}

func (api *Api) MonitorEvents(tctx *logger.TraceContext) (err error) {
	fmt.Println("MonitorEvents")
	var eventRules []db_model.EventRule
	eventRules, err = api.GetControllerEventRules(tctx)
	if err != nil {
		return
	}

	var silenceEventRules []EventRule
	var aggregateEventRules []EventRule
	var actionEventRules []EventRule
	for _, eventRule := range eventRules {
		rule := EventRule{}
		if eventRule.Node != "" {
			re, tmpErr := regexp.Compile(eventRule.Node)
			if tmpErr != nil {
				logger.Errorf(tctx, tmpErr, "Failed compile rule: %s", eventRule.Node)
			}
			rule.ReNode = re
		}
		if eventRule.Check != "" {
			re, tmpErr := regexp.Compile(eventRule.Check)
			if tmpErr != nil {
				logger.Errorf(tctx, tmpErr, "Failed compile rule: %s", eventRule.Check)
			}
			rule.ReCheck = re
		}
		if eventRule.Msg != "" {
			re, tmpErr := regexp.Compile(eventRule.Msg)
			if tmpErr != nil {
				logger.Errorf(tctx, tmpErr, "Failed compile rule: %s", eventRule.Msg)
			}
			rule.ReMsg = re
		}
		if eventRule.Level != "" {
			re, tmpErr := regexp.Compile(eventRule.Level)
			if tmpErr != nil {
				logger.Errorf(tctx, tmpErr, "Failed compile rule: %s", eventRule.Level)
			}
			rule.ReLevel = re
		}

		switch eventRule.Kind {
		case "Silence":
			silenceEventRules = append(silenceEventRules, rule)
		case "Aggregate":
			var spec api_spec.EventRuleAggregateSpec
			if tmpErr := json.Unmarshal([]byte(eventRule.Spec), &spec); tmpErr != nil {
				logger.Errorf(tctx, tmpErr, "Failed Unmarshal")
				continue
			}
			rule.AggregateNode = spec.AggregateNode
			rule.AggregateCheck = spec.AggregateCheck
			rule.Priority = spec.Priority
			aggregateEventRules = append(aggregateEventRules, rule)
		case "Action":
			var spec api_spec.EventRuleActionSpec
			if tmpErr := json.Unmarshal([]byte(eventRule.Spec), &spec); tmpErr != nil {
				logger.Errorf(tctx, tmpErr, "Failed Unmarshal")
				continue
			}
			rule.Actions = spec.Actions
			rule.ContinueNext = spec.ContinueNext
			rule.Priority = spec.Priority
			actionEventRules = append(actionEventRules, rule)
		}
	}

	var getEventsData *api_spec.GetEventsData
	var getIssuedEventsData *api_spec.GetIssuedEventsData
	getEventsData, err = api.tsdbApi.GetEvents(tctx, &api_spec.GetEvents{})
	if err != nil {
		return
	}

	getIssuedEventsData, err = api.tsdbApi.GetIssuedEvents(tctx, &api_spec.GetIssuedEvents{})
	if err != nil {
		return
	}

	issuedEventMap := map[string]api_spec.Event{}
	for _, event := range getIssuedEventsData.Events {
		key := event.Node + "@" + event.Check
		issuedEventMap[key] = event
	}

	nodeEventsMap := map[string][]api_spec.Event{}
	checkEventsMap := map[string][]api_spec.Event{}
	nonAggregatedEvents := []api_spec.Event{}

	var silenceEvent bool
	var aggregateEvent bool
	var aggregateNode bool
	var aggregateCheck bool
	for _, event := range getEventsData.Events {
		key := event.Node + "@" + event.Check
		issuedEvent, ok := issuedEventMap[key]
		if ok {
			if event.Level == issuedEvent.Level {
				reissueDuration := time.Duration(event.ReissueDuration) * time.Second
				if issuedEvent.Time.Add(reissueDuration).Before(event.Time) {
					continue
				}
			}
		}

		silenceEvent = false
		for _, rule := range silenceEventRules {
			if rule.ReCheck != nil {
				if rule.ReCheck.MatchString(event.Check) {
					silenceEvent = true
				}
			}
			if rule.ReNode != nil {
				if rule.ReNode.MatchString(event.Node) {
					silenceEvent = true
				} else {
					silenceEvent = false
				}
			}
			if rule.ReMsg != nil {
				if rule.ReMsg.MatchString(event.Msg) {
					silenceEvent = true
				} else {
					silenceEvent = false
				}
			}
			if silenceEvent {
				break
			}
		}
		if silenceEvent {
			fmt.Println("DEBUG sileceneEvent", event.Node, event.Check)
			event.Silenced = 1
			if tmpErr := api.tsdbApi.IssueEvent(tctx, &api_spec.IssueEvent{Event: event}); tmpErr != nil {
				logger.Warningf(tctx, "Failed IssueEvent: %s", tmpErr.Error())
			}
			continue
		}

		aggregateNode = false
		aggregateCheck = false
		for _, rule := range aggregateEventRules {
			if rule.ReCheck == nil && rule.ReNode == nil && rule.ReMsg == nil {
				aggregateEvent = true
			} else {
				if rule.ReCheck != nil {
					if rule.ReCheck.MatchString(event.Check) {
						aggregateEvent = true
					}
				}
				if rule.ReNode != nil {
					if rule.ReNode.MatchString(event.Node) {
						aggregateEvent = true
					} else {
						aggregateEvent = false
					}
				}
				if rule.ReMsg != nil {
					if rule.ReMsg.MatchString(event.Msg) {
						aggregateEvent = true
					} else {
						aggregateEvent = false
					}
				}
			}
			if aggregateEvent {
				aggregateNode = rule.AggregateNode
				aggregateCheck = rule.AggregateCheck
				break
			}
		}
		if aggregateNode {
			events, ok := nodeEventsMap[event.Node]
			if !ok {
				events = []api_spec.Event{}
			}
			events = append(events, event)
			nodeEventsMap[event.Node] = events
		} else if aggregateCheck {
			events, ok := checkEventsMap[event.Check]
			if !ok {
				events = []api_spec.Event{}
			}
			events = append(events, event)
			checkEventsMap[event.Check] = events
		} else {
			nonAggregatedEvents = append(nonAggregatedEvents, event)
		}
	}

	var actionEvent bool
	for _, event := range nonAggregatedEvents {
		for i, rule := range actionEventRules {
			actionEvent = false
			for _, rule := range actionEventRules {
				if rule.ReCheck != nil {
					if rule.ReCheck.MatchString(event.Check) {
						actionEvent = true
					}
				}
				if rule.ReNode != nil {
					if rule.ReNode.MatchString(event.Node) {
						actionEvent = true
					} else {
						actionEvent = false
					}
				}
				if rule.ReMsg != nil {
					if rule.ReMsg.MatchString(event.Msg) {
						actionEvent = true
					} else {
						actionEvent = false
					}
				}
				if actionEvent {
					break
				}
			}
			if actionEvent {
				fmt.Println("Exec handler", rule.Actions, []api_spec.Event{event})
			}
			if !rule.ContinueNext || i == len(actionEventRules)-1 {
				if tmpErr := api.tsdbApi.IssueEvent(tctx, &api_spec.IssueEvent{Event: event}); tmpErr != nil {
					logger.Warningf(tctx, "Failed IssueEvent: %s", tmpErr.Error())
				}
				break
			}
		}
	}

	for node, events := range nodeEventsMap {
		for i, rule := range actionEventRules {
			actionEvent = false
			for _, rule := range actionEventRules {
				if rule.ReCheck != nil {
					for _, event := range events {
						if rule.ReCheck.MatchString(event.Check) {
							actionEvent = true
							break
						}
					}
				}
				if rule.ReNode != nil {
					if rule.ReNode.MatchString(node) {
						actionEvent = true
					} else {
						actionEvent = false
					}
				}
				if rule.ReMsg != nil {
					for _, event := range events {
						if rule.ReMsg.MatchString(event.Msg) {
							actionEvent = true
							break
						}
					}
				}
				if actionEvent {
					break
				}
			}
			if actionEvent {
				fmt.Println("Exec handler aggregated by node", rule.Actions, events)
			}
			if !rule.ContinueNext || i == len(actionEventRules)-1 {
				for _, event := range events {
					if tmpErr := api.tsdbApi.IssueEvent(tctx, &api_spec.IssueEvent{Event: event}); tmpErr != nil {
						logger.Warningf(tctx, "Failed IssueEvent: %s", tmpErr.Error())
					}
				}
				break
			}
		}
	}

	for check, events := range checkEventsMap {
		for i, rule := range actionEventRules {
			actionEvent = false
			for _, rule := range actionEventRules {
				if rule.ReCheck != nil {
					if rule.ReCheck.MatchString(check) {
						actionEvent = true
					} else {
						actionEvent = false
					}
				}
				if rule.ReNode != nil {
					for _, event := range events {
						if rule.ReNode.MatchString(event.Node) {
							actionEvent = true
							break
						}
					}
				}
				if rule.ReMsg != nil {
					for _, event := range events {
						if rule.ReMsg.MatchString(event.Msg) {
							actionEvent = true
							break
						}
					}
				}
				if actionEvent {
					break
				}
			}
			if actionEvent {
				fmt.Println("Exec handler aggregated by check", rule.Actions, events)
			}
			if !rule.ContinueNext || i == len(actionEventRules)-1 {
				for _, event := range events {
					if tmpErr := api.tsdbApi.IssueEvent(tctx, &api_spec.IssueEvent{Event: event}); tmpErr != nil {
						logger.Warningf(tctx, "Failed IssueEvent: %s", tmpErr.Error())
					}
				}
				break
			}
		}
	}

	return
}

func (api *Api) GetEventRule(tctx *logger.TraceContext, input *api_spec.GetEventRule, user *base_spec.UserAuthority) (data *api_spec.EventRule, err error) {
	data = &api_spec.EventRule{}
	err = api.DB.Where("name = ? AND project = ? AND deleted_at IS NULL", input.Name, user.ProjectName).First(data).Error
	return
}

func (api *Api) GetEventRules(tctx *logger.TraceContext, input *api_spec.GetEventRules, user *base_spec.UserAuthority) (data []api_spec.EventRule, err error) {
	err = api.DB.Table("event_rules").Select("name, kind, `check`, level, project, node, msg, until").
		Where("project = ? AND deleted_at IS NULL", user.ProjectName).Scan(&data).Error
	return
}

func (api *Api) GetFilterEventRules(tctx *logger.TraceContext) (data []db_model.EventRule, err error) {
	err = api.DB.Table("event_rules").Select("name, kind, `check`, level, project, node, msg, until").
		Where("kind = ? AND deleted_at IS NULL", "Filter").Scan(&data).Error
	return
}

func (api *Api) GetControllerEventRules(tctx *logger.TraceContext) (data []db_model.EventRule, err error) {
	err = api.DB.Table("event_rules").Select("name, kind, `check`, level, project, node, msg, until").
		Where("kind != ? AND deleted_at IS NULL", "Filter").Scan(&data).Error
	return
}

func (api *Api) CreateEventRules(tctx *logger.TraceContext, input []api_spec.EventRule, user *base_spec.UserAuthority) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		for _, val := range input {
			var specBytes []byte
			if specBytes, err = json_utils.Marshal(val.Spec); err != nil {
				return
			}
			var tmp db_model.EventRule
			if err = tx.Where("name = ? AND project = ?",
				val.Name, user.ProjectName).
				First(&tmp).Error; err != nil {
				if !gorm.IsRecordNotFoundError(err) {
					return
				}
				tmp = db_model.EventRule{
					Name:    val.Name,
					Project: user.ProjectName,
					Kind:    val.Kind,
					Node:    val.Node,
					Msg:     val.Msg,
					Check:   val.Check,
					Level:   val.Level,
					Until:   val.Until,
					Spec:    string(specBytes),
				}
				if err = tx.Create(&tmp).Error; err != nil {
					fmt.Println("HOGE piyo", specBytes)
					return
				}
				fmt.Println("HOGE piyo2")
			}
		}
		return
	})
	return
}

func (api *Api) UpdateEventRules(tctx *logger.TraceContext, input []api_spec.EventRule, user *base_spec.UserAuthority) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		for _, val := range input {
			var specBytes []byte
			if specBytes, err = json_utils.Marshal(val.Spec); err != nil {
				return
			}
			if err = tx.Model(&db_model.EventRule{}).
				Where("name = ? AND project = ?", val.Name, user.ProjectName).
				Updates(&db_model.EventRule{
					Spec: string(specBytes),
				}).Error; err != nil {
				return
			}
		}
		return
	})
	return
}

func (api *Api) DeleteEventRules(tctx *logger.TraceContext, input []api_spec.EventRule, user *base_spec.UserAuthority) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		for _, val := range input {
			if err = tx.Delete(db_model.EventRule{
				Name:    val.Name,
				Project: user.ProjectName,
			}).Error; err != nil {
				return
			}
		}
		return
	})
	return
}
