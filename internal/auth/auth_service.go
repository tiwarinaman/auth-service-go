package auth

import (
	"auth/internal/jwt"
	"auth/internal/user"
	"auth/pkg/dto"
)

type Service interface {
	Register(dto.RegisterDTO) error
	Login(dto.LoginDTO) (string, error)
}

type authService struct {
	userService user.Service
	jwtService  jwt.Service
}

func NewAuthService(us user.Service, js jwt.Service) Service {
	return &authService{
		userService: us,
		jwtService:  js,
	}
}

func (as *authService) Register(registerDTO dto.RegisterDTO) error {
	return as.userService.Register(registerDTO)
}

func (as *authService) Login(loginDTO dto.LoginDTO) (string, error) {
	user, err := as.userService.Login(loginDTO)
	if err != nil {
		return "", err
	}

	token, err := as.jwtService.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
