package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

type TopicTypePublishData struct {
	Exchange   string          `json:"exchange"`
	RoutingKey string          `json:"routing_key"`
	Publishing amqp.Publishing `json:"publishing"`
}

func (conn Connection) TopicTypePublish(data TopicTypePublishData) (err error) {
	return conn.Publish(PublishData{
		Exchange:   data.Exchange,
		RoutingKey: data.RoutingKey,
		Type:       amqp.ExchangeTopic,
		Args:       nil,
		Durable:    true,
		Publishing: data.Publishing,
	})
}

func (conn Connection) TopicTypeReceive() (err error) {
	return
}
