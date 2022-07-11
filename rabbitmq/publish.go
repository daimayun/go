package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

// PublishData 发布消息队列所提交数据
type PublishData struct {
	Exchange               string          `json:"exchange"`
	Type                   string          `json:"type"`
	QueueName              string          `json:"queue_name"`
	RoutingKey             string          `json:"routing_key"`
	ExchangeDeclareArgs    amqp.Table      `json:"exchange_declare_args"`
	QueueDeclareArgs       amqp.Table      `json:"queue_declare_args"`
	AutoDelete             bool            `json:"auto_delete"`
	Internal               bool            `json:"internal"`
	NoWait                 bool            `json:"no_wait"`
	Exclusive              bool            `json:"exclusive"`
	Mandatory              bool            `json:"mandatory"`
	Immediate              bool            `json:"immediate"`
	Publishing             amqp.Publishing `json:"publishing"`
	ExchangeDeclareDurable bool            `json:"exchange_declare_durable"`
	QueueDeclareDurable    bool            `json:"queue_declare_durable"`
}

// Publish 发布消息队列
func (conn Connection) Publish(data PublishData) (err error) {
	var ch *amqp.Channel
	ch, err = conn.Conn.Channel()
	if err != nil {
		return
	}
	defer ch.Close()
	if data.Type != "" {
		err = ch.ExchangeDeclare(data.Exchange, data.Type, data.ExchangeDeclareDurable, data.AutoDelete, data.Internal, data.NoWait, data.ExchangeDeclareArgs)
		if err != nil {
			return
		}
	} else {
		_, err = ch.QueueDeclare(data.QueueName, data.QueueDeclareDurable, data.AutoDelete, data.Exclusive, data.NoWait, data.QueueDeclareArgs)
		if err != nil {
			return
		}
	}
	err = ch.Publish(data.Exchange, data.RoutingKey, data.Mandatory, data.Immediate, data.Publishing)
	if err != nil {
		return
	}
	return
}
