package main

import (
	"context"
	"log"
	"maintenance-task/pkg/notification/job"
	notificationRepository "maintenance-task/pkg/notification/repository"
	notificationService "maintenance-task/pkg/notification/service"
	taskRepository "maintenance-task/pkg/task/repository"
	taskService "maintenance-task/pkg/task/service"
	userRepository "maintenance-task/pkg/user/repository"
	userService "maintenance-task/pkg/user/service"
	"maintenance-task/shared/database"
	"maintenance-task/shared/queue"
)

func main() {
	consumer, err := queue.GetConsumer()
	if err != nil {
		log.Fatal(err)
	}

	ctx := initWorkerContext()

	err = consumer.Consume(ctx, "notification", job.HandleNotificationJob)
	if err != nil {
		log.Fatal(err)
	}
}

func initWorkerContext() context.Context {
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

	ctx = context.WithValue(ctx, "createNotificationService", notificationService.NewCreateNotificationService(ctx))
	ctx = context.WithValue(ctx, "getUserService", userService.NewGetUserService(ctx))
	ctx = context.WithValue(ctx, "getTaskService", taskService.NewGetTaskService(ctx))

	return ctx
}
