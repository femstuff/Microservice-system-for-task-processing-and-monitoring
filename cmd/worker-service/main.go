package main

import (
	"encoding/json"
	"log"

	"github.com/femstuff/Microservice-system-for-task-processing-and-monitoring/internal/worker-service/entities"
	"github.com/femstuff/Microservice-system-for-task-processing-and-monitoring/internal/worker-service/handlers"
	"github.com/femstuff/Microservice-system-for-task-processing-and-monitoring/internal/worker-service/repository"
	"github.com/femstuff/Microservice-system-for-task-processing-and-monitoring/internal/worker-service/usecases"
	"github.com/go-redis/redis/v8"
	"github.com/streadway/amqp"
)

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	repo := repository.NewRedisTaskRepository(redisClient)
	uc := usecases.NewTaskUseCase(repo)
	handl := handlers.NewTaskHandler(uc)

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
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
		"tasks", false, false, false, false, nil)
	if err != nil {
		log.Fatalf("Ошибка объявления очкреди: %v", err)
	}

	msgs, err := ch.Consume(
		q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("ОШибка регистрации потребителя: %v", err)
	}

	log.Println("Worker service запущен")

	for msg := range msgs {
		var task entities.Task

		log.Printf("Получено сообщение: %s", msg.Body)

		if err := json.Unmarshal(msg.Body, &task); err != nil {
			log.Printf("Ошибка парсинга задачи: %v", err)
			continue
		}

		if err := handl.HandleTask(task); err != nil {
			log.Printf("Ошибка обработки задачи: %v", err)
			return
		}
	}
}
