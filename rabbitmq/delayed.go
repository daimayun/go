package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

const ExchangeXDelayedMessage = "x-delayed-message"

type DelayedTypePublishData struct {
	Exchange   string          `json:"exchange"`
	RoutingKey string          `json:"routing_key"`
	Args       amqp.Table      `json:"args"`
	Publishing amqp.Publishing `json:"publishing"`
}

type DelayedTypeReceiveData struct {
	Exchange   string `json:"exchange"`
	RoutingKey string `json:"routing_key"`
	AutoAck    bool   `json:"auto_ack"`
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

func (conn Connection) DelayedTypeReceive(data DelayedTypeReceiveData) (messages <-chan amqp.Delivery, err error) {
	return conn.Receive(ReceiveData{
		Exchange:   data.Exchange,
		Type:       ExchangeXDelayedMessage,
		RoutingKey: data.RoutingKey,
		Durable:    true,
		AutoAck:    data.AutoAck,
	})
}
