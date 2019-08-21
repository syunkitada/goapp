package autogen

import (
	"encoding/json"

	"github.com/syunkitada/goapp/pkg/authproxy/spec"
	"github.com/syunkitada/goapp/pkg/base/base_model"
)

type QueryResolver interface {
	IssueToken(*spec.IssueToken) (*spec.IssueTokenData, error)
}

type QueryHandler struct {
	resolver QueryResolver
}

func NewQueryHandler(resolver QueryResolver) *QueryHandler {
	return &QueryHandler{
		resolver: resolver,
	}
}

func (handler *QueryHandler) Exec(req *base_model.Request, rep *base_model.Reply) error {
	var err error
	for _, query := range req.Queries {
		switch query.Name {
		case "IssueToken":
			var input spec.IssueToken
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, err := handler.resolver.IssueToken(&input)
			rep.Data["IssueToken"] = data
			return err
		}
	}
	return nil
}
