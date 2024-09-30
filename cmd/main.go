package main

import (
	"NewProj1/internal/client/bot"
	"NewProj1/internal/config"
	"NewProj1/internal/service/utils"
	"log"
	"path/filepath"

	"github.com/joho/godotenv"
)

func init() {
	envPath := filepath.Join("..", ".env")
	if err := godotenv.Load(envPath); err != nil {
		log.Fatal("No .env file found")
	}
}

func main() {
	cfg := config.New()
	utils.AppendMap()
	bot.StartBot(cfg.App, cfg.Postgres, cfg.RabbitMQ)

}
