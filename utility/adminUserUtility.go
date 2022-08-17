package utility

import "github.com/AsliddinTuxtasinov/online-ticket-order/models"

type AdminUserSignUp struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	Password2 string `json:"password2"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func CreateSuperUser(first_name, last_name, email, password string) *models.User {
	return &models.User{Email: email, FirstName: first_name, LirstName: last_name, Password: password, IsAdmin: true}
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
