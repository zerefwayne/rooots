package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
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
	handler := routes.NewCorsConfiguration().Handler(r)

	port := fmt.Sprintf(":%s", os.Getenv("APPLICATION_PORT"))

	if config.IsEnvironmentHeroku() {
		port = fmt.Sprintf(":%s", os.Getenv("PORT"))
	}

	log.Printf("server	listening on port %s\n", port)
	if err := http.ListenAndServe(port, handler); err != nil {
		log.Fatalln(err)
	}
}
