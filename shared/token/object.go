package token

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type UserData struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type GenerateTokenParams struct {
	AccessTokenSecret string
	AccessTokenExpiry time.Duration
}

type Token struct {
	AccessToken string `json:"accessToken"`
}

// JWTToken is
type JWTToken struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	OwnerID  string `json:"ownerID"`
	jwt.StandardClaims
}

// JWTVerifyEmail is
type JWTVerifyEmail struct {
	jwt.StandardClaims
	Email string `json:"email"`
}
