package models

import (
	"nexcommerce/constants/enums"
	"nexcommerce/stores"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Chat struct {
	ID          uuid.UUID        `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Status      enums.ChatStatus `gorm:"type:varchar(15);index"`
	TicketID    uuid.UUID        `gorm:"type:uuid;not null;index"`
	CreatedByID uuid.UUID        `gorm:"type:uuid;not null;index"`
	UpdatedByID uuid.UUID        `gorm:"type:uuid;index"`
	CreatedAt   *time.Time       `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	UpdatedAt   *time.Time       `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	DeletedAt   gorm.DeletedAt

	Ticket    SupportTicket `gorm:"foreignKey:TicketID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedBy User          `gorm:"foreignKey:CreatedByID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UpdatedBy User          `gorm:"foreignKey:UpdatedByID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// Receiver Methods
func (c *Chat) Create() error {
	return stores.GetDb().Create(c).Error
}

func (c *Chat) Update() error {
	return stores.GetDb().Save(c).Error
}

func (c *Chat) Delete() error {
	return stores.GetDb().Delete(c).Error
}

// GetChatByID retrieves chat by ID
func GetChatByID(id uuid.UUID) (*Chat, error) {
	var c Chat
	if err := stores.GetDb().First(&c, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &c, nil
}
