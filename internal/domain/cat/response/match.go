package response

import "time"

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
	ID             string      `json:"id"`
	IssuedBy       MatchIssued `json:"issuedBy"`
	MatchCatDetail MatchDetail `json:"matchCatDetail"`
	UserCatDetail  MatchDetail `json:"userCatDetail"`
	Message        string      `json:"message"`
	CreatedAt      time.Time   `json:"createdAt"`
}
