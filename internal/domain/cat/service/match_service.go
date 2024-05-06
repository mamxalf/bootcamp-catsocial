package service

import (
	"catsocial/internal/domain/cat/model"
	"catsocial/internal/domain/cat/request"
	"catsocial/shared/failure"
	"catsocial/shared/logger"
	"context"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func (u *CatServiceImpl) InsertNewMatch(ctx context.Context, userID uuid.UUID, req request.MatchRequest) (message string, err error) {
	insertModel, err := req.ToModel()
	if err != nil {
		message = "Failed Parse request to model"
		logger.ErrorWithStack(err)
		return
	}

	// Find Cat by User Cat ID
	userCat, err := u.CatRepository.Find(ctx, insertModel.UserCatID)
	// if neither matchCatId / userCatId is not found
	if err != nil {
		err = failure.NotFound("User Cat not found!")
		return
	}

	// Find Cat by Match Cat ID
	matchCat, err := u.CatRepository.Find(ctx, insertModel.MatchCatID)
	// if neither matchCatId / userCatId is not found
	if err != nil {
		err = failure.NotFound("Match Cat not found!")
		return
	}

	//  if userCatId is not belong to the user
	user, err := u.UserRepository.GetUserByID(ctx, userCat.UserID)
	if err != nil {
		return
	}

	if userCat.UserID != user.ID {
		err = failure.NotFound("Not belong user!")
		return
	}

	if matchCat.Sex == userCat.Sex {
		err = failure.BadRequestFromString("Cat Gender is Same!")
		return
	}

	if matchCat.HasMatched && userCat.HasMatched {
		err = failure.BadRequestFromString("Both Has Matched!")
		return
	}
	if matchCat.UserID == user.ID {
		err = failure.BadRequestFromString("From Same owner!")
		return
	}

	// successfully send match request
	insertModel.IssuedUserID = user.ID
	_, err = u.CatRepository.MatchRequest(ctx, &insertModel)
	if err != nil {
		log.Error().Interface("req model", insertModel).Msg("[MatchRequest]")
		message = "Failed to insert match request"
		return
	}

	message = "Successfully Inserted match"
	return
}

func (u *CatServiceImpl) GetAllMatchesData(ctx context.Context) (res []model.MatchDetails, err error) {
	matches, err := u.CatRepository.FindAllMatches(ctx)
	if err != nil {
		return
	}
	if len(matches) == 0 {
		return []model.MatchDetails{}, nil
	}
	res = matches
	return
}

func (u *CatServiceImpl) ApproveCatMatch(ctx context.Context, userID uuid.UUID, matchIDStr string) (message string, err error) {
	if err = u.CatRepository.IsApprove(ctx, matchIDStr, true); err != nil {
		message = "Failed to approve match"
		logger.ErrorWithStack(err)
		return
	}
	if err = u.CatRepository.DeleteAllMatchCat(ctx, userID, matchIDStr); err != nil {
		message = "Failed to approve match"
		logger.ErrorWithStack(err)
		return
	}
	message = "Successfully Approved match"
	return
}

func (u *CatServiceImpl) RejectCatMatch(ctx context.Context, matchIDStr string) (message string, err error) {
	if err = u.CatRepository.IsApprove(ctx, matchIDStr, false); err != nil {
		message = "Failed to reject match"
		logger.ErrorWithStack(err)
		return
	}
	message = "Successfully Rejected match"
	return
}

func (u *CatServiceImpl) DeleteCatMatch(ctx context.Context, userID uuid.UUID, matchIDStr string) (message string, err error) {
	//_, err = u.CatRepository.FindMatchByID(ctx, matchIDStr)
	//if err != nil {
	//	return
	//}
	//if matchCat.IsApproved {
	//	err = failure.BadRequestFromString("Match is Approved!")
	//	return
	//}
	if err = u.CatRepository.DeleteMatch(ctx, userID, matchIDStr); err != nil {
		message = "Failed to delete match"
		logger.ErrorWithStack(err)
		return
	}
	message = "Successfully Deleted match"
	return
}
