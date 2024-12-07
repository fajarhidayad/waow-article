package repositories

import (
	"github.com/fajarhidayad/waow-article/internal/models"
	"gorm.io/gorm"
)

type ArticleRepository interface {
	Create(article *models.Article) error
	FindAll() (*[]models.Article, error)
	FindByID(id string) (*models.Article, error)
	Update(id string, username string, article *models.Article) error
	Delete(id string) error
}

type articleRepositoryImpl struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &articleRepositoryImpl{
		db: db,
	}
}

func (r *articleRepositoryImpl) Create(article *models.Article) error {
	return r.db.Create(article).Error
}
func (r *articleRepositoryImpl) FindAll() (*[]models.Article, error) {
	var articles []models.Article
	err := r.db.Find(&articles).Error
	return &articles, err
}
func (r *articleRepositoryImpl) FindByID(id string) (*models.Article, error) {
	var article models.Article
	err := r.db.Where("id = ?", id).First(&article).Error
	return &article, err
}
func (r *articleRepositoryImpl) Update(id string, username string, data *models.Article) error {
	return r.db.Model(&models.Article{}).Where("id = ?", id).Updates(data).Error
}
func (r *articleRepositoryImpl) Delete(id string) error {
	return r.db.Delete(&models.Article{}, "id = ?", id).Error
}
