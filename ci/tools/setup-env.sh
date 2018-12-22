#!/bin/sh -xe

mkdir -p ~/.goapp/logs
mkdir -p ~/.goapp/tmp
mkdir -p ~/.goapp/etc

cp -r ci/etc/* ~/.goapp/etc/
