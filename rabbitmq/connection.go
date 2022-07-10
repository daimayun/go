package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

func connection() (conn *amqp.Connection, err error) {
	return
}

func NewConnection() (conn *amqp.Connection, err error) {
	return connection()
}
