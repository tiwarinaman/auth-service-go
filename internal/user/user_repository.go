package user

import (
	"auth/model"
	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(user model.User) error
	FindByEmail(email string) (model.User, error)
	FindById(id uint) (model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) Repository {
	return &userRepository{db: db}
}

func (u userRepository) CreateUser(user model.User) error {
	return u.db.Create(&user).Error
}

func (u userRepository) FindByEmail(email string) (model.User, error) {
	var user model.User
	if err := u.db.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (u userRepository) FindById(id uint) (model.User, error) {
	var user model.User
	if err := u.db.First(&user, id).Error; err != nil {
		return user, err
	}
	return user, nil
}
