#!/bin/bash

COMMAND="${@:-start}"
LOG_DIR=~/.goapp/logs

declare -a SERVICES=("authproxy" "resource-api" "resource-controller" "resource-cluster-api" "resource-cluster-controller" "resource-cluster-agent")
declare -a MONITOR_SERVICES=("monitor-proxy" "monitor-alert-manager" "monitor-agent")
declare -a SERVICES2=("resource-api2" "resource-controller2")
declare -a SERVICES3=("resource-api3" "resource-controller3")

start_all() {
    for service in ${SERVICES[@]}
    do
        go run cmd/goapp-godo/main.go goapp-${service} --watch &> ${LOG_DIR}/stdout-${service}.log &
        echo "Started goapp-${service}"
    done

    echo "If you want to logs, you watch /tmp/goapp/logs/*.log"
}

start_monitor() {
    for service in ${MONITOR_SERVICES[@]}
    do
        go run cmd/goapp-godo/main.go goapp-${service} --watch &> ${LOG_DIR}/stdout-${service}.log &
        echo "Started goapp-${service}"
    done

    echo "If you want to logs, you watch ${LOG_DIR}/*.log"
}

start_multi() {
    for service in ${SERVICES[@]}
    do
        go run cmd/goapp-godo/main.go goapp-${service} --watch &> ${LOG_DIR}/stdout-${service}.log &
        echo "Started goapp-${service}"
    done

    for service in ${SERVICES2[@]}
    do
        go run cmd/goapp-godo/main.go goapp-${service} --watch &> ${LOG_DIR}/stdout-${service}2.log &
        echo "Started goapp-${service}"
    done

    for service in ${SERVICES3[@]}
    do
        go run cmd/goapp-godo/main.go goapp-${service} --watch &> ${LOG_DIR}/stdout-${service}3.log &
        echo "Started goapp-${service}"
    done

    echo "If you want to logs, you watch .goapp/logs/*.log"
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
