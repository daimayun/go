package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

type TopicTypePublishData struct {
	Exchange               string          `json:"exchange"`
	RoutingKey             string          `json:"routing_key"`
	Publishing             amqp.Publishing `json:"publishing"`
	ExchangeDeclareArgs    amqp.Table      `json:"exchange_declare_args"`
	ExchangeDeclareDurable bool            `json:"exchange_declare_durable"`
}

type TopicTypeReceiveData struct {
	Exchange            string     `json:"exchange"`
	RoutingKey          string     `json:"routing_key"`
	AutoAck             bool       `json:"auto_ack"`
	ExchangeDeclareArgs amqp.Table `json:"exchange_declare_args"`
	QueueDeclareArgs    amqp.Table `json:"queue_declare_args"`
	QueueBindArgs       amqp.Table `json:"queue_bind_args"`
}

func (conn Connection) TopicTypePublish(data TopicTypePublishData) (err error) {
	var exchangeDeclareArgs amqp.Table = nil
	if len(data.ExchangeDeclareArgs) > 0 {
		exchangeDeclareArgs = data.ExchangeDeclareArgs
	}
	return conn.Publish(PublishData{
		Exchange:               data.Exchange,
		RoutingKey:             data.RoutingKey,
		Type:                   amqp.ExchangeTopic,
		ExchangeDeclareArgs:    exchangeDeclareArgs,
		ExchangeDeclareDurable: data.ExchangeDeclareDurable,
		Publishing:             data.Publishing,
	})
}

func (conn Connection) TopicTypeReceive(data TopicTypeReceiveData) (messages <-chan amqp.Delivery, err error) {
	var exchangeDeclareArgs amqp.Table = nil
	if len(data.ExchangeDeclareArgs) > 0 {
		exchangeDeclareArgs = data.ExchangeDeclareArgs
	}
	var queueDeclareArgs amqp.Table = nil
	if len(data.QueueDeclareArgs) > 0 {
		queueDeclareArgs = data.QueueDeclareArgs
	}
	var queueBindArgs amqp.Table = nil
	if len(data.QueueBindArgs) > 0 {
		queueBindArgs = data.QueueBindArgs
	}
	return conn.Receive(ReceiveData{
		Exchange:               data.Exchange,
		RoutingKey:             data.RoutingKey,
		Type:                   amqp.ExchangeTopic,
		ExchangeDeclareDurable: true,
		QueueDeclareDurable:    false,
		Exclusive:              true,
		ExchangeDeclareArgs:    exchangeDeclareArgs,
		QueueDeclareArgs:       queueDeclareArgs,
		QueueBindArgs:          queueBindArgs,
		AutoAck:                data.AutoAck,
	})
}
