package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

// ReceiveData 接收处理队列所提交数据
type ReceiveData struct {
	Exchange                  string     `json:"exchange"`
	Type                      string     `json:"type"`
	ExchangeDeclareDurable    bool       `json:"exchange_declare_durable"`
	QueueDeclareDurable       bool       `json:"queue_declare_durable"`
	ExchangeDeclareAutoDelete bool       `json:"exchange_declare_auto_delete"`
	QueueDeclareAutoDelete    bool       `json:"queue_declare_auto_delete"`
	Internal                  bool       `json:"internal"`
	NoWait                    bool       `json:"no_wait"`
	Exclusive                 bool       `json:"exclusive"`
	QueueName                 string     `json:"queue_name"`
	RoutingKey                string     `json:"routing_key"`
	Consumer                  string     `json:"consumer"`
	AutoAck                   bool       `json:"auto_ack"`
	NoLocal                   bool       `json:"no_local"`
	QosPrefetchCount          int        `json:"qos_prefetch_count"`
	QosPrefetchSize           int        `json:"qos_prefetch_size"`
	QosGlobal                 bool       `json:"qos_global"`
	ExchangeDeclareArgs       amqp.Table `json:"exchange_declare_args"`
	QueueDeclareArgs          amqp.Table `json:"queue_declare_args"`
	QueueBindArgs             amqp.Table `json:"queue_bind_args"`
	ConsumeArgs               amqp.Table `json:"consume_args"`
}

// Receive 接收消息队列
func (conn Connection) Receive(data ReceiveData) (messages <-chan amqp.Delivery, ch *amqp.Channel, err error) {
	var q amqp.Queue
	ch, err = conn.Conn.Channel()
	if err != nil {
		return
	}
	if data.Type != "" {
		err = ch.ExchangeDeclare(data.Exchange, data.Type, data.ExchangeDeclareDurable, data.ExchangeDeclareAutoDelete, data.Internal, data.NoWait, data.ExchangeDeclareArgs)
		if err != nil {
			return
		}
	}
	q, err = ch.QueueDeclare(data.QueueName, data.QueueDeclareDurable, data.QueueDeclareAutoDelete, data.Exclusive, data.NoWait, data.QueueDeclareArgs)
	if err != nil {
		return
	}
	if data.Type != "" {
		err = ch.QueueBind(q.Name, data.RoutingKey, data.Exchange, data.NoWait, data.QueueBindArgs)
		if err != nil {
			return
		}
	}
	if data.Type == "" {
		err = ch.Qos(data.QosPrefetchCount, data.QosPrefetchSize, data.QosGlobal)
		if err != nil {
			return
		}
	}
	messages, err = ch.Consume(q.Name, data.Consumer, data.AutoAck, data.Exclusive, data.NoLocal, data.NoWait, data.ConsumeArgs)
	return
}
