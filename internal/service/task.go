package service

import (
	"context"
	"log"

	"github.com/amiosamu/todo-list/internal/entity"
	"github.com/amiosamu/todo-list/internal/repo"
)

type TaskService struct {
	TaskRepo repo.Task
}

func (t *TaskService) CreateTask(ctx context.Context, e entity.Task) (string, error) {
	id, err := t.TaskRepo.CreateTask(ctx, e)
	if err != nil {
		log.Fatalf("task service - CreateTask - %v", err)
		return "", err
	}
	return id, nil
}

func (t *TaskService) UpdateTask(ctx context.Context, e entity.UpdateTask, id string) error {
	if err := t.TaskRepo.UpdateTask(ctx, e, id); err != nil {
		log.Fatalf("task service - UpdateTask - %v", e)
	}
	return nil
}

func (t *TaskService) DeleteTask(ctx context.Context, id string) error {
	return t.TaskRepo.DeleteTask(ctx, id)
}

func (t *TaskService) TaskDone(ctx context.Context, id string) error {
	return t.TaskRepo.TaskDone(ctx, id)
}

func (t *TaskService) GetTasksByStatus(ctx context.Context, status string) ([]entity.Task, error) {
	return t.TaskRepo.GetTasksByStatus(ctx, status)
}

func NewTaskService(taskRepo repo.Task) *TaskService {
	return &TaskService{TaskRepo: taskRepo}
}

func (t *TaskService) GetTaskByID(ctx context.Context, id string) (entity.UpdateTask, error) {
	return t.TaskRepo.GetTaskByID(ctx, id)
}
