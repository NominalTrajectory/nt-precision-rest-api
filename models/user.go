package models

import "time"

// TODO: Add validation

type User struct {
	ID        uint      `json:"id,omitempty"`
	Name      string    `json:"name" gorm:"not null"`
	Email     string    `json:"email" gorm:"unique"`
	Pwd       string    `json:"pwd" gorm:"not null"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Credentials struct {
	Email string
	Pwd   string `json:"pwd" gorm:"not null"`
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
