package handlers

import (
	"net/http"

	"github.com/zerefwayne/rooots/server/constants"
	"github.com/zerefwayne/rooots/server/utils"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(constants.REFRESH_TOKEN_COOKIE_NAME)
	if err != nil {
		utils.HandleHttpError(err, w)
		return
	}

	deleteCookie := utils.RemoveCookie(cookie)
	http.SetCookie(w, deleteCookie)
	w.WriteHeader(http.StatusNoContent)
}
