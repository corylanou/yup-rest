package main

import (
	"net/http"
	"log"
	"fmt"
	"github.com/gorilla/mux"
	"julia-allyce/yup-rest/endpoints/users"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func( w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hiiiiiiiiii")
	})
	r.HandleFunc("/users", users.Get).Methods("GET")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
