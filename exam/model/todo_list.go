package model

import "time"

type TodoList struct {
	TaskId       string
	TaskName     string
	TaskTime     time.Time
	TaskComment  string
	TaskPriority string
	UserId       string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
