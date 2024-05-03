package repository

import (
	"catsocial/internal/domain/cat/model"
	"catsocial/shared/failure"
	"catsocial/shared/logger"
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

var catQueries = struct {
	Insertcat string
	getCat    string
}{
	Insertcat: "INSERT INTO cats %s VALUES %s RETURNING id",
	getCat:    "SELECT * FROM cats %s",
}

func (c *CatRepositoryInfra) Insert(ctx context.Context, cat *model.InsertCat) (lastInsertID uuid.UUID, err error) {
	//TODO implement me
	panic("implement me")
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
	//TODO implement me
	panic("implement me")
}
