package routes

import (
	"encurtador-de-link/backend/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, userHandler *handlers.UserHandler) {
	loginGroup := router.Group("/auth")
	{
		loginGroup.POST("/login", userHandler.LoginUser)
	}
	userGroup := router.Group("/users")
	{
		userGroup.GET("/id", userHandler.GetUser)
		userGroup.POST("/", userHandler.CreateUser)
	}
}
