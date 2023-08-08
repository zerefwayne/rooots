package routes

import (
	"github.com/gorilla/mux"
	auth "github.com/zerefwayne/rooots/server/handlers/auth"
)

func LoadStravaAuthRoutes(r *mux.Router) {
	r.HandleFunc("/auth/strava/login", auth.LoginHandler).Methods("GET")
	r.HandleFunc("/auth/strava/exchangeToken", auth.ExchangeTokenHandler).Methods("POST")
}
