package users

import (
	"julia-allyce/yup-rest/common/json_helpers"
	"julia-allyce/yup-rest/models"
	"log"
	"net/http"
)

// Post - Creates a user if it's valid
func Post(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if ok := json_helpers.Read(r, &user); !ok {
		log.Println("Error unmarshalling json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !user.Valid() {
		json_helpers.Write(w, http.StatusBadRequest, user)
		return
	}
	// Here we would save the users

	// If it didn't save, we would semd some type of error response

	// If it's successful, we send this:
	json_helpers.Write(w, http.StatusCreated, user)
}
