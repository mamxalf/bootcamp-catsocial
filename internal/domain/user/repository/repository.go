package repository

import (
	"catsocial/infras"
	"catsocial/internal/domain/user/model"
	"context"

	"github.com/google/uuid"
)

type UserRepository interface {
	Register(ctx context.Context, userRegister *model.UserRegister) (lastInsertId uuid.UUID, err error)
	GetUserByEmail(ctx context.Context, email string) (user *model.User, err error)
}

type UserRepositoryInfra struct {
	DB *infras.PostgresConn
}

func ProvideUserRepositoryInfra(db *infras.PostgresConn) *UserRepositoryInfra {
	r := new(UserRepositoryInfra)
	r.DB = db
	return r
}
