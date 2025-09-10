package handler

import (
	"auth/internal/models"
	"auth/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{authService}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Error:   "Invalid input: " + err.Error(),
		})
		return
	}

	token, err := h.authService.Register(req.Username, req.Email, req.Password, req.FullName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Error:   "Registration failed: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.APIResponse{
		Success: true,
		Message: "User registered successfully",
		Data:    gin.H{"access_token": token},
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Error:   "Invalid input: " + err.Error(),
		})
		return
	}

	token, err := h.authService.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.APIResponse{
			Success: false,
			Error:   "Invalid credentials: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Login successful",
		Data:    gin.H{"access_token": token},
	})
}

func (h *AuthHandler) GetProfile(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.APIResponse{
			Success: false,
			Error:   "No authentication claims found",
		})
		return
	}

	jwtClaims, ok := claims.(*models.JWTClaims)
	if !ok {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Error:   "Invalid claims format",
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Profile retrieved successfully",
		Data: gin.H{
			"user_id":    jwtClaims.UserID,
			"issued_at":  jwtClaims.IssuedAt.Time,
			"expires_at": jwtClaims.ExpiresAt.Time,
		},
	})
}
