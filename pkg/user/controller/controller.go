package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"maintenance-task/pkg/user/middleware"
	"maintenance-task/pkg/user/model"
	"maintenance-task/pkg/user/service"
	"maintenance-task/shared/controller"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func NewUserController(ctx context.Context) controller.Controller {
	return &UserController{
		createUserService:    service.NewCreateUserService(ctx),
		deleteUserService:    service.NewDeleteUserService(ctx),
		getManagerMiddleware: middleware.NewManagerMiddleware(service.NewGetUserService(ctx)),
	}
}

type UserController struct {
	createUserService    *service.CreateUserService
	deleteUserService    *service.DeleteUserService
	getManagerMiddleware *middleware.ManagerMiddleware
}

func (c *UserController) SetRoutes(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		r.Use(c.getManagerMiddleware.ManagerMiddleware)
		r.Post("/", c.CreateUser)
		r.Delete("/{ID}", c.DeleteUser)
	})
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var createUserPayload model.CreateUser

	err := json.NewDecoder(r.Body).Decode(&createUserPayload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !createUserPayload.IsValid() {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Invalid payload")
		return
	}

	ID, err := c.createUserService.CreateUser(createUserPayload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf("User %d created", ID)))
	w.WriteHeader(http.StatusOK)
}

func (c *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ID, err := strconv.Atoi(chi.URLParam(r, "ID"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = c.deleteUserService.DeleteUser(ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
