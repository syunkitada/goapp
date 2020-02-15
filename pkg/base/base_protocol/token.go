package base_protocol

import "github.com/dgrijalva/jwt-go"

type JwtClaims struct {
	Username string
	jwt.StandardClaims
}
