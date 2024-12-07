package dtos

type CreateUserDTO struct {
	Username          string `json:"username" binding:"required"`
	Email             string `json:"email" binding:"required,email"`
	Password          string `json:"password" binding:"required"`
	DisplayName       string `json:"display_name" binding:"required"`
	Bio               string `json:"bio"`
	ProfilePictureURL string `json:"profile_picture_url"`
	Role              string `json:"role" binding:"required"`
}

type UpdateUserDTO struct {
	Username          string `json:"username" binding:"required"`
	Email             string `json:"email" binding:"required,email"`
	Password          string `json:"password"`
	DisplayName       string `json:"display_name" binding:"required"`
	Bio               string `json:"bio"`
	ProfilePictureURL string `json:"profile_picture_url"`
	Role              string `json:"role" binding:"required"`
}
