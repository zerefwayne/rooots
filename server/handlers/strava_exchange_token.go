package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/zerefwayne/rooots/server/config"
	"github.com/zerefwayne/rooots/server/dto"
	strava "github.com/zerefwayne/rooots/server/dto/strava"
	"github.com/zerefwayne/rooots/server/repository"
	"github.com/zerefwayne/rooots/server/utils"
)

func ExchangeTokenHandler(w http.ResponseWriter, r *http.Request) {
	var body dto.ExchangeTokenBody

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		utils.HandleHttpError(err, w)
		return
	}

	STRAVA_CLIENT_ID := os.Getenv("STRAVA_CLIENT_ID")
	STRAVA_CLIENT_SECRET := os.Getenv("STRAVA_CLIENT_SECRET")

	stravaExchangeTokenUri := fmt.Sprintf("https://www.strava.com/api/v3/oauth/token?client_id=%s&client_secret=%s&code=%s&grant_type=authorization_code", STRAVA_CLIENT_ID, STRAVA_CLIENT_SECRET, body.Code)

	request, err := http.Post(stravaExchangeTokenUri, "application/json", nil)
	if err != nil {
		utils.HandleHttpError(err, w)
		return
	}
	defer request.Body.Close()

	var exchangeTokenBody strava.ExchangeTokenResponseBody

	err = json.NewDecoder(request.Body).Decode(&exchangeTokenBody)
	if err != nil {
		utils.HandleHttpError(err, w)
		return
	}

	if exchangeTokenBody.TokenType == "" {
		// Strava token request failed
		utils.HandleHttpError(errors.New("invalid request"), w)
		return
	}

	user, err := repository.FindOrCreateUserByStrava(config.DB, &exchangeTokenBody.Athlete)
	if err != nil {
		utils.HandleHttpError(err, w)
		return
	}

	jwtString, err := config.GenerateJwtToken(&exchangeTokenBody)
	if err != nil {
		utils.HandleHttpError(err, w)
		return
	}

	_, err = config.ValidateJwtToken(jwtString)
	if err != nil {
		utils.HandleHttpError(err, w)
		return
	}

	loginResponse := dto.LoginSuccessResponse{
		AccessToken: exchangeTokenBody.AccessToken,
		Name:        fmt.Sprintf("%s %s", user.FirstName, user.LastName),
	}

	cookie := http.Cookie{
		Name:     "refreshTokenCookie",
		Value:    exchangeTokenBody.RefreshToken,
		Expires:  time.Unix(exchangeTokenBody.ExpiresAt, 0),
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
	}

	fmt.Printf("COOKIE BEING SET %+v\n", &cookie)

	http.SetCookie(w, &cookie)

	jsonResponse, err := json.Marshal(&loginResponse)
	if err != nil {
		utils.HandleHttpError(err, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
