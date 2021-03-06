package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

const ExchangeXDelayedMessage = "x-delayed-message"

type DelayedTypeSendData struct {
	Exchange               string          `json:"exchange"`
	RoutingKey             string          `json:"routing_key"`
	Publishing             amqp.Publishing `json:"publishing"`
	ExchangeDeclareArgs    amqp.Table      `json:"exchange_declare_args"`
	ExchangeDeclareDurable bool            `json:"exchange_declare_durable"` // true
}

type DelayedTypeReceiveData struct {
	Exchange                  string     `json:"exchange"`
	QueueName                 string     `json:"queue_name"`
	RoutingKey                string     `json:"routing_key"`
	AutoAck                   bool       `json:"auto_ack"`
	ExchangeDeclareArgs       amqp.Table `json:"exchange_declare_args"`
	QueueDeclareArgs          amqp.Table `json:"queue_declare_args"`
	QueueBindArgs             amqp.Table `json:"queue_bind_args"`
	ExchangeDeclareAutoDelete bool       `json:"exchange_declare_auto_delete"`
	ExchangeDeclareDurable    bool       `json:"exchange_declare_durable"`  // true
	QueueDeclareDurable       bool       `json:"queue_declare_durable"`     // true
	QueueDeclareAutoDelete    bool       `json:"queue_declare_auto_delete"` // true
}

func (conn Connection) DelayedTypeSend(data DelayedTypeSendData) (err error) {
	var exchangeDeclareArgs amqp.Table = nil
	if len(data.ExchangeDeclareArgs) > 0 {
		exchangeDeclareArgs = data.ExchangeDeclareArgs
	}
	return conn.Send(SendData{
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
		QueueName:                 data.QueueName,
		RoutingKey:                data.RoutingKey,
		ExchangeDeclareAutoDelete: data.ExchangeDeclareAutoDelete,
		ExchangeDeclareDurable:    data.ExchangeDeclareDurable,
		QueueDeclareDurable:       data.QueueDeclareDurable,
		QueueDeclareAutoDelete:    data.QueueDeclareAutoDelete,
		AutoAck:                   data.AutoAck,
		ExchangeDeclareArgs:       exchangeDeclareArgs,
		QueueDeclareArgs:          queueDeclareArgs,
		QueueBindArgs:             queueBindArgs,
	})
}
