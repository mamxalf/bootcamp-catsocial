package service

import (
	"catsocial/internal/domain/cat/request"
	"catsocial/internal/domain/cat/response"

	// "catsocial/shared/failure"
	// "catsocial/shared/token"
	"context"
	// "net/http"
	// "github.com/rs/zerolog/log"
	"github.com/google/uuid"
)

func (u *CatServiceImpl) InsertNewCat(ctx context.Context, req request.InsertCatRequest) (res response.CatResponse, err error) {
	return
}
func (u *CatServiceImpl) GetCatData(ctx context.Context, catID uuid.UUID) (res response.CatResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (u *CatServiceImpl) GetAllCatData(ctx context.Context) (res response.CatResponse, err error) {
	return
}

func (u *CatServiceImpl) UpdateCatData(ctx context.Context, catID uuid.UUID, req request.UpdateCatRequest) (res response.CatResponse, err error) {
	return
}

func (u *CatServiceImpl) DeleteCatData(ctx context.Context, catID uuid.UUID) (res response.CatResponse, err error) {
	return
}
