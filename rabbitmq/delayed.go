package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

const ExchangeXDelayedMessage = "x-delayed-message"

type DelayedTypePublishData struct {
	Exchange               string          `json:"exchange"`
	RoutingKey             string          `json:"routing_key"`
	Publishing             amqp.Publishing `json:"publishing"`
	ExchangeDeclareArgs    amqp.Table      `json:"exchange_declare_args"`
	ExchangeDeclareDurable bool            `json:"exchange_declare_durable"`
}

type DelayedTypeReceiveData struct {
	Exchange                  string     `json:"exchange"`
	RoutingKey                string     `json:"routing_key"`
	AutoAck                   bool       `json:"auto_ack"`
	ExchangeDeclareArgs       amqp.Table `json:"exchange_declare_args"`
	QueueDeclareArgs          amqp.Table `json:"queue_declare_args"`
	QueueBindArgs             amqp.Table `json:"queue_bind_args"`
	ExchangeDeclareAutoDelete bool       `json:"exchange_declare_auto_delete"`
	ExchangeDeclareDurable    bool       `json:"exchange_declare_durable"`
	QueueDeclareDurable       bool       `json:"queue_declare_durable"`
	QueueDeclareAutoDelete    bool       `json:"queue_declare_auto_delete"`
}

func (conn Connection) DelayedTypePublish(data DelayedTypePublishData) (err error) {
	var exchangeDeclareArgs amqp.Table = nil
	if len(data.ExchangeDeclareArgs) > 0 {
		exchangeDeclareArgs = data.ExchangeDeclareArgs
	}
	return conn.Publish(PublishData{
		Exchange:               data.Exchange,
		Type:                   ExchangeXDelayedMessage,
		RoutingKey:             data.RoutingKey,
		ExchangeDeclareArgs:    exchangeDeclareArgs,
		ExchangeDeclareDurable: data.ExchangeDeclareDurable,
		Publishing:             data.Publishing,
	})
}

func (conn Connection) DelayedTypeReceive(data DelayedTypeReceiveData) (messages <-chan amqp.Delivery, ch *amqp.Channel, err error) {
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
		Exchange:                  data.Exchange,
		Type:                      ExchangeXDelayedMessage,
		RoutingKey:                data.RoutingKey,
		ExchangeDeclareAutoDelete: false,
		ExchangeDeclareDurable:    true,
		QueueDeclareDurable:       true,
		QueueDeclareAutoDelete:    true,
		AutoAck:                   data.AutoAck,
		ExchangeDeclareArgs:       exchangeDeclareArgs,
		QueueDeclareArgs:          queueDeclareArgs,
		QueueBindArgs:             queueBindArgs,
	})
}
