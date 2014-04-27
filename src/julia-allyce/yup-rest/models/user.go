package models

type User struct {
	First  string `json:"first"`
	Last   string `json:"last"`
	Email  string `json:"email"`
	Active int    `json:"active"`
}
