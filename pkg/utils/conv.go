package utils

import (
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func StringToInt(s string) (int, error) {
	result, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
