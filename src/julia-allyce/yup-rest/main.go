package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"julia-allyce/yup-rest/endpoints/users"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hiiiiiiiiii")
	})
	r.HandleFunc("/users", users.Get).Methods("GET")
	r.HandleFunc("/users", users.Post).Methods("POST")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
