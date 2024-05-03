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
func (u *CatServiceImpl) GetCatData(ctx context.Context, catID string) (res response.CatResponse, err error) {
	id, err := uuid.Parse(catID)
	if err != nil {
		return
	}
	cat, err := u.CatRepository.Find(ctx, id)
	if err != nil {
		return
	}
	res = response.CatResponse{
		Name:        cat.Name,
		Race:        cat.Race,
		Sex:         cat.Sex,
		AgeInMonth:  cat.Age,
		Description: cat.Descriptions,
		ImageUrls:   []string{"apple", "grape", "banana", "melon"},
	}
	return
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
