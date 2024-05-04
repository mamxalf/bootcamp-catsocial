package repository

import (
	"catsocial/internal/domain/cat/model"
	"catsocial/internal/domain/cat/request"
	"catsocial/shared/failure"
	"catsocial/shared/logger"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"strings"
)

var catQueries = struct {
	Insertcat string
	getCat    string
	deleteCat string
}{
	Insertcat: "INSERT INTO cats %s VALUES %s RETURNING id",
	getCat:    "SELECT * FROM cats WHERE 1=1",
	deleteCat: "DELETE FROM cats WHERE id = $1",
}

func (c *CatRepositoryInfra) Insert(ctx context.Context, cat model.InsertCat) (lastInsertID uuid.UUID, err error) {
	//fieldsStr, valueListStr, args := composeInsertFieldAndParamsCat(*cat)
	//commandQuery := fmt.Sprintf(catQueries.Insertcat, fieldsStr, strings.Join(valueListStr, ","))
	//
	//stmt, err := c.DB.PG.PrepareContext(ctx, commandQuery)
	//if err != nil {
	//	logger.ErrorWithStack(err)
	//	err = failure.InternalError(err)
	//	return
	//}
	//defer stmt.Close()
	//
	//err = stmt.QueryRowContext(ctx, args...).Scan(&lastInsertID)
	query := `INSERT INTO cats (user_id, name, race, sex, age, descriptions, images_url)
              VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err = c.DB.PG.ExecContext(ctx, query, cat.UserID, cat.Name, cat.Race, cat.Sex, cat.Age, cat.Descriptions, pq.Array(cat.Images))
	if err != nil {
		logger.ErrorWithStack(err)
		err = failure.InternalError(err)
		return
	}

	return lastInsertID, nil
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

func (c *CatRepositoryInfra) FindAll(ctx context.Context, userId uuid.UUID, params request.CatQueryParams) (cats []model.Cat, err error) {
	//whereClauses := " LIMIT 10"
	//query := fmt.Sprintf(catQueries.getCat, whereClauses)
	//err = c.DB.PG.SelectContext(ctx, &cats, query)
	//if err != nil {
	//	if errors.Is(err, sql.ErrNoRows) {
	//		err = failure.NotFound("Cat not found!")
	//		return
	//	}
	//	logger.ErrorWithStack(err)
	//	err = failure.InternalError(err)
	//	return
	//}
	//return
	baseQuery := catQueries.getCat
	var args []interface{}
	var conditions []string

	conditions = append(conditions, "user_id = $1")
	args = append(args, userId)

	if params.ID != "" {
		conditions = append(conditions, "id = $2")
		args = append(args, params.ID)
	}
	if params.Race != "" {
		conditions = append(conditions, "race = $3")
		args = append(args, params.Race)
	}
	if params.Sex != "" {
		conditions = append(conditions, "sex = $4")
		args = append(args, params.Sex == "male")
	}
	if params.Search != "" {
		conditions = append(conditions, "name ILIKE $5")
		args = append(args, "%"+params.Search+"%")
	}

	if len(conditions) > 0 {
		baseQuery += " AND " + strings.Join(conditions, " AND ")
	}

	// Adding pagination
	baseQuery += fmt.Sprintf(" LIMIT $%d OFFSET $%d", len(args)+1, len(args)+2)
	args = append(args, params.Limit, params.Offset)

	// Executing the query
	err = c.DB.PG.SelectContext(ctx, &cats, baseQuery, args...)
	if err != nil {
		return nil, err
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
