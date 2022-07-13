package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

// Queue 队列接口
type Queue interface {
	Send(data SendData) (err error)
	Receive(data ReceiveData) (messages <-chan amqp.Delivery, ch *amqp.Channel, err error)
}

// Connection RabbitMQ连接
type Connection struct {
	Conn *amqp.Connection
}

// connection RabbitMQ连接
func connection(url string) (conn *amqp.Connection, err error) {
	return amqp.Dial(url)
}

// NewConnection 实例化RabbitMQ连接
func NewConnection(url string) (conn Connection, err error) {
	var c *amqp.Connection
	c, err = connection(url)
	if err != nil {
		return
	}
	conn = Connection{Conn: c}
	return
}
