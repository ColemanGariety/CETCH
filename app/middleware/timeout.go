package middleware

import (
	"net/http"
	"time"
)

func Timeout(h http.Handler) http.Handler {
	return http.TimeoutHandler(h, 30 * time.Second, "timed out")
}
