package resolver

import (
	"fmt"

	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"

	api_spec "github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

func (resolver *Resolver) ReportNode(tctx *logger.TraceContext, input *api_spec.ReportNode,
	user *base_spec.UserAuthority) (data *api_spec.ReportNodeData, code uint8, err error) {
	fmt.Println("DEBUG reportResource")
	if err = resolver.tsdbApi.ReportNode(tctx, input); err != nil {
		code = base_const.CodeServerInternalError
		fmt.Println("DEBUG error report", err)
		return
	}
	fmt.Println("DEBUG logs:", len(input.Logs))
	fmt.Println("DEBUG metrics:", len(input.Metrics))
	code = base_const.CodeOk
	data = &api_spec.ReportNodeData{}
	return
}
