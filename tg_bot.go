package main

import (
	"fmt"
	"log"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func create_bot(token string) *telego.Bot {
	bot, err := telego.NewBot(token, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		log.Fatal("token is not specified")
	}
	return bot
}

func bot_main_loop(bot *telego.Bot, password *string) {
	updates, _ := bot.UpdatesViaLongPolling(nil)
	defer bot.StopLongPolling()

	for update := range updates {
		if update.Message != nil {
			chatID := update.Message.Chat.ID

			*password = update.Message.Text
			_, _ = bot.SendMessage(tu.Message(
				tu.ID(chatID),
				settings_bot(*password),
			))

		}
	}
}
