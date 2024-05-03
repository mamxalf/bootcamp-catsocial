package repository

import (
	"catsocial/internal/domain/cat/model"
	"context"
	"github.com/google/uuid"
)

// var catQueries = struct {
// 	cat string
// }{
// 	cat: "INSERT INTO cats %s VALUES %s RETURNING id",
// }

func (c *CatRepositoryInfra) Insert(ctx context.Context, cat *model.InsertCat) (lastInsertID uuid.UUID, err error) {
	//TODO implement me
	panic("implement me")
}

func (c *CatRepositoryInfra) Find(ctx context.Context, catID uuid.UUID) (cat model.Cat, err error) {
	//TODO implement me
	panic("implement me")
}

func (c *CatRepositoryInfra) FindAll(ctx context.Context) (cats []model.Cat, err error) {
	//TODO implement me
	panic("implement me")
}

func (c *CatRepositoryInfra) Update(ctx context.Context, catID uuid.UUID, cat *model.Cat) (updatedID uuid.UUID, err error) {
	//TODO implement me
	panic("implement me")
}

func (c *CatRepositoryInfra) Delete(ctx context.Context, catID uuid.UUID) (deletedID uuid.UUID, err error) {
	//TODO implement me
	panic("implement me")
}
