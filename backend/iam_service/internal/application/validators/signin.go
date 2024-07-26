package validators

import (
	"fmt"

	"fivetrace.com/iam_service/internal/application/dtos"
	"fivetrace.com/iam_service/internal/utils"
)

func ValidateSignInDTO(signInDto dtos.SignInDTO) error {
	if !utils.IsValidEmail(signInDto.Email) {
		return fmt.Errorf("invalid email address: %s", signInDto.Email)
	}

	return nil
}
