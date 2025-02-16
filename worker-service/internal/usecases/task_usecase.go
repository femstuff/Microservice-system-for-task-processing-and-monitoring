package usecases

import (
	"worker-service/internal/entities"
	"worker-service/internal/repository"
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
