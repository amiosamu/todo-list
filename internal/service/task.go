package service

import (
	"context"
	"log"

	"github.com/amiosamu/todo-list/internal/entity"
	"github.com/amiosamu/todo-list/internal/repo"
)

type TaskService struct {
	taskRepo repo.Task
}

func (t *TaskService) CreateTask(ctx context.Context, e entity.Task) (string, error) {
	id, err := t.taskRepo.CreateTask(ctx, e)
	if err != nil {
		log.Fatalf("task service - CreateTask - %w", err)
		return "", err
	}
	return id, nil
}

func (t *TaskService) UpdateTask(ctx context.Context, e entity.UpdateTask, id string) error {
	if err := t.taskRepo.UpdateTask(ctx, e, id); err != nil {
		log.Fatalf("task service - UpdateTask - %w", e)
	}
	return nil
}

func (t *TaskService) DeleteTask(ctx context.Context, id string) error {
	return t.taskRepo.DeleteTask(ctx, id)
}

func (t *TaskService) TaskDone(ctx context.Context, id string) error {
	return t.taskRepo.TaskDone(ctx, id)
}

func (t *TaskService) GetTasksByStatus(ctx context.Context, status string) ([]entity.Task, error) {
	return t.taskRepo.GetTasksByStatus(ctx, status)
}

func NewTaskService(taskRepo repo.Task) *TaskService {
	return &TaskService{taskRepo: taskRepo}
}

func (t *TaskService) GetTaskByID(ctx context.Context, id string) (entity.UpdateTask, error) {
	return t.taskRepo.GetTaskByID(ctx, id)
}
