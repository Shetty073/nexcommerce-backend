package models

import (
	"nexcommerce/stores"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRole struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID    uuid.UUID  `gorm:"type:uuid;not null;index"` // Foreign key for UserID
	RoleID    uuid.UUID  `gorm:"type:uuid;not null;index"` // Foreign key for RoleID
	CreatedAt *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	UpdatedAt *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	DeletedAt gorm.DeletedAt

	// Explicit relationships for GORM
	User User `gorm:"foreignKey:UserID"`
	Role Role `gorm:"foreignKey:RoleID"`
}

// CreateUserRole creates a new UserRole
func (ur *UserRole) CreateUserRole() error {
	return stores.GetDb().Create(ur).Error
}

// UpdateUserRole updates an existing UserRole
func (ur *UserRole) UpdateUserRole() error {
	return stores.GetDb().Save(ur).Error
}

// DeleteUserRole soft deletes a UserRole
func (ur *UserRole) DeleteUserRole() error {
	return stores.GetDb().Delete(ur).Error
}

// GetUserRole retrieves a UserRole by ID
func GetUserRole(id uuid.UUID) (*UserRole, error) {
	var userRole UserRole
	// Preload the associated User and Role
	if err := stores.GetDb().Preload("User").Preload("Role").First(&userRole, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &userRole, nil
}
