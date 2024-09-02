package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
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
	for key, value := range map_hashs {
		if GetMd5(password) == key {
			return value
		} else if password == value {
			return key
		}
	}
	return fmt.Sprint("Password/hash \"", password, "\" does not exist ")
}

func GetMd5(text string) string {
	h := md5.New()
	h.Write([]byte(text))
	return hex.EncodeToString(h.Sum(nil))
}
