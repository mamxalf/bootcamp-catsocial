package repository

import (
	"catsocial/internal/domain/cat/model"
	"catsocial/shared/failure"
	"catsocial/shared/logger"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

var catQueries = struct {
	Insertcat string
	getCat    string
	deleteCat string
}{
	Insertcat: "INSERT INTO cats %s VALUES %s RETURNING id",
	getCat:    "SELECT * FROM cats %s",
	deleteCat: "DELETE FROM cats WHERE id = $1",
}

func (c *CatRepositoryInfra) Insert(ctx context.Context, cat *model.InsertCat) (lastInsertID uuid.UUID, err error) {
	fieldsStr, valueListStr, args := composeInsertFieldAndParamsCat(*cat)
	commandQuery := fmt.Sprintf(catQueries.Insertcat, fieldsStr, strings.Join(valueListStr, ","))

	stmt, err := c.DB.PG.PrepareContext(ctx, commandQuery)
	if err != nil {
		logger.ErrorWithStack(err)
		err = failure.InternalError(err)
		return
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx, args...).Scan(&lastInsertID)
	if err != nil {
		logger.ErrorWithStack(err)
		err = failure.InternalError(err)
		return
	}
	return lastInsertID, nil
}

func composeInsertFieldAndParamsCat(cats ...model.InsertCat) (fieldStr string, valueListStr []string, args []interface{}) {
	fields := []string{"user_id", "name", "race", "sex", "age", "descriptions", "images_url"}
	fieldStr = fmt.Sprintf("(%s)", strings.Join(fields, ","))

	args = make([]interface{}, 0, len(cats)*len(fields))

	for _, cat := range cats {
		values := make([]string, len(fields))
		args = append(args, cat.UserID, cat.Name, cat.Race, cat.Sex, cat.Age, cat.Descriptions, cat.Images)
		for j := range fields {
			values[j] = fmt.Sprintf("$%d", len(args)-len(fields)+j+1)
		}
		valueListStr = append(valueListStr, fmt.Sprintf("(%s)", strings.Join(values, ",")))
	}

	return fieldStr, valueListStr, args
}

func (c *CatRepositoryInfra) Find(ctx context.Context, catID uuid.UUID) (cat model.Cat, err error) {
	whereClauses := " WHERE id = $1 LIMIT 1"
	query := fmt.Sprintf(catQueries.getCat, whereClauses)
	err = c.DB.PG.GetContext(ctx, &cat, query, catID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound("Cat not found!")
			return
		}
		logger.ErrorWithStack(err)
		err = failure.InternalError(err)
		return
	}
	return
}

func (c *CatRepositoryInfra) FindAll(ctx context.Context) (cats []model.Cat, err error) {
	whereClauses := " LIMIT 10"
	query := fmt.Sprintf(catQueries.getCat, whereClauses)
	err = c.DB.PG.SelectContext(ctx, &cats, query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound("Cat not found!")
			return
		}
		logger.ErrorWithStack(err)
		err = failure.InternalError(err)
		return
	}
	return
}

func (c *CatRepositoryInfra) Update(ctx context.Context, catID uuid.UUID, cat *model.Cat) (updatedID uuid.UUID, err error) {
	//TODO implement me
	panic("implement me")
}

func (c *CatRepositoryInfra) Delete(ctx context.Context, catID uuid.UUID) (deletedID uuid.UUID, err error) {
	result, err := c.DB.PG.ExecContext(ctx, catQueries.deleteCat, catID)
	if err != nil {
		logger.ErrorWithStack(err)
		err = failure.InternalError(err)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.ErrorWithStack(err)
		err = failure.InternalError(err)
		return
	}

	if rowsAffected == 0 {
		err = failure.NotFound("Cat not found!")
		return
	}

	deletedID = catID
	return
}
