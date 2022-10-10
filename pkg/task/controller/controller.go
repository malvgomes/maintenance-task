package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	notificationModel "maintenance-task/pkg/notification/model"
	"maintenance-task/pkg/task/model"
	"maintenance-task/pkg/task/service"
	"maintenance-task/pkg/task/viewmodel"
	"maintenance-task/pkg/user/middleware"
	userModel "maintenance-task/pkg/user/model"
	userService "maintenance-task/pkg/user/service"
	"maintenance-task/shared/controller"
	"maintenance-task/shared/queue"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func NewTaskController(ctx context.Context) controller.Controller {
	return &TaskController{
		createTaskService: service.NewCreateTaskService(ctx),
		deleteTaskService: service.NewDeleteTaskService(ctx),
		listTasksService:  service.NewListTasksService(ctx),
		updateTaskService: service.NewUpdateTaskService(ctx),
		userMiddleware:    middleware.NewUserMiddleware(userService.NewGetUserService(ctx)),
		queueProducer:     ctx.Value("queueProducer").(queue.Producer),
	}
}

type TaskController struct {
	createTaskService *service.CreateTaskService
	deleteTaskService *service.DeleteTaskService
	listTasksService  *service.ListTasksService
	updateTaskService *service.UpdateTaskService
	userMiddleware    *middleware.UserMiddleware
	queueProducer     queue.Producer
}

func (c *TaskController) SetRoutes(r chi.Router) {
	r.Route("/tasks", func(r chi.Router) {
		r.Use(c.userMiddleware.UserMiddleware)
		r.Post("/", c.CreateTask)
		r.Delete("/{ID}", c.DeleteTask)
		r.Get("/{userID}", c.ListTasks)
		r.Put("/", c.UpdateTask)
	})
}

func (c *TaskController) CreateTask(w http.ResponseWriter, r *http.Request) {
	var createTaskPayload model.CreateTask

	err := json.NewDecoder(r.Body).Decode(&createTaskPayload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !createTaskPayload.IsValid() {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Invalid payload")
		return
	}

	user := r.Context().Value("session_user").(*userModel.User)
	if user.ID != createTaskPayload.UserID {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	ID, err := c.createTaskService.CreateTask(createTaskPayload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	message, err := json.Marshal(notificationModel.CreateNotification{
		TaskID: ID,
		UserID: 123,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = c.queueProducer.Publish("notification", message)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf("Task %d created", ID)))
	w.WriteHeader(http.StatusOK)
}

func (c *TaskController) DeleteTask(w http.ResponseWriter, r *http.Request) {
	ID, err := strconv.Atoi(chi.URLParam(r, "ID"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user := r.Context().Value("session_user").(*userModel.User)
	if user.UserRole != userModel.Manager {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	err = c.deleteTaskService.DeleteTask(ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (c *TaskController) ListTasks(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(r, "userID"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user := r.Context().Value("session_user").(*userModel.User)
	if user.ID != userID || user.UserRole != userModel.Manager {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	tasks, err := c.listTasksService.ListTasks(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	content, err := json.Marshal(viewmodel.MapToTaskListViewmodel(tasks))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(content)
}

func (c *TaskController) UpdateTask(w http.ResponseWriter, r *http.Request) {
	var updateTaskPayload model.UpdateTask

	err := json.NewDecoder(r.Body).Decode(&updateTaskPayload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !updateTaskPayload.IsValid() {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Invalid payload")
		return
	}

	user := r.Context().Value("session_user").(*userModel.User)
	if user.ID != updateTaskPayload.UserID {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	err = c.updateTaskService.UpdateTask(updateTaskPayload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	message, err := json.Marshal(notificationModel.CreateNotification{
		TaskID: updateTaskPayload.ID,
		UserID: user.ID,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = c.queueProducer.Publish("notification", message)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
