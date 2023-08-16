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
