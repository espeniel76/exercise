package models

import (
	"time"

	"gorm.io/gorm"
)

// User represents the users table
type User struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Email     string         `gorm:"type:varchar(255);not null;unique" json:"email"`
	Username  string         `gorm:"type:varchar(50);not null;unique" json:"username"`
	Password  string         `gorm:"type:varchar(255);not null" json:"password"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;column:deleted_at" json:"-"`
}

// TableName specifies the table name for the User model
func (User) TableName() string {
	return "users"
}
