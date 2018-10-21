# go-app


# How to

```
# Authproxy
go run cmd/goapp-godo/main.go authproxy --watch


# Resource
go run cmd/goapp-godo/main.go resource-api --watch
go run cmd/goapp-godo/main.go resource-controller --watch
```

# Test Authproxy

```
export CONFIG_DIR=$PWD/ci/etc
export CONFIG_FILE=test_config.toml
go test pkg/authproxy/test/main_test.go
```
