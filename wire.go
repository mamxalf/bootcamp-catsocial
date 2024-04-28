//go:build wireinject
// +build wireinject

package main

import (
	"catsocial/configs"
	"catsocial/http"
	"catsocial/http/middleware"
	"catsocial/http/router"
	"catsocial/infras"
	"catsocial/internal/domain/health/repository"
	"catsocial/internal/domain/health/service"
	"catsocial/internal/handler/health"
	"github.com/google/wire"
)

var configurations = wire.NewSet(
	configs.Get,
)

var persistences = wire.NewSet(
	infras.ProvidePostgresConn,
)

//var domainUser = wire.NewSet(
//	service.ProvideUserServiceImpl,
//	wire.Bind(new(service.UserService), new(*service.UserServiceImpl)),
//	repository.ProvideUserRepositoryPostgres,
//	wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryPostgres)),
//)

var domainHealth = wire.NewSet(
	service.ProvideHealthServiceImpl,
	wire.Bind(new(service.HealthService), new(*service.HealthServiceImpl)),
	repository.ProvideHealthRepositoryInfra,
	wire.Bind(new(repository.HealthRepository), new(*repository.HealthRepositoryInfra)),
)

// Wiring for all domains.
var domains = wire.NewSet(
	domainHealth,
)

var routing = wire.NewSet(
	wire.Struct(new(router.DomainHandlers), "*"),
	health.ProvideHealthHandler,
	router.ProvideRouter,
)

var authMiddleware = wire.NewSet(
	middleware.ProvideJWTMiddleware,
)

func InitializeService() *http.HTTP {
	wire.Build(
		// configurations
		configurations,
		// persistences
		persistences,
		// middleware
		authMiddleware,
		// domains
		domains,
		// routing
		routing,
		// selected transport layer
		http.ProvideHTTP)
	return &http.HTTP{}
}
