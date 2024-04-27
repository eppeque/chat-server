package models

import (
	"math/rand"
)

type AuthManager struct {
	challenge string
}

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func NewAuthManager() *AuthManager {
	return &AuthManager{challenge: generateChallenge()}
}

func (a *AuthManager) Challenge() string {
	return a.challenge
}

func (a *AuthManager) Register(username, salt, hash string) {
	// TODO: Save the user
}

func generateChallenge() string {
	challenge := make([]byte, 22)

	for i := 0; i < 22; i++ {
		index := rand.Intn(len(chars))
		challenge[i] = chars[index]
	}

	return string(challenge)
}
