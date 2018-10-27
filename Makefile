compile-pb:
	protoc -I pkg/resource/resource_api/resource_api_grpc_pb/ pkg/resource/resource_api/resource_api_grpc_pb/resource_api_grpc_pb.proto --go_out=plugins=grpc:pkg/resource/resource_api/resource_api_grpc_pb
	protoc -I pkg/resource/resource_controller/resource_controller_grpc_pb/ pkg/resource/resource_controller/resource_controller_grpc_pb/resource_controller_grpc_pb.proto --go_out=plugins=grpc:pkg/resource/resource_controller/resource_controller_grpc_pb
	# protoc -I pkg/resource/region/resource_region_grpc_pb/ pkg/resource/region/resource_region_grpc_pb/resource_region_grpc_pb.proto --go_out=plugins=grpc:pkg/resource/region/resource_region_grpc_pb

start-all:
	ci/tools/service.sh start_all

stop-all:
	ci/tools/service.sh stop_all

status:
	ci/tools/service.sh status
