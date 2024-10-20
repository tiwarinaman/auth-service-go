package auth

import (
	"auth/pkg/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	authService Service
}

func NewAuthController(as Service) Controller {
	return Controller{authService: as}
}

func (ac *Controller) Register(c *gin.Context) {
	var regDTO dto.RegisterDTO
	if err := c.ShouldBindJSON(&regDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ac.authService.Register(regDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func (ac *Controller) Login(c *gin.Context) {
	var loginDTO dto.LoginDTO
	if err := c.ShouldBindJSON(&loginDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := ac.authService.Login(loginDTO)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
