package lib

import (
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
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

func IsZeroTime(t time.Time) bool {
	return t.IsZero()
}

func TimeDeletedAt(deletedAt time.Time) *time.Time {
	var timeDeletedAt *time.Time
	if !IsZeroTime(deletedAt) {
		timeDeletedAt = &deletedAt
	}
	return timeDeletedAt
}
