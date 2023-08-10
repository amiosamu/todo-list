package service

import (
	"context"

	"github.com/amiosamu/todo-list/internal/entity"
	"github.com/amiosamu/todo-list/internal/repo"
)

type Dependencies struct {
	Repos *repo.Repos
}

type Task interface {
	CreateTask(ctx context.Context, t entity.Task) (string, error)
	UpdateTask(ctx context.Context, t entity.UpdateTask, id string) error
	DeleteTask(ctx context.Context, id string) error
	TaskDone(ctx context.Context, id string) error
	GetTasksByStatus(ctx context.Context, status string) ([]entity.Task, error)
	GetTaskByID(ctx context.Context, id string) (entity.UpdateTask, error)
}

type Services struct {
	Task Task
}

func NewServices(deps Dependencies) *Services {
	return &Services{
		Task: NewTaskService(deps.Repos.Task),
	}
}
