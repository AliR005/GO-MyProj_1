package main

import (
	"NewProj1/internal/config"
	"NewProj1/internal/domain/bot"
	"NewProj1/internal/service/utils"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	cfg := config.New()
	utils.AppendMap()
	bot.StartBot(cfg.App)
}
