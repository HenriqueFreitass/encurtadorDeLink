package routes

import (
	"encurtador-de-link/backend/handlers"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine, userHandler *handlers.UserHandler, shortenerHandler *handlers.ShortenerHandler) {
	router.GET("/:id", shortenerHandler.RedirectURL)
	router.POST("/shorten", shortenerHandler.ShortenURL)

	userGroup := router.Group("/users")
	{
		userGroup.DELETE("/link/:id", userHandler.DeleteLink)
		userGroup.GET("/link/:id", shortenerHandler.GetUserLinks)
		userGroup.GET("/:id", userHandler.GetUser)
		userGroup.POST("/", userHandler.CreateUser)
		userGroup.POST("/login", userHandler.LoginUser)
	}
}
