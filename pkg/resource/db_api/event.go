package db_api

import (
	"fmt"

	"github.com/syunkitada/goapp/pkg/base/base_client"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/error_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

func (api *Api) GetEvents(tctx *logger.TraceContext, input *spec.GetEvents, user *base_spec.UserAuthority) (data *spec.GetEventsData, err error) {
	client, ok := api.clusterClientMap[input.Cluster]
	if !ok {
		err = error_utils.NewNotFoundError("clusterClient")
		return
	}

	queries := []base_client.Query{
		base_client.Query{
			Name: "GetEvents",
			Data: *input,
		},
	}

	getEventsData, tmpErr := client.ResourceVirtualAdminGetEvents(tctx, queries)
	if tmpErr != nil {
		err = fmt.Errorf("Failed GetEvents: %s", tmpErr.Error())
		return
	}
	data = getEventsData
	return
}

func (api *Api) GetEventRule(tctx *logger.TraceContext, input *spec.GetEventRule, user *base_spec.UserAuthority) (data *spec.GetEventRuleData, err error) {
	data = &spec.GetEventRuleData{}
	// client, ok := api.clusterClientMap[input.Cluster]
	// if !ok {
	// 	err = error_utils.NewNotFoundError("clusterClient")
	// 	return
	// }

	// queries := []base_client.Query{
	// 	base_client.Query{
	// 		Name: "GetEventRule",
	// 		Data: *input,
	// 	},
	// }

	// TODO
	// getEventRuleData, tmpErr := client.ResourceVirtualAdminGetEventRule(tctx, queries)
	// if tmpErr != nil {
	// 	err = fmt.Errorf("Failed GetEventRule: %s", tmpErr.Error())
	// 	return
	// }
	// data = getEventRuleData
	return
}

func (api *Api) GetEventRules(tctx *logger.TraceContext, input *spec.GetEventRules, user *base_spec.UserAuthority) (data *spec.GetEventRulesData, err error) {
	client, ok := api.clusterClientMap[input.Cluster]
	if !ok {
		err = error_utils.NewNotFoundError("clusterClient")
		return
	}

	queries := []base_client.Query{
		base_client.Query{
			Name: "GetEventRules",
			Data: *input,
		},
	}

	getEventRulesData, tmpErr := client.ResourceVirtualAdminGetEventRules(tctx, queries)
	if tmpErr != nil {
		err = fmt.Errorf("Failed GetEvents: %s", tmpErr.Error())
		return
	}
	data = getEventRulesData
	return
}

func (api *Api) CreateEventRules(tctx *logger.TraceContext, input *spec.CreateEventRules, user *base_spec.UserAuthority) (data *spec.CreateEventRulesData, err error) {
	client, ok := api.clusterClientMap[input.Cluster]
	if !ok {
		err = error_utils.NewNotFoundError("clusterClient")
		return
	}

	queries := []base_client.Query{
		base_client.Query{
			Name: "CreateEventRules",
			Data: *input,
		},
	}

	createEventRulesData, tmpErr := client.ResourceVirtualAdminCreateEventRules(tctx, queries)
	if tmpErr != nil {
		err = fmt.Errorf("Failed GetEvents: %s", tmpErr.Error())
		return
	}
	data = createEventRulesData
	return
}

func (api *Api) UpdateEventRules(tctx *logger.TraceContext, input *spec.UpdateEventRules, user *base_spec.UserAuthority) (data *spec.UpdateEventRulesData, err error) {
	client, ok := api.clusterClientMap[input.Cluster]
	if !ok {
		err = error_utils.NewNotFoundError("clusterClient")
		return
	}

	queries := []base_client.Query{
		base_client.Query{
			Name: "UpdateEventRules",
			Data: *input,
		},
	}

	updateEventRulesData, tmpErr := client.ResourceVirtualAdminUpdateEventRules(tctx, queries)
	if tmpErr != nil {
		err = fmt.Errorf("Failed GetEvents: %s", tmpErr.Error())
		return
	}
	data = updateEventRulesData
	return
}

func (api *Api) DeleteEventRules(tctx *logger.TraceContext, input *spec.DeleteEventRules, user *base_spec.UserAuthority) (data *spec.DeleteEventRulesData, err error) {
	fmt.Println("DEBUG getEventRules", input)
	client, ok := api.clusterClientMap[input.Cluster]
	if !ok {
		err = error_utils.NewNotFoundError("clusterClient")
		return
	}

	queries := []base_client.Query{
		base_client.Query{
			Name: "DeleteEventRules",
			Data: *input,
		},
	}

	deleteEventRulesData, tmpErr := client.ResourceVirtualAdminDeleteEventRules(tctx, queries)
	if tmpErr != nil {
		err = fmt.Errorf("Failed GetEventRules: %s", tmpErr.Error())
		return
	}
	data = deleteEventRulesData
	return
}
