package main

import (
	"github.com/femstuff/Microservice-system-for-task-processing-and-monitoring/internal/gateway-service/handlers"
	"github.com/femstuff/Microservice-system-for-task-processing-and-monitoring/internal/gateway-service/repository"
	"github.com/femstuff/Microservice-system-for-task-processing-and-monitoring/internal/gateway-service/usecases"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"log"
)

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	router := gin.Default()

	repo := repository.NewRedisTaskRepository(redisClient)
	uc := usecases.NewTaskUseCase(repo)
	handl := handlers.NewTaskHandler(uc)

	router.POST("/task", handl.CreateTask)
	router.GET("/result", handl.GetTaskResult)
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Ошибка при запуске: %v", err)
	}
}
