package handlers

import (
	"github.com/femstuff/Microservice-system-for-task-processing-and-monitoring/internal/gateway-service/entities"
	"github.com/femstuff/Microservice-system-for-task-processing-and-monitoring/internal/gateway-service/usecases"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TaskHanlder struct {
	taskUseCase *usecases.TaskUseCase
}

func NewTaskHandler(taskUseCase *usecases.TaskUseCase) *TaskHanlder {
	return &TaskHanlder{taskUseCase: taskUseCase}
}

func (t *TaskHanlder) CreateTask(c *gin.Context) {
	var task entities.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	if err := t.taskUseCase.CreateTask(task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании задачи"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Задача успешно создана", "task_id": task.ID})
}

func (t *TaskHanlder) GetTaskResult(c *gin.Context) {
	id := c.Query("id")

	if id == "" {
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
