package model

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
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

type MatchDetails struct {
	ID             int         `db:"id" json:"id"`
	IssuedBy       UserDetails `db:"issued_by"`
	MatchCatDetail CatDetails  `db:"match_cat_detail"`
	UserCatDetail  CatDetails  `db:"user_cat_detail"`
	Message        string      `db:"message" json:"message"`
	CreatedAt      time.Time   `db:"created_at" json:"createdAt"` // ISO 8601 format
}

type UserDetails struct {
	Name      string    `db:"issued_by_name"`
	Email     string    `db:"issued_by_email"`
	CreatedAt time.Time `db:"issued_by_created_at"` // ISO 8601 format
}

type CatDetails struct {
	ID          uuid.UUID      `db:"id"`
	Name        string         `db:"name"`
	Race        string         `db:"race"`
	Sex         string         `db:"sex"`
	Description string         `db:"description"`
	AgeInMonth  int            `db:"age"`
	ImageUrls   pq.StringArray `db:"images_url"` // Assuming this is stored appropriately
	HasMatched  bool           `db:"has_matched"`
	CreatedAt   time.Time      `db:"created_at"` // ISO 8601 format
}

type InsertMatch struct {
	IssuedUserID uuid.UUID `db:"issued_user_id"`
	MatchCatID   uuid.UUID `db:"match_cat_id"`
	UserCatID    uuid.UUID `db:"user_cat_id"`
	Message      string    `db:"message"`
}
