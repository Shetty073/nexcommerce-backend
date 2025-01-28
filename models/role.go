package models

import (
	"nexcommerce/stores"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	ID         uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name       string     `gorm:"type:varchar(50);not null;unique;index"`
	IsElevated bool       `gorm:"default:false;index"`
	CreatedBy  User       `gorm:"foreignKey:CreatedBy;index"`
	UpdatedBy  User       `gorm:"foreignKey:UpdatedBy;index"`
	CreatedAt  *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	UpdatedAt  *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;index"`
	DeletedAt  gorm.DeletedAt
}

// CreateRole creates a new role
func (r *Role) CreateRole() error {
	return stores.GetDb().Create(r).Error
}

// UpdateRole updates an existing role
func (r *Role) UpdateRole() error {
	return stores.GetDb().Save(r).Error
}

// DeleteRole marks a role as inactive
func (r *Role) DeleteRole() error {
	r.IsElevated = false
	return stores.GetDb().Save(r).Error
}

// GetRoleByID retrieves a role by ID
func GetRoleByID(id uuid.UUID) (*Role, error) {
	var role Role
	if err := stores.GetDb().First(&role, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &role, nil
}
