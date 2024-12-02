package dtos

import "time"

type RegisterDto struct {
	Username          string    `json:"username" binding:"required"`
	Email             string    `json:"email" binding:"required,email"`
	Password          string    `json:"password" binding:"required"`
	DisplayName       string    `json:"display_name" binding:"required"`
	Bio               string    `json:"bio"`
	ProfilePictureURL string    `json:"profile_picture_url"`
	RegistrationDate  time.Time `json:"registration_date"`
	Role              string    `json:"role"`
}

type LoginDto struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
