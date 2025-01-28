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
	Product           Product    `gorm:"foreignKey:ProductID;index"`
	Warehouse         Warehouse  `gorm:"foreignKey:WarehouseID;index"`
	CreatedBy         User       `gorm:"foreignKey:CreatedBy;index"`
	UpdatedBy         User       `gorm:"foreignKey:UpdatedBy;index"`
	CreatedAt         *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	UpdatedAt         *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;index"`
	DeletedAt         gorm.DeletedAt
}

// Receiver Methods
func (i *Inventory) Create() error {
	return stores.GetDb().Create(i).Error
}

func (i *Inventory) Update() error {
	return stores.GetDb().Save(i).Error
}

func (i *Inventory) Delete() error {
	return stores.GetDb().Delete(i).Error
}

// GetInventoryByID retrieves an inventory by ID
func GetInventoryByID(id uuid.UUID) (*Inventory, error) {
	var inventory Inventory
	if err := stores.GetDb().First(&inventory, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &inventory, nil
}
