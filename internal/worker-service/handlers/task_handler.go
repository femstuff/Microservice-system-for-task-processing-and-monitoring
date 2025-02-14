package handlers

import (
	"github.com/femstuff/Microservice-system-for-task-processing-and-monitoring/internal/worker-service/entities"
	"github.com/femstuff/Microservice-system-for-task-processing-and-monitoring/internal/worker-service/usecases"
	"log"
)

type TaskHandler struct {
	taskUseCase *usecases.TaskUseCase
}

func NewTaskHandler(taskUseCase *usecases.TaskUseCase) *TaskHandler {
	return &TaskHandler{taskUseCase: taskUseCase}
}

func (h *TaskHandler) HandleTask(task entities.Task) error {
	log.Printf("Получена задача: ID=%s, Data=%d", task.ID, task.Data)
	result := entities.TaskResult{
		ID:     task.ID,
		Result: task.Data * task.Data,
	}
	if err := h.taskUseCase.ProcessTask(result); err != nil {
		log.Printf("Ошибка обработки задачи: %v", err)
		return err
	}
	log.Printf("Результат задачи сохранен в Redis: ID=%s, Result=%d", result.ID, result.Result)
	return nil
}
