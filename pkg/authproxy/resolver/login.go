package resolver

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/authproxy/spec"
	"github.com/syunkitada/goapp/pkg/base/base_model"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

func (resolver *Resolver) Login(tctx *logger.TraceContext, db *gorm.DB, input *spec.Login) (data *spec.LoginData, code uint8, err error) {
	user, code, err := resolver.dbApi.GetUserWithValidatePassword(tctx, db, input.User, input.Password)
	if err != nil {
		return
	}

	token, err := resolver.issueToken(user.Name)
	if err != nil {
		return
	}

	userAuthority, err := resolver.dbApi.GetUserAuthority(tctx, db, input.User)
	if err != nil {
		return
	}

	data = &spec.LoginData{
		Name:      input.User,
		Token:     token,
		Authority: *userAuthority,
	}
	return
}

func (resolver *Resolver) issueToken(userName string) (string, error) {
	claims := base_model.JwtClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    userName,
		},
	}
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token with key
	tokenString, tokenErr := newToken.SignedString([]byte(resolver.appConf.Auth.Secrets[0]))
	return tokenString, tokenErr
}
