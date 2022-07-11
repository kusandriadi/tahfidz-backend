package model

import (
	"time"
)

type User struct {
	Id            int       `json:"id" gorm:"primary_key"`
	CreatedDate   time.Time `json:"createdDate"`
	MarkForDelete bool      `json:"markForDelete"`
	Name          string    `json:"name"`
	Username      string    `json:"username"`
	Password      string    `json:"password"`
	Guardian      string    `json:"guardian"`
	UserPhone     string    `json:"userPhone"`
	GuardianPhone string    `json:"guardianPhone"`
	BirthDate     time.Time `json:"birthDate"`
	City          string    `json:"city"`
	Role          string    `json:"role"`
	LastEducation string    `json:"lastEducation"`
}

type UserCount struct {
	Role  string `json:"role"`
	Total int    `json:"total"`
}
