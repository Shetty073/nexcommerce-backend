package models

import (
	"database/sql"
	"nexcommerce/stores"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SalesOrder struct {
	ID               uuid.UUID        `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	OrderAt          sql.NullTime     `gorm:"type:datetime;not null;index"`
	Status           string           `gorm:"type:enum('pending', 'shipped', 'delivered');not null;index"`
	NoOfUnits        int              `gorm:"type:int;not null"`
	PricePerUnit     float64          `gorm:"type:decimal;not null"`
	TotalAmount      float64          `gorm:"type:decimal;not null"`
	Notes            string           `gorm:"type:varchar(250)"`
	Product          Product          `gorm:"foreignKey:ProductID;index"`
	Inventory        Inventory        `gorm:"foreignKey:InventoryID;index"`
	InventoryProduct InventoryProduct `gorm:"foreignKey:InventoryProductID;index"`
	Customer         User             `gorm:"foreignKey:CustomerID;index"`
	Address          Address          `gorm:"foreignKey:AddressID;index"`
	Currency         Currency         `gorm:"foreignKey:CurrencyID;index"`
	CreatedBy        User             `gorm:"foreignKey:CreatedBy;index"`
	UpdatedBy        User             `gorm:"foreignKey:UpdatedBy;index"`
	CreatedAt        *time.Time       `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	UpdatedAt        *time.Time       `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;index"`
	DeletedAt        gorm.DeletedAt
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
	if err := stores.GetDb().First(&so, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &so, nil
}
