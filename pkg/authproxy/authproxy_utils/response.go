package authproxy_utils

import (
	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_grpc_pb"
	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_model"
	"github.com/syunkitada/goapp/pkg/lib/codes"
	"github.com/syunkitada/goapp/pkg/lib/json_utils"
)

func MergeResponse(
	rep *authproxy_grpc_pb.ActionReply, response *authproxy_model.ActionResponse,
	data map[string]interface{}, err error, statusCode int64) {
	response.Data = data
	if err != nil {
		response.Tctx.Err = err.Error()
	}
	response.Tctx.StatusCode = statusCode
	responseBytes, err := json_utils.Marshal(response)
	if err != nil {
		rep.Tctx.Err = err.Error()
		rep.Tctx.StatusCode = codes.ServerInternalError
		return
	}
	rep.Response = string(responseBytes)
}
