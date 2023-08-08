package routes

import (
	"github.com/gorilla/mux"
	"github.com/zerefwayne/rooots/server/handlers/auth/strava"
)

func LoadStravaAuthRoutes(r *mux.Router) {
	r.HandleFunc("/auth/strava/login", strava.LoginHandler).Methods("GET")
	r.HandleFunc("/auth/strava/exchangeToken", strava.ExchangeTokenHandler).Methods("POST")
}
