package models

import (
	"nexcommerce/stores"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Address struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Line1     string     `gorm:"type:varchar(100);not null"`
	Line2     string     `gorm:"type:varchar(100)"`
	Line3     string     `gorm:"type:varchar(100)"`
	Phone     string     `gorm:"type:varchar(15);index"`
	AltPhone  string     `gorm:"type:varchar(15);index"`
	PinCode   string     `gorm:"type:varchar(10);index"`
	City      string     `gorm:"type:varchar(100)"`
	District  string     `gorm:"type:varchar(100)"`
	State     string     `gorm:"type:varchar(100)"`
	Country   string     `gorm:"type:varchar(100);index"`
	Alias     string     `gorm:"type:varchar(20);index"`
	IsDefault bool       `gorm:"default:false;index"`
	UserID    uuid.UUID  `gorm:"type:uuid;not null;index"`
	CreatedAt *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	UpdatedAt *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	DeletedAt gorm.DeletedAt

	User User `gorm:"foreignKey:UserID"`
}

// CreateAddress creates a new address
func (a *Address) CreateAddress() error {
	return stores.GetDb().Create(a).Error
}

// UpdateAddress updates an existing address
func (a *Address) UpdateAddress() error {
	return stores.GetDb().Save(a).Error
}

// DeleteAddress soft deletes an address
func (a *Address) DeleteAddress() error {
	return stores.GetDb().Delete(a).Error
}

// GetAddressByID retrieves an address by ID
func GetAddressByID(id uuid.UUID) (*Address, error) {
	var address Address
	if err := stores.GetDb().Preload("User").First(&address, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &address, nil
}
