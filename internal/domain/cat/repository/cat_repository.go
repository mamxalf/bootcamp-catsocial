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
	"strconv"
	"strings"
)

var catQueries = struct {
	Insertcat string
	getCat    string
	selectCat string
	deleteCat string
}{
	Insertcat: "INSERT INTO cats %s VALUES %s RETURNING id",
	getCat:    "SELECT * FROM cats WHERE 1=1",
	selectCat: "SELECT * FROM cats %s",
	deleteCat: "DELETE FROM cats WHERE id = $1",
}

func (c *CatRepositoryInfra) Insert(ctx context.Context, cat model.InsertCat) (newCat *model.Cat, err error) {
	query := `INSERT INTO cats (user_id, name, race, sex, age, descriptions, images_url)
              VALUES ($1, $2, $3, $4, $5, $6, $7)
              RETURNING id, user_id, name, race, sex, age, descriptions, images_url;`

	newCat = &model.Cat{}
	err = c.DB.PG.QueryRowxContext(ctx, query, cat.UserID, cat.Name, cat.Race, cat.Sex, cat.Age, cat.Descriptions, pq.Array(cat.Images)).StructScan(newCat)
	if err != nil {
		logger.ErrorWithStack(err)
		err = failure.InternalError(err)
		return nil, err
	}

	return newCat, nil
}

func (c *CatRepositoryInfra) Find(ctx context.Context, userID uuid.UUID, catID uuid.UUID) (cat model.Cat, err error) {
	whereClauses := " WHERE id = $1 AND user_id = $2 LIMIT 1"
	query := fmt.Sprintf(catQueries.selectCat, whereClauses)
	err = c.DB.PG.GetContext(ctx, &cat, query, catID, userID)
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
	baseQuery := catQueries.getCat // Assuming this starts with "SELECT ... FROM cats WHERE 1=1"
	var args []interface{}
	var conditions []string

	// Always include the user_id in the query
	if params.Owned {
		conditions = append(conditions, fmt.Sprintf("user_id = $%d", len(args)+1))
		args = append(args, userId)
	}

	// Check if ID is specified
	if params.ID != "" {
		conditions = append(conditions, fmt.Sprintf("id = $%d", len(args)+1))
		args = append(args, params.ID)
	}

	// Check if Race is specified
	if params.Race != "" {
		conditions = append(conditions, fmt.Sprintf("race = $%d", len(args)+1))
		args = append(args, params.Race)
	}

	// Check if Sex is specified
	if params.Sex != "" {
		conditions = append(conditions, fmt.Sprintf("sex = $%d", len(args)+1))
		args = append(args, params.Sex == "male")
	}

	if params.AgeInMonth != "" {
		operator, age, err := parseAgeFilter(params.AgeInMonth)
		if err != nil {
			return nil, err
		}
		conditions = append(conditions, fmt.Sprintf("age %s $%d", operator, len(args)+1))
		args = append(args, age)
	}

	// Check if Search term is specified
	if params.Search != "" {
		conditions = append(conditions, fmt.Sprintf("name ILIKE $%d", len(args)+1))
		args = append(args, "%"+params.Search+"%")
	}

	// Adding the conditions to the base query
	if len(conditions) > 0 {
		baseQuery += " AND " + strings.Join(conditions, " AND ")
	}

	// Adding pagination with proper indexing
	baseQuery += fmt.Sprintf(" LIMIT $%d OFFSET $%d", len(args)+1, len(args)+2)
	args = append(args, params.Limit, params.Offset)

	// Executing the query
	err = c.DB.PG.SelectContext(ctx, &cats, baseQuery, args...)
	if err != nil {
		return nil, err
	}
	return cats, nil
}

func parseAgeFilter(input string) (operator string, value int, err error) {
	// Remove any spaces and check if the string is empty
	input = strings.TrimSpace(input)
	if input == "" {
		return "", 0, errors.New("input cannot be empty")
	}

	// Determine the operator and the subsequent value
	if strings.HasPrefix(input, ">=") || strings.HasPrefix(input, "=>") {
		value, err = strconv.Atoi(input[2:])
		operator = ">="
	} else if strings.HasPrefix(input, "<=") || strings.HasPrefix(input, "=<") {
		value, err = strconv.Atoi(input[2:])
		operator = "<="
	} else if strings.HasPrefix(input, "<") {
		value, err = strconv.Atoi(input[2:])
		operator = "<"
	} else if strings.HasPrefix(input, ">") {
		value, err = strconv.Atoi(input[2:])
		operator = ">"
	} else if strings.HasPrefix(input, "=") {
		value, err = strconv.Atoi(input[1:])
		operator = "="
	} else {
		return "", 0, errors.New("invalid format or operator")
	}

	if err != nil {
		return "", 0, fmt.Errorf("failed to parse number from input: %v", err)
	}

	return operator, value, nil
}

func (c *CatRepositoryInfra) Update(ctx context.Context, catID uuid.UUID, cat model.Cat) (updatedCat *model.Cat, err error) {
	var setParts []string
	var args []interface{}
	argID := 1

	// Dynamically build the SQL query based on provided fields
	if cat.Name != "" {
		setParts = append(setParts, "name = $"+strconv.Itoa(argID))
		args = append(args, cat.Name)
		argID++
	}
	if cat.Race != "" {
		setParts = append(setParts, "race = $"+strconv.Itoa(argID))
		args = append(args, cat.Race)
		argID++
	}
	if cat.Sex { // Assuming `false` is not a valid update, adjust as needed
		setParts = append(setParts, "sex = $"+strconv.Itoa(argID))
		args = append(args, cat.Sex)
		argID++
	}
	if cat.Age != 0 {
		setParts = append(setParts, "age = $"+strconv.Itoa(argID))
		args = append(args, cat.Age)
		argID++
	}
	if cat.Descriptions != "" {
		setParts = append(setParts, "descriptions = $"+strconv.Itoa(argID))
		args = append(args, cat.Descriptions)
		argID++
	}
	if cat.Images != nil {
		setParts = append(setParts, "images_url = $"+strconv.Itoa(argID))
		args = append(args, pq.Array(cat.Images))
		argID++
	}

	if len(setParts) == 0 {
		return // No updates to perform
	}

	// Construct the full SQL statement
	updateQuery := "UPDATE cats SET " + strings.Join(setParts, ", ") + " WHERE id = $" + strconv.Itoa(argID) + " RETURNING *"
	args = append(args, catID)

	// Execute the query
	updatedCat = &model.Cat{}
	err = c.DB.PG.GetContext(ctx, updatedCat, updateQuery, args...)
	if err != nil {
		logger.ErrorWithStack(err)
		err = failure.InternalError(err)
		return nil, err
	}
	return
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
