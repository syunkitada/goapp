#!/bin/sh -xe

mkdir -p ~/.goapp/logs
mkdir -p ~/.goapp/tmp
mkdir -p ~/.goapp/etc

cp -r ci/etc/goapp/* ~/.goapp/etc/

GO111MODULE=off
go get -u gopkg.in/godo.v2/cmd/godo

# GO111MODULE=off
# go get github.com/golangci/golangci-lint
# 
# go get golang.org/x/lint/golint
# go get github.com/alecthomas/gometalinter
# GO111MODULE=on

OS_RELEASE=""
if grep "Ubuntu" /etc/os-release > /dev/null; then
    OS_RELEASE="Ubuntu"
fi

if [ $OS_RELEASE = "Ubuntu" ]; then
    # install docker
    # https://docs.docker.com/install/linux/docker-ce/ubuntu/
    sudo apt-get -y update
    sudo apt-get -y install \
        apt-transport-https \
        ca-certificates \
        curl \
        gnupg-agent \
        software-properties-common
    curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
    sudo apt-key fingerprint 0EBFCD88
    sudo add-apt-repository \
   "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
       $(lsb_release -cs) \
       stable"
    sudo apt-get -y update
    sudo apt-get -y install docker-ce docker-ce-cli containerd.io

    sudo systemctl start docker

    sudo apt-get -y install mysql-client
fi

# Setup mysql
sudo docker ps | grep mysql || \
    ( \
     ((sudo docker ps --all | grep mysql && sudo docker rm mysql) || echo "mysql not found") && \
     sudo docker run -v "/var/lib/docker-mysql":/var/lib/mysql --net=host --name mysql -e MYSQL_ROOT_PASSWORD=rootpass -d mysql \
    )

mysql -uroot -prootpass -h127.0.0.1 -e "CREATE USER IF NOT EXISTS 'goapp'@'%' IDENTIFIED BY 'goapppass'; GRANT ALL ON *.* TO 'goapp'@'%'; FLUSH PRIVILEGES;"

cat << EOS | tee ~/.my.cnf
[client]
host=127.0.0.1
port=3306
user=goapp
password=goapppass
EOS


# Setup influxdb
# https://hub.docker.com/_/influxdb/
sudo docker ps | grep influxdb || \
    ( \
     ((sudo docker ps --all | grep influxdb && sudo docker rm influxdb) || echo "influxdb not found") && \
     sudo docker run -v "/var/lib/docker-influxdb":/var/lib/influxdb --net=host --name influxdb -e INFLUXDB_ADMIN_ENABLED=true -d influxdb \
    )

curl -XPOST http://localhost:8086/query --data-urlencode "q=CREATE USER goapp WITH PASSWORD 'goapppass'"
curl -XPOST http://localhost:8086/query --data-urlencode "q=GRANT ALL PRIVILEGES TO goapp"
