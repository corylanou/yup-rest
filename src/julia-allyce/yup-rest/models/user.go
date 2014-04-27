package models

type User struct {
	First  string            `json:"first"`
	Last   string            `json:"last"`
	Email  string            `json:"email"`
	Active int               `json:"active"`
	Errors map[string]string `json:"errors,omitempty"`
}

func (user *User) Valid() bool {
	user.Errors = make(map[string]string)
	if len(user.First) == 0 {
		user.Errors["first"] = "can't be blank"
	}
	if len(user.Last) == 0 {
		user.Errors["last"] = "can't be blank"
	}
	if len(user.Email) == 0 {
		user.Errors["Email"] = "can't be blank"
	}
	return len(user.Errors) == 0
}
