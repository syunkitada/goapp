package resolver

import (
	"github.com/syunkitada/goapp/pkg/authproxy/spec"
)

func (resolver *Resolver) IssueToken(input *spec.IssueToken) (*spec.IssueTokenData, error) {
	return &spec.IssueTokenData{Token: "hoge"}, nil
}
