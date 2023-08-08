package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"github.com/zerefwayne/rooots/server/config"
	"github.com/zerefwayne/rooots/server/routes"
)

func main() {
	config.LoadEnvVariables()

	config.ConnectDB()
	defer config.CloseDB()
	config.PingDB()
	config.AutoMigrateModels()

	r := routes.NewRouter()

	handler := cors.Default().Handler(r)

	if err := http.ListenAndServe(":8081", handler); err != nil {
		log.Fatalln(err)
	} else {
		log.Println("server	listening on port :8081")
	}
}
