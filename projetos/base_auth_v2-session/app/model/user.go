package model

type User struct {
	ID       string `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}
type SignUpUserDTO struct {
	Name     string `json:"name" db:"name"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}
type SignInUserDTO struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}
