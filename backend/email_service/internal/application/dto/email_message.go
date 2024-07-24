package dto

import (
	"errors"
	"fmt"

	"luminog.com/email_service/internal/domain"
)

var EmailTypeStrings = []string{
	"confirm_link",
}

type EmailMessageDTO struct {
	EmailType string                 `json:"emailType"`
	To        string                 `json:"to"`
	Data      map[string]interface{} `json:"data"`
}

func getEmailTypeFromString(emailTypeString string) (emailType domain.EmailType, err error) {
	switch emailTypeString {
	case "confirm_link":
		return domain.ConfirmLink, nil
	default:
		return 0, fmt.Errorf("invalid email type")
	}
}

func NewEmailTemplateDataFromDTO(emailMessageDto EmailMessageDTO) (emailTemplateData domain.EmailTemplateData, err error) {
	emailType, err := getEmailTypeFromString(emailMessageDto.EmailType)
	if err != nil {
		return domain.EmailTemplateData{}, errors.Join(fmt.Errorf("couldn't isntantiate EmailTemplateData"), err)
	}

	return domain.EmailTemplateData{
		Type: emailType,
		Data: emailMessageDto.Data,
	}, nil
}
