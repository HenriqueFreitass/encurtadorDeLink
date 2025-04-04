package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"encurtador-de-link/backend/models"
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

func (h *UserHandler) LoginUser(c *gin.Context) {
	var loginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Id       int    `json:"id"`
	}

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	t := sha256.New()
	t.Write([]byte(loginRequest.Password))
	bs := t.Sum(nil)
	hexString := hex.EncodeToString(bs)
	// Tenta autenticar o usuário
	user, err := h.userService.AuthenticateUser(loginRequest.Email, hexString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email ou senha inválidos"})
		return
	}

	// Aqui, você pode gerar um token JWT para o usuário, se desejar
	c.JSON(http.StatusOK, gin.H{"message": "Login realizado com sucesso!", "user": user})
}

func (h *UserHandler) DeleteLink(c *gin.Context) {
	id := c.Param("id")

	result, err := h.userService.DeleteLink(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Link not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": result})
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	user, err := h.userService.GetUserProfile(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.Users
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	createdUser, err := h.userService.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}
