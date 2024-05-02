package model

import "time"

type Users struct {
	Id         string
	FirstName  string
	TelegramId int
	Status     string
	StartTime  time.Time
}
