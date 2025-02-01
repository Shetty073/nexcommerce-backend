package models

import (
	"nexcommerce/stores"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LogisticsUpdate struct {
	ID                   uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	CurrentLocation      string     `gorm:"type:varchar(100);not null"`
	ExpectedDeliveryDate string     `gorm:"type:datetime;not null"`
	Notes                string     `gorm:"type:varchar(500)"`
	LogisticsID          uuid.UUID  `gorm:"type:uuid;not null;index"`
	CreatedByID          uuid.UUID  `gorm:"type:uuid;not null;index"`
	UpdatedByID          uuid.UUID  `gorm:"type:uuid;index"`
	CreatedAt            *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	UpdatedAt            *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	DeletedAt            gorm.DeletedAt

	Logistics Logistics `gorm:"foreignKey:LogisticsID"`
	CreatedBy User      `gorm:"foreignKey:CreatedByID"`
	UpdatedBy User      `gorm:"foreignKey:UpdatedByID"`
}

// Receiver Methods
func (lu *LogisticsUpdate) Create() error {
	return stores.GetDb().Create(lu).Error
}

func (lu *LogisticsUpdate) Update() error {
	return stores.GetDb().Save(lu).Error
}

func (lu *LogisticsUpdate) Delete() error {
	return stores.GetDb().Delete(lu).Error
}

// GetLogisticsUpdateByID retrieves logistics update by ID
func GetLogisticsUpdateByID(id uuid.UUID) (*LogisticsUpdate, error) {
	var lu LogisticsUpdate
	if err := stores.GetDb().Preload("Logistics").Preload("CreatedBy").Preload("UpdatedBy").First(&lu, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &lu, nil
}
