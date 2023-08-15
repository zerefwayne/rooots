package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/zerefwayne/rooots/server/constants"
	"github.com/zerefwayne/rooots/server/utils"
)

type AuthorizationData struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	UserId       uuid.UUID `json:"user_id"`
}

type AuthorizationContextKey struct{}

func Authorize(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jwtCookie, err := r.Cookie(constants.REFRESH_TOKEN_COOKIE_NAME)
		if err != nil {
			utils.HandleHttpError(err, w)
			return
		}

		jwtContent := jwtCookie.Value

		isValid, jwtClaims, err := utils.ValidateJwtToken(jwtContent)
		if !isValid {
			utils.HandleHttpError(err, w)
			return
		}

		if !isValid || err != nil {
			deleteCookie := utils.RemoveCookie(jwtCookie)
			http.SetCookie(w, deleteCookie)
			utils.HandleHttpError(err, w)
			return
		}

		accessToken := r.Header.Get("Authorization")

		userIdUuid, err := uuid.Parse(jwtClaims.UserId)
		if err != nil {
			utils.HandleHttpError(err, w)
			return
		}

		ctx := context.WithValue(r.Context(), AuthorizationContextKey{}, &AuthorizationData{RefreshToken: jwtClaims.RefreshToken, UserId: userIdUuid, AccessToken: accessToken})

		next(w, r.WithContext(ctx))
	}

}
