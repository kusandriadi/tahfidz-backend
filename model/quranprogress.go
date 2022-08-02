package model

import (
	"time"
)

type QuranProgress struct {
	Id            int       `json:"id" gorm:"primary_key" gorm:"column:id"`
	CreatedDate   time.Time `json:"createdDate" gorm:"column:createdDate"`
	MarkForDelete bool      `json:"markForDelete" gorm:"column:markForDelete"`
	Surat         string    `json:"surat" gorm:"column:surah"`
	Ayat          string    `json:"ayat" gorm:"column:ayat"`
	Juz           string    `json:"juz" gorm:"column:juz"`
	UserId        int       `json:"userId" gorm:"column:userId"`
	Method        string    `json:"method" gorm:"column:method"`
}
