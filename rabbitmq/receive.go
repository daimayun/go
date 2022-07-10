package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

type ReceiveData struct {
	Exchange         string     `json:"exchange"`
	Type             string     `json:"type"`
	Durable          bool       `json:"durable"`
	AutoDelete       bool       `json:"auto_delete"`
	Internal         bool       `json:"internal"`
	NoWait           bool       `json:"no_wait"`
	Args             amqp.Table `json:"args"`
	DeleteWhenUnused bool       `json:"delete_when_unused"`
	Exclusive        bool       `json:"exclusive"`
	QueueName        string     `json:"queue_name"`
	RoutingKey       string     `json:"routing_key"`
	Consumer         string     `json:"consumer"`
	AutoAck          bool       `json:"auto_ack"`
	NoLocal          bool       `json:"no_local"`
}

func (conn Connection) Receive() (err error) {
	var ch *amqp.Channel
	ch, err = conn.Conn.Channel()
	if err != nil {
		return
	}
	defer ch.Close()
	return
}
