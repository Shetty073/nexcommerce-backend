package models

import (
	"nexcommerce/stores"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TicketNote struct {
	ID          uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UpdateNote  string     `gorm:"type:varchar(1000)"`
	Image1      string     `gorm:"type:varchar(1024)"`
	Image2      string     `gorm:"type:varchar(1024)"`
	Image3      string     `gorm:"type:varchar(1024)"`
	Image4      string     `gorm:"type:varchar(1024)"`
	TicketID    uuid.UUID  `gorm:"type:uuid;not null;index"`
	CreatedByID uuid.UUID  `gorm:"type:uuid;not null;index"`
	UpdatedByID uuid.UUID  `gorm:"type:uuid;index"`
	CreatedAt   *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	UpdatedAt   *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	DeletedAt   gorm.DeletedAt

	Ticket    SupportTicket `gorm:"foreignKey:TicketID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedBy User          `gorm:"foreignKey:CreatedByID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UpdatedBy User          `gorm:"foreignKey:UpdatedByID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
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
