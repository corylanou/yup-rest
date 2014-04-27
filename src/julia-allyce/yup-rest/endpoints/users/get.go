package users

import (
	"encoding/json"
	"julia-allyce/yup-rest/models"
	"log"
	"net/http"
	"strconv"
)

func Get(w http.ResponseWriter, r *http.Request) {
	user := models.User{
		First:  "Julia",
		Last:   "Poladsky",
		Email:  "a@b.c",
		Active: 1,
	}

	data, err := json.Marshal(user)

	if err != nil {
		log.Printf("Error marshalling json: %s", err)
		return
	}

	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
