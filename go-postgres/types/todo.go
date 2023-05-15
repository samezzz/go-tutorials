package types

import (
	"time"
)

type Todo struct {
	ID         uint      `json:"id" gorm:"primary_key"`
	Title      string    `json:"title"`
	IsFinished bool      `json:"isFinished"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
