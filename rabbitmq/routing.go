package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

type RoutingTypePublishData struct {
	Exchange   string          `json:"exchange"`
	RoutingKey string          `json:"routing_key"`
	Publishing amqp.Publishing `json:"publishing"`
}

type RoutingTypeReceiveData struct {
	Exchange   string `json:"exchange"`
	RoutingKey string `json:"routing_key"`
	AutoAck    bool   `json:"auto_ack"`
}

func (conn Connection) RoutingTypePublish(data RoutingTypePublishData) (err error) {
	return conn.Publish(PublishData{
		Exchange:   data.Exchange,
		RoutingKey: data.RoutingKey,
		Type:       amqp.ExchangeDirect,
		Args:       nil,
		Durable:    true,
		Publishing: data.Publishing,
	})
}

func (conn Connection) RoutingTypeReceive(data RoutingTypeReceiveData) (messages <-chan amqp.Delivery, err error) {
	return conn.Receive(ReceiveData{
		Exchange:   data.Exchange,
		RoutingKey: data.RoutingKey,
		Type:       amqp.ExchangeDirect,
		Durable:    true,
		AutoAck:    data.AutoAck,
	})
}
