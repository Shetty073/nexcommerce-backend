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
	SalesOrder     SalesOrder `gorm:"foreignKey:SalesOrderID;index"`
	CreatedBy      User       `gorm:"foreignKey:CreatedBy;index"`
	UpdatedBy      User       `gorm:"foreignKey:UpdatedBy;index"`
	CreatedAt      *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	UpdatedAt      *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;index"`
	DeletedAt      gorm.DeletedAt
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
	if err := stores.GetDb().First(&l, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &l, nil
}
