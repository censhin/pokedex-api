package main

import (
	"log"
	"net/http"
	"github.com/censhin/pokedex-api/db"
)

func main() {
	log.Println("Connecting to database.")
	db.Connect()
	defer db.Close()

	log.Println("Starting server on port 8000.")
	http.ListenAndServe(":8000", InitRoutes())
}
