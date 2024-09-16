package main

import (
	"NewProj1/internal/config"
	"NewProj1/internal/domain/bot"
	"fmt"
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
	bot.StartBot(cfg.App)
	fmt.Println("test")
}
