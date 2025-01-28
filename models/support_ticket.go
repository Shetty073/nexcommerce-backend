package models

import (
	"nexcommerce/stores"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SupportTicket struct {
	ID               uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	IssueType        string     `gorm:"type:enum('technical', 'billing', 'other');not null;index"`
	IssueDescription string     `gorm:"type:varchar(500);not null"`
	Status           string     `gorm:"type:enum('open', 'closed', 'in-progress');not null;index"`
	Customer         User       `gorm:"foreignKey:CustomerID;index"`
	AssignedTo       User       `gorm:"foreignKey:AssignedTo;index"`
	CreatedBy        User       `gorm:"foreignKey:CreatedBy;index"`
	UpdatedBy        User       `gorm:"foreignKey:UpdatedBy;index"`
	CreatedAt        *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	UpdatedAt        *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;index"`
	DeletedAt        gorm.DeletedAt
}

// Receiver Methods
func (st *SupportTicket) Create() error {
	return stores.GetDb().Create(st).Error
}

func (st *SupportTicket) Update() error {
	return stores.GetDb().Save(st).Error
}

func (st *SupportTicket) Delete() error {
	return stores.GetDb().Delete(st).Error
}

// GetSupportTicketByID retrieves support ticket by ID
func GetSupportTicketByID(id uuid.UUID) (*SupportTicket, error) {
	var st SupportTicket
	if err := stores.GetDb().First(&st, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &st, nil
}
