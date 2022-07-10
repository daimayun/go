package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

// Connection RabbitMQ连接
type Connection struct {
	Conn *amqp.Connection
}

// connection RabbitMQ连接
func connection(url string) (conn *amqp.Connection, err error) {
	return amqp.Dial(url)
}

// NewConnection RabbitMQ连接
func NewConnection(url string) (conn Connection, err error) {
	var c *amqp.Connection
	c, err = connection(url)
	if err != nil {
		return
	}
	conn = Connection{Conn: c}
	return
}
