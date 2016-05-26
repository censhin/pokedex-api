package moves

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

func MovesResource(w http.ResponseWriter, r *http.Request) {
	res := MovesService()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func MoveResource(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	res := MoveService(vars["id"])
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
