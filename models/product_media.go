package models

import (
	"nexcommerce/constants/enums"
	"nexcommerce/stores"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductMedia struct {
	ID          uuid.UUID              `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	File        string                 `gorm:"type:varchar(1024);not null"`
	Type        enums.ProductMediaType `gorm:"type:varchar(15);index"`
	Description string                 `gorm:"type:varchar(20)"`
	ProductID   uuid.UUID              `gorm:"type:uuid;not null;index"`
	CreatedBy   uuid.UUID              `gorm:"type:uuid;not null;index"`
	UpdatedBy   uuid.UUID              `gorm:"type:uuid;index"`
	CreatedAt   *time.Time             `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	UpdatedAt   *time.Time             `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	DeletedAt   gorm.DeletedAt

	Product       Product `gorm:"foreignKey:ProductID"`
	CreatedByUser User    `gorm:"foreignKey:CreatedBy"`
	UpdatedByUser User    `gorm:"foreignKey:UpdatedBy"`
}

// Create creates new product media
func (pm *ProductMedia) Create() error {
	return stores.GetDb().Create(pm).Error
}

// Update updates an existing product media
func (pm *ProductMedia) Update() error {
	return stores.GetDb().Save(pm).Error
}

// Delete soft deletes a product media
func (pm *ProductMedia) Delete() error {
	return stores.GetDb().Delete(pm).Error
}

// GetProductMediaByID retrieves product media by ID
func GetProductMediaByID(id uuid.UUID) (*ProductMedia, error) {
	var pm ProductMedia
	if err := stores.GetDb().Preload("Product").Preload("CreatedByUser").Preload("UpdatedByUser").First(&pm, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &pm, nil
}
