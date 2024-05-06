package router

import (
	"catsocial/internal/handler/cat"
	"catsocial/internal/handler/health"
	"catsocial/internal/handler/user"

	"github.com/go-chi/chi"
)

type DomainHandlers struct {
	HealthHandler health.HealthHandler
	UserHandler   user.UserHandler
	CatHandler    cat.CatHandler
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
		r.DomainHandlers.UserHandler.Router(rc)
		r.DomainHandlers.CatHandler.Router(rc)
	})
}
