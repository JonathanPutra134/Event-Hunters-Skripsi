package dto

type UserRegistrationRequest struct {
	FullName    string
	Email       string
	Password    string
	PhoneNumber string
	Address     string
	Latitude    string
	Longitude   string
}
