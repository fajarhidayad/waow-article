package models

import (
	"github.com/fajarhidayad/waow-article/pkg/common"
	"gorm.io/gorm"
)

type Category struct {
	common.ModelsWithID
	Name        string `json:"name" gorm:"type:varchar(255);not null;unique"`
	Description string `json:"description" gorm:"type:text"`
	Slug        string `json:"slug" gorm:"type:varchar(255);unique"`
}

func (c *Category) BeforeCreate(db *gorm.DB) error {
	c.GenerateUUID()
	return nil
}
