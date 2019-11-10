env:
	ci/tools/setup-env.sh	

gen:
	go run cmd/goapp-resource/main.go ctl generate-code
	go run cmd/goapp-authproxy/main.go ctl generate-code

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
