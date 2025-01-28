package models

import (
	"nexcommerce/stores"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InventoryProduct struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	SKU       string     `gorm:"type:varchar(50);not null;index"`
	Status    string     `gorm:"type:enum('active', 'inactive');not null;index"`
	Inventory Inventory  `gorm:"foreignKey:InventoryID;index"`
	CreatedBy User       `gorm:"foreignKey:CreatedBy;index"`
	UpdatedBy User       `gorm:"foreignKey:UpdatedBy;index"`
	CreatedAt *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	UpdatedAt *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;index"`
	DeletedAt gorm.DeletedAt
}

// Receiver Methods
func (ip *InventoryProduct) Create() error {
	return stores.GetDb().Create(ip).Error
}

func (ip *InventoryProduct) Update() error {
	return stores.GetDb().Save(ip).Error
}

func (ip *InventoryProduct) Delete() error {
	return stores.GetDb().Delete(ip).Error
}

// GetInventoryProductByID retrieves an inventory product by ID
func GetInventoryProductByID(id uuid.UUID) (*InventoryProduct, error) {
	var ip InventoryProduct
	if err := stores.GetDb().First(&ip, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &ip, nil
}
