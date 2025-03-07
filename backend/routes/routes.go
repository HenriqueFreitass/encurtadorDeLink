package routes

import (
	"encurtador-de-link/backend/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, userHandler *handlers.UserHandler, shortenerHandler *handlers.ShortenerHandler) {
	userGroup := router.Group("/users")
	{
		userGroup.GET("/id", userHandler.GetUser)
		userGroup.POST("/", userHandler.CreateUser)
		userGroup.POST("/login", userHandler.LoginUser)
		userGroup.POST("/shorten", shortenerHandler.ShortenURL)
		userGroup.GET("/:id", shortenerHandler.RedirectURL)
	}
}

