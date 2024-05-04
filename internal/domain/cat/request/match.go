package request

import (
	"catsocial/internal/domain/cat/model"
	"catsocial/shared/logger"
	"catsocial/shared/validator"
	"github.com/google/uuid"
)

type MatchRequest struct {
	IssuedUserID uuid.UUID `json:"-"`
	MatchCatID   string    `validate:"required" json:"matchCatId"`
	UserCatID    string    `validate:"required" json:"userCatId"`
	Message      string    `validate:"required" json:"message"`
}

func (r *MatchRequest) Validate() (err error) {
	validate := validator.GetValidator()
	return validate.Struct(r)
}

func (r *MatchRequest) ToModel() (insert model.InsertMatch, err error) {
	matchCatID, err := uuid.Parse(r.MatchCatID)
	if err != nil {
		logger.ErrorWithStack(err)
		return
	}
	userCatID, err := uuid.Parse(r.UserCatID)
	if err != nil {
		logger.ErrorWithStack(err)
		return
	}
	insert = model.InsertMatch{
		MatchCatID: matchCatID,
		UserCatID:  userCatID,
		Message:    r.Message,
	}
	return
}

type MatchApproval struct {
	MatchId string `validate:"required" json:"matchId"`
}
