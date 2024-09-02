package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	var password string = "NULL"
	append_map()
	init_bot(&password)

}

func init_bot(password *string) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
	TOKEN, _ := os.LookupEnv("TOKEN")
	bot := create_bot(TOKEN)
	bot_main_loop(bot, password)
}

func settings_bot(password string) string {
	return GetMd5(password)
}
