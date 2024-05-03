package response

type CatResponse struct {
	Name        string   `json:"name"`
	Race        string   `json:"race"`
	Sex         bool     `json:"sex"`
	AgeInMonth  int      `json:"ageInMonth"`
	Description string   `json:"description"`
	ImageUrls   []string `json:"imageUrls"`
}
