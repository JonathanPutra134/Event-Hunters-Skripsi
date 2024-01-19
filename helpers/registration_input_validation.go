package helpers

import (
	"event-hunters/dto"
	"fmt"
	"regexp"
	"text/template"
)

func InputNullValidation(usrReq dto.UserRegistrationRequest) []string {
	var errorMessages []string

	if usrReq.FullName == "" {
		errorMessages = append(errorMessages, "fullName cannot be empty")
	}
	if usrReq.Email == "" {
		errorMessages = append(errorMessages, "email cannot be empty")
	}
	if usrReq.Password == "" {
		errorMessages = append(errorMessages, "password cannot be empty")
	}
	if usrReq.Address == "" {
		errorMessages = append(errorMessages, "address cannot be empty")
	}
	return errorMessages
}
func SanitizeInput(usrReq *dto.UserRegistrationRequest) {
	allowedCharacters := regexp.MustCompile(`^[a-zA-Z0-9\s]+$`)
	if !allowedCharacters.MatchString(usrReq.FullName) {
		fmt.Println("FULL NAME XSS DETECTED")
		usrReq.FullName = template.HTMLEscapeString(usrReq.FullName)
	}
	if !allowedCharacters.MatchString(usrReq.Address) {
		fmt.Println("ADDRESS XSS DETECTED")
		usrReq.Address = template.HTMLEscapeString(usrReq.Address)
	}
	if !allowedCharacters.MatchString(usrReq.PhoneNumber) {
		fmt.Println("Phone Number XSS DETECTED")
		usrReq.PhoneNumber = template.HTMLEscapeString(usrReq.PhoneNumber)
	}
	fmt.Println("Input successfully Sanitized ")
}
