package bot

import (
	"NewProj1/internal/config"
	"NewProj1/internal/service"
	"fmt"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"log"
)

func StartBot(cfg config.App, cfgDB config.Postgres) {
	bot, err := telego.NewBot(cfg.Token, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		log.Fatal("Token is not specified")
	}

	updates, _ := bot.UpdatesViaLongPolling(nil)
	defer bot.StopLongPolling()

	for update := range updates {
		if update.Message != nil {
			chatID := update.Message.Chat.ID

			password := service.TextProcessing(fmt.Sprint(chatID), update.Message.Text, cfgDB)
			_, _ = bot.SendMessage(tu.Message(
				tu.ID(chatID),
				password,
			))
		}
	}

}
