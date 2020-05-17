env:
	ci/tools/setup-env.sh

mod:
	GO111MODULE=on; go mod tidy; go mod vendor;

gen:
	go run cmd/goapp-authproxy/main.go ctl generate-code
	go run cmd/goapp-home/main.go ctl generate-code
	go run cmd/goapp-resource/main.go ctl generate-code

bootstrap:
	go run cmd/goapp-authproxy/main.go ctl bootstrap
	go run cmd/goapp-home/main.go ctl bootstrap
	go run cmd/goapp-resource/main.go ctl bootstrap

# tests
lint-test:
	golangci-lint run pkg/...
unit-test:
	go test --cover ./pkg/...
unit-test-v:
	go test -v --cover ./pkg/...
senario-test:
	go test --parallel 1 ./ci/senario_test/...
senario-test-v:
	go test -v --parallel 1 ./ci/senario_test/...

# services
start-all:
	ci/tools/service.sh start_docker_services
	ci/tools/service.sh start_authproxy_services
	ci/tools/service.sh start_home_services
	ci/tools/service.sh start_resource_services
stop-all:
	ci/tools/service.sh stop_docker_services
	ci/tools/service.sh stop_authproxy_services
	ci/tools/service.sh stop_home_services
	ci/tools/service.sh stop_resource_services
status:
	ci/tools/service.sh status

start-docker-services:
	ci/tools/service.sh start_docker_services
start-authproxy-services:
	ci/tools/service.sh start_authproxy_services
stop-authproxy-services:
	ci/tools/service.sh stop_authproxy_services
start-home-services:
	ci/tools/service.sh start_home_services
stop-home-services:
	ci/tools/service.sh stop_home_services
start-resource-services:
	ci/tools/service.sh start_resource_services
stop-resource-services:
	ci/tools/service.sh stop_resource_services
