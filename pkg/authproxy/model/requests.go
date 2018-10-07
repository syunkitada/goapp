package model

type AuthRequest struct {
	Username string `binding:"required"`
	Password string `binding:"required"`
	Action   ActionRequest
}

type TokenAuthRequest struct {
	Token  string
	Action ActionRequest
}

type ActionRequest struct {
	ProjectName string
	ServiceName string
	Name        string
}
