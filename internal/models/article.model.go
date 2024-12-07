package models

import (
	"github.com/fajarhidayad/waow-article/pkg/common"
	"gorm.io/gorm"
	"strings"
)

type Article struct {
	common.ModelsWithID
	Title     string `json:"title" gorm:"type:varchar(255);not null;"`
	Content   string `json:"content" gorm:"type:text;not null;"`
	ViewCount int    `json:"view_count" gorm:"not null"`
	Slug      string `json:"slug" gorm:"type:varchar(255);not null;unique;"`

	CategoryID string   `json:"category_id" gorm:"type:uuid;not null;"`
	Category   Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID     string   `json:"author_id" gorm:"type:uuid;not null;"`
	User       User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (a *Article) BeforeCreate(db *gorm.DB) error {
	a.GenerateUUID()
	a.ViewCount = 0
	a.Slug = strings.ToLower(a.Title)
	a.Slug = strings.ReplaceAll(a.Slug, " ", "-")
	return nil
}
