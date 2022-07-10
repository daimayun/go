package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

// PublishData 发布所提交数据
type PublishData struct {
	Exchange   string          `json:"exchange"`
	Type       string          `json:"type"`
	QueueName  string          `json:"queue_name"`
	RoutingKey string          `json:"routing_key"`
	Args       amqp.Table      `json:"args"`
	Durable    bool            `json:"durable"`
	AutoDelete bool            `json:"auto_delete"`
	Internal   bool            `json:"internal"`
	NoWait     bool            `json:"no_wait"`
	Exclusive  bool            `json:"exclusive"`
	Mandatory  bool            `json:"mandatory"`
	Immediate  bool            `json:"immediate"`
	Publishing amqp.Publishing `json:"publishing"`
}

// Publish 发布消息
func (conn Connection) Publish(data PublishData) (err error) {
	var ch *amqp.Channel
	ch, err = conn.Conn.Channel()
	if err != nil {
		return
	}
	defer ch.Close()
	if data.Type != "" {
		err = ch.ExchangeDeclare(data.Exchange, data.Type, data.Durable, data.AutoDelete, data.Internal, data.NoWait, data.Args)
		if err != nil {
			return
		}
	} else {
		_, err = ch.QueueDeclare(data.QueueName, data.Durable, data.AutoDelete, data.Exclusive, data.NoWait, nil)
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
