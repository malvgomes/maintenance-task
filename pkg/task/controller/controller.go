package controller

import (
	"context"
	"encoding/json"
	"log"
	"maintenance-task/pkg/task/model"
	"maintenance-task/pkg/task/service"
	"maintenance-task/pkg/task/viewmodel"
	"maintenance-task/shared/controller"
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
	}
}

type TaskController struct {
	createTaskService *service.CreateTaskService
	deleteTaskService *service.DeleteTaskService
	listTasksService  *service.ListTasksService
	updateTaskService *service.UpdateTaskService
}

func (c *TaskController) SetRoutes(r chi.Router) {
	r.Route("/tasks", func(r chi.Router) {
		r.Post("/", c.CreateTask)
		r.Delete("/{ID:[0-9]+}", c.DeleteTask)
		r.Get("/{userID:[0-9]+}", c.ListTasks)
		r.Put("/", c.UpdateTask)
	})
}

func (c *TaskController) CreateTask(w http.ResponseWriter, r *http.Request) {
	var createTaskPayload model.CreateTask

	err := json.NewDecoder(r.Body).Decode(&createTaskPayload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("An error ocurred:", err)
		return
	}

	if !createTaskPayload.IsValid() {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Invalid payload")
		return
	}

	err = c.createTaskService.CreateTask(createTaskPayload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("An error ocurred:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (c *TaskController) DeleteTask(w http.ResponseWriter, r *http.Request) {
	ID, err := strconv.Atoi(chi.URLParam(r, "ID"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("An error ocurred:", err)
		return
	}

	err = c.deleteTaskService.DeleteTask(ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("An error ocurred:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (c *TaskController) ListTasks(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(r, "userID"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("An error ocurred:", err)
		return
	}

	tasks, err := c.listTasksService.ListTasks(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("An error ocurred:", err)
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
		log.Println("An error ocurred:", err)
		return
	}

	if !updateTaskPayload.IsValid() {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Invalid payload")
		return
	}

	err = c.updateTaskService.UpdateTask(updateTaskPayload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("An error ocurred:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
