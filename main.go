package main

import (
	"os"
	"crypto/md5"
	"encoding/hex"
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

func settings_bot(password string) string{
	return GetMd5(password)
}

func GetMd5(text string) string {
	h := md5.New()
	h.Write([]byte(text))
	return hex.EncodeToString(h.Sum(nil))
}