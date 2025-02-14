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

func (uc *TaskUseCase) ProcessTask(result entities.TaskResult) error {
	return uc.repo.SaveTaskResult(result)
}
