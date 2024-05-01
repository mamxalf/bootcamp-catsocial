package service

import (
	"catsocial/internal/domain/user/request"
	"catsocial/internal/domain/user/response"
	"catsocial/shared/failure"
	"catsocial/shared/token"
	"context"
	"net/http"

	"github.com/rs/zerolog/log"
)

func (u *UserServiceImpl) RegisterNewUser(ctx context.Context, req request.RegisterRequest) (res response.AuthResponse, err error) {
	registerModel, err := req.ToModel()
	if err != nil {
		log.Error().Interface("params", req).Err(err).Msg("[Register - Service]")
		return
	}

	lastInsertId, err := u.UserRepository.Register(ctx, &registerModel)
	if err != nil {
		if failure.GetCode(err) == http.StatusConflict {
			log.Error().Interface("params", req).Err(err).Msg("[RegisterNewUser - Service] Email should unique")
			return
		}
		log.Error().Interface("params", req).Err(err).Msg("[RegisterNewUser - Service] Internal Error")
		return
	}

	generateTokenParams := &token.GenerateTokenParams{
		AccessTokenSecret: u.Config.JwtSecret,
		AccessTokenExpiry: u.Config.JwtExpiry,
	}

	userData := &token.UserData{
		ID:    lastInsertId.String(),
		Name:  req.Name,
		Email: req.Email,
	}

	generatedToken, err := token.GenerateToken(userData, generateTokenParams)
	if err != nil {
		log.Err(err).Msg("[Login - Service] Generate Token Error")
		return
	}

	// TODO: save user sessions if needed

	res = response.AuthResponse{
		AccessToken: generatedToken.AccessToken,
		Email:       req.Email,
		Name:        req.Name,
	}

	return
}
