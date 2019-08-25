package resolver

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/authproxy/spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

func (resolver *Resolver) IssueToken(tctx *logger.TraceContext, db *gorm.DB, input *spec.IssueToken) (data *spec.IssueTokenData, err error) {
	user, err := resolver.dbApi.GetUserWithValidatePassword(tctx, db, input.User, input.Password)
	fmt.Println("DEBUG user", user)
	data = &spec.IssueTokenData{Token: "hoge"}
	return
}
