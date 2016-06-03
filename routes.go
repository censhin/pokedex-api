package main

import (
	"github.com/censhin/pokedex-api/moves"
	"github.com/censhin/pokedex-api/pokemon"
        "github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/pokemon", pokemon.CollectionResource).Methods("GET", "POST", "OPTIONS")
	router.HandleFunc("/pokemon/{id}", pokemon.MemberResource).Methods("GET", "PUT", "DELETE", "OPTIONS")
	router.HandleFunc("/moves", moves.CollectionResource).Methods("GET")
	router.HandleFunc("/moves/{id}", moves.MemberResource).Methods("GET")

	return router
}
