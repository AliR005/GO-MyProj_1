package rabbitmq

import (
	"log"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Publisher struct {
    Channel *amqp.Channel
    Queue   amqp.Queue
}

func (p *Publisher) Publish(body string) {
	err := p.Channel.Publish(
        "",
        p.Queue.Name,
        false,
        false,
        amqp.Publishing{
            ContentType: "text/plain",
            Body:        []byte(body),
        },
    )

    if err != nil {
        log.Fatalf("Failed to publish a message: %s", err)
    }
    
    log.Printf("Message published: %s", body)
}

func (p *Publisher) Close() {
	p.Channel.Close()
}


type Consumer struct {
	Channel *amqp.Channel
	Queue   amqp.Queue
}

func (c *Consumer) Start(bot *telego.Bot, chatID int64) {
	msgs, err := c.Channel.Consume(
		c.Queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %s", err)
	}

	log.Println("Waiting for messages...")

	for msg := range msgs {
		log.Printf("Received a message: %s", msg.Body)
		bot.SendMessage(tu.Message(
			tu.ID(chatID),
			string(msg.Body),
		))
    }
}

