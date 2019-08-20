package base_model

import "github.com/syunkitada/goapp/pkg/lib/logger"

type ReqQuery struct {
	Name string
	Data string
}

type Request struct {
	Tctx    *logger.TraceContext
	Token   string
	Service string
	Queries []ReqQuery
}

type Reply struct {
	TraceId string
	Code    uint8
	Error   string
	Data    map[string]interface{}
}
