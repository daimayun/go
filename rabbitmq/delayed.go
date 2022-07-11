package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

const ExchangeXDelayedMessage = "x-delayed-message"

type DelayedTypePublishData struct {
	Exchange   string          `json:"exchange"`
	RoutingKey string          `json:"routing_key"`
	Args       amqp.Table      `json:"args"`
	Publishing amqp.Publishing `json:"publishing"`
}

func (conn Connection) DelayedTypePublish(data DelayedTypePublishData) (err error) {
	return conn.Publish(PublishData{
		Exchange:   data.Exchange,
		Type:       ExchangeXDelayedMessage,
		RoutingKey: data.RoutingKey,
		Args:       data.Args,
		Durable:    true,
		Publishing: data.Publishing,
	})
}

func (conn Connection) DelayedTypeReceive() (messages <-chan amqp.Delivery, err error) {
	return
}
