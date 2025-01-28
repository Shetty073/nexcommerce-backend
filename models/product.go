package models

import (
	"nexcommerce/stores"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID          uuid.UUID   `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name        string      `gorm:"type:varchar(150);not null;index"`
	Description string      `gorm:"type:varchar(1000)"`
	Price       float64     `gorm:"type:decimal(10,2);not null;index"`
	Market      string      `gorm:"type:enum('global', 'local');not null;index"`
	Brand       Brand       `gorm:"foreignKey:BrandID;index"`
	Currency    Currency    `gorm:"foreignKey:CurrencyID"`
	Category    Category    `gorm:"foreignKey:CategoryID;index"`
	SubCategory SubCategory `gorm:"foreignKey:SubCategoryID;index"`
	CreatedBy   User        `gorm:"foreignKey:CreatedBy;index"`
	UpdatedBy   User        `gorm:"foreignKey:CreatedBy;index"`
	CreatedAt   *time.Time  `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	UpdatedAt   *time.Time  `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;index"`
	DeletedAt   gorm.DeletedAt
}

// Receiver Methods
func (p *Product) Create() error {
	return stores.GetDb().Create(p).Error
}

func (p *Product) Update() error {
	return stores.GetDb().Save(p).Error
}

func (p *Product) Delete() error {
	return stores.GetDb().Delete(p).Error
}

// GetProductByID retrieves a product by ID
func GetProductByID(id uuid.UUID) (*Product, error) {
	var product Product
	if err := stores.GetDb().First(&product, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}
