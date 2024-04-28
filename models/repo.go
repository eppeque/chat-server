package models

import (
	"encoding/json"
	"errors"
	"os"
)

type UserRepository struct {
	Users []*User `json:"users"`
}

const fileName = "users.json"

var UserRepo *UserRepository = newRepoFromFile()

func newRepoFromFile() *UserRepository {
	content, err := os.ReadFile(fileName)

	if err != nil {
		return initEmpty()
	}

	var repo UserRepository
	err = json.Unmarshal(content, &repo)

	if err != nil {
		return initEmpty()
	}

	return &repo
}

func (r *UserRepository) AddUser(user *User) error {
	if r.isUsernameTaken(user.Username) {
		return errors.New("the username is already taken")
	}

	updated := append(r.Users, user)
	r.Users = updated

	r.writeToFile()
	return nil
}

func (r *UserRepository) isUsernameTaken(username string) bool {
	for _, user := range r.Users {
		if user.Username == username {
			return true
		}
	}

	return false
}

func (r *UserRepository) writeToFile() {
	bytes, err := json.Marshal(r)

	if err != nil {
		return
	}

	os.WriteFile(fileName, bytes, 0666)
}

func initEmpty() *UserRepository {
	repo := &UserRepository{[]*User{}}
	repo.writeToFile()
	return repo
}
