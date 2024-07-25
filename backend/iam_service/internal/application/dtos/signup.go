package dtos

type SignupDTO struct {
	OrganizationName string `json:"organization_name"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	Country          string `json:"country"`
}
