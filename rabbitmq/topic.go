package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

type TopicTypePublishData struct {
	Exchange   string          `json:"exchange"`
	RoutingKey string          `json:"routing_key"`
	Publishing amqp.Publishing `json:"publishing"`
}

type TopicTypeReceiveData struct {
	Exchange   string `json:"exchange"`
	RoutingKey string `json:"routing_key"`
	AutoAck    bool   `json:"auto_ack"`
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

func (conn Connection) TopicTypeReceive(data TopicTypeReceiveData) (messages <-chan amqp.Delivery, err error) {
	return conn.Receive(ReceiveData{
		Exchange:   data.Exchange,
		RoutingKey: data.RoutingKey,
		Type:       amqp.ExchangeTopic,
		AutoAck:    data.AutoAck,
	})
}
