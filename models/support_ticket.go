package models

import (
	"nexcommerce/constants/enums"
	"nexcommerce/stores"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SupportTicket struct {
	ID               uuid.UUID                    `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	IssueType        enums.SupportTicketIssueType `gorm:"type:varchar(15);index"`
	IssueDescription string                       `gorm:"type:varchar(500);not null"`
	Status           enums.SupportTicketStatus    `gorm:"type:varchar(15);index"`
	CustomerID       uuid.UUID                    `gorm:"type:uuid;not null;index"`
	AssignedToID     uuid.UUID                    `gorm:"type:uuid;index"`
	CreatedByID      uuid.UUID                    `gorm:"type:uuid;not null;index"`
	UpdatedByID      uuid.UUID                    `gorm:"type:uuid;index"`
	CreatedAt        *time.Time                   `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	UpdatedAt        *time.Time                   `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	DeletedAt        gorm.DeletedAt

	Customer   User `gorm:"foreignKey:CustomerID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	AssignedTo User `gorm:"foreignKey:AssignedToID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedBy  User `gorm:"foreignKey:CreatedByID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UpdatedBy  User `gorm:"foreignKey:UpdatedByID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
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
