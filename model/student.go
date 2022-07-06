package model

import "time"

type Student struct {
	ID          int
	Name        string
	birthDate   time.Time
	CreatedDate time.Time
}
