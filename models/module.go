package models

import (
	"nexcommerce/stores"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Module struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name      string     `gorm:"type:varchar(50);not null;unique"`
	CreatedBy User       `gorm:"foreignKey:CreatedBy;index"`
	UpdatedBy User       `gorm:"foreignKey:UpdatedBy;index"`
	CreatedAt *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	UpdatedAt *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;index"`
	DeletedAt gorm.DeletedAt
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
	if err := stores.GetDb().First(&module, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &module, nil
}
