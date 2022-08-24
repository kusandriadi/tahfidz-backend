package model

import (
	"time"
)

type QuranProgress struct {
	Id            int        `json:"id,omitempty" gorm:"primary_key" gorm:"column:id"`
	CreatedDate   *time.Time `json:"createdDate,omitempty" gorm:"column:createdDate"`
	MarkForDelete bool       `json:"markForDelete,omitempty" gorm:"column:markForDelete"`
	Surat         string     `json:"surat,omitempty" gorm:"column:surat"`
	Ayat          string     `json:"ayat,omitempty" gorm:"column:ayat"`
	Juz           string     `json:"juz,omitempty" gorm:"column:juz"`
	UserId        int        `json:"userId,omitempty" gorm:"column:userId"`
	Method        string     `json:"method,omitempty" gorm:"column:method"`
}

type QuranProgressMethodCount struct {
	Method string `json:"method"`
	Total  int    `json:"total"`
}

type CurrentQuranProgress struct {
	UserId     string `json:"userId"`
	TotalSurat int64  `json:"totalSurat"`
	Surat      int    `json:"surat"`
	Ayat       int    `json:"ayat"`
}
