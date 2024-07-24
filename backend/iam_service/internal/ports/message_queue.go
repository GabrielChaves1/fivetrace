package ports

import "context"

type MessageQueue interface {
	SendMessage(ctx context.Context, message []byte) error
}
