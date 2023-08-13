package routes

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	LoadStravaAuthRoutes(r)

	r.NotFoundHandler = http.NotFoundHandler()

	return r
}

func NewCorsConfiguration() *cors.Cors {
	CLIENT_URL := os.Getenv("CLIENT_URL")

	return cors.New(cors.Options{
		AllowedOrigins: []string{CLIENT_URL},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodOptions,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
}
