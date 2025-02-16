package config

import (
	"os"
)

type Config struct {
	RedisAddr   string
	RabbitMQURL string
	ServerPort  string
}

func LoadConfig() (*Config, error) {
	redisAddress := os.Getenv("REDIS_ADDRESS")
	if redisAddress == "" {
		redisAddress = "localhost:6379"
	}

	rabbitMQURL := os.Getenv("RABBITMQ_URL")
	if rabbitMQURL == "" {
		rabbitMQURL = "amqp://guest:guest@localhost:5672/"
	}

	serverPortStr := os.Getenv("SERVER_PORT")
	if serverPortStr == "" {
		serverPortStr = ":8080"
	}

	return &Config{
		RedisAddr:   redisAddress,
		RabbitMQURL: rabbitMQURL,
		ServerPort:  serverPortStr,
	}, nil
}
