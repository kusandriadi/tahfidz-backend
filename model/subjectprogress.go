package model

import (
	"time"
)

type SubjectProgress struct {
	Id            int        `json:"id,omitempty" gorm:"primary_key" gorm:"column:id"`
	CreatedDate   *time.Time `json:"createdDate,omitempty" gorm:"column:createdDate"`
	MarkForDelete bool       `json:"markForDelete,omitempty" gorm:"column:markForDelete"`
	UserId        int        `json:"userId,omitempty" gorm:"column:userId"`
	SubjectId     int        `json:"subjectId,omitempty" gorm:"column:subjectId"`
	Presence      bool       `json:"presence,omitempty" gorm:"column:presence"`
}
