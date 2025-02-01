package models

import (
	"nexcommerce/constants/enums"
	"nexcommerce/stores"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InventoryProduct struct {
	ID          uuid.UUID                    `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	SKU         string                       `gorm:"type:varchar(50);not null;index"`
	Status      enums.InventoryProductStatus `gorm:"type:varchar(15);index"`
	InventoryID uuid.UUID                    `gorm:"type:uuid;not null;index"`
	CreatedByID uuid.UUID                    `gorm:"type:uuid;not null;index"`
	UpdatedByID uuid.UUID                    `gorm:"type:uuid;index"`
	CreatedAt   *time.Time                   `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	UpdatedAt   *time.Time                   `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	DeletedAt   gorm.DeletedAt

	Inventory Inventory `gorm:"foreignKey:InventoryID"`
	CreatedBy User      `gorm:"foreignKey:CreatedByID"`
	UpdatedBy User      `gorm:"foreignKey:UpdatedByID"`
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
	if err := stores.GetDb().Preload("Inventory").Preload("CreatedBy").Preload("UpdatedBy").First(&ip, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &ip, nil
}
