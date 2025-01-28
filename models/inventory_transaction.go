package models

import (
	"nexcommerce/stores"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InventoryTransaction struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Status    string     `gorm:"type:enum('in', 'out', 'adjustment');not null;index"`
	Quantity  int        `gorm:"type:int;not null"`
	Product   Product    `gorm:"foreignKey:ProductID;index"`
	Warehouse Warehouse  `gorm:"foreignKey:WarehouseID;index"`
	CreatedBy User       `gorm:"foreignKey:CreatedBy;index"`
	UpdatedBy User       `gorm:"foreignKey:UpdatedBy;index"`
	CreatedAt *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	UpdatedAt *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;index"`
	DeletedAt gorm.DeletedAt
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
	if err := stores.GetDb().First(&it, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &it, nil
}
