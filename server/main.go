package main

import (
	"fmt"
	"log"

	"database/sql"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "roootsadmin"
	password = "roootspw"
	dbname   = "roootsdb"
)

func main() {

	pqConnectionString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	log.Println(pqConnectionString)

	db, err := sql.Open("postgres", pqConnectionString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
}
