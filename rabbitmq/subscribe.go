package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

type SubscribeTypePublishData struct {
	Exchange   string          `json:"exchange"`
	Publishing amqp.Publishing `json:"publishing"`
}

type SubscribeTypeReceiveData struct {
	Exchange string `json:"exchange"`
	AutoAck  bool   `json:"auto_ack"`
}

func (conn Connection) SubscribeTypePublish(data SubscribeTypePublishData) (err error) {
	return conn.Publish(PublishData{
		Exchange:   data.Exchange,
		Type:       amqp.ExchangeFanout,
		Args:       nil,
		Durable:    true,
		Publishing: data.Publishing,
	})
}

func (conn Connection) SubscribeTypeReceive(data SubscribeTypeReceiveData) (messages <-chan amqp.Delivery, err error) {
	return conn.Receive(ReceiveData{
		Exchange: data.Exchange,
		Type:     amqp.ExchangeFanout,
		Durable:  true,
		AutoAck:  data.AutoAck,
	})
}
