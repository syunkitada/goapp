env:
	ci/tools/setup-env.sh	

compile-pb:
	protoc -I /usr/local/include/ -I pkg/ pkg/authproxy/authproxy_grpc_pb/authproxy_grpc_pb.proto --go_out=plugins=grpc:${GOPATH}/src

	protoc -I /usr/local/include/ -I pkg/monitor/monitor_api/monitor_api_grpc_pb/ pkg/monitor/monitor_api/monitor_api_grpc_pb/monitor_api_grpc_pb.proto --go_out=plugins=grpc:pkg/monitor/monitor_api/monitor_api_grpc_pb
	protoc -I /usr/local/include/ -I pkg/monitor/monitor_agent/monitor_agent_grpc_pb/ pkg/monitor/monitor_agent/monitor_agent_grpc_pb/monitor_agent_grpc_pb.proto --go_out=plugins=grpc:pkg/monitor/monitor_agent/monitor_agent_grpc_pb
	protoc -I /usr/local/include/ -I pkg/monitor/monitor_alert_manager/monitor_alert_manager_grpc_pb/ pkg/monitor/monitor_alert_manager/monitor_alert_manager_grpc_pb/monitor_alert_manager_grpc_pb.proto --go_out=plugins=grpc:pkg/monitor/monitor_alert_manager/monitor_alert_manager_grpc_pb

	protoc -I /usr/local/include/ -I pkg/ resource/resource_api/resource_api_grpc_pb/resource_api_grpc_pb.proto --go_out=plugins=grpc:${GOPATH}/src

	protoc -I pkg/resource/resource_controller/resource_controller_grpc_pb/ pkg/resource/resource_controller/resource_controller_grpc_pb/resource_controller_grpc_pb.proto --go_out=plugins=grpc:pkg/resource/resource_controller/resource_controller_grpc_pb
	protoc -I pkg/resource/cluster/resource_cluster_api/resource_cluster_api_grpc_pb/ pkg/resource/cluster/resource_cluster_api/resource_cluster_api_grpc_pb/resource_cluster_api_grpc_pb.proto --go_out=plugins=grpc:pkg/resource/cluster/resource_cluster_api/resource_cluster_api_grpc_pb
	protoc -I pkg/resource/cluster/resource_cluster_controller/resource_cluster_controller_grpc_pb/ pkg/resource/cluster/resource_cluster_controller/resource_cluster_controller_grpc_pb/resource_cluster_controller_grpc_pb.proto --go_out=plugins=grpc:pkg/resource/cluster/resource_cluster_controller/resource_cluster_controller_grpc_pb

start-all:
	ci/tools/service.sh start_all

start-monitor:
	ci/tools/service.sh start_monitor

start-multi:
	ci/tools/service.sh start_multi

stop-all:
	ci/tools/service.sh stop_all

status:
	ci/tools/service.sh status
