package handlers

import (
	"net/http"

	"fmt"
	"github.com/syunkitada/goapp/pkg/api/services"
)

// Users provides API access to the user service
type Users struct {
	Service services.UserService
}

// NewUsers creates a users handler
func NewUsers(s services.UserService) *Users {
	return &Users{s}
}

// Handler handles requests to /users
func (u *Users) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		fmt.Print("Get Users")
	case "POST":
	case "PUT":
	case "DELETE":
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
