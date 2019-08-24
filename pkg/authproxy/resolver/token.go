package resolver

import (
	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/authproxy/spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

func (resolver *Resolver) IssueToken(tctx *logger.TraceContext, db *gorm.DB, input *spec.IssueToken) (*spec.IssueTokenData, error) {
	return &spec.IssueTokenData{Token: "hoge"}, nil
}
