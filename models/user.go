package models

import (
	"nexcommerce/stores"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

// Receiver function to get user by ID
func GetUserByID(id uint) (*User, error) {
	var user User
	if err := stores.GetDb().First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Receiver function to create a new user
func (u *User) CreateUser() error {
	return stores.GetDb().Create(u).Error
}

func (u *User) Follow(followedUser *User) error {
	err := stores.GetDb().Model(u).Association("Following").Append(&followedUser)
	if err != nil {
		return err
	}
	// Update the followed user's followers list
	err = stores.GetDb().Model(&followedUser).Association("Followers").Append(u)
	return err
}

func (u *User) Unfollow(followedUser *User) error {
	err := stores.GetDb().Model(u).Association("Following").Delete(&followedUser)
	if err != nil {
		return err
	}
	// Update the followed user's followers list
	err = stores.GetDb().Model(&followedUser).Association("Followers").Delete(u)
	return err
}
