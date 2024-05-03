package request

import "catsocial/shared/validator"

type MatchRequest struct {
	MatchCatId string `validate:"required" json:"matchCatId"`
	UserCatId  string `validate:"required" json:"userCatId"`
	Message    string `validate:"required" json:"message"`
}

func (r *MatchRequest) Validate() (err error) {
	validate := validator.GetValidator()
	return validate.Struct(r)
}

type MatchApproval struct {
	MatchId string `validate:"required" json:"matchId"`
}
