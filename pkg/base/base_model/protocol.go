package base_model

import (
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

type ReqQuery struct {
	Name string
	Data string
}

type Request struct {
	Tctx    *logger.TraceContext
	Token   string
	Service string
	Project string
	Queries []ReqQuery
}

type Response struct {
	TraceId   string
	Code      uint8
	Error     string
	ResultMap map[string]Result
}

type Result struct {
	Code  uint8
	Error string
	Data  interface{}
}
