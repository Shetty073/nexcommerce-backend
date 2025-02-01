package models

import (
	"nexcommerce/constants/enums"
	"nexcommerce/stores"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InventoryTransaction struct {
	ID          uuid.UUID                        `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Status      enums.InventoryTransactionStatus `gorm:"type:varchar(15);index"`
	Quantity    int                              `gorm:"type:int;not null"`
	ProductID   uuid.UUID                        `gorm:"type:uuid;not null;index"`
	WarehouseID uuid.UUID                        `gorm:"type:uuid;not null;index"`
	CreatedByID uuid.UUID                        `gorm:"type:uuid;not null;index"`
	UpdatedByID uuid.UUID                        `gorm:"type:uuid;index"`
	CreatedAt   *time.Time                       `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	UpdatedAt   *time.Time                       `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	DeletedAt   gorm.DeletedAt

	Product   Product   `gorm:"foreignKey:ProductID"`
	Warehouse Warehouse `gorm:"foreignKey:WarehouseID"`
	CreatedBy User      `gorm:"foreignKey:CreatedByID"`
	UpdatedBy User      `gorm:"foreignKey:UpdatedByID"`
}

// Receiver Methods
func (it *InventoryTransaction) Create() error {
	return stores.GetDb().Create(it).Error
}

func (it *InventoryTransaction) Update() error {
	return stores.GetDb().Save(it).Error
}

func (it *InventoryTransaction) Delete() error {
	return stores.GetDb().Delete(it).Error
}

// GetInventoryTransactionByID retrieves inventory transaction by ID
func GetInventoryTransactionByID(id uuid.UUID) (*InventoryTransaction, error) {
	var it InventoryTransaction
	if err := stores.GetDb().Preload("Product").Preload("Warehouse").Preload("CreatedBy").Preload("UpdatedBy").First(&it, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &it, nil
}
