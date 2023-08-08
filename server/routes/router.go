package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	LoadStravaAuthRoutes(r)

	r.NotFoundHandler = http.NotFoundHandler()

	return r
}
