package models

import (
	"nexcommerce/stores"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ServicablePinCode struct {
	ID           uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	PinCode      string     `gorm:"type:varchar(10);not null;unique;index"`
	Country      string     `gorm:"type:varchar(100);index"`
	IsServicable bool       `gorm:"default:false"`
	CreatedBy    User       `gorm:"foreignKey:CreatedBy;index"`
	UpdatedBy    User       `gorm:"foreignKey:CreatedBy;index"`
	CreatedAt    *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	UpdatedAt    *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;index"`
	DeletedAt    gorm.DeletedAt
}

// Receiver Methods
func (spc *ServicablePinCode) Create() error {
	return stores.GetDb().Create(spc).Error
}

func (spc *ServicablePinCode) Update() error {
	return stores.GetDb().Save(spc).Error
}

func (spc *ServicablePinCode) Delete() error {
	return stores.GetDb().Delete(spc).Error
}

// GetServicablePinCodeByID retrieves a servicable pin code by ID
func GetServicablePinCodeByID(id uuid.UUID) (*ServicablePinCode, error) {
	var spc ServicablePinCode
	if err := stores.GetDb().First(&spc, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &spc, nil
}
