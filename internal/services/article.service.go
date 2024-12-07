package services

import (
	"github.com/fajarhidayad/waow-article/internal/dtos"
	"github.com/fajarhidayad/waow-article/internal/models"
	"github.com/fajarhidayad/waow-article/internal/repositories"
	"github.com/fajarhidayad/waow-article/pkg/common"
)

type ArticleService interface {
	Create(username string, article *dtos.CreateArticleDto) (*common.Response, *common.ErrorResponse)
	FindAll() (*common.Response, *common.ErrorResponse)
	FindByID(id string) (*common.Response, *common.ErrorResponse)
	Update(id string, username string, article *dtos.UpdateArticleDto) (*common.Response, *common.ErrorResponse)
	Delete(id string) (*common.Response, *common.ErrorResponse)
}

type articleServiceImpl struct {
	articleRepo repositories.ArticleRepository
	userRepo    repositories.UserRepository
}

func NewArticleService(articleRepo repositories.ArticleRepository, userRepo repositories.UserRepository) ArticleService {
	return &articleServiceImpl{
		articleRepo: articleRepo,
		userRepo:    userRepo,
	}
}

func (s *articleServiceImpl) Create(username string, req *dtos.CreateArticleDto) (*common.Response, *common.ErrorResponse) {
	user, err := s.userRepo.GetUserByUsername(username)
	if err != nil {
		return nil, &common.ErrorResponse{
			Error: err.Error(),
		}
	}

	data := models.Article{
		Title:      req.Title,
		Content:    req.Content,
		CategoryID: req.CategoryID,
		UserID:     user.ID,
	}

	err = s.articleRepo.Create(&data)
	if err != nil {
		return nil, &common.ErrorResponse{
			Error: err.Error(),
		}
	}

	return &common.Response{
		Data:    data.ID,
		Message: "created",
	}, nil
}
func (s *articleServiceImpl) FindAll() (*common.Response, *common.ErrorResponse) {
	articles, err := s.articleRepo.FindAll()
	if err != nil {
		return nil, &common.ErrorResponse{
			Error: err.Error(),
		}
	}

	return &common.Response{
		Data:    articles,
		Message: "success",
	}, nil
}
func (s *articleServiceImpl) FindByID(id string) (*common.Response, *common.ErrorResponse) {
	article, err := s.articleRepo.FindByID(id)
	if err != nil {
		return nil, &common.ErrorResponse{
			Error: err.Error(),
		}
	}

	return &common.Response{
		Data:    article,
		Message: "success",
	}, nil
}
func (s *articleServiceImpl) Update(id string, username string, req *dtos.UpdateArticleDto) (*common.Response, *common.ErrorResponse) {
	user, err := s.userRepo.GetUserByUsername(username)
	if err != nil {
		return nil, &common.ErrorResponse{
			Error: err.Error(),
		}
	}

	article, err := s.articleRepo.FindByID(id)
	if err != nil {
		return nil, &common.ErrorResponse{
			Error: err.Error(),
		}
	}

	article.Title = req.Title
	article.Content = req.Content
	article.CategoryID = req.CategoryID
	article.UserID = user.ID

	err = s.articleRepo.Update(id, username, article)
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
func (s *articleServiceImpl) Delete(id string) (*common.Response, *common.ErrorResponse) {
	err := s.articleRepo.Delete(id)
	if err != nil {
		return nil, &common.ErrorResponse{
			Error: err.Error(),
		}
	}

	return &common.Response{
		Data:    nil,
		Message: "deleted",
	}, nil
}
