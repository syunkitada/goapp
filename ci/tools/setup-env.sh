#!/bin/sh -xe

mkdir -p ~/.goapp/logs
mkdir -p ~/.goapp/tmp
mkdir -p ~/.goapp/etc

cp -r ci/etc/* ~/.goapp/etc/

cd ~/.goapp/tmp
[ -e influxdb_1.7.2_amd64.deb ] || wget https://dl.influxdata.com/influxdb/releases/influxdb_1.7.2_amd64.deb
sudo dpkg -i influxdb_1.7.2_amd64.deb
sudo service influxdb start

influx -execute "CREATE DATABASE goapp_monitor_logs"
influx -execute "CREATE DATABASE goapp_monitor_metrics"
influx -execute "SHOW DATABASES"
influx -execute "create user goapp with password 'goapppass'"
influx -execute "grant all privileges to goapp"
