package dto

import "time"

type TodoDto struct {
	Task     string    `json:"task"`
	TodoDate time.Time `json:"todo_date"`
}
