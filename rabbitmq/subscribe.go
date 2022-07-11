package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

type SubscribeTypePublishData struct {
	Exchange               string          `json:"exchange"`
	ExchangeDeclareDurable bool            `json:"exchange_declare_durable"`
	Publishing             amqp.Publishing `json:"publishing"`
	ExchangeDeclareArgs    amqp.Table      `json:"exchange_declare_args"`
}

type SubscribeTypeReceiveData struct {
	Exchange            string     `json:"exchange"`
	AutoAck             bool       `json:"auto_ack"`
	ExchangeDeclareArgs amqp.Table `json:"exchange_declare_args"`
	QueueDeclareArgs    amqp.Table `json:"queue_declare_args"`
	QueueBindArgs       amqp.Table `json:"queue_bind_args"`
}

func (conn Connection) SubscribeTypePublish(data SubscribeTypePublishData) (err error) {
	var exchangeDeclareArgs amqp.Table = nil
	if len(data.ExchangeDeclareArgs) > 0 {
		exchangeDeclareArgs = data.ExchangeDeclareArgs
	}
	return conn.Publish(PublishData{
		Exchange:               data.Exchange,
		Type:                   amqp.ExchangeFanout,
		ExchangeDeclareArgs:    exchangeDeclareArgs,
		ExchangeDeclareDurable: data.ExchangeDeclareDurable,
		Publishing:             data.Publishing,
	})
}

func (conn Connection) SubscribeTypeReceive(data SubscribeTypeReceiveData) (messages <-chan amqp.Delivery, err error) {
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
		Exchange:            data.Exchange,
		Type:                amqp.ExchangeFanout,
		Durable:             true,
		ExchangeDeclareArgs: exchangeDeclareArgs,
		QueueDeclareArgs:    queueDeclareArgs,
		QueueBindArgs:       queueBindArgs,
		AutoAck:             data.AutoAck,
	})
}
