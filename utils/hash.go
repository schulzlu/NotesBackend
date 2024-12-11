package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}

func ComparePassword(unhashedPassword , hashedPassword string)error{
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(unhashedPassword))

	if err != nil {
		return err
	}

	return nil
}