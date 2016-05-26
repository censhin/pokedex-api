package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Starting server on port 8000.")
	http.ListenAndServe(":8000", InitRoutes())
}
