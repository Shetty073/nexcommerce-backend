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
	CreatedBy  uuid.UUID  `gorm:"type:uuid;not null;index"`
	UpdatedBy  uuid.UUID  `gorm:"type:uuid;not null;index"`
	CreatedAt  *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	UpdatedAt  *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	DeletedAt  gorm.DeletedAt

	UserCreatedBy User `gorm:"foreignKey:CreatedBy"`
	UserUpdatedBy User `gorm:"foreignKey:UpdatedBy"`
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
	if err := stores.GetDb().Preload("UserCreatedBy").Preload("UserUpdatedBy").First(&role, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &role, nil
}
