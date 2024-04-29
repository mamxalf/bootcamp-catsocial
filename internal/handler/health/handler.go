package health

import (
	"catsocial/http/middleware"
	"catsocial/internal/domain/health/service"

	"github.com/go-chi/chi"
)

type HealthHandler struct {
	HealthService service.HealthService
	JWTMiddleware *middleware.JWT
}

func ProvideHealthHandler(healthService service.HealthService, jwt *middleware.JWT) HealthHandler {
	return HealthHandler{
		HealthService: healthService,
		JWTMiddleware: jwt,
	}
}

func (h *HealthHandler) Router(r chi.Router) {
	r.Route("/health", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Get("/ping", h.Ping)
			r.Get("/ping-db", h.PingDB)
		})
	})
}
