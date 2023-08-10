package repo

import (
	"context"
	"github.com/amiosamu/todo-list/internal/entity"
	"github.com/amiosamu/todo-list/internal/repo/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

type Task interface {
	CreateTask(ctx context.Context, t entity.Task) (string, error)
	UpdateTask(ctx context.Context, t entity.UpdateTask, id string) error
	DeleteTask(ctx context.Context, id string) error
	TaskDone(ctx context.Context, id string) error
	GetTasksByStatus(ctx context.Context, status string) ([]entity.Task, error)
	GetTaskByID(ctx context.Context, id string) (entity.UpdateTask, error)
}

type Repos struct {
	Task
}

func NewRepos(mg *mongo.Client) *Repos {
	return &Repos{
		Task: mongodb.NewTaskRepo(mg),
	}
}
