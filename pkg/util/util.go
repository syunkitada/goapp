package util

import (
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/scrypt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/syunkitada/goapp/pkg/model"
	// "github.com/golang/glog"
)

import (
	"github.com/syunkitada/goapp/pkg/config"
)

var (
	Conf = &config.Conf
)

type CustomClaims struct {
	Admin    bool   `json:"admin"`
	Username string `json:"username"`
	// Project  *Project `json:"project"`
	jwt.StandardClaims
}

func GenerateHashFromPassword(username string, password string) (string, error) {
	converted, err := scrypt.Key([]byte(password), []byte(Conf.Admin.Secret+username), 16384, 8, 1, 32)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(converted[:]), nil
}

func GenerateToken(authRequest *model.AuthRequest) (string, error) {
	claims := CustomClaims{
		Admin:    true,
		Username: authRequest.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    authRequest.Username,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token with key
	tokenString, tokenErr := token.SignedString([]byte(Conf.Admin.TokenSecret + authRequest.Username))
	return tokenString, tokenErr
}

func ParseToken(request model.TokenAuthRequest) (jwt.MapClaims, error) {
	token, err := jwt.Parse(request.Token, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			msg := fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			return nil, msg
		}
		return []byte(Conf.Admin.TokenSecret + request.Username), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, nil
}
