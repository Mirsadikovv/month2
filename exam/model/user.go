package model

import "time"

type User struct {
	Id        string
	Name      string
	Age       int
	Email     string
	Phone     string
	CreatedAt time.Time
}
