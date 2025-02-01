package models

import (
	"nexcommerce/stores"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Warehouse struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string     `gorm:"type:varchar(50);not null;index"`
	Phone     string     `gorm:"type:varchar(15);index"`
	Address   string     `gorm:"type:varchar(100)"`
	PinCode   string     `gorm:"type:varchar(10);index"`
	City      string     `gorm:"type:varchar(100)"`
	District  string     `gorm:"type:varchar(100)"`
	State     string     `gorm:"type:varchar(100)"`
	Country   string     `gorm:"type:varchar(100);index"`
	CreatedBy uuid.UUID  `gorm:"type:uuid;not null;index"`
	UpdatedBy uuid.UUID  `gorm:"type:uuid;index"`
	CreatedAt *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	UpdatedAt *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	DeletedAt gorm.DeletedAt

	CreatedByUser User `gorm:"foreignKey:CreatedBy"`
	UpdatedByUser User `gorm:"foreignKey:UpdatedBy"`
}

// CreateWarehouse creates a new warehouse
func (w *Warehouse) CreateWarehouse() error {
	return stores.GetDb().Create(w).Error
}

// UpdateWarehouse updates an existing warehouse
func (w *Warehouse) UpdateWarehouse() error {
	return stores.GetDb().Save(w).Error
}

// DeleteWarehouse soft deletes a warehouse
func (w *Warehouse) DeleteWarehouse() error {
	return stores.GetDb().Delete(w).Error
}

// GetWarehouseByID retrieves a warehouse by ID
func GetWarehouseByID(id uuid.UUID) (*Warehouse, error) {
	var warehouse Warehouse
	if err := stores.GetDb().Preload("CreatedByUser").Preload("UpdatedByUser").First(&warehouse, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &warehouse, nil
}
