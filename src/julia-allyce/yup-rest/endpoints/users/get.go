package users

import (
	"julia-allyce/yup-rest/common/json_helpers"
	"julia-allyce/yup-rest/models"
	"net/http"
)

func Get(w http.ResponseWriter, r *http.Request) {
	user := models.User{
		First:  "Julia",
		Last:   "Poladsky",
		Email:  "a@b.c",
		Active: 1,
	}

	json_helpers.Write(w, user)
}
