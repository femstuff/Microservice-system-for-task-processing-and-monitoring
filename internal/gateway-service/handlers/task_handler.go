package handlers

import (
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

func (h *TaskHanlder) CreateTask(c *gin.Context) {
	var task entities.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	if err := h.taskUseCase.CreateTask(task); err != nil {
		if strings.Contains(err.Error(), "уже существует") {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании задачи"})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Задача успешно создана", "task_id": task.ID})
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
