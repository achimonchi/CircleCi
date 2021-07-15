package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handleHome).Methods("GET")

	// http.Serve(":4000", r)
	// http.Handle("/", r)
	http.ListenAndServe(":4000", r)
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	var user1 User
	user1.Email = "reyhan.jovie@dana.id"
	user1.Name = "Reyhan Jovie Dwiputra"

	json.NewEncoder(w).Encode(user1)
}
