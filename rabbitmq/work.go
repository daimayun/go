package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

type WorkTypePublishData struct {
	QueueName  string          `json:"queue_name"`
	RoutingKey string          `json:"routing_key"`
	Publishing amqp.Publishing `json:"publishing"`
}

type WorkTypeReceiveData struct {
	QueueName        string `json:"queue_name"`
	RoutingKey       string `json:"routing_key"`
	AutoAck          bool   `json:"auto_ack"`
	QosPrefetchCount int    `json:"qos_prefetch_count"`
	QosPrefetchSize  int    `json:"qos_prefetch_size"`
	QosGlobal        bool   `json:"qos_global"`
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

func (conn Connection) WorkTypeReceive(data WorkTypeReceiveData) (messages <-chan amqp.Delivery, err error) {
	return conn.Receive(ReceiveData{
		QueueName:        data.QueueName,
		Durable:          true,
		Args:             nil,
		AutoAck:          data.AutoAck,
		QosPrefetchCount: data.QosPrefetchCount,
		QosPrefetchSize:  data.QosPrefetchSize,
		QosGlobal:        data.QosGlobal,
	})
}
