package models

import (
	"nexcommerce/stores"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Currency struct {
	ID                  uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name                string     `gorm:"type:varchar(20);not null;index"`
	Symbol              string     `gorm:"type:varchar(5);index"`
	ShortCode           string     `gorm:"type:varchar(5);index"`
	DollarExchangeValue float64    `gorm:"type:decimal"`
	CreatedBy           User       `gorm:"foreignKey:CreatedBy;index"`
	UpdatedBy           User       `gorm:"foreignKey:UpdatedBy;index"`
	CreatedAt           *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	UpdatedAt           *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;index"`
	DeletedAt           gorm.DeletedAt
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
	if err := stores.GetDb().First(&currency, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &currency, nil
}
