package model

import (
	"github.com/google/uuid"
	"time"
)

type UserRegister struct {
	Name     string `db:"name"`
	Email    string `db:"email"`
	Password string `db:"password"`
}

type User struct {
	ID        uuid.UUID `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
