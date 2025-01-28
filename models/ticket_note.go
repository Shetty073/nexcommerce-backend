package models

import (
	"nexcommerce/stores"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TicketNote struct {
	ID         uuid.UUID     `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UpdateNote string        `gorm:"type:varchar(1000)"`
	Image1     string        `gorm:"type:varchar(1024)"`
	Image2     string        `gorm:"type:varchar(1024)"`
	Image3     string        `gorm:"type:varchar(1024)"`
	Image4     string        `gorm:"type:varchar(1024)"`
	Ticket     SupportTicket `gorm:"foreignKey:TicketID;index"`
	CreatedBy  User          `gorm:"foreignKey:CreatedBy;index"`
	UpdatedBy  User          `gorm:"foreignKey:UpdatedBy;index"`
	CreatedAt  *time.Time    `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	UpdatedAt  *time.Time    `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;index"`
	DeletedAt  gorm.DeletedAt
}

// Receiver Methods
func (tn *TicketNote) Create() error {
	return stores.GetDb().Create(tn).Error
}

func (tn *TicketNote) Update() error {
	return stores.GetDb().Save(tn).Error
}

func (tn *TicketNote) Delete() error {
	return stores.GetDb().Delete(tn).Error
}

// GetTicketNoteByID retrieves ticket note by ID
func GetTicketNoteByID(id uuid.UUID) (*TicketNote, error) {
	var tn TicketNote
	if err := stores.GetDb().First(&tn, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &tn, nil
}
