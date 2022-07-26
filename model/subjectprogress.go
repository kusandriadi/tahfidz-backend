package model

import (
	"time"
)

type SubjectProgress struct {
	Id            int       `json:"id" gorm:"primary_key" gorm:"column:id"`
	CreatedDate   time.Time `json:"createdDate" gorm:"column:createdDate"`
	MarkForDelete bool      `json:"markForDelete" gorm:"column:markForDelete"`
	UserId        string    `json:"userId" gorm:"column:userId"`
	SubjectId     string    `json:"subjectId" gorm:"column:subjectId"`
}
