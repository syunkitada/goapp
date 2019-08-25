package base_model

import "github.com/dgrijalva/jwt-go"

type JwtClaims struct {
	Username string
	jwt.StandardClaims
}
