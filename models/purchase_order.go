package models

import (
	"database/sql"
	"nexcommerce/constants/enums"
	"nexcommerce/stores"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PurchaseOrder struct {
	ID          uuid.UUID                 `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	OrderAt     sql.NullTime              `gorm:"type:timestamptz;not null;index"`
	Status      enums.PurchaseOrderStatus `gorm:"type:varchar(15);index"`
	NoOfUnits   int                       `gorm:"type:int;not null"`
	CostPerUnit float64                   `gorm:"type:decimal;not null"`
	TotalCost   float64                   `gorm:"type:decimal;not null"`
	Notes       string                    `gorm:"type:varchar(250)"`
	InventoryID uuid.UUID                 `gorm:"type:uuid;not null;index"`
	SupplierID  uuid.UUID                 `gorm:"type:uuid;not null;index"`
	CurrencyID  uuid.UUID                 `gorm:"type:uuid;not null;index"`
	CreatedByID uuid.UUID                 `gorm:"type:uuid;not null;index"`
	UpdatedByID uuid.UUID                 `gorm:"type:uuid;index"`
	CreatedAt   *time.Time                `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	UpdatedAt   *time.Time                `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	DeletedAt   gorm.DeletedAt

	Inventory Inventory `gorm:"foreignKey:InventoryID"`
	Supplier  Supplier  `gorm:"foreignKey:SupplierID"`
	Currency  Currency  `gorm:"foreignKey:CurrencyID"`
	CreatedBy User      `gorm:"foreignKey:CreatedByID"`
	UpdatedBy User      `gorm:"foreignKey:UpdatedByID"`
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
	if err := stores.GetDb().Preload("Inventory").Preload("Supplier").Preload("Currency").Preload("CreatedBy").Preload("UpdatedBy").First(&po, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &po, nil
}

func (u *User) GetAllPurchaseOrders() ([]PurchaseOrder, error) {
	var purchaseOrders []PurchaseOrder
	if err := stores.GetDb().Where("created_by_id = ?", u.ID).Find(&purchaseOrders).Error; err != nil {
		return nil, err
	}
	return purchaseOrders, nil
}
