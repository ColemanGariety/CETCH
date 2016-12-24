package utils

import (
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"context"

	"github.com/JacksonGariety/wetch/models"
)

func AuthorizeClaims(protectedPage http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

		cookie, err := r.Cookie("Auth")
		if err != nil {
			protectedPage.ServeHTTP(w, r)
			return
		}

		// Return a Token using the cookie
		token, err := jwt.ParseWithClaims(cookie.Value, &models.Claims{}, func(token *jwt.Token) (interface{}, error){
			// Make sure token's signature wasn't changed
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected siging method")
			}
			return []byte("secret"), nil
		})
		if err != nil {
			protectedPage.ServeHTTP(w, r)
			return
		}

		if claims, ok := token.Claims.(*models.Claims); ok && token.Valid {
			ctx := context.WithValue(r.Context(), "foo", *claims)
			protectedPage(w, r.WithContext(ctx))
		} else {
			protectedPage.ServeHTTP(w, r)
			return
		}
	})
}
