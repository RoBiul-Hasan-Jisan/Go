package handlers

import (
	"errors"
	"net/http"

	"taskmanager/internal/services"
	"taskmanager/internal/utils"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

type registerRequest struct {
	Name     string `json:"name" binding:"required,min=2"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type loginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// authResponse is what gets returned in the "data" field for register/login.
type authResponse struct {
	Token string      `json:"token"`
	User  interface{} `json:"user"`
}

// Register handles POST /api/auth/register
func (h *AuthHandler) Register(c *gin.Context) {
	var req registerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	user, token, err := h.authService.Register(req.Name, req.Email, req.Password)
	if err != nil {
		if errors.Is(err, services.ErrEmailTaken) {
			utils.Error(c, http.StatusConflict, err.Error())
			return
		}
		utils.Error(c, http.StatusInternalServerError, "Failed to register user")
		return
	}

	utils.Success(c, http.StatusCreated, "User registered successfully", authResponse{
		Token: token,
		User:  user,
	})
}

// Login handles POST /api/auth/login
func (h *AuthHandler) Login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	user, token, err := h.authService.Login(req.Email, req.Password)
	if err != nil {
		utils.Error(c, http.StatusUnauthorized, err.Error())
		return
	}

	utils.Success(c, http.StatusOK, "Login successful", authResponse{
		Token: token,
		User:  user,
	})
}
