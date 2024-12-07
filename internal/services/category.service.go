package services

import (
	"github.com/fajarhidayad/waow-article/internal/dtos"
	"github.com/fajarhidayad/waow-article/internal/models"
	"github.com/fajarhidayad/waow-article/internal/repositories"
	"github.com/fajarhidayad/waow-article/pkg/common"
)

type CategoryService interface {
	CreateCategory(category *dtos.CreateCategoryDto) (*common.Response, *common.ErrorResponse)
	FindAllCategories() (*common.Response, *common.ErrorResponse)
	FindCategoryByID(id string) (*common.Response, *common.ErrorResponse)
	UpdateCategory(id string, category *dtos.UpdateCategoryDto) (*common.Response, *common.ErrorResponse)
	DeleteCategory(id string) (*common.Response, *common.ErrorResponse)
}

type categoryServiceImpl struct {
	categoryRepo repositories.CategoryRepository
}

func NewCategoryService(categoryRepo repositories.CategoryRepository) CategoryService {
	return &categoryServiceImpl{
		categoryRepo: categoryRepo,
	}
}

func (service *categoryServiceImpl) CreateCategory(req *dtos.CreateCategoryDto) (*common.Response, *common.ErrorResponse) {
	category := models.Category{
		Name:        req.Name,
		Description: req.Description,
	}
	err := service.categoryRepo.Create(&category)
	if err != nil {
		return nil, &common.ErrorResponse{
			Error: err.Error(),
		}
	}

	return &common.Response{
		Data:    category.ID,
		Message: "success",
	}, nil
}

func (service *categoryServiceImpl) FindAllCategories() (*common.Response, *common.ErrorResponse) {
	categories, err := service.categoryRepo.FindAll()
	if err != nil {
		return nil, &common.ErrorResponse{
			Error: err.Error(),
		}
	}

	return &common.Response{
		Data:    categories,
		Message: "success",
	}, nil
}

func (service *categoryServiceImpl) FindCategoryByID(id string) (*common.Response, *common.ErrorResponse) {
	category, err := service.categoryRepo.FindByID(id)
	if err != nil {
		return nil, &common.ErrorResponse{
			Error: err.Error(),
		}
	}

	return &common.Response{
		Data:    category,
		Message: "success",
	}, nil
}

func (service *categoryServiceImpl) UpdateCategory(id string, req *dtos.UpdateCategoryDto) (*common.Response, *common.ErrorResponse) {
	category := models.Category{
		Name:        req.Name,
		Description: req.Description,
	}

	err := service.categoryRepo.Update(id, &category)
	if err != nil {
		return nil, &common.ErrorResponse{
			Error: err.Error(),
		}
	}

	return &common.Response{
		Data:    nil,
		Message: "success",
	}, nil
}

func (service *categoryServiceImpl) DeleteCategory(id string) (*common.Response, *common.ErrorResponse) {
	err := service.categoryRepo.Delete(id)
	if err != nil {
		return nil, &common.ErrorResponse{
			Error: err.Error(),
		}
	}

	return &common.Response{
		Message: "success",
		Data:    nil,
	}, nil
}
