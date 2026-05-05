package models

type User struct {
	Id             string `json:"id"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	Role           string `json:"role"`
	Gender         string `json:"gender"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Street         string `json:"street"`
	PostalCode     string `json:"postal_code"`
	City           string `json:"city"`
	Country        string `json:"country"`
	PhoneNumber    string `json:"phone_number"`
	Company        bool   `json:"company"`
	CompanyName    string `json:"company_name"`
	CompanyUstIdNr string `json:"company_ustidnr"`
}