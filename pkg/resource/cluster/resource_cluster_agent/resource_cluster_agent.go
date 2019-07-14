package resource_cluster_agent

import (
	"fmt"
	"time"

	"github.com/golang/glog"
	"google.golang.org/grpc"

	"github.com/syunkitada/goapp/pkg/base"
	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_agent/compute_drivers"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_agent/metrics_plugins"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_agent/metrics_plugins/system_metrics_reader"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_agent/resource_cluster_agent_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_api/resource_cluster_api_client"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_model/resource_cluster_model_api"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

type ResourceClusterAgentServer struct {
	base.BaseApp
	conf                        *config.Config
	cluster                     *config.ResourceClusterConfig
	resourceClusterModelApi     *resource_cluster_model_api.ResourceClusterModelApi
	resourceClusterApiClient    *resource_cluster_api_client.ResourceClusterApiClient
	role                        string
	labels                      []string
	resourceLabels              []string
	metricsReaderMap            map[string]metrics_plugins.MetricsReader
	computeDriver               compute_drivers.ComputeDriver
	computeConfirmRetryCount    int
	computeConfirmRetryInterval time.Duration
}

func NewResourceClusterAgentServer(conf *config.Config) *ResourceClusterAgentServer {
	cluster, ok := conf.Resource.ClusterMap[conf.Resource.Node.ClusterName]
	if !ok {
		glog.Fatal(fmt.Errorf("Cluster(%v) is not found in ClusterMap", conf.Resource.Node.ClusterName))
	}

	resourceLabels := []string{}
	if conf.Resource.Node.Compute.Enable {
		resourceLabels = append(resourceLabels, resource_model.ResourceKindCompute)
	}

	fmt.Println("DEBUG CPU", conf.Resource.Node.Compute.Libvirt.AvailableCpus)

	computeDriver := compute_drivers.Load(conf)

	metricsReaderMap := map[string]metrics_plugins.MetricsReader{}
	metricsReaderMap["system"] = system_metrics_reader.New(&conf.Resource.Node.Metrics.System)

	cluster.AgentApp.Name = "resource.cluster.agent"
	server := ResourceClusterAgentServer{
		BaseApp:                     base.NewBaseApp(conf, &cluster.AgentApp),
		conf:                        conf,
		resourceClusterModelApi:     resource_cluster_model_api.NewResourceClusterModelApi(conf),
		resourceClusterApiClient:    resource_cluster_api_client.NewResourceClusterApiClient(conf, nil),
		labels:                      cluster.AgentApp.Labels,
		resourceLabels:              resourceLabels,
		metricsReaderMap:            metricsReaderMap,
		computeDriver:               computeDriver,
		computeConfirmRetryCount:    conf.Resource.Node.Compute.ConfirmRetryCount,
		computeConfirmRetryInterval: time.Duration(conf.Resource.Node.Compute.ConfirmRetryInterval) * time.Second,
	}

	server.RegisterDriver(&server)
	return &server
}

func (srv *ResourceClusterAgentServer) RegisterGrpcServer(grpcServer *grpc.Server) error {
	resource_cluster_agent_grpc_pb.RegisterResourceClusterAgentServer(grpcServer, srv)
	return nil
}
