package resource_cluster_agent

// import (
// 	"fmt"
// 	"net"
// 	"path/filepath"
// 	"time"
//
// 	"google.golang.org/grpc"
//
// 	"github.com/syunkitada/goapp/pkg/base"
// 	"github.com/syunkitada/goapp/pkg/config"
// 	"github.com/syunkitada/goapp/pkg/lib/logger"
// 	"github.com/syunkitada/goapp/pkg/lib/os_utils"
// 	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_agent/compute_drivers"
// 	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_agent/metrics_plugins"
// 	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_agent/metrics_plugins/system_metrics_reader"
// 	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_agent/resource_cluster_agent_grpc_pb"
// 	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_api/resource_cluster_api_client"
// 	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_model_api"
// 	"github.com/syunkitada/goapp/pkg/resource/resource_model"
// )
//
// type ResourceClusterAgentServer struct {
// 	base.BaseApp
// 	conf                        *config.Config
// 	cluster                     *config.ResourceClusterConfig
// 	resourceClusterModelApi     *resource_cluster_model_api.ResourceClusterModelApi
// 	resourceClusterApiClient    *resource_cluster_api_client.ResourceClusterApiClient
// 	role                        string
// 	labels                      []string
// 	resourceLabels              []string
// 	metricsReaderMap            map[string]metrics_plugins.MetricsReader
// 	computeDriver               compute_drivers.ComputeDriver
// 	computeConfirmRetryCount    int
// 	computeConfirmRetryInterval time.Duration
// 	computeVmsDir               string
// 	computeImagesDir            string
//
// 	vmNetnsServiceIp      net.IP
// 	vmNetnsGatewayStartIp net.IP
// 	vmNetnsGatewayEndIp   net.IP
// 	vmNetnsStartIp        net.IP
// 	vmNetnsEndIp          net.IP
// }
//
// func NewResourceClusterAgentServer(conf *config.Config) *ResourceClusterAgentServer {
// 	cluster, ok := conf.Resource.ClusterMap[conf.Resource.Node.ClusterName]
// 	if !ok {
// 		logger.StdoutFatalf("Cluster(%v) is not found in ClusterMap", conf.Resource.Node.ClusterName)
// 	}
//
// 	resourceLabels := []string{}
// 	if conf.Resource.Node.Compute.Enable {
// 		resourceLabels = append(resourceLabels, resource_model.ResourceKindCompute)
// 	}
//
// 	if conf.Resource.Node.Compute.VmNetnsGatewayStartIp == "" {
// 		conf.Resource.Node.Compute.VmNetnsGatewayStartIp = "169.254.1.1"
// 	}
// 	vmNetnsGatewayStartIp := net.ParseIP(conf.Resource.Node.Compute.VmNetnsGatewayStartIp)
//
// 	if conf.Resource.Node.Compute.VmNetnsGatewayEndIp == "" {
// 		conf.Resource.Node.Compute.VmNetnsGatewayEndIp = "169.254.1.100"
// 	}
// 	vmNetnsGatewayEndIp := net.ParseIP(conf.Resource.Node.Compute.VmNetnsGatewayEndIp)
//
// 	if conf.Resource.Node.Compute.VmNetnsServiceIp == "" {
// 		conf.Resource.Node.Compute.VmNetnsServiceIp = "169.254.1.200"
// 	}
// 	vmNetnsServiceIp := net.ParseIP(conf.Resource.Node.Compute.VmNetnsServiceIp)
//
// 	if conf.Resource.Node.Compute.VmNetnsStartIp == "" {
// 		conf.Resource.Node.Compute.VmNetnsStartIp = "169.254.32.1"
// 	}
// 	vmNetnsStartIp := net.ParseIP(conf.Resource.Node.Compute.VmNetnsStartIp)
//
// 	if conf.Resource.Node.Compute.VmNetnsEndIp == "" {
// 		conf.Resource.Node.Compute.VmNetnsEndIp = "169.254.63.254"
// 	}
// 	vmNetnsEndIp := net.ParseIP(conf.Resource.Node.Compute.VmNetnsEndIp)
//
// 	conf.Resource.Node.Compute.ConfigDir = conf.Path("resource/compute")
//
// 	fmt.Println("DEBUG CPU", conf.Resource.Node.Compute.Libvirt.AvailableCpus)
// 	computeDir := filepath.Join(conf.Default.VarDir, "compute")
// 	computeVmsDir := filepath.Join(computeDir, "vms")
// 	computeImagesDir := filepath.Join(computeDir, "images")
// 	os_utils.MustMkdir(computeDir, 0755)
// 	os_utils.MustMkdir(computeVmsDir, 0755)
// 	os_utils.MustMkdir(computeImagesDir, 0755)
// 	conf.Resource.Node.Compute.VarDir = computeDir
// 	conf.Resource.Node.Compute.VmsDir = computeVmsDir
// 	conf.Resource.Node.Compute.ImagesDir = computeImagesDir
//
// 	computeDriver := compute_drivers.Load(conf)
//
// 	metricsReaderMap := map[string]metrics_plugins.MetricsReader{}
// 	metricsReaderMap["system"] = system_metrics_reader.New(&conf.Resource.Node.Metrics.System)
//
// 	cluster.AgentApp.Name = "resource.cluster.agent"
// 	server := ResourceClusterAgentServer{
// 		BaseApp:                     base.NewBaseApp(conf, &cluster.AgentApp),
// 		conf:                        conf,
// 		resourceClusterModelApi:     resource_cluster_model_api.NewResourceClusterModelApi(conf),
// 		resourceClusterApiClient:    resource_cluster_api_client.NewResourceClusterApiClient(conf, nil),
// 		labels:                      cluster.AgentApp.Labels,
// 		resourceLabels:              resourceLabels,
// 		metricsReaderMap:            metricsReaderMap,
// 		computeDriver:               computeDriver,
// 		computeConfirmRetryCount:    conf.Resource.Node.Compute.ConfirmRetryCount,
// 		computeConfirmRetryInterval: time.Duration(conf.Resource.Node.Compute.ConfirmRetryInterval) * time.Second,
// 		computeVmsDir:               computeVmsDir,
// 		computeImagesDir:            computeImagesDir,
//
// 		vmNetnsServiceIp:      vmNetnsServiceIp,
// 		vmNetnsGatewayStartIp: vmNetnsGatewayStartIp,
// 		vmNetnsGatewayEndIp:   vmNetnsGatewayEndIp,
// 		vmNetnsStartIp:        vmNetnsStartIp,
// 		vmNetnsEndIp:          vmNetnsEndIp,
// 	}
//
// 	server.RegisterDriver(&server)
//
// 	return &server
// }
//
// func (srv *ResourceClusterAgentServer) RegisterGrpcServer(grpcServer *grpc.Server) error {
// 	resource_cluster_agent_grpc_pb.RegisterResourceClusterAgentServer(grpcServer, srv)
// 	return nil
// }
