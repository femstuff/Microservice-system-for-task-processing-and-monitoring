package usecases

import (
	"github.com/femstuff/Microservice-system-for-task-processing-and-monitoring/internal/worker-service/entities"
	"github.com/femstuff/Microservice-system-for-task-processing-and-monitoring/internal/worker-service/repository"
)

type TaskUseCase struct {
	repo repository.TaskRepository
}

func NewTaskUseCase(repo repository.TaskRepository) *TaskUseCase {
	return &TaskUseCase{repo: repo}
}

func (uc *TaskUseCase) ProcessTask(task entities.Task) error {
	res := entities.TaskResult{
		ID:     task.ID,
		Result: task.Data * task.Data,
	}

	return uc.repo.SaveTaskResult(res)
}
