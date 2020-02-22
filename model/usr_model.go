package model

import "time"

// User struct
type User struct {
	UserID    uint64    `gorm:"column:user_id;primary_key;auto_increment;unique" json:"user_id"`
	Username  string    `gorm:"column:username;size:100;unique;not null" json:"username"`
	Email     string    `gorm:"column:email;size:100;unique;not null" json:"email"`
	Password  string    `gorm:"column:password;size:255;not null" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeleteAt  time.Time `gorm:"default:null" json:"delete_at"`
}
