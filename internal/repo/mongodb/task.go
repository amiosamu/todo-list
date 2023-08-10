package mongodb

import (
	"context"
	"errors"
	"fmt"
	"github.com/amiosamu/todo-list/internal/entity"
	"github.com/amiosamu/todo-list/internal/repo/repoerrors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepo struct {
	DB *mongo.Collection
}

func (t *TaskRepo) GetTaskByID(ctx context.Context, id string) (entity.UpdateTask, error) {
	var task entity.UpdateTask

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return task, err
	}

	filter := bson.D{
		{"_id", objectID},
	}

	err = t.DB.FindOne(ctx, filter).Decode(&task)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return task, repoerrors.TaskNotFound
		}
		return task, err
	}

	return task, nil
}

func (t *TaskRepo) CreateTask(ctx context.Context, task entity.Task) (string, error) {
	res, err := t.DB.InsertOne(ctx, task)
	if err != nil {
		return "", fmt.Errorf("task repository - CreateTask - %w", err)
	}
	id := res.InsertedID.(primitive.ObjectID).Hex()
	return id, nil
}

func (t *TaskRepo) UpdateTask(ctx context.Context, task entity.UpdateTask, id string) error {
	fmt.Println(task)
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("task repository - UpdateTask - %w", err)
	}
	filter := bson.D{{"_id", objectID}}
	updateFields := bson.D{
		{"title", task.Title},
		{"activeAt", task.ActiveAt},
		{"status", task.Status},
	}
	update := bson.D{
		{"$set", updateFields},
	}
	_, err = t.DB.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("task repository - Update - %w", err)
	}
	return nil
}

func (t *TaskRepo) DeleteTask(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("task repository - ObjectIDFromHex - %w", repoerrors.TaskNotFound)
	}
	filter := bson.D{
		{"_id", objectID},
	}

	res, err := t.DB.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("task repository - DeleteTask - %w", err)
	}
	if res.DeletedCount != 1 {
		return repoerrors.TaskNotFound
	}
	return nil
}

func (t *TaskRepo) TaskDone(ctx context.Context, id string) error {
	task, err := t.GetTaskByID(ctx, id)
	if err != nil {
		return err
	}
	if task.Status == "done" {
		return repoerrors.TaskAlreadyDone
	}
	task.Status = "done"
	if err := t.UpdateTask(ctx, task, id); err != nil {
		return fmt.Errorf("task repository - TaskDone - %w", err)
	}
	return nil
}

func (t *TaskRepo) GetTasksByStatus(ctx context.Context, status string) ([]entity.Task, error) {
	filter := bson.D{
		{"status", status},
	}
	cursor, err := t.DB.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("task repository - GetTasksByStatus - Find - %w", err)
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, ctx)

	var tasks []entity.Task
	for cursor.Next(ctx) {
		var task entity.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, fmt.Errorf("task repository - GetTasksByStatus - Decode - %w", err)
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func NewTaskRepo(mg *mongo.Client) *TaskRepo {
	return &TaskRepo{mg.Database("todo-list").Collection("tasks")}
}
