package core

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang/glog"

	"github.com/syunkitada/goapp/pkg/authproxy/model"
	"github.com/syunkitada/goapp/pkg/authproxy/model/model_api"
)

type CustomClaims struct {
	UserName        string
	RoleName        string
	ProjectName     string
	ProjectRoleName string
	ExpirationDate  string
	jwt.StandardClaims
}

func (authproxy *Authproxy) IssueToken(c *gin.Context) {
	var authRequest model.AuthRequest

	if err := c.ShouldBindWith(&authRequest, binding.JSON); err != nil {
		glog.Warningf("Invalid AuthRequest: Failed ShouldBindJSON: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid AuthRequest",
		})
		c.Abort()
		return
	}

	user, userErr := model_api.GetAuthUser(&authRequest)
	if userErr != nil {
		glog.Error("Failed GetAuthUser", userErr)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed IssueToken",
		})
		c.Abort()
		return
	}

	if token, err := authproxy.GenerateToken(user); err != nil {
		glog.Error("Failed GenerateToken", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed IssueToken",
		})
		c.Abort()
		return
	} else {
		if token != "" {
			glog.Info("Success Login: ", authRequest)
			c.JSON(http.StatusOK, gin.H{
				"token": token,
			})
			return
		} else {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Invalid authRequest",
			})
			c.Abort()
			return
		}
	}
}

func (authproxy *Authproxy) GenerateToken(user *model.CustomUser) (string, error) {
	claims := CustomClaims{
		UserName:        user.Name,
		RoleName:        user.RoleName,
		ProjectName:     user.ProjectName,
		ProjectRoleName: user.ProjectRoleName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    user.Name,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token with key
	tokenString, tokenErr := token.SignedString([]byte(Conf.Admin.TokenSecret + user.Name))
	return tokenString, tokenErr
}

func (authproxy *Authproxy) ParseToken(request model.TokenAuthRequest) (jwt.MapClaims, error) {
	token, err := jwt.Parse(request.Token, func(token *jwt.Token) (interface{}, error) {
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
