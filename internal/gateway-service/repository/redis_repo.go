package repository

import (
	"context"
	"github.com/femstuff/Microservice-system-for-task-processing-and-monitoring/internal/gateway-service/entities"
	"github.com/go-redis/redis/v8"
	"log"
)

type TaskRepo interface {
	SaveTask(task entities.Task) error
	GetTaskResult(id string) (*entities.TaskResult, error)
}

type RedisTaskRepo struct {
	client *redis.Client
}

func NewRedisTaskRepository(client *redis.Client) *RedisTaskRepo {
	return &RedisTaskRepo{client: client}
}

func (r *RedisTaskRepo) SaveTask(task entities.Task) error {
	ctx := context.Background()
	err := r.client.Set(ctx, "task_"+task.ID, task.Data, 0).Err()
	if err != nil {
		log.Printf("Ошибка при сохранении задачи в Redis: %v", err)
	}
	return err
}

func (r *RedisTaskRepo) GetTaskResult(id string) (*entities.TaskResult, error) {
	ctx := context.Background()
	val, err := r.client.Get(ctx, "task_"+id).Int()
	if err != nil {
		return nil, err
	}

	return &entities.TaskResult{
		ID:     id,
		Result: val,
	}, nil
}
