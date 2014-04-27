package users

import ( 
	"net/http"
	"encoding/json"
	"log"
	"strconv"
	"julia-allyce/yup-rest/models"
)

func Get ( w http.ResponseWriter, r *http.Request ) {
	user := models.User{
		First: "Julia",
		Last: "Poladsky",
		Email: "a@b.c",
		Active: 1,
	}

	data, err := json.Marshal(user)

	if err != nil {
		log.Printf("Error marshalling json: %w", err)
		return
	}

	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}