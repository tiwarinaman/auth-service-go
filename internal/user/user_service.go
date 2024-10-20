package user

import (
	"auth/model"
	"auth/pkg/dto"
	"auth/pkg/utils"
	"errors"
)

type Service interface {
	Register(dto dto.RegisterDTO) error
	Login(dto dto.LoginDTO) (model.User, error)
}

type userService struct {
	userRepo Repository
}

func NewUserService(ur Repository) Service {
	return &userService{userRepo: ur}
}

func (us *userService) Register(registerDTO dto.RegisterDTO) error {
	hashedPassword, err := utils.HashPassword(registerDTO.Password)
	if err != nil {
		return err
	}

	newUser := model.User{
		Email:    registerDTO.Email,
		Password: hashedPassword,
	}

	return us.userRepo.CreateUser(newUser)
}

func (us *userService) Login(loginDTO dto.LoginDTO) (model.User, error) {
	user, err := us.userRepo.FindByEmail(loginDTO.Email)
	if err != nil || !utils.CheckPasswordHash(loginDTO.Password, user.Password) {
		return model.User{}, errors.New("invalid credentials")
	}
	return user, nil
}
