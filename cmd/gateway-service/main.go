package main

import (
	"context"
	"log"

	"github.com/femstuff/Microservice-system-for-task-processing-and-monitoring/internal/gateway-service/handlers"
	"github.com/femstuff/Microservice-system-for-task-processing-and-monitoring/internal/gateway-service/repository"
	"github.com/femstuff/Microservice-system-for-task-processing-and-monitoring/internal/gateway-service/usecases"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("Ошибка подключения к Redis: %v", err)
	}

	taskRepo := repository.NewRedisTaskRepository(redisClient)
	taskUseCase := usecases.NewTaskUseCase(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskUseCase)

	router := gin.Default()
	router.POST("/task", taskHandler.CreateTask)
	router.GET("/result", taskHandler.GetTaskResult)

	log.Println("Gateway запущен на :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}
