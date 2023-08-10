package todo_list

import (
	"errors"
	"net/http"
	"github.com/amiosamu/todo-list/internal/entity"
	"github.com/amiosamu/todo-list/internal/repo/repoerrors"
	"github.com/amiosamu/todo-list/internal/service"
	"github.com/gin-gonic/gin"
	"fmt"
)

type taskRoutes struct {
	taskService service.Task
}

func newTaskRoutes(c *gin.RouterGroup, taskService service.Task) {
	r := &taskRoutes{
		taskService: taskService,
	}
	c.POST("/", r.create)
	c.PUT("/:id", r.update)
	c.DELETE("/:id", r.delete)
	c.PUT("/:id/done", r.done)
	c.GET("/", r.getByStatus)
}

type CreateTaskRequest struct {
	Title    string `json:"title" binding:"required"`
	ActiveAt string `json:"activeAt" binding:"required"`
	Status   string `json:"status" binding:"required"`
}

type CreateTaskResponse struct {
	ID   string `json:"id"`
	Code int    `json:"code"`
}

type UpdateTaskResponse struct {
	Title    string `json:"title"`
	ActiveAt string `json:"activeAt"`
	Status   string `json:"status"`
}

type statusResponse struct {
	Status string `json:"status"`
}




func (r *taskRoutes) create(ctx *gin.Context) {
	var request CreateTaskRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task := &entity.Task{
		Title:    request.Title,
		ActiveAt: request.ActiveAt,
		Status:   request.Status,
	}
	id, err := r.taskService.CreateTask(ctx.Request.Context(), *task)
	if err != nil {
		resp := CreateTaskResponse{
			ID:   id,
			Code: http.StatusInternalServerError,
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, resp)
		return
	}
	resp := CreateTaskResponse{
		ID:   id,
		Code: http.StatusCreated,
	}
	ctx.JSON(http.StatusCreated, resp)
}

func (r *taskRoutes) update(ctx *gin.Context) {

}

func (r *taskRoutes) delete(ctx *gin.Context) {
	taskID := ctx.Param("id")
	if err := r.taskService.DeleteTask(ctx, taskID); err != nil {
		ctx.JSON(http.StatusInternalServerError, statusResponse{"Could not delete the task"})
		return
	}
	ctx.JSON(http.StatusOK, statusResponse{
		Status: "Successfully removed task",
	})
}

func (r *taskRoutes) done(ctx *gin.Context) {
    taskID := ctx.Param("id")

    // Call the TaskDone function from your repository
    err := r.taskService.TaskDone(ctx, taskID)
    if err !=  nil{
        if errors.Is(err, repoerrors.TaskAlreadyDone) {
            ctx.JSON(http.StatusConflict, gin.H{"error": "Task is already marked as done"})
        } else {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "Task marked as done"})
}


func (r *taskRoutes) getByStatus(ctx *gin.Context) {
	status := ctx.DefaultQuery("status", "active")
	tasks, err := r.taskService.GetTasksByStatus(ctx, status)
	fmt.Println(tasks)
	fmt.Println(status)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

	var filteredTasks []entity.Task
		if status == "active" {
			for _, t := range tasks {
				if t.Status == "active"{
					filteredTasks=append(filteredTasks, t)
				}
			}
		} else if status == "done" {
			for _, t := range tasks {
				if t.Status == "done" {
					filteredTasks = append(filteredTasks, t)
				}
			}
		}
	ctx.JSON(http.StatusOK, filteredTasks)
}
