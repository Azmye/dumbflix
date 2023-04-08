package models

import "time"

type Movie struct {
	ID          int       `json:"id"`
	Title       string    `json:"title" gorm:"type VARCHAR(255)"`
	Thumbnail   string    `json:"thumbnail" gorm:"type VARCHAR(255)"`
	Year        int       `json:"year" gorm:"type VARCHAR(255)"`
	CategoryID  int       `json:"category_id"`
	Category    Category  `json:"category"`
	Description string    `json:"description" gorm:"type VARCHAR(255)"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}
