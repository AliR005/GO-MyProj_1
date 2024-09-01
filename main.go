package main

import (
	"os"
)

func main(){
	var password string = "NULL"
	init_bot(&password)



}

func init_bot(password *string){
	TOKEN := os.Getenv("TOKEN_BOT")
	bot := create_bot(TOKEN)
	bot_main_loop(bot, password)
}
