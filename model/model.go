package model

type Objective struct {
	ID          int     `json:"id,omitempty"`
	Title       *string `json:"title" gorm:"not null"`
	Description string  `json:"description"`
}
