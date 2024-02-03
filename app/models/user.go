package models

import (
	"time"
)

type User struct {
	ID        int `gorm:"primaryKey"`
	FirstName string
	LastName  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserFilter struct {
	User
}
