package models

import (
	"nexcommerce/stores"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Module struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string     `gorm:"type:varchar(50);not null;unique"`
	CreatedBy uuid.UUID  `gorm:"type:uuid;not null;index"`
	UpdatedBy uuid.UUID  `gorm:"type:uuid;not null;index"`
	CreatedAt *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	UpdatedAt *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	DeletedAt gorm.DeletedAt

	UserCreatedBy User `gorm:"foreignKey:CreatedBy"`
	UserUpdatedBy User `gorm:"foreignKey:UpdatedBy"`
}

// CreateModule creates a new module
func (m *Module) CreateModule() error {
	return stores.GetDb().Create(m).Error
}

// UpdateModule updates an existing module
func (m *Module) UpdateModule() error {
	return stores.GetDb().Save(m).Error
}

// DeleteModule marks a module as inactive
func (m *Module) DeleteModule() error {
	m.Name = m.Name + "_inactive" // Example of making it non-functional
	return stores.GetDb().Save(m).Error
}

// GetModuleByID retrieves a module by ID
func GetModuleByID(id uuid.UUID) (*Module, error) {
	var module Module
	// Preload the associated UserCreatedBy and UserUpdatedBy relationships
	if err := stores.GetDb().Preload("UserCreatedBy").Preload("UserUpdatedBy").First(&module, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &module, nil
}
