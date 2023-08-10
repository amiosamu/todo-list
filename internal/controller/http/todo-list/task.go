package todo_list

import (
	"errors"
	"net/http"
	"github.com/amiosamu/todo-list/internal/entity"
	"github.com/amiosamu/todo-list/internal/repo/repoerrors"
	"github.com/amiosamu/todo-list/internal/service"
	"github.com/gin-gonic/gin"
	_ "github.com/amiosamu/todo-list/docs"
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

type createTaskRequest struct {
	Title    string `json:"title" binding:"required"`
	ActiveAt string `json:"activeAt" binding:"required"`
	Status   string `json:"status" binding:"required"`
}

type createTaskResponse struct {
	ID   string `json:"id"`
	Code int    `json:"code"`
}

type updateTaskRequest struct{
	Title string `json:"title"`
	ActiveAt string `json:"activeAt"`
	Status string `json:"status,omitempty"`
}

type updateTaskResponse struct {
	Title    string `json:"title"`
	ActiveAt string `json:"activeAt"`
	Status   string `json:"status"`
}

type statusResponse struct {
	Status string `json:"status"`
}



// @Summary Create task
// @Description Create task
// @Tags tasks
// @Accept json
// @Produce json
// @Param request body createTaskRequest true "Task Request"
// @Success 200 {object} createTaskResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/todo-list/tasks/ [post]
func (r *taskRoutes) create(ctx *gin.Context) {
	var request createTaskRequest
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
		resp := createTaskResponse{
			ID:   id,
			Code: http.StatusInternalServerError,
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, resp)
		return
	}
	resp := createTaskResponse{
		ID:   id,
		Code: http.StatusCreated,
	}
	ctx.JSON(http.StatusCreated, resp)
}

// @Summary Update task
// @Description Update task
// @Tags tasks
// @Produce json
// @Param request body updateTaskRequest true "Task Request"
// @Success 200 {object} updateTaskResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure 400 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/todo-list/tasks/{id}/done [put]
func (r *taskRoutes) update(ctx *gin.Context) {
	var task entity.UpdateTask
	taskID := ctx.Param("id")
	if err := ctx.ShouldBindJSON(&task); err != nil{
		ctx.JSON(http.StatusBadRequest, statusResponse{"Bad Request"})
		return
	}

	resp := entity.UpdateTask{
		Title: task.Title,
		ActiveAt: task.ActiveAt,
	}
	if err := r.taskService.UpdateTask(ctx, resp, taskID); err != nil{
		if errors.Is(err, repoerrors.TaskNotFound){
			ctx.JSON(http.StatusNotFound, statusResponse{"Task not found"})
			return
		}else{
			ctx.JSON(http.StatusInternalServerError, statusResponse{"Could not update the task"})
			return
		}
	}
	ctx.JSON(http.StatusOK, resp)
}
// @Summary Delete task
// @Description Delete task
// @Tags tasks
// @Produce json
// @Success 200 {object} statusResponse
// @Failure 404 {object} errorResponse
// @Failure 400 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/todo-list/tasks/{id} [delete]
func (r *taskRoutes) delete(ctx *gin.Context) {
	taskID := ctx.Param("id")
	if err := r.taskService.DeleteTask(ctx, taskID); err != nil {
		if errors.Is(err, repoerrors.TaskNotFound){
			ctx.JSON(http.StatusNotFound, statusResponse{"Task not found"})
			return
		} else {
			ctx.JSON(http.StatusBadRequest, statusResponse{"Bad Request"})
			return
		}
	}
	ctx.JSON(http.StatusOK, statusResponse{
		Status: "Successfully removed task",
	})
}


// @Summary Complete task
// @Description Complete task
// @Tags tasks
// @Produce json
// @Success 200 {object} statusResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure 409 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/todo-list/tasks/{id}/done [put]
func (r *taskRoutes) done(ctx *gin.Context) {
    taskID := ctx.Param("id")
    err := r.taskService.TaskDone(ctx, taskID)
    if err !=  nil{
        if errors.Is(err, repoerrors.TaskAlreadyDone) {
            ctx.JSON(http.StatusConflict, gin.H{"error": "Task is already marked as done"})
			return
        } else if errors.Is(err, repoerrors.TaskNotFound){
			ctx.JSON(http.StatusNotFound,gin.H{"error": "Task not found"})
			return
		} else {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
        }
    }
    ctx.JSON(http.StatusOK, gin.H{"message": "Task marked as done"})
}

// @Summary Get Tasks By Status
// @Description Get Tasks By Status
// @Tags tasks
// @Produce json
// @Success 200 {object} statusResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure 409 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/todo-list/tasks/{id}/done [get]
func (r *taskRoutes) getByStatus(ctx *gin.Context) {
	status := ctx.DefaultQuery("status", "active")
	tasks, err := r.taskService.GetTasksByStatus(ctx, status)
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
