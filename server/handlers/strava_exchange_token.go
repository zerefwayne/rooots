package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/zerefwayne/rooots/server/config"
	"github.com/zerefwayne/rooots/server/dto/strava"
	"github.com/zerefwayne/rooots/server/repository"
	"github.com/zerefwayne/rooots/server/utils"
)

func getStravaExchangeTokenUri(code string) string {
	STRAVA_CLIENT_ID := os.Getenv("STRAVA_CLIENT_ID")
	STRAVA_CLIENT_SECRET := os.Getenv("STRAVA_CLIENT_SECRET")

	grantType := "authorization_code"

	return fmt.Sprintf("https://www.strava.com/api/v3/oauth/token?client_id=%s&client_secret=%s&code=%s&grant_type=%s", STRAVA_CLIENT_ID, STRAVA_CLIENT_SECRET, code, grantType)
}

func isInvalidExchangeTokenBody(response *strava.ExchangeTokenResponseBody) bool {
	return response.TokenType == ""
}

func ExchangeTokenHandler(w http.ResponseWriter, r *http.Request) {
	var body strava.ExchangeTokenBody
	err := utils.DecodeJson(r.Body, &body)
	if err != nil {
		utils.HandleHttpError(err, w)
		return
	}

	stravaExchangeTokenUri := getStravaExchangeTokenUri(body.Code)
	request, err := http.Post(stravaExchangeTokenUri, "application/json", nil)
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

	user, err := repository.FindOrCreateUserByStrava(config.DB, &exchangeTokenBody)
	if err != nil {
		utils.HandleHttpError(err, w)
		return
	}

	newCookie := utils.GetCookie(config.REFRESH_TOKEN_COOKIE_NAME, exchangeTokenBody.RefreshToken, time.Unix(exchangeTokenBody.ExpiresAt, 0))
	http.SetCookie(w, newCookie)

	loginResponse := strava.LoginSuccessResponse{
		AccessToken: exchangeTokenBody.AccessToken,
		Name:        fmt.Sprintf("%s %s", user.FirstName, user.LastName),
		UserId:      user.Id,
	}

	if err := utils.RespondWithJson(w, loginResponse, http.StatusOK); err != nil {
		utils.HandleHttpError(err, w)
		return
	}
}
