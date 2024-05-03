package request

import (
	"catsocial/internal/domain/cat/model"
	"catsocial/shared/validator"
)

type InsertCatRequest struct {
	Name        string   `validate:"required" json:"name"`
	Race        string   `validate:"required" json:"race"`
	Sex         bool     `validate:"required" json:"sex"`
	AgeInMonth  int      `validate:"required" json:"ageInMonth"`
	Description string   `validate:"required" json:"description"`
	ImageUrls   []string `validate:"required" json:"imageUrls"`
}

func (r *InsertCatRequest) Validate() (err error) {
	validate := validator.GetValidator()
	return validate.Struct(r)
}

func (r *InsertCatRequest) ToModel() (cat model.InsertCat, err error) {
	cat = model.InsertCat{
		Name:         r.Name,
		Race:         r.Race,
		Sex:          r.Sex,
		Age:          r.AgeInMonth,
		Descriptions: r.Description,
		Images:       r.ImageUrls,
	}
	return
}

type UpdateCatRequest struct {
	Name        string   `validate:"required" json:"name"`
	Race        string   `validate:"required" json:"race"`
	Sex         bool     `validate:"required" json:"sex"`
	AgeInMonth  int      `validate:"required" json:"ageInMonth"`
	Description string   `validate:"required" json:"description"`
	ImageUrls   []string `validate:"required" json:"imageUrls"`
}

func (r *UpdateCatRequest) Validate() (err error) {
	validate := validator.GetValidator()
	return validate.Struct(r)
}

func (r *UpdateCatRequest) ToModel() (cat model.InsertCat, err error) {
	cat = model.InsertCat{
		Name:         r.Name,
		Race:         r.Race,
		Sex:          r.Sex,
		Age:          r.AgeInMonth,
		Descriptions: r.Description,
		Images:       r.ImageUrls,
	}
	return
}
