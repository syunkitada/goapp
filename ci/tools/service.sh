#!/bin/bash

COMMAND="${@:-start}"
LOG_DIR=~/.goapp/logs

declare -a SERVICES=("authproxy")
declare -a RESOURCE_SERVICES=("resource-api" "resource-controller" "resource-cluster-api" "resource-cluster-controller" "resource-cluster-agent")
declare -a MONITOR_SERVICES=("monitor-api" "monitor-alert-manager" "monitor-agent")

start_all() {
    for service in ${SERVICES[@]}
    do
        go run cmd/goapp-godo/main.go goapp-${service} --watch &> ${LOG_DIR}/stdout-${service}.log &
        echo "Started goapp-${service}"
    done

    start_resource
}

start_resource() {
    for service in ${RESOURCE_SERVICES[@]}
    do
        go run cmd/goapp-godo/main.go goapp-${service} --watch &> ${LOG_DIR}/stdout-${service}.log &
        echo "Started goapp-${service}"
    done

    echo "If you want to logs, you watch ${LOG_DIR}/*.log"
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
