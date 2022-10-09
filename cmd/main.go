package main

import (
	"context"
	"log"
	notificationRepository "maintenance-task/pkg/notification/repository"
	taskController "maintenance-task/pkg/task/controller"
	taskRepository "maintenance-task/pkg/task/repository"
	userController "maintenance-task/pkg/user/controller"
	userRepository "maintenance-task/pkg/user/repository"
	"maintenance-task/shared/database"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)

	ctx := initContext()

	uC := userController.NewUserController(ctx)
	uC.SetRoutes(router)

	tC := taskController.NewTaskController(ctx)
	tC.SetRoutes(router)

	log.Println("Listening on port 3000")
	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatal(err)
	}
}

func initContext() context.Context {
	ctx := context.Background()

	db, err := database.GetDatabase()
	if err != nil {
		log.Fatal(err)
	}

	ctx = context.WithValue(ctx, "database", db)

	userRepo := userRepository.GetUserRepository(ctx)
	notifRepo := notificationRepository.GetNotificationRepository(ctx)
	taskRepo := taskRepository.GetTaskRepository(ctx)

	ctx = context.WithValue(ctx, "userRepository", userRepo)
	ctx = context.WithValue(ctx, "notificationRepository", notifRepo)
	ctx = context.WithValue(ctx, "taskRepository", taskRepo)

	return ctx
}
