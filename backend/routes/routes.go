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
		router.POST("/shorten", shortenerHandler.ShortenURL)
		router.GET("/:id", shortenerHandler.RedirectURL)
	}
}

