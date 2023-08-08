package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"github.com/zerefwayne/rooots/server/config"
	"github.com/zerefwayne/rooots/server/models"
	"github.com/zerefwayne/rooots/server/repository"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	STRAVA_CLIENT_ID := os.Getenv("STRAVA_CLIENT_ID")
	STRAVA_REDIRECT_URI := os.Getenv("STRAVA_REDIRECT_URI")
	STRAVA_SCOPE := os.Getenv("STRAVA_SCOPE")

	stravaLoginRedirectUrl := fmt.Sprintf("https://www.strava.com/oauth/authorize?client_id=%s&response_type=code&redirect_uri=%s&approval_prompt=force&scope=%s", STRAVA_CLIENT_ID, STRAVA_REDIRECT_URI, STRAVA_SCOPE)

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, stravaLoginRedirectUrl)
}

func FindOrCreateUserByStrava(athlete *models.SummaryAthlete) (*models.User, error) {
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
	AuthenticationToken string `json:"authenticationToken"`
}

func ExchangeTokenHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Exchange token requested!")

	var body ExchangeTokenBody

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {

		return
	}

	log.Printf("Authentication code is: %s\n", body.Code)

	STRAVA_CLIENT_ID := os.Getenv("STRAVA_CLIENT_ID")
	STRAVA_CLIENT_SECRET := os.Getenv("STRAVA_CLIENT_SECRET")

	stravaExchangeTokenUri := fmt.Sprintf("https://www.strava.com/api/v3/oauth/token?client_id=%s&client_secret=%s&code=%s&grant_type=authorization_code", STRAVA_CLIENT_ID, STRAVA_CLIENT_SECRET, body.Code)

	request, err := http.Post(stravaExchangeTokenUri, "application/json", nil)
	if err != nil {
		HandleHttpError(err, w)
		return
	}
	defer request.Body.Close()

	var exchangeTokenBody models.ExchangeTokenResponseBody

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

	log.Printf("Found User: %+v\n", user)

	jwtString, err := config.GenerateJwtToken(&exchangeTokenBody)
	if err != nil {
		HandleHttpError(err, w)
		return
	}

	log.Println("GeneratedUserToken:", jwtString)

	ok, err := config.ValidateJwtToken(jwtString)
	if err != nil {
		HandleHttpError(err, w)
		return
	}

	log.Println("ValidUserToken:", jwtString, ok)

	authToken := LoginSuccessResponse{AuthenticationToken: jwtString}

	jsonResponse, err := json.Marshal(&authToken)
	if err != nil {
		HandleHttpError(err, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.SetupDB()
	config.AutoMigrateDB()

	router := mux.NewRouter()
	router.HandleFunc("/auth/strava/login", LoginHandler).Methods("GET")
	router.HandleFunc("/auth/strava/exchangeToken", ExchangeTokenHandler).Methods("POST")

	router.NotFoundHandler = http.NotFoundHandler()

	handler := cors.Default().Handler(router)

	if err := http.ListenAndServe(":8081", handler); err != nil {
		panic(err)
	}
}
