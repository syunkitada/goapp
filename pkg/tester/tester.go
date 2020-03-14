package tester

import (
	"testing"

	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/tester/config"

	authproxy_api_server "github.com/syunkitada/goapp/pkg/authproxy/authproxy_api/server"
	authproxy_config "github.com/syunkitada/goapp/pkg/authproxy/config"
	authproxy_ctl "github.com/syunkitada/goapp/pkg/authproxy/ctl"

	resource_cluster_api_server "github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_api/server"
	resource_cluster_controller_server "github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_controller/server"
	resource_config "github.com/syunkitada/goapp/pkg/resource/config"
	resource_ctl "github.com/syunkitada/goapp/pkg/resource/ctl"
	resource_api_server "github.com/syunkitada/goapp/pkg/resource/resource_api/server"
	resource_controller_server "github.com/syunkitada/goapp/pkg/resource/resource_controller/server"
)

type Tester struct {
	t        *testing.T
	baseConf *base_config.Config
	mainConf *config.Config

	authproxyCtl       *authproxy_ctl.Ctl
	authproxyApiServer *authproxy_api_server.Server

	resourceCtl                  *resource_ctl.Ctl
	resourceApiServer            *resource_api_server.Server
	resourceControllerServer     *resource_controller_server.Server
	resourceClusterApiMap        map[string]*resource_cluster_api_server.Server
	resourceClusterControllerMap map[string]*resource_cluster_controller_server.Server
}

func NewTester(t *testing.T, baseConf *base_config.Config, mainConf *config.Config) *Tester {
	var err error
	authproxyConf := &authproxy_config.Config{Authproxy: mainConf.Authproxy}
	authproxyCtl := authproxy_ctl.NewCtl(baseConf, authproxyConf)
	if err = authproxyCtl.Bootstrap(true); err != nil {
		t.Fatalf("Failed authproxyCtl.Bootstrap: err=%s\n", err.Error())
	}
	authproxyApiServer := authproxy_api_server.New(baseConf, authproxyConf)

	resourceConf := &resource_config.Config{Resource: mainConf.Resource}
	resourceCtl := resource_ctl.NewCtl(baseConf, resourceConf)
	if err = resourceCtl.Bootstrap(true); err != nil {
		t.Fatalf("Failed resourceCtl.Bootstrap: err=%s\n", err.Error())
	}
	resourceClusterApiMap := map[string]*resource_cluster_api_server.Server{}
	resourceClusterControllerMap := map[string]*resource_cluster_controller_server.Server{}
	for clusterName := range mainConf.Resource.ClusterMap {
		clusterMainConf := *resourceConf
		clusterMainConf.Resource.ClusterName = clusterName
		resourceClusterApi := resource_cluster_api_server.New(baseConf, &clusterMainConf)
		resourceClusterController := resource_cluster_controller_server.New(baseConf, &clusterMainConf)
		resourceClusterApiMap[clusterName] = resourceClusterApi
		resourceClusterControllerMap[clusterName] = resourceClusterController
	}

	resourceApiServer := resource_api_server.New(baseConf, resourceConf)
	resourceControllerServer := resource_controller_server.New(baseConf, resourceConf)

	tester := Tester{
		t:        t,
		baseConf: baseConf,
		mainConf: mainConf,

		authproxyCtl:       authproxyCtl,
		authproxyApiServer: authproxyApiServer,

		resourceCtl:                  resourceCtl,
		resourceApiServer:            resourceApiServer,
		resourceControllerServer:     resourceControllerServer,
		resourceClusterApiMap:        resourceClusterApiMap,
		resourceClusterControllerMap: resourceClusterControllerMap,
	}

	return &tester
}
