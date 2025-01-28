package models

import (
	"nexcommerce/stores"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Chat struct {
	ID        uuid.UUID     `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Status    string        `gorm:"type:enum('active', 'closed');not null;index"`
	Ticket    SupportTicket `gorm:"foreignKey:TicketID;index"`
	CreatedBy User          `gorm:"foreignKey:CreatedBy;index"`
	UpdatedBy User          `gorm:"foreignKey:UpdatedBy;index"`
	CreatedAt *time.Time    `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	UpdatedAt *time.Time    `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;index"`
	DeletedAt gorm.DeletedAt
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
