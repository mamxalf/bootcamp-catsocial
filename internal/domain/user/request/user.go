package request

import (
	"catsocial/internal/domain/user/model"
	"catsocial/shared/utils"
	"catsocial/shared/validator"
	"github.com/rs/zerolog/log"
)

type RegisterRequest struct {
	Name     string `validate:"required,min=5,max=50" json:"name" example:"Test Name"`
	Email    string `validate:"required,email" json:"email" example:"test@example.com"`
	Password string `validate:"required,alphanum,min=5,max=15" json:"password,omitempty" example:"s3Cr3Tk3y"`
	//ConfirmPassword string `validate:"required_with=Password,eqfield=Password" json:"confirmPassword"`
}

func (r *RegisterRequest) Validate() (err error) {
	validate := validator.GetValidator()
	return validate.Struct(r)
}

func (r *RegisterRequest) ToModel() (register model.UserRegister, err error) {
	hashPassword, err := utils.HashPassword(r.Password)
	if err != nil {
		log.Err(err).Msg("[Hash Password]")
		return
	}
	register = model.UserRegister{
		Name:     r.Name,
		Email:    r.Email,
		Password: hashPassword,
	}
	return
}
