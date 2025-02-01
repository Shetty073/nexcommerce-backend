package models

import (
	"nexcommerce/stores"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Supplier struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name        string     `gorm:"type:varchar(100);not null;index"`
	Description string     `gorm:"type:varchar(250)"`
	Phone       string     `gorm:"type:varchar(15);index"`
	Address     string     `gorm:"type:varchar(100)"`
	PinCode     string     `gorm:"type:varchar(10);index"`
	City        string     `gorm:"type:varchar(100)"`
	District    string     `gorm:"type:varchar(100)"`
	State       string     `gorm:"type:varchar(100)"`
	Country     string     `gorm:"type:varchar(100);index"`
	CreatedByID uuid.UUID  `gorm:"type:uuid;not null;index"`
	UpdatedByID uuid.UUID  `gorm:"type:uuid;index"`
	CreatedAt   *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	UpdatedAt   *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	DeletedAt   gorm.DeletedAt

	CreatedBy User `gorm:"foreignKey:CreatedByID"`
	UpdatedBy User `gorm:"foreignKey:UpdatedByID"`
}

// Receiver Methods
func (s *Supplier) Create() error {
	return stores.GetDb().Create(s).Error
}

func (s *Supplier) Update() error {
	return stores.GetDb().Save(s).Error
}

func (s *Supplier) Delete() error {
	return stores.GetDb().Delete(s).Error
}

// GetSupplierByID retrieves a supplier by ID
func GetSupplierByID(id uuid.UUID) (*Supplier, error) {
	var supplier Supplier
	if err := stores.GetDb().Preload("CreatedBy").Preload("UpdatedBy").First(&supplier, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &supplier, nil
}
