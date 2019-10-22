package metrics_plugins

import (
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

type MetricsReader interface {
	GetName() string
	Read(tctx *logger.TraceContext) error
	Report() ([]resource_model.Metric, []resource_model.Alert)
	Reported()
}
