package resolver

import (
	"github.com/syunkitada/goapp/pkg/authproxy/config"
	"github.com/syunkitada/goapp/pkg/authproxy/db_api"
	"github.com/syunkitada/goapp/pkg/base/base_config"
)

type Resolver struct {
	dbApi *db_api.Api
}

func New(baseConf *base_config.Config, mainConf *config.Config) *Resolver {
	return &Resolver{
		dbApi: db_api.New(baseConf, mainConf),
	}
}
