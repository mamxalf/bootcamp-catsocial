package model

import (
	"github.com/google/uuid"
	"time"
)

type Match struct {
	ID           int       `db:"id"`
	IssuedUserID uuid.UUID `db:"issued_user_id"`
	MatchCatID   uuid.UUID `db:"match_cat_id"`
	UserCatID    uuid.UUID `db:"user_cat_id"`
	Message      string    `db:"message"`
	IsApproved   bool      `db:"is_approved"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

type InsertMatch struct {
	IssuedUserID uuid.UUID `db:"issued_user_id"`
	MatchCatID   uuid.UUID `db:"match_cat_id"`
	UserCatID    uuid.UUID `db:"user_cat_id"`
	Message      string    `db:"message"`
}
