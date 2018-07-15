package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/syunkitada/goapp/pkg/api/handlers"
	"github.com/syunkitada/goapp/pkg/api/services"
)

// API holds the api handlers
type API struct {
	encryptionKey []byte
	AllowedHosts  []string
	AclService    services.ACLService

	Hello  *handlers.Hello
	Tokens *handlers.Tokens
	Users  *handlers.Users
}

// NewAPI creates a new API
func NewAPI() *API {
	// TODO: Use generated key from README
	encryptionKey := []byte("secret")

	allowedHosts := []string{
		"localhost",
	}

	aclService := services.NewACLService()
	tokenService := services.NewTokenService()
	userService := services.NewUserService()
	helloService := services.NewHelloService()

	return &API{
		encryptionKey: encryptionKey,
		AllowedHosts:  allowedHosts,
		AclService:    aclService,
		Tokens:        handlers.NewTokens(tokenService),
		Hello:         handlers.NewHello(helloService),
		Users:         handlers.NewUsers(userService),
	}
}

// Middleware

// Authenticate provides Authentication middleware for handlers
func (a *API) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("Authenticate")
		var tokenString string

		// Get token from the Authorization header
		// format: Authorization: Bearer <token>
		tokens, ok := r.Header["Authorization"]
		if ok && len(tokens) >= 1 {
			tokenString = tokens[0]
			tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		}

		// If the token is empty...
		if tokenString == "" {
			// If we get here, the required token is missing
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		// Now parse the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				msg := fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				return nil, msg
			}
			return a.encryptionKey, nil
		})

		if err != nil {
			fmt.Println("Invalid Token: ", err)
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Everything worked! Set the user in the context.
			context.Set(r, "user", claims["user"])
			next.ServeHTTP(w, r)
			return
		}
		http.Error(w, "Error parsing token", http.StatusUnauthorized)
		return
	})
}

// Authorize provides authorization middleware for our handlers
func (a *API) Authorize(permissions ...services.Permission) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Print("Authorize")

			// TODO: Get User Information from Request
			user := &services.User{
				ID:        1,
				FirstName: "Admin",
				LastName:  "User",
				Roles:     []string{services.AdministratorRole},
			}

			// if user == nil {
			// 	http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			// }

			for _, permission := range permissions {
				if err := a.AclService.CheckPermission(user, permission); err != nil {
					http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
					return
				}
			}

			next.ServeHTTP(w, r)
		})
	}
}

// SecureHeaders adds secure headers to the API
func (a *API) SecureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check AllowedHosts
		var err error
		if len(a.AllowedHosts) > 0 {
			isGoodHost := false
			for _, allowedHost := range a.AllowedHosts {
				if strings.EqualFold(allowedHost, r.Host) {
					isGoodHost = true
					break
				}
			}
			if !isGoodHost {
				http.Error(w, fmt.Sprintf("Bad host name: %s", r.Host), http.StatusForbidden)
				return
			}
		}
		// If there was an error, do not continue request
		if err != nil {
			http.Error(w, "Failed to check allowed hosts", http.StatusInternalServerError)
			return
		}

		// Add X-XSS-Protection header
		// Enables XSS filtering. Rather than sanitizing the page, the browser will prevent rendering of the page if an attack is detected.
		w.Header().Add("X-XSS-Protection", "1; mode=blockFilter")

		// Add Content-Type header
		// Content type tells the browser what type of content you are sending. If you do not include it, the browser will try to guess the type and may get it wrong.
		// w.Header().Add("Content-Type", "application/json")

		// Add X-Content-Type-Options header
		// Content Sniffing is the inspecting the content of a byte stream to attempt to deduce the file format of the data within it.
		// Browsers will do this to try to guess at the content type you are sending.
		// By setting this header to “nosniff”, it prevents IE and Chrome from content sniffing a response away from its actual content type. This reduces exposure to drive-by download attacks.
		w.Header().Add("X-Content-Type-Options", "nosniff")

		// Prevent page from being displayed in an iframe
		w.Header().Add("X-Frame-Options", "DENY")

		next.ServeHTTP(w, r)
	})
}
