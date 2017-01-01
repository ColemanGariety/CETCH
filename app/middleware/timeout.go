package middleware

import (
	"net/http"
	"time"
)

func Timeout(h http.Handler) http.Handler {
	return http.TimeoutHandler(h, 10*time.Second, "504 gateway timeout")
}
