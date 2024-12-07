package repositories

import (
	"errors"
	"golang.org/x/crypto/bcrypt"

	"github.com/fajarhidayad/waow-article/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(*models.User) error
	GetUsers() (*[]models.User, error)
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

func (u *userRepository) GetUsers() (*[]models.User, error) {
	var users []models.User

	err := u.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return &users, nil
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
		return nil, errors.New("username not found")
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

func (u *userRepository) UpdateUser(id string, req *models.User) error {
	tx := u.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	user, err := u.GetUserById(id)
	if err != nil {
		tx.Rollback()
		return err
	}

	if req.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			tx.Rollback()
			return err
		}
		req.Password = string(hash)
	} else {
		req.Password = user.Password
	}

	err = tx.Model(&user).Updates(req).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (u *userRepository) DeleteUser(id string) error {
	err := u.db.Where("id = ?", id).Delete(&models.User{}).Error
	return err
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
