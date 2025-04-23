package models

import "time"

type User struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name" gorm:"size:255;not null" validate:"required"`
	Email     string    `json:"email" gorm:"size:255;not null;unique" validate:"required,email"`
	Password  string    `json:"-" gorm:"size:255;not null" validate:"required,min=6"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RegisterRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}
