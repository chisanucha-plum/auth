package services

import (
	"auth/internal/models"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func GenerateJWT(userID int) (string, error) {
	claims := &models.JWTClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return "", errors.New("JWT secret is not set in environment variables")
	}

	return token.SignedString([]byte(jwtSecret))

}

func ValidateJWT(tokenString string) (*models.JWTClaims, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return nil, errors.New("JWT secret is not set in environment variables")
	}

	token, err := jwt.ParseWithClaims(tokenString, &models.JWTClaims{}, func(t *jwt.Token) (interface{}, error) {
		// enforce HMAC signing method
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*models.JWTClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

type AuthService struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{db: db}
}

func (s *AuthService) Register(username, email, password, fullName string) (string, error) {
	// Check if user already exists
	var existingUser models.User
	if err := s.db.Where("username = ? OR email = ?", username, email).First(&existingUser).Error; err == nil {
		return "", errors.New("user already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	// Create user
	user := models.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
		FullName: fullName,
		IsActive: true,
	}

	if err := s.db.Create(&user).Error; err != nil {
		return "", err
	}

	// Generate JWT token for the newly registered user
	token, err := GenerateJWT(int(user.ID))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthService) Login(username, password string) (string, error) {
	var user models.User
	if err := s.db.Where("username = ?", username).First(&user).Error; err != nil {
		return "", errors.New("user not found")
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid password")
	}

	// Generate JWT token
	token, err := GenerateJWT(int(user.ID))
	if err != nil {
		return "", err
	}

	return token, nil
}
