package model

type AuthRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type TokenAuthRequest struct {
	Project string
	Token   string
}
