package main

import (
	"context"
	"log"
  
	"github.com/femstuff/Microservice-system-for-task-processing-and-monitoring/internal/shared/config"
	"github.com/femstuff/Microservice-system-for-task-processing-and-monitoring/internal/gateway-service/handlers"
	"github.com/femstuff/Microservice-system-for-task-processing-and-monitoring/internal/gateway-service/repository"
	"github.com/femstuff/Microservice-system-for-task-processing-and-monitoring/internal/gateway-service/usecases"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
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

	taskRepo := repository.NewRedisTaskRepository(redisClient)
	taskUseCase := usecases.NewTaskUseCase(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskUseCase)

	router := gin.Default()
	router.POST("/task", taskHandler.CreateTask)
	router.GET("/result", taskHandler.GetTaskResult)

	log.Printf("Gateway запущен на %v \n", cfg.ServerPort)
	if err := router.Run(cfg.ServerPort); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}
