package middleware

import (
	"maintenance-task/pkg/user/model"
	"maintenance-task/pkg/user/service"
	"net/http"
)

type ManagerMiddleware struct {
	getUserService *service.GetUserService
}

func NewManagerMiddleware(getUserService *service.GetUserService) *ManagerMiddleware {
	return &ManagerMiddleware{
		getUserService: getUserService,
	}
}

func (s *ManagerMiddleware) ManagerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		user, err := s.getUserService.GetUser(username, password)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if user == nil || user.UserRole != model.Manager {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
