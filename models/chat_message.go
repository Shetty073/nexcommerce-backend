package models

import (
	"nexcommerce/stores"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ChatMessage struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Message   string     `gorm:"type:varchar(1000);not null"`
	CreatedBy User       `gorm:"foreignKey:CreatedBy;index"`
	UpdatedBy User       `gorm:"foreignKey:UpdatedBy;index"`
	CreatedAt *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	UpdatedAt *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;index"`
	DeletedAt gorm.DeletedAt
}

// Receiver Methods
func (cm *ChatMessage) Create() error {
	return stores.GetDb().Create(cm).Error
}

func (cm *ChatMessage) Update() error {
	return stores.GetDb().Save(cm).Error
}

func (cm *ChatMessage) Delete() error {
	return stores.GetDb().Delete(cm).Error
}

// GetChatMessageByID retrieves chat message by ID
func GetChatMessageByID(id uuid.UUID) (*ChatMessage, error) {
	var cm ChatMessage
	if err := stores.GetDb().First(&cm, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &cm, nil
}
