package middleware

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"

	"github.com/JacksonGariety/cetch/app/models"
	"github.com/JacksonGariety/cetch/app/utils"
)

var sessionHash = os.Getenv("session_hash")

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var ctx context.Context
		claims, ok := isAuthentic(r)
		if ok {
			user := models.User{Name: claims.Username}
			models.Find(&user)
			ctx = context.WithValue(r.Context(), "data", &utils.Props{
				"authorized":          ok,
				"authorized_username": claims.Username,
				"userpath":            user.Userpath(),
				"admin":               user.Admin,
				"current_user":        user,
			})
		} else {
			ctx = context.WithValue(r.Context(), "data", &utils.Props{
				"authorized":          false,
				"authorized_username": "",
				"userpath":            "",
				"admin":               false,
			})
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// unauthorized users are redirected to signup
func Protect(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data := r.Context().Value("data").(*utils.Props)
		if (*data)["authorized"].(bool) {
			next.ServeHTTP(w, r)
		} else {
			utils.NotAuthorized(w, r)
		}
	})
}

// non-admin users recieve 403
func Forbid(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data := r.Context().Value("data").(*utils.Props)
		if (*data)["authorized"].(bool) && (*data)["admin"].(bool) {
			next.ServeHTTP(w, r)
		} else {
			utils.Forbidden(w, r)
		}
	})
}

// authorized users are redirected to home
func Retain(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data := r.Context().Value("data").(*utils.Props)
		if (*data)["authorized"].(bool) {
			http.Redirect(w, r, "/", 307)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func isAuthentic(r *http.Request) (*models.Claims, bool) {
	// If no Auth cookie is set then return a 404 not found
	cookie, err := r.Cookie("Auth")
	if err != nil {
		return &models.Claims{}, false
	}

	// Return a Token using the cookie
	token, err := jwt.ParseWithClaims(cookie.Value, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Make sure token's signature wasn't changed
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected siging method")
		}
		return []byte(sessionHash), nil
	})
	if err != nil {
		return &models.Claims{}, false
	}

	// Grab the tokens claims and pass it into the original request
	if claims, ok := token.Claims.(*models.Claims); ok && token.Valid {
		return claims, true
	} else {
		return claims, false
	}
}
