package models

import (
	"nexcommerce/stores"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductMedia struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	File        string     `gorm:"type:varchar(1024);not null"`
	Type        string     `gorm:"type:enum('image', 'video');not null;index"`
	Description string     `gorm:"type:varchar(20)"`
	Product     Product    `gorm:"foreignKey:ProductID;index"`
	CreatedBy   User       `gorm:"foreignKey:CreatedBy;index"`
	UpdatedBy   User       `gorm:"foreignKey:CreatedBy;index"`
	CreatedAt   *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	UpdatedAt   *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;index"`
	DeletedAt   gorm.DeletedAt
}

// Receiver Methods
func (pm *ProductMedia) Create() error {
	return stores.GetDb().Create(pm).Error
}

func (pm *ProductMedia) Update() error {
	return stores.GetDb().Save(pm).Error
}

func (pm *ProductMedia) Delete() error {
	return stores.GetDb().Delete(pm).Error
}

// GetProductMediaByID retrieves product media by ID
func GetProductMediaByID(id uuid.UUID) (*ProductMedia, error) {
	var pm ProductMedia
	if err := stores.GetDb().First(&pm, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &pm, nil
}
