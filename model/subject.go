package model

import (
	"time"
)

type Subject struct {
	Id            int        `json:"id,omitempty" gorm:"primary_key" gorm:"column:id"`
	CreatedDate   *time.Time `json:"createdDate,omitempty" gorm:"column:createdDate"`
	MarkForDelete bool       `json:"markForDelete,omitempty" gorm:"column:markForDelete"`
	Name          string     `json:"name,omitempty" gorm:"column:name"`
	Book          string     `json:"book,omitempty" gorm:"column:book"`
	Author        string     `json:"author,omitempty" gorm:"column:author"`
	Type          string     `json:"type,omitempty" gorm:"column:type"`
	Duration      int64      `json:"duration,omitempty" gorm:"column:duration"`
}

type SubjectCount struct {
	Type  string `json:"type"`
	Total int    `json:"total"`
}
