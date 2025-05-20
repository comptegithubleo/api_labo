package models

type User struct {
	ID       int
	UUID     string
	Username string `json:"username" validate:"required,alphanum"`
	Password string `json:"password" validate:"required,min=8,max=50"`
	Email    string `json:"email" validate:"required,email"`
}
