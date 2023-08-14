package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/zerefwayne/rooots/server/config"
	strava "github.com/zerefwayne/rooots/server/dto/strava"
	"github.com/zerefwayne/rooots/server/repository"
	"github.com/zerefwayne/rooots/server/utils"
)

func getStravaRefreshTokenUri(refreshToken string) string {
	STRAVA_CLIENT_ID := os.Getenv("STRAVA_CLIENT_ID")
	STRAVA_CLIENT_SECRET := os.Getenv("STRAVA_CLIENT_SECRET")

	grantType := "refresh_token"

	return fmt.Sprintf("https://www.strava.com/api/v3/oauth/token?client_id=%s&client_secret=%s&refresh_token=%s&grant_type=%s", STRAVA_CLIENT_ID, STRAVA_CLIENT_SECRET, refreshToken, grantType)
}

func RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	refreshTokenCookie, err := r.Cookie(config.REFRESH_TOKEN_COOKIE_NAME)
	if err != nil {
		utils.HandleHttpError(err, w)
		return
	}

	stravaRefreshTokenUri := getStravaRefreshTokenUri(refreshTokenCookie.Value)

	request, err := http.Post(stravaRefreshTokenUri, "application/json", nil)
	if err != nil {
		utils.HandleHttpError(err, w)
		return
	}
	defer request.Body.Close()

	var exchangeTokenBody strava.ExchangeTokenResponseBody
	err = utils.DecodeJson(request.Body, &exchangeTokenBody)
	if err != nil {
		utils.HandleHttpError(err, w)
		return
	}

	if isInvalidExchangeTokenBody(&exchangeTokenBody) {
		utils.HandleHttpError(errors.New("invalid request"), w)
		return
	}

	user, err := repository.UpdateRefreshToken(config.DB, refreshTokenCookie.Value, exchangeTokenBody.RefreshToken)
	if err != nil {
		utils.HandleHttpError(err, w)
		return
	}

	newCookie := utils.GetCookie(config.REFRESH_TOKEN_COOKIE_NAME, exchangeTokenBody.RefreshToken, time.Unix(exchangeTokenBody.ExpiresAt, 0))
	http.SetCookie(w, newCookie)

	loginResponse := strava.LoginSuccessResponse{
		AccessToken: exchangeTokenBody.AccessToken,
		UserId:      user.Id,
	}

	if err := utils.RespondWithJson(w, loginResponse, http.StatusOK); err != nil {
		utils.HandleHttpError(err, w)
		return
	}
}
