package services

import (
	"errors"

	"github.com/fajarhidayad/waow-article/internal/dtos"
	"github.com/fajarhidayad/waow-article/internal/models"
	"github.com/fajarhidayad/waow-article/internal/repositories"
	"github.com/fajarhidayad/waow-article/pkg/auth"
	"github.com/fajarhidayad/waow-article/pkg/common"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(*dtos.RegisterRequest) (*common.Response, error)
	Login(*dtos.LoginRequest) (*auth.TokenResponse, error)
	GetUser(id string) (*models.User, error)
}

type authService struct {
	userRepository repositories.UserRepository
}

func NewAuthService(userRepository repositories.UserRepository) AuthService {
	return &authService{
		userRepository: userRepository,
	}
}

func (service *authService) Register(user *dtos.RegisterRequest) (*common.Response, error) {
	usernameExist := service.userRepository.IsUsernameExist(user.Username)
	if usernameExist {
		return nil, errors.New("username already exist")
	}

	emailExist := service.userRepository.IsEmailExist(user.Email)
	if emailExist {
		return nil, errors.New("email already exist")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	newUser := &models.User{
		Password:          string(hashedPassword),
		Username:          user.Username,
		Email:             user.Email,
		DisplayName:       user.DisplayName,
		Bio:               user.Bio,
		ProfilePictureURL: user.ProfilePictureURL,
		RegistrationDate:  user.RegistrationDate,
		Role:              user.Role,
	}

	err = service.userRepository.CreateUser(newUser)
	if err != nil {
		return nil, err
	}

	response := &common.Response{
		Data:    newUser.ID,
		Message: "Success",
	}

	return response, nil
}

func (service *authService) Login(req *dtos.LoginRequest) (*auth.TokenResponse, error) {
	user, err := service.userRepository.GetUserByUsername(req.Username)
	if err != nil {
		if err.Error() == "Username not found" {
			return nil, errors.New("Invalid credentials")
		}
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("Invalid credentials")
	}

	response, err := auth.GenerateToken(user.ID, user.Username, user.Role)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *authService) GetUser(id string) (*models.User, error) {
	user, err := service.userRepository.GetUserById(id)
	return user, err
}
