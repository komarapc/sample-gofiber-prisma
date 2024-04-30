package lib

import (
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func ConvertStringToInt(str string) int {
	number, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return number
}
