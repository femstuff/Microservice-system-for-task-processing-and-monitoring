package usecases

import (
	"fmt"

	"api-gateway/internal/entities"
	"api-gateway/internal/repository"
)

type TaskUseCase struct {
	repo repository.TaskRepo
}

func NewTaskUseCase(repo repository.TaskRepo) *TaskUseCase {
	return &TaskUseCase{repo: repo}
}

func (uc *TaskUseCase) CreateTask(task entities.Task) error {
	exists, err := uc.repo.TaskExists(task.ID)
	if err != nil {
		return err
	}

	if exists {
		return fmt.Errorf(err.Error(), "Задача с таким id = %s, уже существует", task.ID)
	}

	return uc.repo.SaveTask(task)
}

func (uc *TaskUseCase) GetTaskResult(id string) (*entities.TaskResult, error) {
	return uc.repo.GetTaskResult(id)
}
