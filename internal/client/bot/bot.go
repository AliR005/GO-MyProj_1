package bot

import (
	"NewProj1/internal/config"
	"NewProj1/internal/rabbitmq"
	"NewProj1/internal/service"
	"fmt"
	"log"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func StartBot(cfg config.App, cfgDB config.Postgres, cfgMQ config.RabbitMQ) {
	bot, err := telego.NewBot(cfg.Token, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		log.Fatal("Token is not specified")
	}

	publisher := rabbitmq.NewPublisher(cfgMQ)
	consumer := rabbitmq.NewConsumer(cfgMQ)

	updates, _ := bot.UpdatesViaLongPolling(nil)
	defer bot.StopLongPolling()

	for update := range updates {
		if update.Message != nil {
			chatID := update.Message.Chat.ID

			password := service.TextProcessing(fmt.Sprint(chatID), update.Message.Text, cfgDB)
			publisher.Publish(password)
			consumer.Start(bot, chatID)

			_, _ = bot.SendMessage(tu.Message(
				tu.ID(chatID),
				"Ваш запрос отправлен:)",
			))
		}
	}

}
