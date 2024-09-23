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

	var chatID int64
	publisher := rabbitmq.NewPublisher(cfgMQ)
	defer publisher.Close()
	consumer := rabbitmq.NewConsumer(cfgMQ)
	defer consumer.Close()
	

	updates, _ := bot.UpdatesViaLongPolling(nil)
	defer bot.StopLongPolling()

	for update := range updates {
		if update.Message != nil {
			chatID = update.Message.Chat.ID

			password := service.TextProcessing(fmt.Sprint(chatID), update.Message.Text, cfgDB)
			publisher.Publish(password)
			

			_, _ = bot.SendMessage(tu.Message(
				tu.ID(chatID),
				"Ваш запрос отправлен:)",
			))
			go consumer.Start(bot, chatID)
		}
	}
	
}
