package response

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	"time"
)

type MatchIssued struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

type MatchDetail struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Race        string    `json:"race"`
	Sex         string    `json:"sex"`
	Description string    `json:"description"`
	AgeInMonth  int       `json:"ageInMonth"`
	ImageURLs   []string  `json:"imageUrls"`
	HasMatched  bool      `json:"hasMatched"`
	CreatedAt   time.Time `json:"createdAt"`
}

type MatchList struct {
	ID             int         `json:"id"`
	IssuedBy       UserDetails `json:"issued_by"`
	MatchCatDetail CatDetails  `json:"match_cat_detail"`
	UserCatDetail  CatDetails  `json:"user_cat_detail"`
	Message        string      `json:"message"`
	CreatedAt      time.Time   `json:"created_at"` // ISO 8601 format
}

type UserDetails struct {
	Name      string    `json:"issued_by_name"`
	Email     string    `json:"issued_by_email"`
	CreatedAt time.Time `json:"issued_by_created_at"` // ISO 8601 format
}

type CatDetails struct {
	ID          uuid.UUID      `json:"id"`
	Name        string         `json:"name"`
	Race        string         `json:"race"`
	Sex         string         `json:"sex"`
	Description string         `json:"description"`
	AgeInMonth  int            `json:"age"`
	ImageUrls   pq.StringArray `json:"images_url"` // Assuming this is stored appropriately
	HasMatched  bool           `json:"has_matched"`
	CreatedAt   time.Time      `json:"created_at"` // ISO 8601 format
}

type InsertMatch struct {
	IssuedUserID uuid.UUID `json:"issued_user_id"`
	MatchCatID   uuid.UUID `json:"match_cat_id"`
	UserCatID    uuid.UUID `json:"user_cat_id"`
	Message      string    `json:"message"`
}
