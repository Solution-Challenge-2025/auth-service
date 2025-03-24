package models

import (
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"-"`
	Role     Role   `gorm:"default:user" json:"role"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = 0 // Assuming auto-incrementing ID starts from 0
	return nil
}
