package model

type SignUpUserDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInUserDTO struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
type UserEntityDomain struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type UserAcces struct {
	User  UserEntityDomain `json:"user"`
	Token string           `json:"token"`
}
