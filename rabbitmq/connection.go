package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

// connection RabbitMQ连接
func connection(url string) (conn *amqp.Connection, err error) {
	return amqp.Dial(url)
}

// NewConnection RabbitMQ连接
func NewConnection(url string) (conn *amqp.Connection, err error) {
	return connection(url)
}
