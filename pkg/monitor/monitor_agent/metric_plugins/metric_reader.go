package metric_plugins

import (
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

type MetricReader interface {
	GetName() string
	Read(tctx *logger.TraceContext) error
}
