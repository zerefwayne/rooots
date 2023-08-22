package routes

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf(
			"%s %s %s",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)

		next.ServeHTTP(w, r)
	})
}

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	LoadStravaAuthRoutes(r)
	LoadStravaApiRoutes(r)
	LoadAuthRoutes(r)

	r.NotFoundHandler = http.NotFoundHandler()

	r.Use(loggingMiddleware)

	return r
}

func NewCorsConfiguration() *cors.Cors {
	CLIENT_URL := os.Getenv("CLIENT_URL")

	log.Println(CLIENT_URL)

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
