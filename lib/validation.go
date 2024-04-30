package lib

import (
	"fmt"
	"regexp"
	"strings"
)

type ValidationRule func(interface{}) bool

type ValidationRules map[string]ValidationRule

type ValidationResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ValidateRequest(request map[string]interface{}, rules ValidationRules) []ValidationResponse {
	var errors []ValidationResponse
	for key, value := range request {
		if rule, ok := rules[key]; ok {
			if !rule(value) {
				lowerKey := strings.ToLower(key)
				errors = append(errors, ValidationResponse{Field: lowerKey, Message: fmt.Sprintf("Invalid value for %s", lowerKey)})
			}
		}
	}
	return errors
}

func ValidateEmail(email string) bool {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	validateRegex, _ := regexp.MatchString(regex, email)
	return validateRegex
}
