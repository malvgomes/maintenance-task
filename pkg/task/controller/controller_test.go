package controller_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	taskController "maintenance-task/pkg/task/controller"
	"maintenance-task/pkg/task/model"
	userModel "maintenance-task/pkg/user/model"
	"maintenance-task/shared/controller"
	mockRepository "maintenance-task/shared/mock/task/repository"
	mockUserRepository "maintenance-task/shared/mock/user/repository"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
)

func TestCreateTask(t *testing.T) {
	input := model.CreateTask{
		UserID:  123,
		Summary: "summary",
	}

	t.Run("Success", func(t *testing.T) {
		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(input)
		assert.NoError(t, err)

		routerContext, taskCtrlr, taskMock, finish := getMockedController(t)
		defer finish()

		taskMock.EXPECT().CreateTask(input).Return(nil)

		router := chi.NewRouter()
		taskCtrlr.SetRoutes(router)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/tasks", &buf)
		r = r.WithContext(routerContext)

		taskCtrlr.(*taskController.TaskController).CreateTask(w, r)
		result := w.Result()
		defer result.Body.Close()

		_, err = io.ReadAll(result.Body)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, result.StatusCode)
	})

	t.Run("Failure", func(t *testing.T) {
		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(input)
		assert.NoError(t, err)

		routerContext, taskCtrlr, taskMock, finish := getMockedController(t)
		defer finish()

		taskMock.EXPECT().CreateTask(input).Return(errors.New("create failure"))

		router := chi.NewRouter()
		taskCtrlr.SetRoutes(router)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/tasks", &buf)

		r = r.WithContext(routerContext)

		taskCtrlr.(*taskController.TaskController).CreateTask(w, r)
		result := w.Result()
		defer result.Body.Close()

		_, err = io.ReadAll(result.Body)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, result.StatusCode)
	})

	t.Run("Bad Input", func(t *testing.T) {
		routerContext, taskCtrlr, _, finish := getMockedController(t)
		defer finish()

		router := chi.NewRouter()
		taskCtrlr.SetRoutes(router)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/tasks", nil)
		r = r.WithContext(routerContext)

		taskCtrlr.(*taskController.TaskController).CreateTask(w, r)
		result := w.Result()
		defer result.Body.Close()

		_, err := io.ReadAll(result.Body)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, result.StatusCode)
	})

	t.Run("Bad Request", func(t *testing.T) {
		invalidInput := model.CreateTask{}

		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(invalidInput)
		assert.NoError(t, err)

		routerContext, taskCtrlr, _, finish := getMockedController(t)
		defer finish()

		router := chi.NewRouter()
		taskCtrlr.SetRoutes(router)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/tasks", &buf)
		r = r.WithContext(routerContext)

		taskCtrlr.(*taskController.TaskController).CreateTask(w, r)
		result := w.Result()
		defer result.Body.Close()

		_, err = io.ReadAll(result.Body)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, result.StatusCode)
	})
}

func TestDeleteTask(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		routerContext, taskCtrlr, mock, finish := getMockedController(t)
		defer finish()

		mock.EXPECT().DeleteTask(123).Return(nil)

		router := chi.NewRouter()
		taskCtrlr.SetRoutes(router)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodDelete, "/tasks/123", nil)

		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("ID", "123")

		r = r.WithContext(context.WithValue(routerContext, chi.RouteCtxKey, rctx))

		taskCtrlr.(*taskController.TaskController).DeleteTask(w, r)
		result := w.Result()
		defer result.Body.Close()

		_, err := io.ReadAll(result.Body)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNoContent, result.StatusCode)
	})

	t.Run("Failure", func(t *testing.T) {
		routerContext, taskCtrlr, mock, finish := getMockedController(t)
		defer finish()

		mock.EXPECT().DeleteTask(123).Return(errors.New("delete failure"))

		router := chi.NewRouter()
		taskCtrlr.SetRoutes(router)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/tasks/123", nil)

		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("ID", "123")

		r = r.WithContext(context.WithValue(routerContext, chi.RouteCtxKey, rctx))

		taskCtrlr.(*taskController.TaskController).DeleteTask(w, r)
		result := w.Result()
		defer result.Body.Close()

		_, err := io.ReadAll(result.Body)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, result.StatusCode)
	})

	t.Run("Invalid input", func(t *testing.T) {
		_, taskCtrlr, _, finish := getMockedController(t)
		defer finish()

		router := chi.NewRouter()
		taskCtrlr.SetRoutes(router)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/tasks/test", nil)

		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("ID", "test")

		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))

		taskCtrlr.(*taskController.TaskController).DeleteTask(w, r)
		result := w.Result()
		defer result.Body.Close()

		_, err := io.ReadAll(result.Body)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, result.StatusCode)
	})
}

func TestListTasks(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		routerContext, taskCtrlr, mock, finish := getMockedController(t)
		defer finish()

		loc, err := time.LoadLocation("America/Sao_Paulo")
		assert.NoError(t, err)

		date := time.Date(2022, 9, 9, 11, 12, 13, 0, loc)

		mock.EXPECT().ListTasks(123).Return([]*model.Task{{
			ID:        1,
			UserID:    123,
			Summary:   "summary",
			CreatedAt: date,
			UpdatedAt: nil,
		}}, nil)

		router := chi.NewRouter()
		taskCtrlr.SetRoutes(router)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/tasks/123", nil)

		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("userID", "123")

		r = r.WithContext(context.WithValue(routerContext, chi.RouteCtxKey, rctx))

		taskCtrlr.(*taskController.TaskController).ListTasks(w, r)
		result := w.Result()
		defer result.Body.Close()

		body, err := io.ReadAll(result.Body)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, result.StatusCode)
		assert.Equal(t, `[{"id":1,"userId":123,"summary":"summary","createdAt":"2022-09-09T11:12:13-03:00","updatedAt":null}]`, string(body))
	})

	t.Run("Failure", func(t *testing.T) {
		routerContext, taskCtrlr, mock, finish := getMockedController(t)
		defer finish()
		mock.EXPECT().ListTasks(123).Return(nil, errors.New("list failure"))

		router := chi.NewRouter()
		taskCtrlr.SetRoutes(router)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/tasks/123", nil)

		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("userID", "123")

		r = r.WithContext(context.WithValue(routerContext, chi.RouteCtxKey, rctx))

		taskCtrlr.(*taskController.TaskController).ListTasks(w, r)
		result := w.Result()
		defer result.Body.Close()

		_, err := io.ReadAll(result.Body)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, result.StatusCode)
	})

	t.Run("Invalid input", func(t *testing.T) {
		_, taskCtrlr, _, finish := getMockedController(t)
		defer finish()

		router := chi.NewRouter()
		taskCtrlr.SetRoutes(router)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/tasks/test", nil)

		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("userID", "test")

		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))

		taskCtrlr.(*taskController.TaskController).ListTasks(w, r)
		result := w.Result()
		defer result.Body.Close()

		_, err := io.ReadAll(result.Body)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, result.StatusCode)
	})
}

func TestUpdateTask(t *testing.T) {
	input := model.UpdateTask{
		ID:      123,
		UserID:  123,
		Summary: "summary",
	}

	t.Run("Success", func(t *testing.T) {
		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(input)
		assert.NoError(t, err)

		routerContext, taskCtrlr, mock, finish := getMockedController(t)
		defer finish()

		mock.EXPECT().UpdateTask(input).Return(nil)

		router := chi.NewRouter()
		taskCtrlr.SetRoutes(router)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPut, "/tasks", &buf)
		r = r.WithContext(routerContext)

		taskCtrlr.(*taskController.TaskController).UpdateTask(w, r)
		result := w.Result()
		defer result.Body.Close()

		_, err = io.ReadAll(result.Body)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNoContent, result.StatusCode)
	})

	t.Run("Failure", func(t *testing.T) {
		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(input)
		assert.NoError(t, err)

		routerContext, taskCtrlr, mock, finish := getMockedController(t)
		defer finish()

		mock.EXPECT().UpdateTask(input).Return(errors.New("update failure"))

		router := chi.NewRouter()
		taskCtrlr.SetRoutes(router)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPut, "/tasks", &buf)
		r = r.WithContext(routerContext)

		taskCtrlr.(*taskController.TaskController).UpdateTask(w, r)
		result := w.Result()
		defer result.Body.Close()

		_, err = io.ReadAll(result.Body)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, result.StatusCode)
	})

	t.Run("Bad Input", func(t *testing.T) {
		_, taskCtrlr, _, finish := getMockedController(t)
		defer finish()

		router := chi.NewRouter()
		taskCtrlr.SetRoutes(router)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPut, "/tasks", nil)

		taskCtrlr.(*taskController.TaskController).UpdateTask(w, r)
		result := w.Result()
		defer result.Body.Close()

		_, err := io.ReadAll(result.Body)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, result.StatusCode)
	})

	t.Run("Bad Request", func(t *testing.T) {
		invalidInput := model.UpdateTask{}

		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(invalidInput)
		assert.NoError(t, err)

		routerContext, taskCtrlr, _, finish := getMockedController(t)
		defer finish()

		router := chi.NewRouter()
		taskCtrlr.SetRoutes(router)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPut, "/tasks", &buf)
		r = r.WithContext(routerContext)

		taskCtrlr.(*taskController.TaskController).UpdateTask(w, r)
		result := w.Result()
		defer result.Body.Close()

		_, err = io.ReadAll(result.Body)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, result.StatusCode)
	})
}

func getMockedController(t *testing.T) (
	context.Context,
	controller.Controller,
	*mockRepository.MockTaskRepository,
	func(),
) {
	t.Helper()

	ctrl := gomock.NewController(t)

	userRepositoryMock := mockUserRepository.NewMockUserRepository(ctrl)
	taskRepositoryMock := mockRepository.NewMockTaskRepository(ctrl)

	finish := func() {
		ctrl.Finish()
	}

	ctx := context.WithValue(context.Background(), "userRepository", userRepositoryMock)
	ctx = context.WithValue(ctx, "taskRepository", taskRepositoryMock)

	routerContext := context.WithValue(context.Background(), "session_user",
		&userModel.User{ID: 123, UserRole: userModel.Manager})

	return routerContext, taskController.NewTaskController(ctx), taskRepositoryMock, finish

}
