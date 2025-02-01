package models

import (
	"nexcommerce/stores"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MarketingCampaign struct {
	ID             uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name           string     `gorm:"type:varchar(100);not null;index"`
	Description    string     `gorm:"type:varchar(1000)"`
	HasStoreBanner bool       `gorm:"type:boolean;default:false;index"`
	StoreBannerImg string     `gorm:"type:varchar(1024)"`
	StartDate      string     `gorm:"type:timestamptz;not null;index"`
	EndDate        string     `gorm:"type:timestamptz;not null;index"`
	ProductID      uuid.UUID  `gorm:"type:uuid;not null;index"`
	CreatedByID    uuid.UUID  `gorm:"type:uuid;not null;index"`
	UpdatedByID    uuid.UUID  `gorm:"type:uuid;index"`
	CreatedAt      *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	UpdatedAt      *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	DeletedAt      gorm.DeletedAt

	Product   Product `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedBy User    `gorm:"foreignKey:CreatedByID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UpdatedBy User    `gorm:"foreignKey:UpdatedByID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// Receiver Methods
func (mc *MarketingCampaign) Create() error {
	return stores.GetDb().Create(mc).Error
}

func (mc *MarketingCampaign) Update() error {
	return stores.GetDb().Save(mc).Error
}

func (mc *MarketingCampaign) Delete() error {
	return stores.GetDb().Delete(mc).Error
}

// GetMarketingCampaignByID retrieves marketing campaign by ID
func GetMarketingCampaignByID(id uuid.UUID) (*MarketingCampaign, error) {
	var mc MarketingCampaign
	if err := stores.GetDb().First(&mc, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &mc, nil
}
