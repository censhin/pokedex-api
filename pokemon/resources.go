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
	case "POST":
		PostCollection(w, r)
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
	case "DELETE":
		DeleteMember(w, r)
	default:
		http.Error(w, "Method not allowed.", 405)
	}
}

func GetCollection(w http.ResponseWriter, r *http.Request) {
	res := CollectionService()
	if res.Code > 399 {
		http.Error(w, res.Msg, res.Code)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res.Body)
	}
}

func PostCollection(w http.ResponseWriter, r *http.Request) {
	body := Pokemon{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Could not parse JSON to Pokemon.", 422)
	}

	res := PostCollectionService(body)
	if res.Code > 399 {
		http.Error(w, res.Msg, res.Code)
	} else {
		w.WriteHeader(res.Code)
	}
}

func GetMember(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	res := MemberService(vars["id"])
	if res.Code > 399 {
		http.Error(w, res.Msg, res.Code)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res.Body)
	}
}

func PutMember(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	body := make(map[string]interface{})
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Could not parse JSON to map.", 422)
	}

	res := PutMemberService(vars["id"], body)
	if res.Code > 399 {
		http.Error(w, res.Msg, res.Code)
	} else {
		w.WriteHeader(res.Code)
	}
}

func DeleteMember(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	res := DeleteMemberService(vars["id"])
	if res.Code > 399 {
		http.Error(w, res.Msg, res.Code)
	} else {
		w.WriteHeader(res.Code)
	}
}
