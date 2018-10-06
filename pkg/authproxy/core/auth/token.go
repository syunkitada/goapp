package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/syunkitada/goapp/pkg/authproxy/model"
	"github.com/syunkitada/goapp/pkg/authproxy/model/model_api"
	"github.com/syunkitada/goapp/pkg/config"
)

type CustomClaims struct {
	Username string
	jwt.StandardClaims
}

type Token struct {
	Conf     *config.Config
	ModelApi *model_api.ModelApi
}

func NewToken(conf *config.Config, modelApi *model_api.ModelApi) *Token {
	token := Token{
		Conf:     conf,
		ModelApi: modelApi,
	}
	return &token
}

func (token *Token) AuthAndIssueToken(authRequest *model.AuthRequest) (string, error) {
	user, userErr := token.ModelApi.GetAuthUser(authRequest)
	if userErr != nil {
		return "", errors.New("Failed GetAuthUser")
	}

	if token, err := token.Generate(user); err != nil {
		return "", err
	} else if token == "" {
		return "", errors.New("Failed GenerateToken")
	} else {
		return token, nil
	}
}

func (token *Token) Generate(user *model.User) (string, error) {
	claims := CustomClaims{
		Username: user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    user.Name,
		},
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token with key
	tokenString, tokenErr := newToken.SignedString([]byte(token.Conf.Admin.TokenSecret))
	return tokenString, tokenErr
}

func (token *Token) ParseToken(request model.TokenAuthRequest) (jwt.MapClaims, error) {
	parsedToken, err := jwt.Parse(request.Token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			msg := fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
			return nil, msg
		}
		return []byte(token.Conf.Admin.TokenSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		return claims, nil
	}

	return nil, nil
}