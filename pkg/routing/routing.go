package routing

import (
	"auth/internal/auth"
	"auth/internal/jwt"
	"auth/internal/user"
	"auth/pkg/database"
	"auth/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func Init() {
	router = gin.Default()
}

func GetRouter() *gin.Engine {
	return router
}

func RegisterRoutes() {

	userRepo := user.NewUserRepository(database.Connection())
	userService := user.NewUserService(userRepo)
	jwtService := jwt.NewJWTService()
	authService := auth.NewAuthService(userService, jwtService)

	authCtrl := auth.NewAuthController(authService)

	router.POST("/api/v1/register", authCtrl.Register)
	router.POST("/api/v1/login", authCtrl.Login)

	protectedRoute := router.Group("api/v1/")
	protectedRoute.Use(middleware.JWTAuthMiddleware(jwtService))
	protectedRoute.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

}
