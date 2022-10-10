package middleware

import (
	"context"
	"maintenance-task/pkg/user/service"
	"net/http"
)

type UserMiddleware struct {
	getUserService *service.GetUserService
}

func NewUserMiddleware(getUserService *service.GetUserService) *UserMiddleware {
	return &UserMiddleware{
		getUserService: getUserService,
	}
}

func (s *UserMiddleware) UserMiddleware(next http.Handler) http.Handler {
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

		if user == nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		ctx := context.WithValue(r.Context(), "session_user", user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
