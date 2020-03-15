# Quick Starting

## Requirements

```
$ cat /etc/os-release
NAME="Ubuntu"
VERSION="18.04.3 LTS (Bionic Beaver)"
ID=ubuntu
ID_LIKE=debian
PRETTY_NAME="Ubuntu 18.04.3 LTS"
VERSION_ID="18.04"
HOME_URL="https://www.ubuntu.com/"
SUPPORT_URL="https://help.ubuntu.com/"
BUG_REPORT_URL="https://bugs.launchpad.net/ubuntu/"
PRIVACY_POLICY_URL="https://www.ubuntu.com/legal/terms-and-policies/privacy-policy"
VERSION_CODENAME=bionic
UBUNTU_CODENAME=bionic

$ go version
go version go1.12.15 linux/amd64
```

## Setup pkg

```
$ mkdir -p $GOPATH/src/github.com/syunkitada/
$ git clone git@github.com:syunkitada/goapp.git $GOPATH/src/github.com/syunkitada/goapp
$ cd $GOPATH/src/github.com/syunkitada/goapp
$ make env
```

## Edit config

```
# Allow your server ip
$ vim ~/.goapp/etc/config.yaml
```

## Setup DB

```
$ go run cmd/goapp-authproxy/main.go ctl bootstrap
$ go run cmd/goapp-resource/main.go ctl bootstrap
```

## Start Services

```
# Enable root if you want to use qemu
$ sudo -E zsh

$ make start-all

$ make status

$ make status
ci/tools/service.sh status
4275 pts/5    Sl     0:00 go run cmd/goapp-godo/main.go goapp-authproxy-api --watch
6973 pts/5    Sl     0:00 go run cmd/goapp-godo/main.go goapp-resource-api --watch
6986 pts/5    Sl     0:00 go run cmd/goapp-godo/main.go goapp-resource-controller --watch
7016 pts/5    Sl     0:00 go run cmd/goapp-godo/main.go goapp-resource-cluster-api --watch
7040 pts/5    Sl     0:00 go run cmd/goapp-godo/main.go goapp-resource-cluster-controller --watch
7080 pts/5    Sl     0:00 go run cmd/goapp-godo/main.go goapp-resource-cluster-agent --watch
```

## Setup Dashboard

```
$ cd $GOPATH/src/github.com/syunkitada/goapp/dashboard
$ cd dashboard
$ yarn install
```

## Start Dashboard

```
$ export REACT_APP_AUTHPROXY_URL="https://${YOUR_SERVER_HOST}:8000";
$ yarn start
```
