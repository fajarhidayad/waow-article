package repositories

import (
	"errors"

	"github.com/fajarhidayad/waow-article/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(*models.User) error
	GetUsers() ([]*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	IsEmailExist(email string) bool
	IsUsernameExist(username string) bool
	GetUserById(id string) (*models.User, error)
	UpdateUser(id string, user *models.User) error
	DeleteUser(id string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) CreateUser(user *models.User) error {
	return u.db.Create(user).Error
}

func (u *userRepository) GetUsers() ([]*models.User, error) {
	return nil, nil
}

func (u *userRepository) GetUserByEmail(email string) (*models.User, error) {
	var user *models.User
	err := u.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	if user.ID == "" {
		return nil, errors.New("Email not found")
	}

	return user, nil
}

func (u *userRepository) GetUserByUsername(username string) (*models.User, error) {
	var user *models.User
	err := u.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}

	if user.ID == "" {
		return nil, errors.New("Username not found")
	}

	return user, nil
}

func (u *userRepository) GetUserById(id string) (*models.User, error) {
	var user *models.User
	err := u.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	if user.ID == "" {
		return nil, errors.New("User not found")
	}

	return user, nil
}

func (u *userRepository) UpdateUser(id string, user *models.User) error {
	return nil
}

func (u *userRepository) DeleteUser(id string) error {
	return nil
}

func (u *userRepository) IsEmailExist(email string) bool {
	var user models.User
	u.db.Where("email = ?", email).First(&user)
	if user.ID == "" {
		return false
	}
	return true
}

func (u *userRepository) IsUsernameExist(username string) bool {
	var user models.User
	u.db.Where("username = ?", username).First(&user)
	if user.ID == "" {
		return false
	}
	return true
}
