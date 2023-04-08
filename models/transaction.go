package models

import "time"

type Transaction struct {
	ID        int       `json:"id"`
	StartDate time.Time `json:"start_date"`
	DueDate   time.Time `json:"due_date"`
	UserID    int       `json:"user_id"`
	User      User      `json:"user"`
	Status    string    `json:"status" gorm:"type: VARCHAR(25)"`
}
