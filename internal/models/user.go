package models

import "gorm.io/gorm"

type User struct {
	gorm.Model        // ID, CreatedAt, UpdatedAt, DeletedAt
	Username   string `gorm:"uniqueIndex;not null" json:"username" binding:"required"`
	Email      string `gorm:"uniqueIndex" json:"email" binding:"omitempty,email"`
	Password   string `gorm:"not null" json:"-"`
	FullName   string `json:"full_name"`
	IsActive   bool   `gorm:"default:true" json:"is_active"`
}
