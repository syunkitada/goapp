#!/bin/bash -e

COMMAND="${@:-start}"
LOG_DIR=~/.goapp/logs

declare -a DOCKER_SERVICES=("mysql" "influxdb")
declare -a AUTHPROXY_SERVICES=("authproxy-api")
declare -a HOME_SERVICES=("home-api" "home-controller")
declare -a RESOURCE_SERVICES=("resource-api" "resource-controller" "resource-cluster-api" "resource-cluster-controller" "resource-cluster-agent")

start_docker_services() {
    for service in ${DOCKER_SERVICES[@]}
    do
        sudo docker ps | grep ${service} || sudo docker start ${service}
        echo "Started ${service}"
    done
}
stop_docker_services() {
    for service in ${DOCKER_SERVICES[@]}
    do
        sudo docker ps | grep ${service} || sudo docker stop ${service}
        echo "Stoped ${service}"
    done
}

stop_service() {
    service=$1
    for pid in `ps ax | grep ${service} | egrep -v 'make|service.sh|grep' | awk '{print $1}'`
    do
        kill -9 $pid
    done
}

start_service() {
    service=$1
    psax=`ps ax`
    echo $psax | grep $service || go run cmd/goapp-godo/main.go goapp-${service} --watch &> ${LOG_DIR}/stdout-${service}.log &
    echo "Started ${service}"
}

start_authproxy_services() {
    for service in ${AUTHPROXY_SERVICES[@]}
    do
        start_service $service
    done
}
stop_authproxy_services() {
    stop_service "authproxy"
}

start_home_services() {
    for service in ${HOME_SERVICES[@]}
    do
        start_service $service
    done
}
stop_home_services() {
    stop_service "home"
}

start_resource_services() {
    for service in ${RESOURCE_SERVICES[@]}
    do
        start_service $service
    done
}
stop_resource_services() {
    stop_service "resource"
}


status() {
    ps ax | grep "go run " | grep -v grep || echo "NotFound Processes"
}

$COMMAND
