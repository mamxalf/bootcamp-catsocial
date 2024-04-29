package service

import (
	"catsocial/internal/domain/user/request"
	"catsocial/internal/domain/user/response"
	"catsocial/shared/failure"
	"catsocial/shared/token"
	"context"
	"github.com/rs/zerolog/log"
	"net/http"
)

func (u *UserServiceImpl) RegisterNewUser(ctx context.Context, req request.RegisterRequest) (res response.RegisterResponse, err error) {
	registerModel, err := req.ToModel()
	if err != nil {
		log.Error().Interface("params", req).Err(err).Msg("[Register - Service]")
		return
	}

	lastInsertId, err := u.UserRepository.Register(ctx, &registerModel)
	if err != nil {
		if failure.GetCode(err) != http.StatusNotFound {
			log.Error().Interface("params", req).Err(err).Msg("[RegisterNewUser - Service]")
		}
	}

	generateTokenParams := &token.GenerateTokenParams{
		AccessTokenSecret: u.Config.JwtSecret,
		AccessTokenExpiry: u.Config.JwtExpiry,
	}

	userData := &token.UserData{
		ID:       lastInsertId.String(),
		Username: req.Name,
		Email:    req.Email,
	}

	generatedToken, err := token.GenerateToken(userData, generateTokenParams)
	if err != nil {
		log.Err(err).Msg("[Login - Service] Generate Token Error")
		return
	}

	// TODO: save user sessions if needed

	res = response.RegisterResponse{
		AccessToken: generatedToken.AccessToken,
		Email:       req.Email,
		Name:        req.Name,
	}

	return
}
