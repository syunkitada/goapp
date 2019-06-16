package authproxy_model

import (
	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_grpc_pb"
	"github.com/syunkitada/goapp/pkg/authproxy/index_model"
)

type AuthRequest struct {
	Username string `binding:"required"`
	Password string `binding:"required"`
	Action   ActionRequest
}

type TokenAuthRequest struct {
	Token  string
	Action ActionRequest
}

type ActionRequest struct {
	ProjectName string
	ServiceName string
	Queries     []Query
}

type Query struct {
	Kind      string
	StrParams map[string]string
	NumParams map[string]int64
}

type ActionResponse struct {
	Tctx  authproxy_grpc_pb.TraceContext
	Index index_model.Index
	Data  map[string]interface{}
}
