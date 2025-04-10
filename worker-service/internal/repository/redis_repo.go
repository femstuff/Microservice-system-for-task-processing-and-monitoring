package repository

import (
	"context"

	"worker-service/internal/entities"

	"github.com/go-redis/redis/v8"
)

type TaskRepository interface {
	SaveTaskResult(result entities.TaskResult) error
}

type RedisTaskRepository struct {
	client *redis.Client
}

func NewRedisTaskRepository(client *redis.Client) *RedisTaskRepository {
	return &RedisTaskRepository{client: client}
}

func (r *RedisTaskRepository) SaveTaskResult(result entities.TaskResult) error {
	ctx := context.Background()
	return r.client.Set(ctx, "task_"+result.ID, result.Result, 0).Err()
}
