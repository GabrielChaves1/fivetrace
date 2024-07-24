package ports

import "context"

type AuthTokenTable interface {
	PutToken(ctx context.Context, sub, token string) error
	GetToken(ctx context.Context, sub string) (string, error)
	DeleteToken(ctx context.Context, sub string) error
}