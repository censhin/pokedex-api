package pokemon

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func CollectionResource(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetCollection(w, r)
	default:
		http.Error(w, "Method not allowed.", 405)
	}
}

func MemberResource(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetMember(w, r)
	default:
		http.Error(w, "Method not allowed.", 405)
	}
}

func GetCollection(w http.ResponseWriter, r *http.Request) {
	res := CollectionService()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func GetMember(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	res := MemberService(vars["id"])
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
