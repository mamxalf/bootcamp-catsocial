package service

import (
	"catsocial/configs"
	"catsocial/internal/domain/user/repository"
	"catsocial/internal/domain/user/request"
	"catsocial/internal/domain/user/response"
	"context"
)

type UserService interface {
	RegisterNewUser(ctx context.Context, req request.RegisterRequest) (res response.RegisterResponse, err error)
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	Config         *configs.Config
}

// ProvideUserServiceImpl is the provider for this service.
func ProvideUserServiceImpl(
	userRepository repository.UserRepository,
	config *configs.Config,
) *UserServiceImpl {
	s := new(UserServiceImpl)
	s.UserRepository = userRepository
	s.Config = config
	return s
}
