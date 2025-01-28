package models

import (
	"nexcommerce/stores"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Promotion struct {
	ID           uuid.UUID         `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Code         string            `gorm:"type:varchar(10);not null;index"`
	DiscountRate float64           `gorm:"type:decimal"`
	DiscountFlat float64           `gorm:"type:decimal"`
	Description  string            `gorm:"type:varchar(100)"`
	Campaign     MarketingCampaign `gorm:"foreignKey:CampaignID;index"`
	CreatedBy    User              `gorm:"foreignKey:CreatedBy;index"`
	UpdatedBy    User              `gorm:"foreignKey:UpdatedBy;index"`
	CreatedAt    *time.Time        `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	UpdatedAt    *time.Time        `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;index"`
	DeletedAt    gorm.DeletedAt
}

// Receiver Methods
func (p *Promotion) Create() error {
	return stores.GetDb().Create(p).Error
}

func (p *Promotion) Update() error {
	return stores.GetDb().Save(p).Error
}

func (p *Promotion) Delete() error {
	return stores.GetDb().Delete(p).Error
}

// GetPromotionByID retrieves promotion by ID
func GetPromotionByID(id uuid.UUID) (*Promotion, error) {
	var p Promotion
	if err := stores.GetDb().First(&p, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &p, nil
}
