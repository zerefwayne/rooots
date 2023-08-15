package routes

import (
	"github.com/gorilla/mux"
	"github.com/zerefwayne/rooots/server/handlers"
	"github.com/zerefwayne/rooots/server/middleware"
)

func LoadStravaAuthRoutes(r *mux.Router) {
	r.HandleFunc("/auth/strava/login", handlers.StravaLoginHandler).Methods("GET")
	r.HandleFunc("/auth/strava/exchangeToken", handlers.ExchangeTokenHandler).Methods("POST")
	r.HandleFunc("/auth/strava/refreshToken", middleware.Authorize(handlers.RefreshTokenHandler)).Methods("GET")
}

func LoadStravaApiRoutes(r *mux.Router) {
	r.HandleFunc("/strava/user", middleware.Authorize(handlers.GetUser)).Methods("GET")
	r.HandleFunc("/strava/activities", middleware.Authorize(handlers.GetActivities)).Methods("GET")
}

func LoadAuthRoutes(r *mux.Router) {
	r.HandleFunc("/auth/logout", handlers.LogoutHandler).Methods("GET")
}
