package user

import (
	"catsocial/http/middleware"
	"catsocial/internal/domain/user/service"

	"github.com/go-chi/chi"
)

type UserHandler struct {
	UserService   service.UserService
	JWTMiddleware *middleware.JWT
}

func ProvideUserHandler(userService service.UserService, jwtMiddleware *middleware.JWT) UserHandler {
	return UserHandler{
		UserService:   userService,
		JWTMiddleware: jwtMiddleware,
	}
}

func (h *UserHandler) Router(r chi.Router) {
	r.Route("/user", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Post("/register", h.Register)
			r.Post("/login", h.Login)
		})
	})
}
