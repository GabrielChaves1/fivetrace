package validators

import (
	"fmt"
	"regexp"
	"strings"

	"luminog.com/email_service/internal/application/dto"
)

func ValidateEmailMessageDTO(emailMessageDto dto.EmailMessageDTO) error {
	if !isValidEmail(emailMessageDto.To) {
		return fmt.Errorf("invalid email address: %s", emailMessageDto.To)
	}

	if !isValidEmailType(emailMessageDto.EmailType) {
		return fmt.Errorf("invalid email type: %s (should be: %v)", emailMessageDto.EmailType, strings.Join(dto.EmailTypeStrings, ", "))
	}

	if emailMessageDto.Data == nil {
		return fmt.Errorf("invalid data: %v", emailMessageDto.Data)
	}

	return nil
}

func isValidEmail(email string) bool {
	const emailRegex = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

func isValidEmailType(emailType string) bool {
	for _, value := range dto.EmailTypeStrings {
		if value == emailType {
			return true
		}
	}

	return false
}
