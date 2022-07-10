package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

type SubscribeTypePublishData struct {
	Exchange   string          `json:"exchange"`
	Publishing amqp.Publishing `json:"publishing"`
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

func (conn Connection) SubscribeTypeReceive() (err error) {
	return
}
