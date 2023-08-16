package todo_list_test

import (
	"bytes"
	"encoding/json"
	"github.com/amiosamu/todo-list/internal/controller/http/todo-list"
	"github.com/amiosamu/todo-list/internal/entity"
	service_mocks "github.com/amiosamu/todo-list/internal/mocks/servicemocks"
	"github.com/amiosamu/todo-list/internal/repo/repoerrors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTaskRoutes_create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTaskService := service_mocks.NewMockTask(ctrl)
	r := gin.Default()
	todo_list.NewTaskRoutes(r.Group("/api/todo-list/tasks"), mockTaskService)

	t.Run("Successful task creation", func(t *testing.T) {
		mockTaskService.EXPECT().CreateTask(gomock.Any(), gomock.Eq(entity.Task{
			Title:    "test",
			ActiveAt: "2023-08-16",
			Status:   "active",
		})).Return("taskID", nil)

		reqBody := todo_list.CreateTaskRequest{
			Title:    "test",
			ActiveAt: "2023-08-16",
			Status:   "active",
		}
		body, _ := json.Marshal(reqBody)

		req, _ := http.NewRequest("POST", "/api/todo-list/tasks/", bytes.NewReader(body))
		recorder := httptest.NewRecorder()
		r.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusCreated, recorder.Code)
	})
	t.Run("Title too long", func(t *testing.T) {
		mockTaskService.EXPECT().CreateTask(gomock.Any(), gomock.Any()).Return("", repoerrors.TaskTitleTooLong)

		reqBody := todo_list.CreateTaskRequest{
			Title: "This is a very long title" +
				"This is a very long title" +
				"This is a very long title" +
				"This is a very long title" +
				"This is a very long title",
			ActiveAt: "2023-08-16",
			Status:   "active",
		}
		body, _ := json.Marshal(reqBody)

		req, _ := http.NewRequest("POST", "/api/todo-list/tasks/", bytes.NewReader(body))
		recorder := httptest.NewRecorder()
		r.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})
}

func TestTaskRoutes_delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTaskService := service_mocks.NewMockTask(ctrl)
	r := gin.Default()
	todo_list.NewTaskRoutes(r.Group("/api/todo-list/tasks"), mockTaskService)

	t.Run("Task not found", func(t *testing.T) {
		taskID := "invalid-id"
		mockTaskService.EXPECT().DeleteTask(gomock.Any(), gomock.Eq(taskID)).Return(repoerrors.TaskNotFound)

		req, _ := http.NewRequest("DELETE", "/api/todo-list/tasks/"+taskID, nil)
		recorder := httptest.NewRecorder()
		r.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusNotFound, recorder.Code)
	})

	t.Run("Successful task deletion", func(t *testing.T) {
		taskID := "64d74607f9c13a703bd73fcb"
		mockTaskService.EXPECT().DeleteTask(gomock.Any(), gomock.Eq(taskID)).Return(nil)

		req, _ := http.NewRequest("DELETE", "/api/todo-list/tasks/"+taskID, nil)
		recorder := httptest.NewRecorder()
		r.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}

func TestTaskRoutes_done(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTaskService := service_mocks.NewMockTask(ctrl)
	r := gin.Default()
	todo_list.NewTaskRoutes(r.Group("/api/todo-list/tasks"), mockTaskService)

	t.Run("Task not found", func(t *testing.T) {
		taskID := "invalid-id"
		mockTaskService.EXPECT().TaskDone(gomock.Any(), gomock.Eq(taskID)).Return(repoerrors.TaskNotFound)

		req, _ := http.NewRequest("PUT", "/api/todo-list/tasks/"+taskID+"/done", nil)
		recorder := httptest.NewRecorder()
		r.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusNotFound, recorder.Code)
	})

	t.Run("Task already done", func(t *testing.T) {
		taskID := "64d74607f9c13a703bd73fcb"
		mockTaskService.EXPECT().TaskDone(gomock.Any(), gomock.Eq(taskID)).Return(repoerrors.TaskAlreadyDone)

		req, _ := http.NewRequest("PUT", "/api/todo-list/tasks/"+taskID+"/done", nil)
		recorder := httptest.NewRecorder()
		r.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusConflict, recorder.Code)
	})

	t.Run("Successful task completion", func(t *testing.T) {
		taskID := "64d46cb51afc77c3d0e32851"
		mockTaskService.EXPECT().TaskDone(gomock.Any(), gomock.Eq(taskID)).Return(nil)

		req, _ := http.NewRequest("PUT", "/api/todo-list/tasks/"+taskID+"/done", nil)
		recorder := httptest.NewRecorder()
		r.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}

func TestTaskRoutes_update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTaskService := service_mocks.NewMockTask(ctrl)
	r := gin.Default()
	todo_list.NewTaskRoutes(r.Group("/api/todo-list/tasks"), mockTaskService)

	t.Run("Invalid request body", func(t *testing.T) {
		taskID := "valid-task-id"

		req, _ := http.NewRequest("PUT", "/api/todo-list/tasks/"+taskID, bytes.NewReader([]byte("invalid-json")))
		recorder := httptest.NewRecorder()
		r.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("Task not found", func(t *testing.T) {
		taskID := "invalid-id"
		mockTaskService.EXPECT().UpdateTask(gomock.Any(), gomock.Any(), gomock.Eq(taskID)).Return(repoerrors.TaskNotFound)

		reqBody := todo_list.UpdateTaskRequest{
			Title: "updated-title",
		}
		body, _ := json.Marshal(reqBody)

		req, _ := http.NewRequest("PUT", "/api/todo-list/tasks/"+taskID, bytes.NewReader(body))
		recorder := httptest.NewRecorder()
		r.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusNotFound, recorder.Code)
	})

	t.Run("Successful task update", func(t *testing.T) {
		taskID := "64d46cb51afc77c3d0e32851"
		mockTaskService.EXPECT().UpdateTask(gomock.Any(), gomock.Any(), gomock.Eq(taskID)).Return(nil)

		reqBody := todo_list.UpdateTaskRequest{
			Title: "updated-title",
		}
		body, _ := json.Marshal(reqBody)

		req, _ := http.NewRequest("PUT", "/api/todo-list/tasks/"+taskID, bytes.NewReader(body))
		recorder := httptest.NewRecorder()
		r.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}
