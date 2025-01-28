package models

import (
	"nexcommerce/stores"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name        string     `gorm:"type:varchar(50);not null;index"`
	Description string     `gorm:"type:varchar(100)"`
	CreatedBy   User       `gorm:"foreignKey:CreatedBy;index"`
	UpdatedBy   User       `gorm:"foreignKey:UpdatedBy;index"`
	CreatedAt   *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	UpdatedAt   *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;index"`
	DeletedAt   gorm.DeletedAt
}

// CreateCategory creates a new category
func (c *Category) CreateCategory() error {
	return stores.GetDb().Create(c).Error
}

// UpdateCategory updates an existing category
func (c *Category) UpdateCategory() error {
	return stores.GetDb().Save(c).Error
}

// DeleteCategory soft deletes a category
func (c *Category) DeleteCategory() error {
	return stores.GetDb().Delete(c).Error
}

// GetCategoryByID retrieves a category by ID
func GetCategoryByID(id uuid.UUID) (*Category, error) {
	var category Category
	if err := stores.GetDb().First(&category, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}
