package services

import (
	"github.com/fajarhidayad/waow-article/internal/dtos"
	"github.com/fajarhidayad/waow-article/internal/models"
	"github.com/fajarhidayad/waow-article/internal/repositories"
	"github.com/fajarhidayad/waow-article/pkg/common"
	"golang.org/x/crypto/bcrypt"
)

type AdminService interface {
	CreateUser(reqBody *dtos.CreateUserDTO) (*common.Response, *common.ErrorResponse)
	FindUsers() (*common.Response, *common.ErrorResponse)
	FindUserByID(id string) (*common.Response, *common.ErrorResponse)
	UpdateUser(id string, reqBody *dtos.UpdateUserDTO) (*common.Response, *common.ErrorResponse)
	DeleteUser(id string) (*common.Response, *common.ErrorResponse)
}

type adminServiceImpl struct {
	userRepo repositories.UserRepository
}

func NewAdminService(userRepo repositories.UserRepository) AdminService {
	return &adminServiceImpl{
		userRepo: userRepo,
	}
}

func (s *adminServiceImpl) CreateUser(reqBody *dtos.CreateUserDTO) (*common.Response, *common.ErrorResponse) {
	usernameExist := s.userRepo.IsUsernameExist(reqBody.Username)
	if usernameExist {
		return nil, &common.ErrorResponse{
			Error: "username already exist",
		}
	}

	emailExist := s.userRepo.IsEmailExist(reqBody.Email)
	if emailExist {
		return nil, &common.ErrorResponse{
			Error: "email already exist",
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(reqBody.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, &common.ErrorResponse{
			Error: err.Error(),
		}
	}

	newUser := models.User{
		Username:          reqBody.Username,
		Email:             reqBody.Email,
		Password:          string(hashedPassword),
		DisplayName:       reqBody.DisplayName,
		Bio:               reqBody.Bio,
		ProfilePictureURL: reqBody.ProfilePictureURL,
		Role:              reqBody.Role,
	}

	err = s.userRepo.CreateUser(&newUser)
	if err != nil {
		return nil, &common.ErrorResponse{
			Error: err.Error(),
		}
	}

	return &common.Response{
		Data:    newUser.ID,
		Message: "success",
	}, nil
}

func (s *adminServiceImpl) FindUsers() (*common.Response, *common.ErrorResponse) {
	users, err := s.userRepo.GetUsers()
	if err != nil {
		return nil, &common.ErrorResponse{
			Error: err.Error(),
		}
	}

	return &common.Response{
		Data:    users,
		Message: "success",
	}, nil
}
func (s *adminServiceImpl) FindUserByID(id string) (*common.Response, *common.ErrorResponse) {
	user, err := s.userRepo.GetUserById(id)
	if err != nil {
		return nil, &common.ErrorResponse{
			Error: err.Error(),
		}
	}

	return &common.Response{
		Data:    user,
		Message: "success",
	}, nil
}

func (s *adminServiceImpl) UpdateUser(id string, reqBody *dtos.UpdateUserDTO) (*common.Response, *common.ErrorResponse) {
	user := &models.User{
		Username:          reqBody.Username,
		Email:             reqBody.Email,
		Password:          reqBody.Password,
		DisplayName:       reqBody.DisplayName,
		Bio:               reqBody.Bio,
		ProfilePictureURL: reqBody.ProfilePictureURL,
		Role:              reqBody.Role,
	}
	err := s.userRepo.UpdateUser(id, user)
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

func (s *adminServiceImpl) DeleteUser(id string) (*common.Response, *common.ErrorResponse) {
	user, err := s.userRepo.GetUserById(id)
	if err != nil {
		return nil, &common.ErrorResponse{
			Error: err.Error(),
		}
	}
	if user == nil {
		return nil, &common.ErrorResponse{
			Error: "user not found",
		}
	}

	err = s.userRepo.DeleteUser(id)
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
