package models

import (
	"nexcommerce/stores"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Permission struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name      string     `gorm:"type:varchar(50);not null;unique"`
	CreatedBy User       `gorm:"foreignKey:CreatedBy;index"`
	UpdatedBy User       `gorm:"foreignKey:UpdatedBy;index"`
	CreatedAt *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	UpdatedAt *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;index"`
	DeletedAt gorm.DeletedAt
}

// CreatePermission creates a new permission
func (p *Permission) CreatePermission() error {
	return stores.GetDb().Create(p).Error
}

// UpdatePermission updates an existing permission
func (p *Permission) UpdatePermission() error {
	return stores.GetDb().Save(p).Error
}

// DeletePermission marks a permission as inactive
func (p *Permission) DeletePermission() error {
	p.Name = p.Name + "_inactive" // Example to signify it is inactive
	return stores.GetDb().Save(p).Error
}

// GetPermissionByID retrieves a permission by ID
func GetPermissionByID(id uuid.UUID) (*Permission, error) {
	var permission Permission
	if err := stores.GetDb().First(&permission, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &permission, nil
}
