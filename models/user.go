package models

type User struct {
	Username string `json:"username"`
	Salt     string `json:"salt"`
	Hash     string `json:"hash"`
}

func NewUser(username, salt, hash string) *User {
	return &User{username, salt, hash}
}
