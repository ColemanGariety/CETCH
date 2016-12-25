package models

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
    Username string `json:"username"`
    jwt.StandardClaims
}

func ClaimsCreate(username string) (string, time.Time) {
	expireToken := time.Now().Add(time.Hour * 8760).Unix() // 24 hours * 365 days = 8760 hours/year
	expireCookie := time.Now().Add(time.Hour * 8760)

	claims := Claims {
		username,
		jwt.StandardClaims {
			ExpiresAt: expireToken,
			Issuer:    "localhost:8080",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, _ := token.SignedString([]byte("secret"))

	return signedToken, expireCookie
}
