package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

type WorkTypePublishData struct {
	QueueName  string          `json:"queue_name"`
	RoutingKey string          `json:"routing_key"`
	Publishing amqp.Publishing `json:"publishing"`
}

func (conn Connection) WorkTypePublish(data WorkTypePublishData) (err error) {
	return conn.Publish(PublishData{
		QueueName:  data.QueueName,
		RoutingKey: data.RoutingKey,
		Args:       nil,
		Durable:    true,
		Publishing: data.Publishing,
	})
}

func (conn Connection) WorkTypeReceive() (err error) {
	return
}
