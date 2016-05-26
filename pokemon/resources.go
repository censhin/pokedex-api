package pokemon

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func PokemonsResource(w http.ResponseWriter, r *http.Request) {
	res := PokemonsService()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func PokemonResource(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	res := PokemonService(vars["id"])
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
