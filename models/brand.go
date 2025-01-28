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
	CreatedBy   User       `gorm:"foreignKey:CreatedBy;index"`
	UpdatedBy   User       `gorm:"foreignKey:UpdatedBy;index"`
	CreatedAt   *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	UpdatedAt   *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;index"`
	DeletedAt   gorm.DeletedAt
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
	err := stores.GetDb().First(&brand, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &brand, nil
}
