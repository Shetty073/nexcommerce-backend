package models

import (
	"database/sql"
	"nexcommerce/constants/enums"
	"nexcommerce/stores"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SalesOrder struct {
	ID                 uuid.UUID              `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	OrderAt            sql.NullTime           `gorm:"type:timestamptz;not null;index"`
	Status             enums.SalesOrderStatus `gorm:"type:varchar(15);index"`
	NoOfUnits          int                    `gorm:"type:int;not null"`
	PricePerUnit       float64                `gorm:"type:decimal;not null"`
	TotalAmount        float64                `gorm:"type:decimal;not null"`
	Notes              string                 `gorm:"type:varchar(250)"`
	ProductID          uuid.UUID              `gorm:"type:uuid;not null;index"`
	InventoryID        uuid.UUID              `gorm:"type:uuid;not null;index"`
	InventoryProductID uuid.UUID              `gorm:"type:uuid;not null;index"`
	CustomerID         uuid.UUID              `gorm:"type:uuid;not null;index"`
	AddressID          uuid.UUID              `gorm:"type:uuid;not null;index"`
	CurrencyID         uuid.UUID              `gorm:"type:uuid;not null;index"`
	CreatedByID        uuid.UUID              `gorm:"type:uuid;not null;index"`
	UpdatedByID        uuid.UUID              `gorm:"type:uuid;index"`
	CreatedAt          *time.Time             `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	UpdatedAt          *time.Time             `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	DeletedAt          gorm.DeletedAt

	Product          Product          `gorm:"foreignKey:ProductID"`
	Inventory        Inventory        `gorm:"foreignKey:InventoryID"`
	InventoryProduct InventoryProduct `gorm:"foreignKey:InventoryProductID"`
	Customer         User             `gorm:"foreignKey:CustomerID"`
	Address          Address          `gorm:"foreignKey:AddressID"`
	Currency         Currency         `gorm:"foreignKey:CurrencyID"`
	CreatedBy        User             `gorm:"foreignKey:CreatedByID"`
	UpdatedBy        User             `gorm:"foreignKey:UpdatedByID"`
}

// Receiver Methods
func (so *SalesOrder) Create() error {
	return stores.GetDb().Create(so).Error
}

func (so *SalesOrder) Update() error {
	return stores.GetDb().Save(so).Error
}

func (so *SalesOrder) Delete() error {
	return stores.GetDb().Delete(so).Error
}

// GetSalesOrderByID retrieves sales order by ID
func GetSalesOrderByID(id uuid.UUID) (*SalesOrder, error) {
	var so SalesOrder
	if err := stores.GetDb().Preload("Product").Preload("Inventory").Preload("InventoryProduct").
		Preload("Customer").Preload("Address").Preload("Currency").
		Preload("CreatedBy").Preload("UpdatedBy").
		First(&so, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &so, nil
}
