package models

import "time"

// TODO: Add validation

type User struct {
	ID        uint      `json:"id,omitempty" gorm:"primaryKey;not null"`
	Name      string    `json:"name" gorm:"not null;"`
	Email     string    `json:"email" gorm:"not null;unique"`
	Pwd       string    `json:"-" gorm:"not null"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Credentials struct {
	Email string
	Pwd   string
}

type UserProfile struct {
	ID           uint
	Nickname     string
	Greeting     string
	ProfilePhoto string
}
