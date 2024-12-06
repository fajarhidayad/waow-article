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
	Username          string    `json:"username" gorm:"unique;type:varchar(100)"`
	Email             string    `json:"email" gorm:"unique;type:varchar(100)"`
	Password          string    `json:"password" gorm:"type:varchar(100)"`
	DisplayName       string    `json:"display_name" gorm:"type:varchar(255)"`
	Bio               string    `json:"bio" gorm:"type:text"`
	ProfilePictureURL string    `json:"profile_picture_url" gorm:"type:varchar(255)"`
	RegistrationDate  time.Time `json:"registration_date"`
	Role              string    `json:"role" gorm:"type:varchar(20)"`
}

func (u *User) BeforeCreate(db *gorm.DB) error {
	u.GenerateUUID()
	u.RegistrationDate = time.Now()
	if u.Role == "" {
		u.Role = ROLE_USER
	}
	return nil
}
