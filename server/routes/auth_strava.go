package routes

import (
	"github.com/gorilla/mux"
	"github.com/zerefwayne/rooots/server/handlers"
	"github.com/zerefwayne/rooots/server/middleware"
)

func LoadStravaAuthRoutes(r *mux.Router) {
	r.HandleFunc("/api/auth/strava/login", handlers.StravaLoginHandler).Methods("GET")
	r.HandleFunc("/api/auth/strava/exchangeToken", handlers.ExchangeTokenHandler).Methods("POST")
	r.HandleFunc("/api/auth/strava/refreshToken", middleware.Authorize(handlers.RefreshTokenHandler)).Methods("GET")
}

func LoadStravaApiRoutes(r *mux.Router) {
	r.HandleFunc("/api/strava/user", middleware.Authorize(handlers.GetUser)).Methods("GET")
	r.HandleFunc("/api/strava/activities", middleware.Authorize(handlers.GetActivities)).Methods("GET")
	r.HandleFunc("/api/strava/activities/{activityId}", middleware.Authorize(handlers.GetActivity)).Methods("GET")
}

func LoadAuthRoutes(r *mux.Router) {
	r.HandleFunc("/api/auth/logout", handlers.LogoutHandler).Methods("GET")
}
