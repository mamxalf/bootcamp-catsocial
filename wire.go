//go:build wireinject
// +build wireinject

package main

import (
	"catsocial/configs"
	"catsocial/http"
	"catsocial/http/middleware"
	"catsocial/http/router"
	"catsocial/infras"
	catRepository "catsocial/internal/domain/cat/repository"
	catService "catsocial/internal/domain/cat/service"
	healthRepository "catsocial/internal/domain/health/repository"
	healthService "catsocial/internal/domain/health/service"
	userRepository "catsocial/internal/domain/user/repository"
	userService "catsocial/internal/domain/user/service"
	"catsocial/internal/handler/cat"
	"catsocial/internal/handler/health"
	"catsocial/internal/handler/user"

	"github.com/google/wire"
)

var configurations = wire.NewSet(
	configs.Get,
)

var persistences = wire.NewSet(
	infras.ProvidePostgresConn,
)

var domainUser = wire.NewSet(
	userService.ProvideUserServiceImpl,
	wire.Bind(new(userService.UserService), new(*userService.UserServiceImpl)),
	userRepository.ProvideUserRepositoryInfra,
	wire.Bind(new(userRepository.UserRepository), new(*userRepository.UserRepositoryInfra)),
)

var domainHealth = wire.NewSet(
	healthService.ProvideHealthServiceImpl,
	wire.Bind(new(healthService.HealthService), new(*healthService.HealthServiceImpl)),
	healthRepository.ProvideHealthRepositoryInfra,
	wire.Bind(new(healthRepository.HealthRepository), new(*healthRepository.HealthRepositoryInfra)),
)
var domainCat = wire.NewSet(
	catService.ProvideCatServiceImpl,
	wire.Bind(new(catService.CatService), new(*catService.CatServiceImpl)),
	catRepository.ProvideCatRepositoryInfra,
	wire.Bind(new(catRepository.CatRepository), new(*catRepository.CatRepositoryInfra)),
)

// Wiring for all domains.
var domains = wire.NewSet(
	domainHealth,
	domainUser,
	domainCat,
)

var routing = wire.NewSet(
	wire.Struct(new(router.DomainHandlers), "*"),
	health.ProvideHealthHandler,
	user.ProvideUserHandler,
	cat.ProvideCatHandler,
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
