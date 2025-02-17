package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"worker-service/config"
	"worker-service/internal/entities"
	"worker-service/internal/handlers"
	"worker-service/internal/repository"
	"worker-service/internal/usecases"

	"github.com/go-redis/redis/v8"
	"github.com/streadway/amqp"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Ошибка при загрузке конфигурации: %v", err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr: cfg.RedisAddr,
	})

	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("Ошибка подключения к Redis: %v", err)
	}

	repo := repository.NewRedisTaskRepository(redisClient)
	uc := usecases.NewTaskUseCase(repo)
	handler := handlers.NewTaskHandler(uc)

	time.Sleep(10 * time.Second)
	conn, err := amqp.Dial(cfg.RabbitMQURL)
	if err != nil {
		log.Fatalf("Ошибка подключения к RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Ошибка создания канала RabbitMQ: %v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"tasks",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Ошибка объявления очереди: %v", err)
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Ошибка регистрации потребителя: %v", err)
	}

	log.Println("Worker Service запущен и ожидает задачи...")

	for msg := range msgs {
		var task entities.Task

		log.Printf("Получено сообщение: %s", msg.Body)

		if err := json.Unmarshal(msg.Body, &task); err != nil {
			log.Printf("Ошибка парсинга задачи: %v", err)
			continue
		}

		if err := handler.HandleTask(task); err != nil {
			log.Printf("Ошибка обработки задачи: %v", err)
			msg.Nack(false, true)
		} else {
			msg.Ack(false)
		}
	}
}
