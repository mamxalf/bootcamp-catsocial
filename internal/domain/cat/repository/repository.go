package repository

import (
	"catsocial/infras"
	"catsocial/internal/domain/cat/model"
	"context"

	"github.com/google/uuid"
)

type CatRepository interface {
	// Insert Cat CRUD Interface
	Insert(ctx context.Context, cat *model.InsertCat) (lastInsertID uuid.UUID, err error)
	Find(ctx context.Context, catID uuid.UUID) (cat model.Cat, err error)
	FindAll(ctx context.Context) (cats []model.Cat, err error)
	Update(ctx context.Context, catID uuid.UUID, cat *model.Cat) (updatedID uuid.UUID, err error)
	Delete(ctx context.Context, catID uuid.UUID) (deletedID uuid.UUID, err error)

	// Match Request Interface
	MatchRequest(ctx context.Context, requestMatch *model.InsertMatch) (match *model.Match, err error)
	FindAllMatches(ctx context.Context) (matches []model.Match, err error)
	Approve(ctx context.Context, matchID uuid.UUID) (err error)
	Reject(ctx context.Context, matchID uuid.UUID) (err error)
	DeleteMatch(ctx context.Context, matchID uuid.UUID) (err error)
}

type CatRepositoryInfra struct {
	DB *infras.PostgresConn
}

func ProvideCatRepositoryInfra(db *infras.PostgresConn) *CatRepositoryInfra {
	r := new(CatRepositoryInfra)
	r.DB = db
	return r
}
