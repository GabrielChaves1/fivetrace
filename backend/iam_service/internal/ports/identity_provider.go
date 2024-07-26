package ports

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

type IdentityProviderApplication struct {
	ClientName string
	ClientId   string
}

type IdentityProvider interface {
	ConfirmSignUp(ctx context.Context, email string) error
	SignUpUser(email, password, organizationName, country string) (string, error)
	SignInUser(ctx context.Context, email, password string) (string, error)
	CreateApplication(ctx context.Context, name string) (*IdentityProviderApplication, error)
	
	UpdateUserAttributes(sub string,attributes []types.AttributeType) error
	GetUser(ctx context.Context, accessToken string) (*cognitoidentityprovider.GetUserOutput, error)
}
