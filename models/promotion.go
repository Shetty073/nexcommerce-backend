package models

import (
	"nexcommerce/stores"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Promotion struct {
	ID           uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Code         string     `gorm:"type:varchar(10);not null;index"`
	DiscountRate float64    `gorm:"type:decimal"`
	DiscountFlat float64    `gorm:"type:decimal"`
	Description  string     `gorm:"type:varchar(100)"`
	CampaignID   uuid.UUID  `gorm:"type:uuid;not null;index"`
	CreatedByID  uuid.UUID  `gorm:"type:uuid;not null;index"`
	UpdatedByID  uuid.UUID  `gorm:"type:uuid;index"`
	CreatedAt    *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	UpdatedAt    *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	DeletedAt    gorm.DeletedAt

	Campaign  MarketingCampaign `gorm:"foreignKey:CampaignID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedBy User              `gorm:"foreignKey:CreatedByID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UpdatedBy User              `gorm:"foreignKey:UpdatedByID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
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
