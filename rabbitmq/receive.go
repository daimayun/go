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

func (conn Connection) Receive(data ReceiveData) (messages <-chan amqp.Delivery, err error) {
	var (
		ch *amqp.Channel
		q  amqp.Queue
	)
	ch, err = conn.Conn.Channel()
	if err != nil {
		return
	}
	defer ch.Close()
	if data.Type != "" {
		err = ch.ExchangeDeclare(data.Exchange, data.Type, data.Durable, data.AutoDelete, data.Internal, data.NoWait, data.Args)
		if err != nil {
			return
		}
	}
	q, err = ch.QueueDeclare(data.QueueName, data.Durable, data.DeleteWhenUnused, data.Exclusive, data.NoWait, data.Args)
	if err != nil {
		return
	}
	if data.Type != "" {
		err = ch.QueueBind(q.Name, data.RoutingKey, data.Exchange, data.NoWait, data.Args)
		if err != nil {
			return
		}
	}
	messages, err = ch.Consume(q.Name, data.Consumer, data.AutoAck, data.Exclusive, data.NoLocal, data.NoWait, data.Args)
	return
}