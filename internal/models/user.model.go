package models

import (
	"time"

	"github.com/fajarhidayad/waow-article/pkg/common"
	"gorm.io/gorm"
)

const (
	ROLE_ADMIN = "ADMIN"
	ROLE_USER  = "USER"
)

type User struct {
	common.ModelsWithID
	Username          string    `json:"username" gorm:"unique"`
	Email             string    `json:"email" gorm:"unique"`
	Password          string    `json:"password"`
	DisplayName       string    `json:"display_name"`
	Bio               string    `json:"bio" gorm:"type:text"`
	ProfilePictureURL string    `json:"profile_picture_url"`
	RegistrationDate  time.Time `json:"registration_date"`
	Role              string    `json:"role"`
}

func (u *User) BeforeCreate(db *gorm.DB) error {
	u.GenerateUUID()
	u.RegistrationDate = time.Now()
	if u.Role == "" {
		u.Role = ROLE_USER
	}
	return nil
}
