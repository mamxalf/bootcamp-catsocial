package repository

import (
	"catsocial/internal/domain/cat/model"
	"context"
	"github.com/google/uuid"
)

func (c *CatRepositoryInfra) MatchRequest(ctx context.Context, requestMatch *model.InsertMatch) (match *model.Match, err error) {
	//TODO implement me
	panic("implement me")
}

func (c *CatRepositoryInfra) FindAllMatches(ctx context.Context) (matches []model.Match, err error) {
	//TODO implement me
	panic("implement me")
}

func (c *CatRepositoryInfra) Approve(ctx context.Context, matchID uuid.UUID) (err error) {
	//TODO implement me
	panic("implement me")
}

func (c *CatRepositoryInfra) Reject(ctx context.Context, matchID uuid.UUID) (err error) {
	//TODO implement me
	panic("implement me")
}

func (c *CatRepositoryInfra) DeleteMatch(ctx context.Context, matchID uuid.UUID) (err error) {
	//TODO implement me
	panic("implement me")
}
