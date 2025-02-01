package models

import (
	"nexcommerce/stores"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoleModulePermission struct {
	ID           uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	RoleID       uuid.UUID  `gorm:"type:uuid;not null;index"` // Foreign key for RoleID
	ModuleID     uuid.UUID  `gorm:"type:uuid;not null;index"` // Foreign key for ModuleID
	PermissionID uuid.UUID  `gorm:"type:uuid;not null;index"` // Foreign key for PermissionID
	CreatedAt    *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	UpdatedAt    *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	DeletedAt    gorm.DeletedAt

	Role       Role       `gorm:"foreignKey:RoleID"`
	Module     Module     `gorm:"foreignKey:ModuleID"`
	Permission Permission `gorm:"foreignKey:PermissionID"`
}

// CreateRoleModulePermission creates a new RoleModulePermission
func (rmp *RoleModulePermission) CreateRoleModulePermission() error {
	return stores.GetDb().Create(rmp).Error
}

// UpdateRoleModulePermission updates an existing RoleModulePermission
func (rmp *RoleModulePermission) UpdateRoleModulePermission() error {
	return stores.GetDb().Save(rmp).Error
}

// DeleteRoleModulePermission soft deletes a RoleModulePermission
func (rmp *RoleModulePermission) DeleteRoleModulePermission() error {
	return stores.GetDb().Delete(rmp).Error
}

// GetRoleModulePermission retrieves a RoleModulePermission by ID
func GetRoleModulePermission(id uuid.UUID) (*RoleModulePermission, error) {
	var rmp RoleModulePermission
	// Preload the associated Role, Module, and Permission
	if err := stores.GetDb().Preload("Role").Preload("Module").Preload("Permission").First(&rmp, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &rmp, nil
}
