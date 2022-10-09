package controller_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	userController "maintenance-task/pkg/user/controller"
	"maintenance-task/pkg/user/model"
	"maintenance-task/shared/controller"
	mockRepository "maintenance-task/shared/mock/user/repository"
	"maintenance-task/shared/pointer"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-chi/chi"

	"github.com/golang/mock/gomock"
)

func TestCreateUser(t *testing.T) {
	input := model.CreateUser{
		Username:  "userName",
		Password:  "password",
		FirstName: "firstName",
		LastName:  pointer.String("lastName"),
		UserRole:  model.Manager,
	}

	t.Run("Success", func(t *testing.T) {
		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(input)
		assert.NoError(t, err)

		userCtrlr, mock, finish := getMockedController(t)
		defer finish()

		mock.EXPECT().CreateUser(input).Return(nil)

		router := chi.NewRouter()
		userCtrlr.SetRoutes(router)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/users", &buf)

		userCtrlr.(*userController.UserController).CreateUser(w, r)
		result := w.Result()
		defer result.Body.Close()

		_, err = io.ReadAll(result.Body)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, result.StatusCode)

	})

	t.Run("Failure", func(t *testing.T) {
		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(input)
		assert.NoError(t, err)

		userCtrlr, mock, finish := getMockedController(t)
		defer finish()

		mock.EXPECT().CreateUser(input).Return(errors.New("create failure"))

		router := chi.NewRouter()
		userCtrlr.SetRoutes(router)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/users", &buf)

		userCtrlr.(*userController.UserController).CreateUser(w, r)
		result := w.Result()
		defer result.Body.Close()

		_, err = io.ReadAll(result.Body)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, result.StatusCode)
	})

	t.Run("Bad input", func(t *testing.T) {
		userCtrlr, _, finish := getMockedController(t)
		defer finish()

		router := chi.NewRouter()
		userCtrlr.SetRoutes(router)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/users", nil)

		userCtrlr.(*userController.UserController).CreateUser(w, r)
		result := w.Result()
		defer result.Body.Close()

		_, err := io.ReadAll(result.Body)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, result.StatusCode)
	})

	t.Run("Bad Request", func(t *testing.T) {
		invalidInput := model.CreateUser{}

		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(invalidInput)
		assert.NoError(t, err)

		userCtrlr, _, finish := getMockedController(t)
		defer finish()

		router := chi.NewRouter()
		userCtrlr.SetRoutes(router)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/users", &buf)

		userCtrlr.(*userController.UserController).CreateUser(w, r)
		result := w.Result()
		defer result.Body.Close()

		_, err = io.ReadAll(result.Body)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, result.StatusCode)
	})
}

func TestDeleteUser(t *testing.T) {
	input := model.DeleteUser{
		Username: "userName",
	}

	t.Run("Success", func(t *testing.T) {
		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(input)
		assert.NoError(t, err)

		userCtrlr, mock, finish := getMockedController(t)
		defer finish()

		mock.EXPECT().DeleteUser("userName").Return(nil)

		router := chi.NewRouter()
		userCtrlr.SetRoutes(router)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodDelete, "/users", &buf)

		userCtrlr.(*userController.UserController).DeleteUser(w, r)
		result := w.Result()
		defer result.Body.Close()

		_, err = io.ReadAll(result.Body)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, result.StatusCode)

	})

	t.Run("Failure", func(t *testing.T) {
		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(input)
		assert.NoError(t, err)

		userCtrlr, mock, finish := getMockedController(t)
		defer finish()

		mock.EXPECT().DeleteUser("userName").Return(errors.New("create failure"))

		router := chi.NewRouter()
		userCtrlr.SetRoutes(router)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodDelete, "/users", &buf)

		userCtrlr.(*userController.UserController).DeleteUser(w, r)
		result := w.Result()
		defer result.Body.Close()

		_, err = io.ReadAll(result.Body)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, result.StatusCode)
	})

	t.Run("Bad input", func(t *testing.T) {
		userCtrlr, _, finish := getMockedController(t)
		defer finish()

		router := chi.NewRouter()
		userCtrlr.SetRoutes(router)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodDelete, "/users", nil)

		userCtrlr.(*userController.UserController).DeleteUser(w, r)
		result := w.Result()
		defer result.Body.Close()

		_, err := io.ReadAll(result.Body)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, result.StatusCode)
	})

	t.Run("Bad Request", func(t *testing.T) {
		invalidInput := model.CreateUser{}

		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(invalidInput)
		assert.NoError(t, err)

		userCtrlr, _, finish := getMockedController(t)
		defer finish()

		router := chi.NewRouter()
		userCtrlr.SetRoutes(router)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodDelete, "/users", &buf)

		userCtrlr.(*userController.UserController).DeleteUser(w, r)
		result := w.Result()
		defer result.Body.Close()

		_, err = io.ReadAll(result.Body)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, result.StatusCode)
	})

}

func getMockedController(t *testing.T) (
	controller.Controller,
	*mockRepository.MockUserRepository,
	func(),
) {
	t.Helper()

	ctrl := gomock.NewController(t)

	finish := func() {
		ctrl.Finish()
	}

	userRepositoryMock := mockRepository.NewMockUserRepository(ctrl)

	ctx := context.WithValue(context.Background(), "userRepository", userRepositoryMock)

	return userController.NewUserController(ctx), userRepositoryMock, finish

}
