package models

import "time"

// TODO: Add validation

type User struct {
	ID        uint      `json:"id,omitempty" gorm:"primaryKey;not null"`
	Name      string    `json:"name" gorm:"not null;"`
	Email     string    `json:"email" gorm:"not null;unique"`
	Password  string    `json:"password" gorm:"not null"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Credentials struct {
	Email    string
	Password string `json:"password" gorm:"not null"`
}

type UserProfile struct {
	ID           uint
	Nickname     string
	Greeting     string
	ProfilePhoto string
}

type Token struct {
	Token string `json:"token"`
}
