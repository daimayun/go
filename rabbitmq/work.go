package rabbitmq

import (
	"errors"
	amqp "github.com/rabbitmq/amqp091-go"
	"strings"
)

type WorkTypeSendData struct {
	QueueName           string          `json:"queue_name"`
	RoutingKey          string          `json:"routing_key"`
	QueueDeclareArgs    amqp.Table      `json:"queue_declare_args"`
	QueueDeclareDurable bool            `json:"queue_declare_durable"` // true
	Publishing          amqp.Publishing `json:"publishing"`
}

type WorkTypeReceiveData struct {
	QueueName           string     `json:"queue_name"`
	RoutingKey          string     `json:"routing_key"`
	AutoAck             bool       `json:"auto_ack"`
	QosPrefetchCount    int        `json:"qos_prefetch_count"`
	QosPrefetchSize     int        `json:"qos_prefetch_size"`
	QosGlobal           bool       `json:"qos_global"`
	QueueDeclareArgs    amqp.Table `json:"queue_declare_args"`
	ConsumeArgs         amqp.Table `json:"consume_args"`
	QueueDeclareDurable bool       `json:"queue_declare_durable"` // true
}

func (conn Connection) WorkTypeSend(data WorkTypeSendData) (err error) {
	if strings.Trim(data.RoutingKey, " ") == "" {
		err = errors.New("routing key cannot be empty")
		return
	}
	var queueDeclareArgs amqp.Table = nil
	if len(data.QueueDeclareArgs) > 0 {
		queueDeclareArgs = data.QueueDeclareArgs
	}
	return conn.Send(SendData{
		QueueName:           data.QueueName,
		RoutingKey:          data.RoutingKey,
		QueueDeclareArgs:    queueDeclareArgs,
		QueueDeclareDurable: data.QueueDeclareDurable,
		Publishing:          data.Publishing,
	})
}

func (conn Connection) WorkTypeReceive(data WorkTypeReceiveData) (messages <-chan amqp.Delivery, ch *amqp.Channel, err error) {
	var queueDeclareArgs amqp.Table = nil
	if len(data.QueueDeclareArgs) > 0 {
		queueDeclareArgs = data.QueueDeclareArgs
	}
	var consumeArgs amqp.Table = nil
	if len(data.ConsumeArgs) > 0 {
		consumeArgs = data.ConsumeArgs
	}
	return conn.Receive(ReceiveData{
		QueueName:           data.QueueName,
		QueueDeclareDurable: data.QueueDeclareDurable,
		QueueDeclareArgs:    queueDeclareArgs,
		ConsumeArgs:         consumeArgs,
		AutoAck:             data.AutoAck,
		QosPrefetchCount:    data.QosPrefetchCount,
		QosPrefetchSize:     data.QosPrefetchSize,
		QosGlobal:           data.QosGlobal,
	})
}
