package controller

import "github.com/go-chi/chi"

type Controller interface {
	SetRoutes(r chi.Router)
}
