package service

import (
	"catsocial/internal/domain/cat/model"
	"catsocial/internal/domain/cat/request"
	"catsocial/internal/domain/cat/response"

	"catsocial/shared/failure"
	// "catsocial/shared/token"
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func (u *CatServiceImpl) InsertNewCat(ctx context.Context, req request.InsertCatRequest) (res response.CatResponse, err error) {
	cat := model.InsertCat{
		Name:         req.Name,
		Race:         req.Race,
		Sex:          req.Sex,
		Age:          req.AgeInMonth,
		Descriptions: req.Description,
		Images:       req.ImageUrls,
	}

	_, err = u.CatRepository.Insert(ctx, &cat)
	if err != nil {
		if failure.GetCode(err) == http.StatusConflict {
			log.Error().Interface("params", req).Err(err).Msg("[InsertNewCat - Service] Cat should be unique")
			return
		}
		log.Error().Interface("params", req).Err(err).Msg("[InsertNewCat - Service] Internal Error")
		return
	}

	res = response.CatResponse{
		Name:        req.Name,
		Race:        req.Race,
		Sex:         req.Sex,
		AgeInMonth:  req.AgeInMonth,
		Description: req.Description,
		ImageUrls:   req.ImageUrls,
	}

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

func (u *CatServiceImpl) GetAllCatData(ctx context.Context) (res []response.CatResponse, err error) {
	var catList []response.CatResponse
	cats, err := u.CatRepository.FindAll(ctx)
	if err != nil {
		return
	}
	for _, cat := range cats {
		catList = append(catList, response.CatResponse{
			Name:        cat.Name,
			Race:        cat.Race,
			Sex:         cat.Sex,
			AgeInMonth:  cat.Age,
			Description: cat.Descriptions,
			ImageUrls:   []string{"apple", "grape", "banana", "melon"},
		})

	}
	res = catList
	return
}

func (u *CatServiceImpl) UpdateCatData(ctx context.Context, catID uuid.UUID, req request.UpdateCatRequest) (res response.CatResponse, err error) {
	return
}

func (u *CatServiceImpl) DeleteCatData(ctx context.Context, catIDStr string) (res response.CatResponse, err error) {
	catID, err := uuid.Parse(catIDStr)
	if err != nil {
		return
	}
	_, err = u.CatRepository.Delete(ctx, catID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteCatData - Service] Internal Error")
		return
	}
	return
}
