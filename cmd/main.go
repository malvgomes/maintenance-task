package main

import (
	"context"
	"log"
	userController "maintenance-task/pkg/user/controller"
	"maintenance-task/pkg/user/repository"
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

	userRepository := repository.GetUserRepository(ctx)

	ctx = context.WithValue(ctx, "userRepository", userRepository)

	return ctx
}
