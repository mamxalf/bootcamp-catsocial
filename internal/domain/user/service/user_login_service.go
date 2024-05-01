package service

import (
	"catsocial/internal/domain/user/request"
	"catsocial/internal/domain/user/response"
	"catsocial/shared/token"
	"catsocial/shared/utils"
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
)

func (u *UserServiceImpl) LoginUser(ctx context.Context, req request.LoginRequest) (res response.AuthResponse, err error) {
	user, err := u.UserRepository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		log.Err(err).Msg("[Login - Service]")
		return
	}

	err = utils.CheckPasswordHash(req.Password, user.Password)
	if err != nil {
		log.Err(err).Msg("[Login - Service] Wrong Password")
		err = fmt.Errorf("invalid Password")
		return
	}

	generateTokenParams := &token.GenerateTokenParams{
		AccessTokenSecret: u.Config.JwtSecret,
		AccessTokenExpiry: u.Config.JwtExpiry,
	}

	userData := &token.UserData{
		ID:    user.ID.String(),
		Name:  user.Name,
		Email: user.Email,
	}
	generatedToken, err := token.GenerateToken(userData, generateTokenParams)
	if err != nil {
		log.Err(err).Msg("[Login - Service] Generate Token Error")
		return
	}

	res = response.AuthResponse{
		AccessToken: generatedToken.AccessToken,
		Email:       user.Email,
		Name:        user.Name,
	}

	return
}
