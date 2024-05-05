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
	"github.com/lib/pq"
	"time"
)

var matchQueries = struct {
	insertMatch string
	updateMatch string
	deleteMatch string
	getMatch    string
}{
	insertMatch: "INSERT INTO matches %s VALUES %s RETURNING *",
	updateMatch: "UPDATE matches SET %s WHERE %s",
	deleteMatch: "DELETE FROM matches WHERE %s",
	getMatch:    "SELECT * FROM matches WHERE %s",
}

func (c *CatRepositoryInfra) MatchRequest(ctx context.Context, insertMatch *model.InsertMatch) (match *model.Match, err error) {
	// Initialize 'match' as a pointer to a new 'model.Match' struct
	match = &model.Match{}

	query := `INSERT INTO matches (issued_user_id, match_cat_id, user_cat_id, message, is_approved)
              VALUES ($1, $2, $3, $4, $5)
              RETURNING id, issued_user_id, match_cat_id, user_cat_id, message, is_approved, created_at, updated_at;`

	// Execute the query and scan the result directly into 'match'
	err = c.DB.PG.QueryRowxContext(ctx, query,
		insertMatch.IssuedUserID, insertMatch.MatchCatID, insertMatch.UserCatID, insertMatch.Message, true).StructScan(match)
	if err != nil {
		logger.ErrorWithStack(err)
		err = failure.InternalError(err)
		return nil, err
	}

	return match, nil
}

//func composeInsertFieldAndParamsMatchRequest(inserts ...model.InsertMatch) (fieldStr string, valueListStr []string, args []any) {
//	fields := []string{"issued_user_id", "match_cat_id", "user_cat_id", "message"}
//	fieldStr = fmt.Sprintf("(%s)", strings.Join(fields, ","))
//
//	args = make([]any, 0, len(inserts)*len(fields))
//
//	for i, reg := range inserts {
//		values := make([]string, len(fields))
//		args = append(args, reg.IssuedUserID, reg.MatchCatID, reg.UserCatID, reg.Message)
//		for j := range fields {
//			values[j] = fmt.Sprintf("$%d", i*len(fields)+j+1)
//		}
//		valueListStr = append(valueListStr, fmt.Sprintf("(%s)", strings.Join(values, ",")))
//	}
//
//	return fieldStr, valueListStr, args
//}

func (c *CatRepositoryInfra) FindAllMatches(ctx context.Context) (matches []model.MatchDetails, err error) {
	type FlatMatch struct {
		model.MatchDetails
		IssuedByName    string    `db:"issued_by_name"`
		IssuedByEmail   string    `db:"issued_by_email"`
		IssuedByCreated time.Time `db:"issued_by_created_at"`

		MatchCatID          uuid.UUID      `db:"match_cat_id"`
		MatchCatName        string         `db:"match_cat_name"`
		MatchCatRace        string         `db:"match_cat_race"`
		MatchCatSex         string         `db:"match_cat_sex"`
		MatchCatDescription string         `db:"match_cat_description"`
		MatchCatAge         int            `db:"match_cat_age"`
		MatchCatImages      pq.StringArray `db:"match_cat_images"`
		MatchCatHasMatched  bool           `db:"match_cat_has_matched"`
		MatchCatCreatedAt   time.Time      `db:"match_cat_created_at"`

		UserCatID          uuid.UUID      `db:"user_cat_id"`
		UserCatName        string         `db:"user_cat_name"`
		UserCatRace        string         `db:"user_cat_race"`
		UserCatSex         string         `db:"user_cat_sex"`
		UserCatDescription string         `db:"user_cat_description"`
		UserCatAge         int            `db:"user_cat_age"`
		UserCatImages      pq.StringArray `db:"user_cat_images"`
		UserCatHasMatched  bool           `db:"user_cat_has_matched"`
		UserCatCreatedAt   time.Time      `db:"user_cat_created_at"`
	}

	query := "SELECT \n    m.id,\n    u.name AS issued_by_name,\n    u.email AS issued_by_email,\n    u.created_at AS issued_by_created_at,\n    mc.id AS match_cat_id,\n    mc.name AS match_cat_name,\n    mc.race AS match_cat_race,\n    mc.sex AS match_cat_sex,\n    mc.descriptions AS match_cat_description,\n    mc.age AS match_cat_age,\n    mc.images_url AS match_cat_images,\n    mc.has_matched AS match_cat_has_matched,\n    mc.created_at AS match_cat_created_at,\n    uc.id AS user_cat_id,\n    uc.name AS user_cat_name,\n    uc.race AS user_cat_race,\n    uc.sex AS user_cat_sex,\n    uc.descriptions AS user_cat_description,\n    uc.age AS user_cat_age,\n    uc.images_url AS user_cat_images,\n    uc.has_matched AS user_cat_has_matched,\n    uc.created_at AS user_cat_created_at,\n    m.message,\n    m.created_at\nFROM matches m\nJOIN users u ON m.issued_user_id = u.id\nJOIN cats mc ON m.match_cat_id = mc.id\nJOIN cats uc ON m.user_cat_id = uc.id;\n"
	var flatMatches []FlatMatch
	// Using sqlx to query and map the results directly into a slice of Match structs
	err = c.DB.PG.SelectContext(ctx, &flatMatches, query)
	if err != nil {
		// Log and return error, avoid panic to handle the error gracefully
		logger.ErrorWithStack(err) // Assuming you have a logger set up
		return nil, err
	}

	matches = make([]model.MatchDetails, len(flatMatches))
	for i, fm := range flatMatches {
		matches[i] = model.MatchDetails{
			ID: fm.ID,
			IssuedBy: model.UserDetails{
				Name:      fm.IssuedByName,
				Email:     fm.IssuedByEmail,
				CreatedAt: fm.IssuedByCreated,
			},
			MatchCatDetail: model.CatDetails{
				ID:          fm.MatchCatID,
				Name:        fm.MatchCatName,
				Race:        fm.MatchCatRace,
				Sex:         fm.MatchCatSex,
				Description: fm.MatchCatDescription,
				AgeInMonth:  fm.MatchCatAge,
				ImageUrls:   fm.MatchCatImages,
				HasMatched:  fm.MatchCatHasMatched,
				CreatedAt:   fm.MatchCatCreatedAt,
			},
			UserCatDetail: model.CatDetails{
				ID:          fm.UserCatID,
				Name:        fm.UserCatName,
				Race:        fm.UserCatRace,
				Sex:         fm.UserCatSex,
				Description: fm.UserCatDescription,
				AgeInMonth:  fm.UserCatAge,
				ImageUrls:   fm.UserCatImages,
				HasMatched:  fm.UserCatHasMatched,
				CreatedAt:   fm.UserCatCreatedAt,
			},
			Message:   fm.Message,
			CreatedAt: fm.CreatedAt,
		}
	}

	return matches, nil
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

func (c *CatRepositoryInfra) DeleteMatch(ctx context.Context, userID uuid.UUID, matchID uuid.UUID) (err error) {
	whereClause := "id = $1 AND issued_user_id = $2"
	commandQuery := fmt.Sprintf(matchQueries.deleteMatch, whereClause)

	_, err = c.DB.PG.ExecContext(ctx, commandQuery, matchID, userID)
	if err != nil {
		logger.ErrorWithStack(err)
		return failure.InternalError(err)
	}

	return
}

func (c *CatRepositoryInfra) FindMatchByUserCatID(ctx context.Context, userCatID uuid.UUID) (cat model.Match, err error) {
	whereClause := "user_cat_id = $1"
	commandQuery := fmt.Sprintf(matchQueries.deleteMatch, whereClause)
	err = c.DB.PG.GetContext(ctx, &cat, commandQuery, userCatID)
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

func (c *CatRepositoryInfra) FindMatchByMatchCatID(ctx context.Context, matchCatID uuid.UUID) (cat model.Match, err error) {
	whereClause := "match_cat_id = $1"
	commandQuery := fmt.Sprintf(matchQueries.deleteMatch, whereClause)
	err = c.DB.PG.GetContext(ctx, &cat, commandQuery, matchCatID)
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
