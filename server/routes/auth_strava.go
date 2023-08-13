package routes

import (
	"github.com/gorilla/mux"
	"github.com/zerefwayne/rooots/server/handlers"
)

func LoadStravaAuthRoutes(r *mux.Router) {
	r.HandleFunc("/auth/strava/login", handlers.LoginHandler).Methods("GET")
	r.HandleFunc("/auth/strava/exchangeToken", handlers.ExchangeTokenHandler).Methods("POST")
	r.HandleFunc("/auth/strava/refreshToken", handlers.RefreshTokenHandler).Methods("GET")
}

func LoadStravaApiRoutes(r *mux.Router) {
	r.HandleFunc("/strava/user", handlers.GetUser).Methods("GET")
}
