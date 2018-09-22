package dashboard

import (
	"github.com/syunkitada/goapp/pkg/config"
)

var (
	Conf = &config.Conf
)

type Dashboard struct{}

func NewDashboard() *Dashboard {
	dashboard := Dashboard{}
	return &dashboard
}
