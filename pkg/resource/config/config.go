package config

import (
	"github.com/syunkitada/goapp/pkg/base/base_config"
)

type Config struct {
	Resource ResourceConfig
}

type ResourceConfig struct {
	Api         base_config.AppConfig
	Controller  ResourceControllerConfig
	ClusterName string
	ClusterMap  map[string]ResourceClusterConfig
}

type ResourceControllerConfig struct {
	base_config.AppConfig
	SyncRegionServiceTimeout int
}

type ResourceClusterConfig struct {
	Region     string
	Datacenter string
	Kind       string
	Weight     int
	Api        base_config.AppConfig
	Controller base_config.AppConfig
}

type ResourceClusterApiConfig struct {
	base_config.AppConfig
	RegionName string
}

var BaseConf = base_config.Config{}

var MainConf = Config{
	Resource: ResourceConfig{
		Api: base_config.AppConfig{
			Name:                 "ResourceApi",
			NodeDownTimeDuration: 60,
		},
		Controller: ResourceControllerConfig{
			AppConfig: base_config.AppConfig{
				Name: "ReosurceController",
			},
			SyncRegionServiceTimeout: 10,
		},
	},
}
