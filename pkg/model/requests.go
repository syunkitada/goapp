package model

type AuthRequest struct {
	Username string
	Password string
	Project  string
}

type TokenAuthRequest struct {
	Username string
	Token    string
}
