package utils

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

const (
	MinPasswordLenght = 8
)

var (
	ErrPasswordTooShort = errors.New("password must be at least 8 characters long")
	ErrHashingFailed    = errors.New("failed to hash password")
)

func HashPassword(password string) (string, error) {
	if len(password) < MinPasswordLenght {
		return "", ErrPasswordTooShort
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", ErrHashingFailed
	}
	return string(hashedPassword), nil
}

func ComparePassword(hashedPassword string, candidatePassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candidatePassword))
}
