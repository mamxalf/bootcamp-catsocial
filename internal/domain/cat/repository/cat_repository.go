package repository

import (
	"catsocial/internal/domain/cat/model"
	// "catsocial/shared/failure"
	// "catsocial/shared/logger"
	"context"
	// "fmt"
	// "strings"

	"github.com/google/uuid"
	// "github.com/lib/pq"
	// "github.com/rs/zerolog/log"
)

// var catQueries = struct {
// 	cat string
// }{
// 	cat: "INSERT INTO cats %s VALUES %s RETURNING id",
// }

func (repo *CatRepositoryInfra) InsertCat(ctx context.Context, cat *model.InsertCat) (lastInsertID uuid.UUID, err error) {
	return
}

func (repo *CatRepositoryInfra) Find(ctx context.Context, catID uuid.UUID) (cat model.Cat, err error) {
	return
}

func (repo *CatRepositoryInfra) FindAll(ctx context.Context) (cats []model.Cat, err error) {
	return
}

func (repo *CatRepositoryInfra) Update(ctx context.Context, catID uuid.UUID, cat *model.Cat) (updatedID uuid.UUID, err error) {
	return
}

func (repo *CatRepositoryInfra) Delete(ctx context.Context, catID uuid.UUID) (deletedID uuid.UUID, err error) {
	return
}
