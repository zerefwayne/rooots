package main

import (
	"encoding/json"
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

type ExchangeTokenBody struct {
	Code string
}

func ExchangeTokenHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Exchange token requested!")

	var body ExchangeTokenBody

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("Authentication code is: %s\n", body.Code)

	STRAVA_CLIENT_ID := os.Getenv("STRAVA_CLIENT_ID")
	STRAVA_CLIENT_SECRET := os.Getenv("STRAVA_CLIENT_SECRET")

	stravaExchangeTokenUri := fmt.Sprintf("https://www.strava.com/api/v3/oauth/token?client_id=%s&client_secret=%s&code=%s&grant_type=authorization_code", STRAVA_CLIENT_ID, STRAVA_CLIENT_SECRET, body.Code)

	request, err := http.Post(stravaExchangeTokenUri, "application/json", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer request.Body.Close()

	var exchangeTokenBody models.ExchangeTokenResponseBody

	err = json.NewDecoder(request.Body).Decode(&exchangeTokenBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	foundUser, err := repository.FindUserByStravaId(config.DB, exchangeTokenBody.Athlete.Id)
	if err != nil {
		// Cannot find user
		newUser, err := repository.CreateUserByStravaLogin(config.DB, &exchangeTokenBody.Athlete)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		jsonResponse, jsonError := json.Marshal(newUser)
		if jsonError != nil {
			http.Error(w, jsonError.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
		return
	}

	jsonResponse, jsonError := json.Marshal(&foundUser)
	if jsonError != nil {
		http.Error(w, jsonError.Error(), http.StatusBadRequest)
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
