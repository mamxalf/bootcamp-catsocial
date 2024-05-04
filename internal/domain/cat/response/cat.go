package response

import "github.com/google/uuid"

type CatResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Race        string    `json:"race"`
	Sex         bool      `json:"sex"`
	AgeInMonth  int       `json:"ageInMonth"`
	Description string    `json:"description"`
	ImageUrls   []string  `json:"imageUrls"`
}
