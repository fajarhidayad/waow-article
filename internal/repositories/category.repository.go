package repositories

import (
	"github.com/fajarhidayad/waow-article/internal/models"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(category *models.Category) error
	FindAll() (*[]models.Category, error)
	FindByID(id string) (*models.Category, error)
	Update(id string, category *models.Category) error
	Delete(id string) error
}

type categoryRepositoryImpl struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepositoryImpl{
		db: db,
	}
}

func (r *categoryRepositoryImpl) Create(category *models.Category) error {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(category).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (r *categoryRepositoryImpl) FindAll() (*[]models.Category, error) {
	var categories *[]models.Category
	err := r.db.Find(&categories).Error
	return categories, err
}

func (r *categoryRepositoryImpl) FindByID(id string) (*models.Category, error) {
	var category *models.Category
	err := r.db.Where("id = ?", id).First(&category).Error
	return category, err
}

func (r *categoryRepositoryImpl) Update(id string, data *models.Category) error {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	category, err := r.FindByID(id)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Model(&category).Updates(data).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (r *categoryRepositoryImpl) Delete(id string) error {
	err := r.db.Where("id = ?", id).Delete(&models.Category{}).Error
	return err
}
