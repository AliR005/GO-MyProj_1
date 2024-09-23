package rabbitmq


import (
	"log"
	"NewProj1/internal/config"
	amqp "github.com/rabbitmq/amqp091-go"
)



func NewConsumer(config config.RabbitMQ) *Consumer {
	conn, err := amqp.Dial(config.URI)
	if err != nil {
		log.Fatalf("unable to open connect to RabbitMQ server. Error: %s", err)
	}

	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf("failed to open a channel. Error: %s", err)
	}
	queue, err := channel.QueueDeclare(
		config.QueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("failed to declare a queue. Error: %s", err)
	}


	return &Consumer{
		Channel: channel,
		Queue:   queue,
	}
}

