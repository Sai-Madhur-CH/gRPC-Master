package models

import (
	"time"
)

// Default struct
type Default struct {
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// User struct
type User struct {
	ID           string `json:"id,omitempty" gorm:"primaryKey"`
	Name         string `json:"name,omitempty" gorm:"type:varchar;not null;"`
	Email        string `json:"email,imitempty" gorm:"type:varchar;not null;"`
	Phone        string `json:"phone,imitempty" gorm:"type:varchar;not null;"`
	HashPassword string `json:"hash_password,imitempty" gorm:"type:varchar;not null;"`
	Default
}

// TableName - to set schema
func (u *User) TableName() string {
	return "dbo.users"
}
