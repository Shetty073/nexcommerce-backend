package models

import (
	"nexcommerce/stores"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Username     string     `gorm:"unique;not null;index"`
	Email        string     `gorm:"unique;not null;index"`
	Password     string     `gorm:"not null"`
	FirstName    string     `gorm:"type:varchar(50)"`
	LastName     string     `gorm:"type:varchar(50)"`
	DateOfBirth  string     `gorm:"type:date"`
	Gender       string     `gorm:"type:varchar(12)"`
	MobileNumber string     `gorm:"type:varchar(15);index"`
	IsStaff      bool       `gorm:"default:false;index"`
	IsCustomer   bool       `gorm:"default:false;index"`
	LastLoginAt  *time.Time `gorm:"type:timestamp;index"`
	Status       string     `gorm:"type:enum('active','inactive','banned');default:'active';index"`
	CreatedAt    *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index"`
	UpdatedAt    *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;index"`
	DeletedAt    gorm.DeletedAt
}

// GetUserByID retrieves a user by ID
func GetUserByID(id uuid.UUID) (*User, error) {
	var user User
	if err := stores.GetDb().First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateUser creates a new user
func (u *User) CreateUser() error {
	u.ID = uuid.New()
	return stores.GetDb().Create(u).Error
}

// UpdateUser updates an existing user
func (u *User) UpdateUser() error {
	return stores.GetDb().Save(u).Error
}

// DeleteUser marks the user as inactive
func (u *User) DeleteUser() error {
	return stores.GetDb().Delete(u).Error
}

// GetAllUsers retrieves all users
func GetAllUsers() ([]User, error) {
	var users []User
	if err := stores.GetDb().Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// GetUserByEmail retrieves a user by email
func GetUserByEmail(email string) (*User, error) {
	var user User
	if err := stores.GetDb().Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByUsername retrieves a user by username
func GetUserByUsername(username string) (*User, error) {
	var user User
	if err := stores.GetDb().Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetAllSupportTickets retrieves all support tickets for a user, regardless of role (customer, assigned, etc.)
func (u *User) GetAllSupportTickets() ([]SupportTicket, error) {
	var supportTickets []SupportTicket
	query := stores.GetDb()

	if u.IsCustomer {
		query = query.Where("customer_id = ?", u.ID)
	} else {
		query = query.Where("assigned_to = ? OR created_by = ? OR updated_by = ?", u.ID, u.ID, u.ID)
	}

	if err := query.Find(&supportTickets).Error; err != nil {
		return nil, err
	}

	return supportTickets, nil
}
