package controllers

import (
	"net/http"

	"github.com/JacksonGariety/wetch/models"
	"github.com/JacksonGariety/wetch/utils"
)

func ProfileShow(w http.ResponseWriter, r *http.Request){
	claims, ok := r.Context().Value("foo").(models.Claims)
	if !ok {
		http.Redirect(w, r, "/login", 307)
		return
	}

	utils.Render(w, "profile.html", claims)
}
