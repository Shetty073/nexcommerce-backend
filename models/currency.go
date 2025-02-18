package models

import (
	"nexcommerce/stores"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Currency struct {
	ID                  uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name                string     `gorm:"type:varchar(20);not null;index"`
	Symbol              string     `gorm:"type:varchar(5);index"`
	ShortCode           string     `gorm:"type:varchar(5);index"`
	Country             string     `gorm:"type:varchar(100);index"`
	DollarExchangeValue float64    `gorm:"type:decimal"`
	CreatedBy           uuid.UUID  `gorm:"type:uuid;not null;index"`
	UpdatedBy           uuid.UUID  `gorm:"type:uuid;index"`
	CreatedAt           *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	UpdatedAt           *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	DeletedAt           gorm.DeletedAt

	CreatedByUser User `gorm:"foreignKey:CreatedBy"`
	UpdatedByUser User `gorm:"foreignKey:UpdatedBy"`
}

// CreateCurrency creates a new currency
func (c *Currency) CreateCurrency() error {
	return stores.GetDb().Create(c).Error
}

// UpdateCurrency updates an existing currency
func (c *Currency) UpdateCurrency() error {
	return stores.GetDb().Save(c).Error
}

// DeleteCurrency soft deletes a currency
func (c *Currency) DeleteCurrency() error {
	return stores.GetDb().Delete(c).Error
}

// GetCurrencyByID retrieves a currency by ID
func GetCurrencyByID(id uuid.UUID) (*Currency, error) {
	var currency Currency
	if err := stores.GetDb().Preload("CreatedByUser").Preload("UpdatedByUser").First(&currency, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &currency, nil
}
