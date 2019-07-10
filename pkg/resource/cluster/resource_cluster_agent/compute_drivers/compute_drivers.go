package compute_drivers

import "github.com/syunkitada/goapp/pkg/lib/logger"

type ComputeDriver interface {
	GetName() string
	Create(tctx *logger.TraceContext) error
	ConfirmCreate(tctx *logger.TraceContext) error
	Delete(tctx *logger.TraceContext) error
	ConfirmDelete(tctx *logger.TraceContext) error
}
