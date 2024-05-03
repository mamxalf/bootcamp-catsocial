package model

import (
	"time"

	"github.com/google/uuid"
)

type Cat struct {
	ID           uuid.UUID `db:"id"`
	UserID       uuid.UUID `db:"user_id"`
	Name         string    `db:"name"`
	Race         string    `db:"race"`
	Sex          bool      `db:"sex"`
	Age          int       `db:"age"`
	Descriptions string    `db:"descriptions"`
	Images       []string  `db:"images_url"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}
type InsertCat struct {
	UserID       uuid.UUID `db:"user_id"`
	Name         string    `db:"name"`
	Race         string    `db:"race"`
	Sex          bool      `db:"sex"`
	Age          int       `db:"age"`
	Descriptions string    `db:"descriptions"`
	Images       []string  `db:"images_url"`
}
