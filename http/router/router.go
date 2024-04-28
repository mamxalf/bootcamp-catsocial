package router

import (
	"catsocial/internal/handler/health"

	"github.com/go-chi/chi"
)

type DomainHandlers struct {
	HealthHandler health.HealthHandler
}

type Router struct {
	DomainHandlers DomainHandlers
}

func ProvideRouter(domainHandlers DomainHandlers) Router {
	return Router{
		DomainHandlers: domainHandlers,
	}
}

// SetupRoutes sets up all routing for this server.
func (r *Router) SetupRoutes(mux *chi.Mux) {
	mux.Route("/v1", func(rc chi.Router) {
		r.DomainHandlers.HealthHandler.Router(rc)
	})
}
