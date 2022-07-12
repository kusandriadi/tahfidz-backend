package model

import (
	"time"
)

type User struct {
	Id            int        `json:"id" gorm:"primary_key" gorm:"column:id"`
	CreatedDate   time.Time  `json:"createdDate" gorm:"column:createdDate"`
	MarkForDelete bool       `json:"markForDelete" gorm:"column:markForDelete"`
	Name          string     `json:"name" gorm:"column:name"`
	Username      string     `json:"username" gorm:"column:username"`
	Password      string     `json:"password,omitempty" gorm:"column:password"`
	Guardian      string     `json:"guardian,omitempty" gorm:"column:guardian"`
	UserPhone     string     `json:"userPhone,omitempty" gorm:"column:userPhone"`
	GuardianPhone string     `json:"guardianPhone,omitempty" gorm:"column:guardianPhone"`
	BirthDate     *time.Time `json:"birthDate,omitempty" gorm:"column:birthDate"`
	City          string     `json:"city,omitempty" gorm:"column:city"`
	Role          string     `json:"role" gorm:"column:role"`
	LastEducation string     `json:"lastEducation,omitempty" gorm:"column:lastEducation"`
}

type UserCount struct {
	Role  string `json:"role"`
	Total int    `json:"total"`
}

func (User) TableName() string {
	return "user"
}
