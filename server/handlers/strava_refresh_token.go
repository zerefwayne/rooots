package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	strava "github.com/zerefwayne/rooots/server/dto/strava"
	"github.com/zerefwayne/rooots/server/utils"
)

func RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("refreshTokenCookie")
	if err != nil {
		utils.HandleHttpError(err, w)
		return
	}

	log.Printf("REFRESH TOKEN recieved from client: %+v\n", cookie)

	STRAVA_CLIENT_ID := os.Getenv("STRAVA_CLIENT_ID")
	STRAVA_CLIENT_SECRET := os.Getenv("STRAVA_CLIENT_SECRET")
	refreshToken := cookie.Value

	stravaRefreshTokenUri := fmt.Sprintf("https://www.strava.com/api/v3/oauth/token?client_id=%s&client_secret=%s&grant_type=refresh_token&refresh_token=%s", STRAVA_CLIENT_ID, STRAVA_CLIENT_SECRET, refreshToken)

	request, err := http.Post(stravaRefreshTokenUri, "application/json", nil)
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

	loginResponse := strava.LoginSuccessResponse{
		AccessToken: exchangeTokenBody.AccessToken,
	}

	newCookie := http.Cookie{
		Name:     "refreshTokenCookie",
		Value:    exchangeTokenBody.RefreshToken,
		Expires:  time.Unix(exchangeTokenBody.ExpiresAt, 0),
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
	}

	fmt.Printf("NEW COOKIE BEING SET %+v\n", &newCookie)
	http.SetCookie(w, &newCookie)

	jsonResponse, err := json.Marshal(&loginResponse)
	if err != nil {
		utils.HandleHttpError(err, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
