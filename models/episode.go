package models

import "time"

type Episode struct {
	ID        int       `json:"id,"`
	Title     string    `json:"title" gorm:"type: VARCHAR(255)"`
	Thumbnail string    `json:"thumbnail" gorm:"type: VARCHAR(255)"`
	VideoLink string    `json:"video_link" gorm:"type: VARCHAR(255)"`
	MovieID   int       `json:"movie_id"`
	Movie     Movie     `json:"movie"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
