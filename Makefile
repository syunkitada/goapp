compile-pb:
	protoc -I pkg/health/grpc_pb/ pkg/health/grpc_pb/grpc_pb.proto --go_out=plugins=grpc:pkg/health/grpc_pb
	protoc -I pkg/resource/grpc_pb/ pkg/resource/grpc_pb/grpc_pb.proto --go_out=plugins=grpc:pkg/resource/grpc_pb
	protoc -I pkg/resource/region/grpc_pb/ pkg/resource/region/grpc_pb/grpc_pb.proto --go_out=plugins=grpc:pkg/resource/region/grpc_pb
