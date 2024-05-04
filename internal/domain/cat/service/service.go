package service

import (
	"catsocial/configs"
	catRepository "catsocial/internal/domain/cat/repository"
	"catsocial/internal/domain/cat/request"
	"catsocial/internal/domain/cat/response"
	userRepository "catsocial/internal/domain/user/repository"
	"context"

	"github.com/google/uuid"
)

type CatService interface {
	// Cat Service Interface
	InsertNewCat(ctx context.Context, req request.InsertCatRequest) (res response.CatResponse, err error)
	GetCatData(ctx context.Context, catID string) (res response.CatResponse, err error)
	GetAllCatData(ctx context.Context) (res []response.CatResponse, err error)
	UpdateCatData(ctx context.Context, catID uuid.UUID, req request.UpdateCatRequest) (res response.CatResponse, err error)
	DeleteCatData(ctx context.Context, catID uuid.UUID) (res response.CatResponse, err error)

	// Match Service Interface
	InsertNewMatch(ctx context.Context, req request.MatchRequest) (message string, err error)
	GetAllMatchesData(ctx context.Context) (res []response.MatchList, err error)
	ApproveCatMatch(ctx context.Context, matchID string) (message string, err error)
	RejectCatMatch(ctx context.Context, matchID string) (message string, err error)
	DeleteCatMatch(ctx context.Context, matchID string) (message string, err error)
}
type CatServiceImpl struct {
	CatRepository  catRepository.CatRepository
	UserRepository userRepository.UserRepository
	Config         *configs.Config
}

// ProvideCatServiceImpl is the provider for this service.
func ProvideCatServiceImpl(
	catRepository catRepository.CatRepository,
	userRepository userRepository.UserRepository,
	config *configs.Config,
) *CatServiceImpl {
	s := new(CatServiceImpl)
	s.CatRepository = catRepository
	s.UserRepository = userRepository
	s.Config = config
	return s
}
