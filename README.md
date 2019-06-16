# goapp


# Getting Started
## Requirements

```
$ go version
go version go1.12.5 linux/amd64

$ env | grep GO
GOPATH=/home/hoge/go
GO111MODULE=on

$ yarn version
yarn version v1.16.0
```

## Install

```
mkdir -p $GOPATH/src/github.com/syunkitada/
git clone git@github.com:syunkitada/goapp.git $GOPATH/src/github.com/syunkitada/goapp
cd $GOPATH/src/github.com/syunkitada/goapp
make env
go run cmd/goapp-adminctl/main.go bootstrap
```

## Edit config
```
# Allow your server ip
$ vim ~/.goapp/etc/config.toml
[Authproxy.HttpServer]
AllowedHosts = ["127.0.0.1:8000", "192.168.10.121:3000", "192.168.10.121:8000"]
```

## Start apps

```
$ make start-all

$ make status
ci/tools/service.sh status
30694 pts/7    Sl     0:00 go run cmd/goapp-godo/main.go goapp-authproxy --watch
30695 pts/7    Sl     0:00 go run cmd/goapp-godo/main.go goapp-resource-api --watch
30696 pts/7    Sl     0:00 go run cmd/goapp-godo/main.go goapp-resource-controller --watch
30697 pts/7    Sl     0:00 go run cmd/goapp-godo/main.go goapp-resource-cluster-api --watch
30698 pts/7    Sl     0:00 go run cmd/goapp-godo/main.go goapp-resource-cluster-controller --watch
30699 pts/7    Sl     0:00 go run cmd/goapp-godo/main.go goapp-resource-cluster-agent --watch
```

## Start dashboard

```
cd $GOPATH/src/github.com/syunkitada/goapp/dashboard
cd dashboard
export REACT_APP_AUTHPROXY_URL="https://192.168.10.103:8000";
yarn install
yarn start
```
