package request

import (
	"catsocial/internal/domain/cat/model"
	"catsocial/shared/validator"
	"github.com/google/uuid"
)

type InsertCatRequest struct {
	UserID      uuid.UUID
	Name        string   `validate:"required,min=1,max=30"`                                                                                                                 // Not null, minLength 1, maxLength 30
	Race        string   `validate:"required,oneof='Persian' 'Maine Coon' 'Siamese' 'Ragdoll' 'Bengal' 'Sphynx' 'British Shorthair' 'Abyssinian' 'Scottish Fold' 'Birman'"` // Enum
	Sex         string   `validate:"required,oneof='male' 'female'"`                                                                                                        // Enum
	AgeInMonth  int      `validate:"required,min=1,max=120082"`                                                                                                             // Min 1, Max 120082
	Description string   `validate:"required,min=1,max=200"`                                                                                                                // Not null, minLength 1, maxLength 200
	ImageUrls   []string `validate:"required,min=1,dive,required,url"`                                                                                                      // Not null, minItems 1, items must be valid URLs
}

func (r *InsertCatRequest) Validate() (err error) {
	validate := validator.GetValidator()
	return validate.Struct(r)
}

func (r *InsertCatRequest) ToModel() (cat model.InsertCat, err error) {
	var sex bool
	if r.Sex == "male" {
		sex = true
	} else {
		sex = false
	}
	cat = model.InsertCat{
		UserID:       r.UserID,
		Name:         r.Name,
		Race:         r.Race,
		Sex:          sex,
		Age:          r.AgeInMonth,
		Descriptions: r.Description,
		Images:       r.ImageUrls,
	}
	return
}

type UpdateCatRequest struct {
	Name        string   `validate:"required" json:"name"`
	Race        string   `validate:"required" json:"race"`
	Sex         string   `validate:"required,oneof='male' 'female'"`
	AgeInMonth  int      `validate:"required" json:"ageInMonth"`
	Description string   `validate:"required" json:"description"`
	ImageUrls   []string `validate:"required" json:"imageUrls"`
}

func (r *UpdateCatRequest) Validate() (err error) {
	validate := validator.GetValidator()
	return validate.Struct(r)
}

func (r *UpdateCatRequest) ToModel() (cat model.Cat, err error) {
	var sex bool
	if r.Sex == "male" {
		sex = true
	} else {
		sex = false
	}
	cat = model.Cat{
		Name:         r.Name,
		Race:         r.Race,
		Sex:          sex,
		Age:          r.AgeInMonth,
		Descriptions: r.Description,
		Images:       r.ImageUrls,
	}
	return
}

type CatQueryParams struct {
	ID         string `json:"id" validate:"omitempty,uuid"`
	Limit      int    `json:"limit" validate:"omitempty,min=1"`
	Offset     int    `json:"offset" validate:"omitempty,min=0"`
	Race       string `json:"race" validate:"omitempty,oneof='Persian' 'Maine Coon' 'Siamese' 'Ragdoll' 'Bengal' 'Sphynx' 'British Shorthair' 'Abyssinian' 'Scottish Fold' 'Birman'"`
	Sex        string `json:"sex" validate:"omitempty,oneof='male' 'female'"`
	HasMatched bool   `json:"hasMatched" validate:"omitempty"`
	AgeInMonth string `json:"ageInMonth" validate:"omitempty,ageInMonthValidator"`
	Owned      bool   `json:"owned" validate:"omitempty"`
	Search     string `json:"search" validate:"omitempty,min=1"`
}

func (r *CatQueryParams) Validate() (err error) {
	validate := validator.GetValidator()
	return validate.Struct(r)
}
