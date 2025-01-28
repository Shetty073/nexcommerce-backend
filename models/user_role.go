package models

import (
	"nexcommerce/stores"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRole struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	User      User       `gorm:"foreignKey:UserID;index"`
	Role      Role       `gorm:"foreignKey:RoleID;index"`
	CreatedAt *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	UpdatedAt *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;index"`
	DeletedAt gorm.DeletedAt
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
	if err := stores.GetDb().First(&userRole, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &userRole, nil
}
