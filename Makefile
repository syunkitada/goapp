env:
	ci/tools/setup-env.sh

mod:
	GO111MODULE=on; go mod tidy; go mod vendor;

gen:
	go run cmd/goapp-resource/main.go ctl generate-code
	go run cmd/goapp-authproxy/main.go ctl generate-code

lint-test:
	golangci-lint run pkg/...

# services
start-all:
	ci/tools/service.sh start_docker_services
	ci/tools/service.sh start_authproxy_services
	ci/tools/service.sh start_resource_services
stop-all:
	ci/tools/service.sh stop_docker_services
	ci/tools/service.sh stop_authproxy_services
	ci/tools/service.sh stop_resource_services
status:
	ci/tools/service.sh status
start-docker-services:
	ci/tools/service.sh start_docker_services
start-authproxy-services:
	ci/tools/service.sh start_authproxy_services
stop-authproxy-services:
	ci/tools/service.sh stop_authproxy_services
start-resource-services:
	ci/tools/service.sh start_resource_services
stop-resource-services:
	ci/tools/service.sh stop_resource_services
