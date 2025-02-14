package handlers

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	"net/http"
	"strings"

	"github.com/femstuff/Microservice-system-for-task-processing-and-monitoring/internal/gateway-service/entities"
	"github.com/femstuff/Microservice-system-for-task-processing-and-monitoring/internal/gateway-service/usecases"
	"github.com/gin-gonic/gin"
)

type TaskHanlder struct {
	taskUseCase *usecases.TaskUseCase
}

func NewTaskHandler(taskUseCase *usecases.TaskUseCase) *TaskHanlder {
	return &TaskHanlder{taskUseCase: taskUseCase}
}

func (t *TaskHanlder) CreateTask(c *gin.Context) {
	var task entities.Task
	log.Print("CreateTask вызван")

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	if err := t.taskUseCase.CreateTask(task); err != nil {
		if strings.Contains(err.Error(), "уже существует") {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании задачи"})
		}
		return
	}

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Printf("Ошибка подключения к RabbitMQ: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка RabbitMQ"})
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Printf("Ошибка создания канала RabbitMQ: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка RabbitMQ"})
		return
	}
	defer ch.Close()

	body, _ := json.Marshal(task)
	err = ch.Publish(
		"",
		"tasks",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		log.Printf("Ошибка отправки в RabbitMQ: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка отправки задачи"})
		return
	}

	log.Println("Задача отправлена в RabbitMQ:", string(body))
	c.JSON(http.StatusCreated, gin.H{"message": "Задача успешно создана и отправлена", "task_id": task.ID})
}

func (t *TaskHanlder) GetTaskResult(c *gin.Context) {
	id := c.Query("id")

	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Параметр id не может быть пустым"})
		return
	}

	res, err := t.taskUseCase.GetTaskResult(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении результата"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"task_id": res.ID,
		"result":  res.Result,
	})
}
