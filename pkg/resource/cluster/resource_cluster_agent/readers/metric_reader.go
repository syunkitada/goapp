package readers

import (
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

type MetricReader interface {
	GetName() string
	Read(tctx *logger.TraceContext) error
	Report() ([]spec.ResourceMetric, []spec.ResourceEvent)
	Reported()
}
