package repository

import (
	"catsocial/infras"
	"catsocial/internal/domain/cat/model"
	"catsocial/internal/domain/cat/request"
	"context"

	"github.com/google/uuid"
)

type CatRepository interface {
	// Insert Cat CRUD Interface
	Insert(ctx context.Context, cat model.InsertCat) (newCat *model.Cat, err error)
	Find(ctx context.Context, catID uuid.UUID) (cat model.Cat, err error)
	FindAll(ctx context.Context, userId uuid.UUID, params request.CatQueryParams) (cats []model.Cat, err error)
	Update(ctx context.Context, catID uuid.UUID, cat model.Cat) (updatedCat *model.Cat, err error)
	Delete(ctx context.Context, catID uuid.UUID) (deletedID uuid.UUID, err error)
	Approve(ctx context.Context, catID uuid.UUID) (err error)

	// Match Request Interface
	MatchRequest(ctx context.Context, insertMatch *model.InsertMatch) (match *model.Match, err error)
	FindAllMatches(ctx context.Context) (matches []model.MatchDetails, err error)
	IsApprove(ctx context.Context, matchID string, isApprove bool) (err error)
	DeleteMatch(ctx context.Context, userID uuid.UUID, matchID string) (err error)
	FindMatchByUserCatID(ctx context.Context, userCatID uuid.UUID) (cat model.Match, err error)
	FindMatchByMatchCatID(ctx context.Context, matchCatID uuid.UUID) (cat model.Match, err error)
	FindMatchByID(ctx context.Context, ID string) (cat model.Match, err error)
	DeleteAllMatchCat(ctx context.Context, userID uuid.UUID, matchID string) (err error)
}

type CatRepositoryInfra struct {
	DB *infras.PostgresConn
}

func ProvideCatRepositoryInfra(db *infras.PostgresConn) *CatRepositoryInfra {
	r := new(CatRepositoryInfra)
	r.DB = db
	return r
}
