package auth

import (
	"github.com/syunkitada/goapp/pkg/config"
)

var (
	Conf = &config.Conf
)

type Auth struct{}

func NewAuth() *Auth {
	auth := Auth{}
	return &auth
}
