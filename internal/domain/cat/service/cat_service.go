package service

import (
	"catsocial/internal/domain/cat/request"
	"catsocial/internal/domain/cat/response"
	"catsocial/shared/logger"

	"catsocial/shared/failure"
	// "catsocial/shared/token"
	"context"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func (u *CatServiceImpl) InsertNewCat(ctx context.Context, req request.InsertCatRequest) (message string, err error) {
	cat, err := req.ToModel()
	if err != nil {
		logger.ErrorWithStack(err)
		err = failure.BadRequestFromString("doesn't pass validation")
		return
	}

	_, err = u.CatRepository.Insert(ctx, cat)
	if err != nil {
		logger.ErrorWithStack(err)
		err = failure.BadRequestFromString("can't insert new cat")
		return
	}

	message = "successfully insert new cat"
	return
}
func (u *CatServiceImpl) GetCatData(ctx context.Context, userID uuid.UUID, catID string) (res response.CatResponse, err error) {
	id, err := uuid.Parse(catID)
	if err != nil {
		return
	}
	cat, err := u.CatRepository.Find(ctx, userID, id)
	if err != nil {
		return
	}
	res = response.CatResponse{
		ID:          cat.ID,
		Name:        cat.Name,
		Race:        cat.Race,
		Sex:         cat.Sex,
		AgeInMonth:  cat.Age,
		Description: cat.Descriptions,
		ImageUrls:   []string{"apple", "grape", "banana", "melon"},
	}
	return
}

func (u *CatServiceImpl) GetAllCatData(ctx context.Context, userId uuid.UUID, params request.CatQueryParams) (res []response.CatResponse, err error) {
	var catList []response.CatResponse
	cats, err := u.CatRepository.FindAll(ctx, userId, params)
	if err != nil {
		return
	}
	for _, cat := range cats {
		catList = append(catList, response.CatResponse{
			ID:          cat.ID,
			Name:        cat.Name,
			Race:        cat.Race,
			Sex:         cat.Sex,
			AgeInMonth:  cat.Age,
			Description: cat.Descriptions,
			ImageUrls:   cat.Images,
		})

	}
	res = catList
	return
}

func (u *CatServiceImpl) UpdateCatData(ctx context.Context, catID uuid.UUID, req request.UpdateCatRequest) (res response.CatResponse, err error) {
	cat, err := req.ToModel()
	if err != nil {
		logger.ErrorWithStack(err)
		err = failure.BadRequestFromString("doesn't pass validation")
		return
	}
	_, err = u.CatRepository.Update(ctx, catID, cat)
	if err != nil {
		logger.ErrorWithStack(err)
		return
	}
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
