package ports

type QueueMessageConsumer interface {
	Consume(event interface{}) error
}