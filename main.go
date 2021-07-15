package main

import (
	"encoding/json"
	"net/http"

	"github.com/achimonchi/CircleCi/controller"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handleHome).Methods("GET")

	// http.Serve(":4000", r)
	// http.Handle("/", r)
	http.ListenAndServe(":4000", r)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	data := controller.GetData()

	json.NewEncoder(w).Encode(data)
}
