package repository

import (
	"catsocial/internal/domain/cat/model"
	"catsocial/shared/failure"
	"catsocial/shared/logger"
	"context"
	"fmt"
	"github.com/google/uuid"
	"strings"
)

var matchQueries = struct {
	insertMatch string
	updateMatch string
	deleteMatch string
}{
	insertMatch: "INSERT INTO matches %s VALUES %s RETURNING id",
	updateMatch: "UPDATE matches SET %s WHERE %s",
	deleteMatch: "DELETE FROM matches WHERE %s",
}

func (c *CatRepositoryInfra) MatchRequest(ctx context.Context, insertMatch *model.InsertMatch) (id uuid.UUID, err error) {
	fieldsStr, valueListStr, args := composeInsertFieldAndParamsMatchRequest(*insertMatch)
	commandQuery := fmt.Sprintf(matchQueries.insertMatch, fieldsStr, strings.Join(valueListStr, ","))

	stmt, err := c.DB.PG.PrepareContext(ctx, commandQuery)
	if err != nil {
		logger.ErrorWithStack(err)
		err = failure.InternalError(err)
		return
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx, args...).Scan(&id)
	if err != nil {
		logger.ErrorWithStack(err)
		err = failure.InternalError(err)
		return
	}

	return
}

func composeInsertFieldAndParamsMatchRequest(inserts ...model.InsertMatch) (fieldStr string, valueListStr []string, args []any) {
	fields := []string{"issued_user_id", "match_cat_id", "message"}
	fieldStr = fmt.Sprintf("(%s)", strings.Join(fields, ","))

	args = make([]any, 0, len(inserts)*len(fields))

	for i, reg := range inserts {
		values := make([]string, len(fields))
		args = append(args, reg.IssuedUserID, reg.MatchCatID, reg.Message)
		for j := range fields {
			values[j] = fmt.Sprintf("$%d", i*len(fields)+j+1)
		}
		valueListStr = append(valueListStr, fmt.Sprintf("(%s)", strings.Join(values, ",")))
	}

	return fieldStr, valueListStr, args
}

func (c *CatRepositoryInfra) FindAllMatches(ctx context.Context) (matches []model.Match, err error) {
	//TODO implement me
	panic("implement me")
}

func (c *CatRepositoryInfra) IsApprove(ctx context.Context, matchID uuid.UUID, isApprove bool) (err error) {
	updateClause := "is_approved = $1"
	whereClause := "id = $2"
	commandQuery := fmt.Sprintf(matchQueries.updateMatch, updateClause, whereClause)

	_, err = c.DB.PG.ExecContext(ctx, commandQuery, isApprove, matchID)
	if err != nil {
		logger.ErrorWithStack(err)
		return failure.InternalError(err)
	}

	return
}

func (c *CatRepositoryInfra) DeleteMatch(ctx context.Context, matchID uuid.UUID) (err error) {
	whereClause := "id = $1"
	commandQuery := fmt.Sprintf(matchQueries.deleteMatch, whereClause)

	_, err = c.DB.PG.ExecContext(ctx, commandQuery, matchID)
	if err != nil {
		logger.ErrorWithStack(err)
		return failure.InternalError(err)
	}

	return
}
