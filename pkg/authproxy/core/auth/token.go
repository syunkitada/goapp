package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/golang/glog"

	"github.com/syunkitada/goapp/pkg/authproxy/model"
	"github.com/syunkitada/goapp/pkg/authproxy/model/model_api"
)

type CustomClaims struct {
	Username string
	jwt.StandardClaims
}

func AuthAndIssueToken(authRequest *model.AuthRequest) (string, error) {
	user, userErr := model_api.GetAuthUser(authRequest)
	if userErr != nil {
		return "", errors.New("Failed GetAuthUser")
	}

	if token, err := GenerateToken(user); err != nil {
		return "", err
	} else if token == "" {
		return "", errors.New("Failed GenerateToken")
	} else {
		return token, nil
	}
}

func GenerateToken(user *model.User) (string, error) {
	claims := CustomClaims{
		Username: user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    user.Name,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token with key
	tokenString, tokenErr := token.SignedString([]byte(Conf.Admin.TokenSecret))
	return tokenString, tokenErr
}

func ParseToken(request model.TokenAuthRequest) (jwt.MapClaims, error) {
	token, err := jwt.Parse(request.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			msg := fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			return nil, msg
		}
		return []byte(Conf.Admin.TokenSecret), nil
	})
	glog.Info("DEBUGaaaa")

	if err != nil {
		return nil, err
	}
	glog.Info("DEBUGaaaabbb")

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	glog.Info("DEBUGaaaabbbaaaee")

	return nil, nil
}
