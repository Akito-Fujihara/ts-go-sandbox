package model

import "time"

type Todo struct {
	ID        uint      `json:"id  param:"id""`
	Name      string    `json:"name"`
	Task      string    `json:"task"`
	Done      bool      `json:"done" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
