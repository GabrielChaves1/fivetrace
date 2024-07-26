package validators

import (
	"fmt"
	"regexp"

	"fivetrace.com/iam_service/internal/application/dtos"
)

func ValidateSignupDTO(signupDto *dtos.SignupDTO) error {
	if !isValidEmail(signupDto.Email) {
		return fmt.Errorf("invalid email address: %s", signupDto.Email)
	}

	if signupDto.OrganizationName == "" {
		return fmt.Errorf("organization name is required")
	}

	if signupDto.Country == "" {
		return fmt.Errorf("country is required")
	}

	return nil
}

func isValidEmail(email string) bool {
	const emailRegex = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}
