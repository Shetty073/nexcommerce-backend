package models

import (
	"nexcommerce/stores"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Inventory struct {
	ID                uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name              string     `gorm:"type:varchar(150);not null;index"`
	Description       string     `gorm:"type:varchar(250)"`
	StockLevel        int        `gorm:"type:int;not null;index"`
	MinimumStockLevel int        `gorm:"type:int;not null;index"`
	ProductID         uuid.UUID  `gorm:"type:uuid;not null;index"`
	WarehouseID       uuid.UUID  `gorm:"type:uuid;not null;index"`
	CreatedByID       uuid.UUID  `gorm:"type:uuid;not null;index"`
	UpdatedByID       uuid.UUID  `gorm:"type:uuid;not null;index"`
	CreatedAt         *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	UpdatedAt         *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	DeletedAt         gorm.DeletedAt

	Product   Product   `gorm:"foreignKey:ProductID"`
	Warehouse Warehouse `gorm:"foreignKey:WarehouseID"`
	CreatedBy User      `gorm:"foreignKey:CreatedByID"`
	UpdatedBy User      `gorm:"foreignKey:UpdatedByID"`
}

// Create creates a new inventory record
func (i *Inventory) Create() error {
	return stores.GetDb().Create(i).Error
}

// Update updates an existing inventory record
func (i *Inventory) Update() error {
	return stores.GetDb().Save(i).Error
}

// Delete soft deletes an inventory record
func (i *Inventory) Delete() error {
	return stores.GetDb().Delete(i).Error
}

// GetInventoryByID retrieves an inventory record by ID
func GetInventoryByID(id uuid.UUID) (*Inventory, error) {
	var inventory Inventory
	if err := stores.GetDb().Preload("Product").Preload("Warehouse").Preload("CreatedBy").Preload("UpdatedBy").First(&inventory, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &inventory, nil
}
