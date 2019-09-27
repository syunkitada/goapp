package resolver

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/spec"
)

func (resolver *Resolver) GetCluster(tctx *logger.TraceContext, input *spec.GetCluster) (data *spec.GetClusterData, code uint8, err error) {
	var cluster *spec.Cluster
	if cluster, err = resolver.dbApi.GetCluster(tctx, input); err != nil {
		if gorm.IsRecordNotFoundError(err) {
			code = base_const.CodeOkNotFound
			return
		}
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOk
	data = &spec.GetClusterData{Cluster: *cluster}
	return
}

func (resolver *Resolver) GetClusters(tctx *logger.TraceContext, input *spec.GetClusters) (data *spec.GetClustersData, code uint8, err error) {
	var clusters []spec.Cluster
	if clusters, err = resolver.dbApi.GetClusters(tctx, input); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOk
	data = &spec.GetClustersData{Clusters: clusters}
	return
}

func (resolver *Resolver) CreateCluster(tctx *logger.TraceContext, input *spec.CreateCluster) (data *spec.CreateClusterData, code uint8, err error) {
	var specs []spec.Cluster
	if specs, err = resolver.ConvertToClusterSpecs(input.Spec); err != nil {
		code = base_const.CodeClientBadRequest
		return
	}
	if err = resolver.dbApi.CreateClusters(tctx, specs); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOkCreated
	data = &spec.CreateClusterData{}
	return
}

func (resolver *Resolver) UpdateCluster(tctx *logger.TraceContext, input *spec.UpdateCluster) (data *spec.UpdateClusterData, code uint8, err error) {
	var specs []spec.Cluster
	if specs, err = resolver.ConvertToClusterSpecs(input.Spec); err != nil {
		code = base_const.CodeClientBadRequest
		return
	}
	if err = resolver.dbApi.UpdateClusters(tctx, specs); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOkUpdated
	data = &spec.UpdateClusterData{}
	return
}

func (resolver *Resolver) DeleteCluster(tctx *logger.TraceContext, input *spec.DeleteCluster) (data *spec.DeleteClusterData, code uint8, err error) {
	if err = resolver.dbApi.DeleteCluster(tctx, input); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOkDeleted
	data = &spec.DeleteClusterData{}
	return
}

func (resolver *Resolver) DeleteClusters(tctx *logger.TraceContext, input *spec.DeleteClusters) (data *spec.DeleteClustersData, code uint8, err error) {
	var specs []spec.Cluster
	if specs, err = resolver.ConvertToClusterSpecs(input.Spec); err != nil {
		code = base_const.CodeClientBadRequest
		return
	}
	if err = resolver.dbApi.DeleteClusters(tctx, specs); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOkDeleted
	data = &spec.DeleteClustersData{}
	return
}

func (resolver *Resolver) ConvertToClusterSpecs(specStr string) (data []spec.Cluster, err error) {
	var baseSpecs []base_spec.Spec
	if err = json.Unmarshal([]byte(specStr), &baseSpecs); err != nil {
		return
	}

	specs := []spec.Cluster{}
	for _, base := range baseSpecs {
		if base.Kind != "Cluster" {
			continue
		}
		var specBytes []byte
		if specBytes, err = json.Marshal(base.Spec); err != nil {
			return
		}
		var specData spec.Cluster
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
