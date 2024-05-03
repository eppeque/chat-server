package models

import (
	"encoding/json"
	"errors"
	"os"
	"sync"
)

type UserRepository struct {
	Users map[string]*User `json:"users"`
	mu    sync.Mutex
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

func (r *UserRepository) GetUser(username string) *User {
	return r.Users[username]
}

func (r *UserRepository) AddUser(user *User) error {
	if r.isUsernameTaken(user.Username) {
		return errors.New("the username is already taken")
	}

	r.mu.Lock()
	r.Users[user.Username] = user
	r.mu.Unlock()

	r.writeToFile()
	return nil
}

func (r *UserRepository) isUsernameTaken(username string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, user := range r.Users {
		if user.Username == username {
			return true
		}
	}

	return false
}

func (r *UserRepository) writeToFile() {
	r.mu.Lock()
	bytes, err := json.Marshal(r)
	r.mu.Unlock()

	if err != nil {
		return
	}

	os.WriteFile(fileName, bytes, 0666)
}

func initEmpty() *UserRepository {
	repo := &UserRepository{Users: make(map[string]*User)}
	repo.writeToFile()
	return repo
}
