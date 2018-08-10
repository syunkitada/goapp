compile-pb:
	protoc -I pkg/health/grpc_pb/ pkg/health/grpc_pb/grpc_pb.proto --go_out=plugins=grpc:pkg/health/grpc_pb
