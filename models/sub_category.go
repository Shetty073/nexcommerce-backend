package models

import (
	"nexcommerce/stores"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SubCategory struct {
	ID          uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name        string     `gorm:"type:varchar(50);not null;index"`
	Description string     `gorm:"type:varchar(100)"`
	CategoryID  uuid.UUID  `gorm:"type:uuid;not null;index"`
	CreatedBy   uuid.UUID  `gorm:"type:uuid;not null;index"`
	UpdatedBy   uuid.UUID  `gorm:"type:uuid;index"`
	CreatedAt   *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	UpdatedAt   *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	DeletedAt   gorm.DeletedAt

	Category      Category `gorm:"foreignKey:CategoryID"`
	CreatedByUser User     `gorm:"foreignKey:CreatedBy"`
	UpdatedByUser User     `gorm:"foreignKey:UpdatedBy"`
}

// CreateSubCategory creates a new sub-category
func (sc *SubCategory) CreateSubCategory() error {
	return stores.GetDb().Create(sc).Error
}

// UpdateSubCategory updates an existing sub-category
func (sc *SubCategory) UpdateSubCategory() error {
	return stores.GetDb().Save(sc).Error
}

// DeleteSubCategory soft deletes a sub-category
func (sc *SubCategory) DeleteSubCategory() error {
	return stores.GetDb().Delete(sc).Error
}

// GetSubCategoryByID retrieves a sub-category by ID
func GetSubCategoryByID(id uuid.UUID) (*SubCategory, error) {
	var subCategory SubCategory
	if err := stores.GetDb().Preload("Category").Preload("CreatedByUser").Preload("UpdatedByUser").First(&subCategory, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &subCategory, nil
}
