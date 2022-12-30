package entity

import (
	"time"
)

type Todo struct {
	ID         string    `json:"id"`
	Task       string    `json:"task"`
	TodoDate   time.Time `json:"todo_date"`
	CreateDate time.Time `json:"create_date"`
}

// func (t *Todo) BeforeCreate(tx *gorm.DB) error {
// 	t.ID = uuid.New().String()
// 	t.CreateDate = time.Now()
// 	return nil
// }
