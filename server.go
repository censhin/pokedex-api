package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Starting server on port 8001.")
	http.ListenAndServe(":8001", InitRoutes())
}
