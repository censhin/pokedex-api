package main

import (
	"github.com/censhin/pokedex-api/moves"
	"github.com/censhin/pokedex-api/pokemon"
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/pokemon", pokemon.PokemonsResource)
	router.HandleFunc("/pokemon/{id}", pokemon.PokemonResource)
	router.HandleFunc("/moves", moves.MovesResource)
	router.HandleFunc("/moves/{id}", moves.MoveResource)

	return router
}
