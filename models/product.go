package models

import (
	"nexcommerce/constants/enums"
	"nexcommerce/stores"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID            uuid.UUID               `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name          string                  `gorm:"type:varchar(150);not null;index"`
	Description   string                  `gorm:"type:varchar(1000)"`
	Price         float64                 `gorm:"type:decimal(10,2);not null;index"`
	Market        enums.ProductMarketType `gorm:"type:varchar(15);index"`
	BrandID       uuid.UUID               `gorm:"type:uuid;not null;index"`
	CurrencyID    uuid.UUID               `gorm:"type:uuid;not null;index"`
	CategoryID    uuid.UUID               `gorm:"type:uuid;not null;index"`
	SubCategoryID uuid.UUID               `gorm:"type:uuid;index"`
	CreatedBy     uuid.UUID               `gorm:"type:uuid;not null;index"`
	UpdatedBy     uuid.UUID               `gorm:"type:uuid;index"`
	CreatedAt     *time.Time              `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	UpdatedAt     *time.Time              `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	DeletedAt     gorm.DeletedAt

	Brand         Brand       `gorm:"foreignKey:BrandID"`
	Currency      Currency    `gorm:"foreignKey:CurrencyID"`
	Category      Category    `gorm:"foreignKey:CategoryID"`
	SubCategory   SubCategory `gorm:"foreignKey:SubCategoryID"`
	CreatedByUser User        `gorm:"foreignKey:CreatedBy"`
	UpdatedByUser User        `gorm:"foreignKey:UpdatedBy"`
}

// Create creates a new product
func (p *Product) Create() error {
	return stores.GetDb().Create(p).Error
}

// Update updates an existing product
func (p *Product) Update() error {
	return stores.GetDb().Save(p).Error
}

// Delete soft deletes a product
func (p *Product) Delete() error {
	return stores.GetDb().Delete(p).Error
}

// GetProductByID retrieves a product by ID
func GetProductByID(id uuid.UUID) (*Product, error) {
	var product Product
	if err := stores.GetDb().Preload("Brand").Preload("Currency").Preload("Category").Preload("SubCategory").Preload("CreatedByUser").Preload("UpdatedByUser").First(&product, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}
