package db_api

import (
	"fmt"
	"time"

	"github.com/syunkitada/goapp/pkg/lib/logger"
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

	// TODO Filter Ignore Events

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
		switch event.Handler {
		case "Mail":
			fmt.Println("TODO Add Event To SendMail Queue")
		case "Alert":
			fmt.Println("TODO Add Event To Alert Queue")
		}
		if tmpErr := api.tsdbApi.IssueEvent(tctx, &api_spec.IssueEvent{Event: event}); tmpErr != nil {
			logger.Warningf(tctx, "Failed IssueEvent: %s", tmpErr.Error())
		}
	}

	// TODO Handle Event Queue For Aggregation

	return
}
