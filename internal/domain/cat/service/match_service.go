package service

import (
	"catsocial/internal/domain/cat/request"
	"catsocial/internal/domain/cat/response"
	"catsocial/shared/logger"
	"context"
	"github.com/google/uuid"
)

func (u *CatServiceImpl) InsertNewMatch(ctx context.Context, req request.MatchRequest) (message string, err error) {
	insertModel, err := req.ToModel()
	if err != nil {
		message = "Failed Parse request to model"
		logger.ErrorWithStack(err)
		return
	}
	_, err = u.CatRepository.MatchRequest(ctx, &insertModel)
	if err != nil {
		message = "Failed to insert match request"
		logger.ErrorWithStack(err)
		return
	}

	message = "Successfully Inserted match"
	return
}

func (u *CatServiceImpl) GetAllMatchesData(ctx context.Context) (res []response.MatchList, err error) {
	//TODO implement me
	panic("implement me")
}

func (u *CatServiceImpl) ApproveCatMatch(ctx context.Context, matchIDStr string) (message string, err error) {
	matchID, err := uuid.Parse(matchIDStr)
	if err != nil {
		message = "Failed to parse match id"
		logger.ErrorWithStack(err)
		return
	}
	if err := u.CatRepository.IsApprove(ctx, matchID, true); err != nil {
		message = "Failed to approve match"
		logger.ErrorWithStack(err)
		return
	}
	message = "Successfully Approved match"
	return
}

func (u *CatServiceImpl) RejectCatMatch(ctx context.Context, matchIDStr string) (message string, err error) {
	matchID, err := uuid.Parse(matchIDStr)
	if err != nil {
		message = "Failed to parse match id"
		logger.ErrorWithStack(err)
		return
	}
	if err := u.CatRepository.IsApprove(ctx, matchID, false); err != nil {
		message = "Failed to reject match"
		logger.ErrorWithStack(err)
		return
	}
	message = "Successfully Rejected match"
	return
}

func (u *CatServiceImpl) DeleteCatMatch(ctx context.Context, matchIDStr string) (message string, err error) {
	matchID, err := uuid.Parse(matchIDStr)
	if err != nil {
		message = "Failed to parse match id"
		logger.ErrorWithStack(err)
		return
	}
	if err := u.CatRepository.DeleteMatch(ctx, matchID); err != nil {
		message = "Failed to delete match"
		logger.ErrorWithStack(err)
		return
	}
	message = "Successfully Deleted match"
	return
}
