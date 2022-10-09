package controller

import (
	"context"
	"encoding/json"
	"log"
	"maintenance-task/pkg/user/model"
	"maintenance-task/pkg/user/service"
	"maintenance-task/shared/controller"
	"net/http"

	"github.com/go-chi/chi"
)

func NewUserController(ctx context.Context) controller.Controller {
	return &userController{createUserService: service.NewCreateUserService(ctx)}
}

type userController struct {
	createUserService *service.CreateUserService
}

func (c *userController) SetRoutes(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		r.Post("/", c.createUser)
	})
}

func (c *userController) createUser(w http.ResponseWriter, r *http.Request) {
	var createUserPayload model.CreateUser

	err := json.NewDecoder(r.Body).Decode(&createUserPayload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("An error ocurred:", err)
		return
	}

	if !createUserPayload.IsValid() {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Invalid payload")
		return
	}

	err = c.createUserService.CreateUser(createUserPayload)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		log.Println("An error ocurred:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
