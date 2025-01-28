package models

import (
	"database/sql"
	"nexcommerce/stores"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PurchaseOrder struct {
	ID          uuid.UUID    `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	OrderAt     sql.NullTime `gorm:"type:datetime;not null;index"`
	Status      string       `gorm:"type:enum('pending', 'completed', 'cancelled');not null;index"`
	NoOfUnits   int          `gorm:"type:int;not null"`
	CostPerUnit float64      `gorm:"type:decimal;not null"`
	TotalCost   float64      `gorm:"type:decimal;not null"`
	Notes       string       `gorm:"type:varchar(250)"`
	Inventory   Inventory    `gorm:"foreignKey:InventoryID;index"`
	Supplier    Supplier     `gorm:"foreignKey:SupplierID;index"`
	Currency    Currency     `gorm:"foreignKey:CurrencyID;index"`
	CreatedBy   User         `gorm:"foreignKey:CreatedBy;index"`
	UpdatedBy   User         `gorm:"foreignKey:UpdatedBy;index"`
	CreatedAt   *time.Time   `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	UpdatedAt   *time.Time   `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;index"`
	DeletedAt   gorm.DeletedAt
}

// Receiver Methods
func (po *PurchaseOrder) Create() error {
	return stores.GetDb().Create(po).Error
}

func (po *PurchaseOrder) Update() error {
	return stores.GetDb().Save(po).Error
}

func (po *PurchaseOrder) Delete() error {
	return stores.GetDb().Delete(po).Error
}

// GetPurchaseOrderByID retrieves purchase order by ID
func GetPurchaseOrderByID(id uuid.UUID) (*PurchaseOrder, error) {
	var po PurchaseOrder
	if err := stores.GetDb().First(&po, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &po, nil
}

func (u *User) GetAllPurchaseOrders() ([]PurchaseOrder, error) {
	var purchaseOrders []PurchaseOrder
	if err := stores.GetDb().Where("created_by = ?", u.ID).Find(&purchaseOrders).Error; err != nil {
		return nil, err
	}
	return purchaseOrders, nil
}
