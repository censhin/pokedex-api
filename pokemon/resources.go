package pokemon

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"gopkg.in/mgo.v2"
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
	case "PUT":
		PutMember(w, r)
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

func PutMember(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	body := make(map[string]interface{})
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Could not parse JSON to Pokemon.", 422)
	}

	err = PutMemberService(vars["id"], body)
	if err == mgo.ErrNotFound {
		http.Error(w, "Not Found", 404)
	} else if err != nil {
		http.Error(w, "Bad Request", 400)
	} else {
		w.WriteHeader(204)
	}
}
