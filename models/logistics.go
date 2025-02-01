package models

import (
	"nexcommerce/stores"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Logistics struct {
	ID             uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	TrackingNumber string     `gorm:"type:varchar(50);not null;index"`
	SalesOrderID   uuid.UUID  `gorm:"type:uuid;not null;index"`
	CreatedByID    uuid.UUID  `gorm:"type:uuid;not null;index"`
	UpdatedByID    uuid.UUID  `gorm:"type:uuid;index"`
	CreatedAt      *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	UpdatedAt      *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	DeletedAt      gorm.DeletedAt

	SalesOrder SalesOrder `gorm:"foreignKey:SalesOrderID"`
	CreatedBy  User       `gorm:"foreignKey:CreatedByID"`
	UpdatedBy  User       `gorm:"foreignKey:UpdatedByID"`
}

// Receiver Methods
func (l *Logistics) Create() error {
	return stores.GetDb().Create(l).Error
}

func (l *Logistics) Update() error {
	return stores.GetDb().Save(l).Error
}

func (l *Logistics) Delete() error {
	return stores.GetDb().Delete(l).Error
}

// GetLogisticsByID retrieves logistics by ID
func GetLogisticsByID(id uuid.UUID) (*Logistics, error) {
	var l Logistics
	if err := stores.GetDb().Preload("SalesOrder").Preload("CreatedBy").Preload("UpdatedBy").First(&l, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &l, nil
}
