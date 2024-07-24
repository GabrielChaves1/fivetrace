package adapters

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	
	"luminog.com/iam_serviceluminog.com/iam_service/internal/domain"
	"luminog.com/iam_serviceluminog.com/iam_service/internal/ports"
	"luminog.com/iam_serviceluminog.com/iam_service/internal/utils"
)

type CognitoIdentityProvider struct {
	client       *cognitoidentityprovider.Client
	clientId     string
	clientSecret string
	userPoolId   string
}

type CognitoIdentityProviderConfig struct {
	ClientId     string
	ClientSecret string
	UserPoolId   string
}

func NewCognitoIdentityProvider(client *cognitoidentityprovider.Client, config *CognitoIdentityProviderConfig) ports.IdentityProvider {
	return &CognitoIdentityProvider{
		client:       client,
		clientId:     config.ClientId,
		clientSecret: config.ClientSecret,
		userPoolId:   config.UserPoolId,
	}
}

func (a *CognitoIdentityProvider) SignUpUser(ctx context.Context, email, password string) (string, error) {
	input := &cognitoidentityprovider.SignUpInput{
		ClientId:   aws.String(a.clientId),
		SecretHash: aws.String(utils.ComputeSecretHash(a.clientId, a.clientSecret, email)),
		Username:   aws.String(email),
		Password:   aws.String(password),
		UserAttributes: []types.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String(email),
			},
			{
				Name:  aws.String("custom:role"),
				Value: aws.String(domain.Manager.String()),
			},
		},
	}

	data, err := a.client.SignUp(context.TODO(), input)

	if err != nil {
		return "", err
	}

	return *data.UserSub, nil
}

func (a *CognitoIdentityProvider) ConfirmSignUp(ctx context.Context, sub string) error {
	confirmInput := &cognitoidentityprovider.AdminConfirmSignUpInput{
		UserPoolId: &a.userPoolId,
		Username:   aws.String(sub),
	}

	_, err := a.client.AdminConfirmSignUp(context.Background(), confirmInput)

	if err != nil {
		return err
	}

	input := &cognitoidentityprovider.AdminUpdateUserAttributesInput{
		UserPoolId: &a.userPoolId,
		Username:   aws.String(sub),
		UserAttributes: []types.AttributeType{
			{
				Name:  aws.String("email_verified"),
				Value: aws.String("true"),
			},
		},
	}

	_, err = a.client.AdminUpdateUserAttributes(ctx, input)

	if err != nil {
		return err
	}

	return nil
}

func (a *CognitoIdentityProvider) SignInUser(ctx context.Context, email, password string) (string, error) {
	input := &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: types.AuthFlowTypeUserPasswordAuth,
		ClientId: aws.String(a.clientId),
		AuthParameters: map[string]string{
			"USERNAME":    email,
			"PASSWORD":    password,
			"SECRET_HASH": utils.ComputeSecretHash(a.clientId, a.clientSecret, email),
		},
	}

	res, err := a.client.InitiateAuth(ctx, input)

	if err != nil {
		return "", err
	}

	type AuthenticationResult struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		IdToken      string `json:"id_token"`
		TokenType    string `json:"token_type"`
	}

	authResult := AuthenticationResult{
		AccessToken:  *res.AuthenticationResult.AccessToken,
		RefreshToken: *res.AuthenticationResult.RefreshToken,
		IdToken:      *res.AuthenticationResult.IdToken,
		TokenType:    *res.AuthenticationResult.TokenType,
	}

	authResultJSON, err := json.Marshal(authResult)

	if err != nil {
		return "", err
	}

	return string(authResultJSON), nil
}

func (a *CognitoIdentityProvider) CreateApplication(ctx context.Context, name string) (*ports.IdentityProviderApplication, error) {
	input := &cognitoidentityprovider.CreateUserPoolClientInput{
		UserPoolId:     &a.userPoolId,
		ClientName:     aws.String(name),
		GenerateSecret: *aws.Bool(true),
		AllowedOAuthFlows: []types.OAuthFlowType{
			types.OAuthFlowTypeClientCredentials,
		},
		AllowedOAuthScopes: []string{
			"fivelogs/write",
		},
		AllowedOAuthFlowsUserPoolClient: *aws.Bool(true),
	}

	data, err := a.client.CreateUserPoolClient(ctx, input)

	if err != nil {
		return nil, err
	}

	return &ports.IdentityProviderApplication{
		ClientName: *data.UserPoolClient.ClientName,
		ClientId:   *data.UserPoolClient.ClientId,
	}, nil
}

func (a *CognitoIdentityProvider) GetUser(ctx context.Context, accessToken string) (*cognitoidentityprovider.GetUserOutput, error) {
	input := &cognitoidentityprovider.GetUserInput{
		AccessToken: aws.String(accessToken),
	}

	data, err := a.client.GetUser(ctx, input)

	if err != nil {
		return nil, err
	}

	return data, nil
}
