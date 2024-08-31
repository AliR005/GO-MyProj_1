package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func main(){
	TOKEN := os.Getenv("TOKEN_BOT")
	bot := create_bot(TOKEN)

	updates , _ := bot.UpdatesViaLongPolling(nil)
	defer bot.StopLongPolling()

	for update := range updates {
		if update.Message != nil {
			chatID := update.Message.Chat.ID

			sentMessage, _ := bot.SendMessage(
				tu.Message(
					tu.ID(chatID),
					update.Message.Text,
				),
			)

			fmt.Printf("Sent Message: %v\n", sentMessage)
		}
	}

}

func create_bot(token string) *telego.Bot{
	bot, err := telego.NewBot(token, telego.WithDefaultDebugLogger())
	if err != nil{
		fmt.Println(err)
		log.Fatal("token is not specified")
	}
	return bot
}