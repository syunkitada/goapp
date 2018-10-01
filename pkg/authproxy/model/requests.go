package model

type AuthRequest struct {
	Username string `binding:"required"`
	Password string `binding:"required"`
}

type TokenAuthRequest struct {
	Project string
	Token   string
}
