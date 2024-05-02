package model

import "time"

type Session struct {
	Id         string
	UserId     string
	Data       string
	ExiparedAt time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
