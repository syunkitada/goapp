package base_app

import (
	"encoding/json"
	"strings"

	"github.com/syunkitada/goapp/pkg/base/base_client"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/base/base_spec_model"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

func (app *BaseApp) SyncService(tctx *logger.TraceContext, syncRoot bool) (err error) {
	queries := []base_client.Query{}
	var data *base_spec.GetServicesData
	if data, err = app.dbApi.GetServices(tctx, &base_spec.GetServices{}); err != nil {
		return
	}

	serviceMap := map[string]base_spec_model.ServiceRouter{}
	for _, service := range data.Services {
		var queryMap map[string]base_spec_model.QueryModel
		if err = json.Unmarshal([]byte(service.QueryMap), &queryMap); err != nil {
			return
		}
		serviceMap[service.Name] = base_spec_model.ServiceRouter{
			Token:     service.Token,
			Endpoints: strings.Split(service.Endpoints, ","),
			QueryMap:  queryMap,
		}

		if service.SyncRootCluster {
			var token string
			token, err = app.dbApi.IssueToken(service.Name)
			if err != nil {
				return
			}
			queries = append(queries, base_client.Query{
				Name: "UpdateService",
				Data: base_spec.UpdateService{
					Name:            service.Name,
					Token:           token,
					Scope:           service.Scope,
					Endpoints:       app.appConf.Endpoints,
					ProjectRoles:    strings.Split(service.ProjectRoles, ","),
					QueryMap:        queryMap,
					SyncRootCluster: false,
				},
			})
		}
	}
	app.serviceMap = serviceMap

	if syncRoot && len(queries) > 0 {
		if _, err = app.rootClient.UpdateServices(tctx, queries); err != nil {
			return
		}
	}
	return
}
