package repository

import (
	"catsocial/internal/domain/user/model"
	"catsocial/shared/failure"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"strings"
)

var registerQueries = struct {
	registerUser string
}{
	registerUser: "INSERT INTO users %s VALUES %s",
}

func (repo *UserRepositoryInfra) Register(ctx context.Context, userRegister *model.UserRegister) (lastInsertId uuid.UUID, err error) {
	fieldsStr, valueListStr, args := composeInsertFieldAndParamsUser(*userRegister)
	commandQuery := fmt.Sprintf(registerQueries.registerUser, fieldsStr, strings.Join(valueListStr, ","))

	stmt, err := repo.DB.PG.PrepareContext(ctx, commandQuery)
	if err != nil {
		log.Error().Err(err).Msg("[UserRepository - exec] failed prepare query")
		err = failure.InternalError(err)
		return
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx, args...).Scan(&lastInsertId)
	if err != nil {
		log.Error().Err(err).Msg("[UserRepository - execInsert] failed to execute query and scan id")
		err = failure.InternalError(err)
		return
	}
	return
}

func composeInsertFieldAndParamsUser(registers ...model.UserRegister) (fieldStr string, valueListStr []string, args []any) {
	fields := []string{"name", "email", "password"}
	fieldStr = fmt.Sprintf("(%s)", strings.Join(fields, ","))

	args = make([]any, 0, len(registers)*len(fields))

	for i, reg := range registers {
		values := make([]string, len(fields))
		args = append(args, reg.Name, reg.Email, reg.Password)
		for j := range fields {
			values[j] = fmt.Sprintf("$%d", i*len(fields)+j+1)
		}
		valueListStr = append(valueListStr, fmt.Sprintf("(%s)", strings.Join(values, ",")))
	}

	return fieldStr, valueListStr, args
}
