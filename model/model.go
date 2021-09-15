package model

// type User struct {
// 	ID          uint `json:"id,omitempty"`
// 	UserProfile UserProfile
// }

// type UserProfile struct {
// 	ID     uint `json:"id,omitempty"`
// 	UserID uint `json:"userId"`
// }

type Objective struct {
	ID          uint    `json:"id,omitempty"`
	Title       *string `json:"title" gorm:"not null"`
	Description *string `json:"description"`
	// KeyResults  []KeyResult `json:"keyResults"`
}

// type KeyResult struct {
// 	ID          uint    `json:"id,omitempty"`
// 	Title       *string `json:"title" gorm:"not null"`
// 	Description *string `json:"description"`
// }
