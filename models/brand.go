package models

import (
	"nexcommerce/stores"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Brand struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name        string     `gorm:"type:varchar(150);not null;unique;index"`
	LogoImg     string     `gorm:"type:varchar(1024)"`
	BannerImg   string     `gorm:"type:varchar(1024)"`
	Description string     `gorm:"type:varchar(1000)"`
	CreatedBy   uuid.UUID  `gorm:"type:uuid;not null;index"`
	UpdatedBy   uuid.UUID  `gorm:"type:uuid;index"`
	CreatedAt   *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	UpdatedAt   *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	DeletedAt   gorm.DeletedAt

	CreatedByUser User `gorm:"foreignKey:CreatedBy"`
	UpdatedByUser User `gorm:"foreignKey:UpdatedBy"`
}

// CreateBrand creates a new brand
func (b *Brand) CreateBrand() error {
	return stores.GetDb().Create(b).Error
}

// UpdateBrand updates an existing brand
func (b *Brand) UpdateBrand() error {
	return stores.GetDb().Save(b).Error
}

// DeleteBrand marks a brand as inactive
func (b *Brand) DeleteBrand() error {
	b.Name = b.Name + "_inactive"
	return stores.GetDb().Save(b).Error
}

// GetBrand retrieves a brand by its ID
func (b *Brand) GetBrand(id uuid.UUID) (*Brand, error) {
	var brand Brand
	if err := stores.GetDb().Preload("CreatedByUser").Preload("UpdatedByUser").First(&brand, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &brand, nil
}
