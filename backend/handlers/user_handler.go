package handlers

import (
	"encurtador-de-link/backend/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(UserService *service.UserService) *UserHandler {
	return &UserHandler{userService: UserService}
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	user, err := h.userService.GetUserProfile(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "USer not found"})
		return

	}

	c.JSON(http.StatusOK, user)
}
