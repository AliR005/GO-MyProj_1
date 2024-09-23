package config

import (
	"os"
)

type (
	Config struct {
		App      App
		Postgres Postgres
		RabbitMQ RabbitMQ
	}

	App struct {
		Token  string
		Driver string
	}

	Postgres struct {
		Host     string
		Port     string
		Username string
		Password string
		DBName   string
		SSL      string
	}

	RabbitMQ struct {
		URI      string
    	QueueName string
	}
)

func New() *Config {
	return &Config{
		App: App{
			Token:  getEnv("BOT_TOKEN", ""),
			Driver: getEnv("BOT_DB_DRIVER", "")},

		Postgres: Postgres{
			Host:     getEnv("BOT_POSTGRES_HOST", ""),
			Port:     getEnv("BOT_POSTGRES_PORT", "5432"),
			Username: getEnv("BOT_POSTGRES_USERNAME", ""),
			Password: getEnv("BOT_POSTGRES_PASSWORD", ""),
			DBName:   getEnv("BOT_POSTGRES_DBNAME", "messageTime"),
			SSL:      getEnv("BOT_POSTGRES_SSL_MODE", "disable"),
		},

		RabbitMQ: RabbitMQ{
			URI:      getEnv("RABBITMQ_URI", "amqp://guest:guest@localhost:5672/"),
            QueueName: getEnv("RABBITMQ_QUEUE", "telegram_queue"),
		},
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
