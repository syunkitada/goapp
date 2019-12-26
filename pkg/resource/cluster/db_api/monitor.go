package db_api

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/json_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/db_model"
	api_spec "github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

func (api *Api) MonitorEvents(tctx *logger.TraceContext) (err error) {
	fmt.Println("MonitorEvents")
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

	// TODO Ignore Events

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
		if tmpErr := api.tsdbApi.IssueEvent(tctx, &api_spec.IssueEvent{Event: event}); tmpErr != nil {
			logger.Warningf(tctx, "Failed IssueEvent: %s", tmpErr.Error())
		}
	}

	// TODO Handle Event Queue For Aggregation

	return
}

func (api *Api) GetEventRule(tctx *logger.TraceContext, input *api_spec.GetEventRule, user *base_spec.UserAuthority) (data *api_spec.EventRule, err error) {
	data = &api_spec.EventRule{}
	err = api.DB.Where("name = ? AND project = ? AND deleted_at IS NULL", input.Name, user.ProjectName).First(data).Error
	return
}

func (api *Api) GetEventRules(tctx *logger.TraceContext, input *api_spec.GetEventRules, user *base_spec.UserAuthority) (data []api_spec.EventRule, err error) {
	err = api.DB.Debug().Table("event_rules").Select("name, kind, `check`, level, project, node, msg, until").
		Where("project = ? AND deleted_at IS NULL", user.ProjectName).Scan(&data).Error
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
