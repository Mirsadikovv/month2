package model

import "time"

type User struct {
	Id        string
	Name      string
	Email     string
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
