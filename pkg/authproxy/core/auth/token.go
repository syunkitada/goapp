package auth

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_model"
	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_model/authproxy_model_api"
	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

type CustomClaims struct {
	Username string
	jwt.StandardClaims
}

type Token struct {
	conf              *config.Config
	authproxyModelApi *authproxy_model_api.AuthproxyModelApi
}

func NewToken(conf *config.Config, authproxyModelApi *authproxy_model_api.AuthproxyModelApi) *Token {
	token := Token{
		conf:              conf,
		authproxyModelApi: authproxyModelApi,
	}
	return &token
}

func (token *Token) AuthAndIssueToken(tctx *logger.TraceContext, authRequest *authproxy_model.AuthRequest) (string, error) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var user *authproxy_model.User
	user, err = token.authproxyModelApi.GetAuthUser(tctx, authRequest)
	if err != nil {
		return "", err
	}

	var tokenStr string
	if tokenStr, err = token.Generate(user); err != nil {
		return "", err
	} else if tokenStr == "" {
		return "", fmt.Errorf("Failed GenerateToken")
	} else {
		return tokenStr, nil
	}
}

func (token *Token) Generate(user *authproxy_model.User) (string, error) {
	claims := CustomClaims{
		Username: user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    user.Name,
		},
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token with key
	tokenString, tokenErr := newToken.SignedString([]byte(token.conf.Admin.TokenSecret))
	return tokenString, tokenErr
}

func (token *Token) ParseToken(request authproxy_model.TokenAuthRequest) (jwt.MapClaims, error) {
	parsedToken, err := jwt.Parse(request.Token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			msg := fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
			return nil, msg
		}
		return []byte(token.conf.Admin.TokenSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		return claims, nil
	}

	return nil, nil
}
