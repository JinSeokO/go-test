package model

import "time"

type TodoModel struct {
	Id int64
	Title string
	Content string
	CreatedAt time.Time
	UpdatedAt time.Time
}