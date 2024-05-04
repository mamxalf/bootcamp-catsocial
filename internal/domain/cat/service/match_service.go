package service

import (
	"catsocial/internal/domain/cat/request"
	"catsocial/internal/domain/cat/response"
	"catsocial/shared/failure"
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

	if insertModel.IssuedUserID != insertModel.UserCatID {
		err = failure.BadRequestFromString("cat is not issued by user")
		return
	}

	// Find Cat by User Cat ID
	userCat, err := u.CatRepository.Find(ctx, insertModel.UserCatID)
	// if neither matchCatId / userCatId is not found
	if err != nil {
		return
	}

	//  if userCatId is not belong to the user
	user, err := u.UserRepository.GetUserByID(ctx, userCat.UserID)
	if err != nil {
		return
	}

	// Find Cat by Match Cat ID
	matchCat, err := u.CatRepository.Find(ctx, insertModel.MatchCatID)
	// if neither matchCatId / userCatId is not found
	if err != nil {
		return
	}

	if matchCat.UserID == user.ID {
		err = failure.BadRequestFromString("From Same owner!")
		return
	}

	if matchCat.Sex == userCat.Sex {
		err = failure.BadRequestFromString("Cat Gender is Same!")
		return
	}

	// if both matchCatId & userCatId already matched
	matchCatUser, err := u.CatRepository.FindMatchByMatchCatID(ctx, matchCat.UserID)
	if err != nil {
		return
	}
	userCatUser, err := u.CatRepository.FindMatchByUserCatID(ctx, userCat.UserID)
	if err != nil {
		return
	}
	if matchCatUser.IsApproved && userCatUser.IsApproved {
		err = failure.BadRequestFromString("Both Cat Already Matched!")
		return
	}

	// successfully send match request
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
	if err = u.CatRepository.IsApprove(ctx, matchID, true); err != nil {
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
	if err = u.CatRepository.IsApprove(ctx, matchID, false); err != nil {
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
	if err = u.CatRepository.DeleteMatch(ctx, matchID); err != nil {
		message = "Failed to delete match"
		logger.ErrorWithStack(err)
		return
	}
	message = "Successfully Deleted match"
	return
}
