package service

import (
	"catsocial/configs"
	"catsocial/internal/domain/cat/repository"
	"catsocial/internal/domain/cat/request"
	"catsocial/internal/domain/cat/response"
	"context"

	"github.com/google/uuid"
)

type CatService interface {
	InsertNewCat(ctx context.Context, req request.InsertCatRequest) (res response.CatResponse, err error)
	GetCatData(ctx context.Context, catID uuid.UUID) (res response.CatResponse, err error)
	GetAllCatData(ctx context.Context) (res response.CatResponse, err error)
	UpdateCatData(ctx context.Context, catID uuid.UUID, req request.UpdateCatRequest) (res response.CatResponse, err error)
	DeleteCatData(ctx context.Context, catID uuid.UUID) (res response.CatResponse, err error)
}
type CatServiceImpl struct {
	CatRepository repository.CatRepository
	Config        *configs.Config
}

// ProvideCatServiceImpl is the provider for this service.
func ProvideCatServiceImpl(
	catRepository repository.CatRepository,
	config *configs.Config,
) *CatServiceImpl {
	s := new(CatServiceImpl)
	s.CatRepository = catRepository
	s.Config = config
	return s
}
