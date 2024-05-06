package repository

import (
	"catsocial/internal/domain/user/model"
	"catsocial/shared/failure"
	"catsocial/shared/logger"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
)

var userQueries = struct {
	getUser string
}{
	getUser: "SELECT * FROM users %s",
}

func (repo *UserRepositoryInfra) GetUserByEmail(ctx context.Context, email string) (user model.User, err error) {
	whereClauses := " WHERE email = $1 LIMIT 1"
	query := fmt.Sprintf(userQueries.getUser, whereClauses)
	err = repo.DB.PG.GetContext(ctx, &user, query, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound("User not found!")
			return
		}
		logger.ErrorWithStack(err)
		err = failure.InternalError(err)
		return
	}
	return
}

func (repo *UserRepositoryInfra) GetUserByID(ctx context.Context, id uuid.UUID) (user model.User, err error) {
	whereClauses := " WHERE id = $1 LIMIT 1"
	query := fmt.Sprintf(userQueries.getUser, whereClauses)
	err = repo.DB.PG.GetContext(ctx, &user, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound("User not found!")
			return
		}
		logger.ErrorWithStack(err)
		err = failure.InternalError(err)
		return
	}
	return
}
