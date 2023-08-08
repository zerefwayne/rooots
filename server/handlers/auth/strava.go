package strava

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/zerefwayne/rooots/server/config"
	"github.com/zerefwayne/rooots/server/models"
	strava "github.com/zerefwayne/rooots/server/models/strava"
	"github.com/zerefwayne/rooots/server/repository"
)

func FindOrCreateUserByStrava(athlete *strava.SummaryAthlete) (*models.User, error) {
	foundUser, err := repository.FindUserByStravaId(config.DB, athlete.Id)
	if err != nil {
		// Cannot find user
		createdUser, createUserErr := repository.CreateUserByStravaLogin(config.DB, athlete)
		return createdUser, createUserErr
	}

	return foundUser, err
}

func HandleHttpError(err error, w http.ResponseWriter) {
	fmt.Println("Error: ", err.Error())
	http.Error(w, err.Error(), http.StatusBadRequest)
}

type ExchangeTokenBody struct {
	Code string
}

type LoginSuccessResponse struct {
	AuthenticationToken string       `json:"authenticationToken"`
	User                *models.User `json:"user"`
}

func ExchangeTokenHandler(w http.ResponseWriter, r *http.Request) {
	var body ExchangeTokenBody

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		HandleHttpError(err, w)
		return
	}

	STRAVA_CLIENT_ID := os.Getenv("STRAVA_CLIENT_ID")
	STRAVA_CLIENT_SECRET := os.Getenv("STRAVA_CLIENT_SECRET")

	stravaExchangeTokenUri := fmt.Sprintf("https://www.strava.com/api/v3/oauth/token?client_id=%s&client_secret=%s&code=%s&grant_type=authorization_code", STRAVA_CLIENT_ID, STRAVA_CLIENT_SECRET, body.Code)

	request, err := http.Post(stravaExchangeTokenUri, "application/json", nil)
	if err != nil {
		HandleHttpError(err, w)
		return
	}
	defer request.Body.Close()

	var exchangeTokenBody strava.ExchangeTokenResponseBody

	err = json.NewDecoder(request.Body).Decode(&exchangeTokenBody)
	if err != nil {
		HandleHttpError(err, w)
		return
	}

	if exchangeTokenBody.TokenType == "" {
		// Strava token request failed
		HandleHttpError(errors.New("invalid request"), w)
		return
	}

	user, err := FindOrCreateUserByStrava(&exchangeTokenBody.Athlete)
	if err != nil {
		HandleHttpError(err, w)
		return
	}

	jwtString, err := config.GenerateJwtToken(&exchangeTokenBody)
	if err != nil {
		HandleHttpError(err, w)
		return
	}

	_, err = config.ValidateJwtToken(jwtString)
	if err != nil {
		HandleHttpError(err, w)
		return
	}

	authToken := LoginSuccessResponse{AuthenticationToken: jwtString, User: user}

	jsonResponse, err := json.Marshal(&authToken)
	if err != nil {
		HandleHttpError(err, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	STRAVA_CLIENT_ID := os.Getenv("STRAVA_CLIENT_ID")
	STRAVA_REDIRECT_URI := os.Getenv("STRAVA_REDIRECT_URI")
	STRAVA_SCOPE := os.Getenv("STRAVA_SCOPE")

	stravaLoginRedirectUrl := fmt.Sprintf("https://www.strava.com/oauth/authorize?client_id=%s&response_type=code&redirect_uri=%s&approval_prompt=force&scope=%s", STRAVA_CLIENT_ID, STRAVA_REDIRECT_URI, STRAVA_SCOPE)

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, stravaLoginRedirectUrl)
}