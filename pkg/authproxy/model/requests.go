package model

type AuthRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Project  string `form:"project" json:"project" binding:"required"`
}

type TokenAuthRequest struct {
	Username string
	Token    string
}
