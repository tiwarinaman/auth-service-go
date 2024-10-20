package jwt

import (
	"auth/pkg/config"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Service interface {
	GenerateToken(userID uint) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() Service {
	cfg := config.Get()
	return &jwtService{
		secretKey: cfg.JWT.Secret,
		issuer:    cfg.JWT.Issuer,
	}
}

func (js *jwtService) GenerateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
		"iss":    js.issuer,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(js.secretKey))
}

func (js *jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(js.secretKey), nil
	})
}
