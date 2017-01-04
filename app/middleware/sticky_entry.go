package middleware

import (
	"net/http"

	"github.com/JacksonGariety/cetch/app/models"
	"github.com/JacksonGariety/cetch/app/utils"
)

func StickyEntry(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		data := *r.Context().Value("data").(*utils.Props)
		if data["current_user"] != nil {
			user := data["current_user"].(models.User)
			data["sticky_entry"] = user.CurrentEntry()
		}
		next.ServeHTTP(w, r)
	})
}
