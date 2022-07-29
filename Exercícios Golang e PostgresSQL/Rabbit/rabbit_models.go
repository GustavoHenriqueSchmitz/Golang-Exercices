package models

import "github.com/streadway/amqp"

type ConnectionSettings struct {
	Connection *amqp.Connection
	Host       string
	Port       string
	User       string
	Pass       string
}

type QueueSettings struct {
	Name       string
	Durable    bool
	AutoDelete bool
	Exclusive  bool
	NoWait     bool
	Arguments  amqp.Table
}

type ConsumerSettings struct {
	QueueName string
	Name      string
	AutoAck   bool
	Exclusive bool
	NoLocal   bool
	NoWait    bool
	Arguments amqp.Table
}

type PublishSettings struct {
	Exchange   string
	RoutingKey string
	Mandatory  bool
	Immediate  bool
	Publish    struct {
		ContentType string
		Body        []byte
	}
}

type ExchangeSettings struct {
	Name       string
	Type       string
	Durable    bool
	AutoDelete bool
	Internal   bool
	NoWait     bool
	Arguments  amqp.Table
}

type BindingSettings struct {
	QueueName  string
	RoutingKey string
	Exchange   string
	NoWait     bool
	Arguments  amqp.Table
}
