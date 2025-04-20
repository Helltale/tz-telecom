package models

type UserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	Password  string `json:"password"`
	IsMarried bool   `json:"is_married"`
}
