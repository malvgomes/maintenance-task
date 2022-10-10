package job

import (
	"context"
	"encoding/json"
	"log"
	"maintenance-task/pkg/notification/model"
	notificationService "maintenance-task/pkg/notification/service"
	taskService "maintenance-task/pkg/task/service"
	userModel "maintenance-task/pkg/user/model"
	userService "maintenance-task/pkg/user/service"
	"sync"
	"time"
)

func HandleNotificationJob(ctx context.Context, message []byte) {
	createNotificationService := ctx.Value("createNotificationService").(*notificationService.CreateNotificationService)
	getUserService := ctx.Value("getUserService").(*userService.GetUserService)
	getTaskService := ctx.Value("getTaskService").(*taskService.GetTaskService)

	var notification model.CreateNotification

	if err := json.Unmarshal(message, &notification); err != nil {
		log.Println("An error occurred trying to parse the notification payload", err)
		return
	}

	var wg sync.WaitGroup

	user, err := getUserService.GetUserByID(notification.UserID)
	if err != nil {
		log.Println("An error occurred trying fetch user by ID", err)
		return
	}

	task, err := getTaskService.GetTask(notification.TaskID)
	if err != nil {
		log.Println("An error occurred trying fetch task by ID", err)
		return
	}

	managers, err := getUserService.GetUsersByRole(userModel.Manager)
	if err != nil {
		log.Println("An error occurred trying fetch users by role", err)
		return
	}

	var done = make(chan bool)

	for _, manager := range managers {
		wg.Add(1)

		go func(u, m *userModel.User) {
			if u.ID != m.ID {
				_, err = createNotificationService.CreateNotification(model.CreateNotification{
					UserID: m.ID,
					TaskID: task.ID,
				})

				date := task.CreatedAt
				if task.UpdatedAt != nil {
					date = *task.UpdatedAt
				}

				log.Printf("To manager %s. At %s, %s completed the following task: %s", m.Username,
					date.Format(time.RFC822), u.FirstName, task.Summary)
			}

			wg.Done()
		}(user, manager)
	}

	go func() {
		wg.Wait()
		done <- true
	}()

	<-done
}
