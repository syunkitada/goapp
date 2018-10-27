#!/bin/bash

COMMAND="${@:-start}"

declare -a SERVICES=("authproxy" "resource-api" "resource-controller")

start_all() {
    mkdir -p /tmp/goapp/logs
    for service in ${SERVICES[@]}
    do
        go run cmd/goapp-godo/main.go goapp-${service} --watch &> /tmp/goapp/logs/${service}.log &
        echo "Started goapp-${service}"
    done

    echo "If you want to logs, you watch /tmp/goapp/logs/*.log"
}

stop_all() {
    for pid in `ps ax | grep goapp | grep -v grep | grep -v vim | awk '{print $1}'`
    do
        kill $pid
    done
}

status() {
    ps ax | grep "go run " | grep -v grep || echo "NotFound Processes"
}

$COMMAND
