package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

type RoutingTypePublishData struct {
	Exchange   string          `json:"exchange"`
	RoutingKey string          `json:"routing_key"`
	Publishing amqp.Publishing `json:"publishing"`
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

func (conn Connection) RoutingTypeReceive() (err error) {
	return
}
