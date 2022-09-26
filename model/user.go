package model

import (
	"time"
)

type User struct {
	Id            int        `json:"id,omitempty" gorm:"column:id;primary_key"`
	CreatedDate   *time.Time `json:"createdDate,omitempty" gorm:"column:createdDate"`
	MarkForDelete bool       `json:"markForDelete,omitempty" gorm:"column:markForDelete"`
	Name          string     `json:"name,omitempty" gorm:"column:name"`
	Username      string     `json:"username,omitempty" gorm:"column:username"`
	Address       string     `json:"address,omitempty" gorm:"column:address"`
	Password      string     `json:"password,omitempty" gorm:"column:password"`
	Guardian      string     `json:"guardian,omitempty" gorm:"column:guardian"`
	UserPhone     string     `json:"userPhone,omitempty" gorm:"column:userPhone"`
	GuardianPhone string     `json:"guardianPhone,omitempty" gorm:"column:guardianPhone"`
	BirthDate     *time.Time `json:"birthDate,omitempty" gorm:"column:birthDate"`
	City          string     `json:"city,omitempty" gorm:"column:city"`
	Role          string     `json:"role,omitempty" gorm:"column:role"`
	LastEducation string     `json:"lastEducation,omitempty" gorm:"column:lastEducation"`
}

type UserCount struct {
	Role  string `json:"role"`
	Total int    `json:"total"`
}

func (User) TableName() string {
	return "user"
}
