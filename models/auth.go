package models

import (
	"errors"
	"fmt"
	"math/rand"

	"golang.org/x/crypto/sha3"
)

type AuthManager struct {
	challenge    string
	username     string
	tempUsername string
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
		a.tempUsername = username
		return user.Salt, nil
	}

	return "", errors.New("no user found with the given username")
}

func (a *AuthManager) Confirm(challenge string) bool {
	user := UserRepo.GetUser(a.tempUsername)
	result := fmt.Sprintf("%s$2b$12$%s%s", a.challenge, user.Salt, user.Hash)
	result = hashChallenge(result)
	fmt.Println(challenge, result)

	if challenge == result {
		a.username = user.Username
		a.tempUsername = ""
		return true
	}

	return false
}

func generateChallenge() string {
	challenge := make([]byte, 22)

	for i := 0; i < 22; i++ {
		index := rand.Intn(len(chars))
		challenge[i] = chars[index]
	}

	return string(challenge)
}

func hashChallenge(challenge string) string {
	hash := sha3.New256()
	hash.Write([]byte(challenge))
	sha3 := hash.Sum(nil)

	return fmt.Sprintf("%x", sha3)
}
