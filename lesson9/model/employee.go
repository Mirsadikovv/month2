package model

import "time"

type Employee struct {
	Id             string
	FirstName      string
	LastName       string
	DepartmentName string
	CreatedAt      time.Time
}
