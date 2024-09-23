package rabbitmq

import (
	"log"

	"NewProj1/internal/config"
	amqp "github.com/rabbitmq/amqp091-go"
)


func NewPublisher(config config.RabbitMQ) (*Publisher) {
    conn, err := amqp.Dial(config.URI)
	if err != nil {
		log.Fatalf("unable to open connect to RabbitMQ server. Error: %s", err)
	}

    channel, err := conn.Channel()
 	if err != nil {
		log.Fatalf("failed to open channel. Error: %s", err)
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

    return &Publisher{
        Channel: channel,
        Queue:   queue,
    }
}


