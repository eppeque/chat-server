package models

import (
	"errors"
	"math/rand"
)

type AuthManager struct {
	challenge string
	username  string
}

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func NewAuthManager() *AuthManager {
	return &AuthManager{challenge: generateChallenge()}
}

func (a *AuthManager) Register(username, salt, hash string) error {
	user := NewUser(username, salt, hash)
	err := UserRepo.AddUser(user)

	if err != nil {
		return err
	}

	a.username = username
	return nil
}

func (a *AuthManager) Login(username string) (string, error) {
	user := UserRepo.GetUser(username)

	if user != nil {
		return user.Salt, nil
	}

	return "", errors.New("no user found with the given username")
}

func generateChallenge() string {
	challenge := make([]byte, 22)

	for i := 0; i < 22; i++ {
		index := rand.Intn(len(chars))
		challenge[i] = chars[index]
	}

	return string(challenge)
}
