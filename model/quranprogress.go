package model

import (
	"time"
)

type QuranProgress struct {
	Id            int        `json:"id,omitempty" gorm:"primary_key" gorm:"column:id"`
	CreatedDate   *time.Time `json:"createdDate,omitempty" gorm:"column:createdDate"`
	MarkForDelete bool       `json:"markForDelete,omitempty" gorm:"column:markForDelete"`
	Surat         string     `json:"surat,omitempty" gorm:"column:surat"`
	Ayat          int        `json:"ayat,omitempty" gorm:"column:ayat"`
	Juz           int        `json:"juz,omitempty" gorm:"column:juz"`
	UserId        int        `json:"userId,omitempty" gorm:"column:userId"`
	Name          string     `json:"name,omitempty" gorm:"->"`
	Method        string     `json:"method,omitempty" gorm:"column:method"`
}

type QuranProgressMethodCount struct {
	Method string `json:"method,omitempty"`
	Total  int    `json:"total,omitempty"`
}

type AllUserQuranProgress struct {
	UserId int    `json:"userId,omitempty"`
	Name   string `json:"name,omitempty"`
	Total  int    `json:"total,omitempty"`
}

type CurrentQuranProgress struct {
	UserId     int    `json:"userId,omitempty"`
	TotalSurat int    `json:"totalSurat,omitempty"`
	Surat      string `json:"surat,omitempty"`
	Ayat       int    `json:"ayat,omitempty"`
}

func (QuranProgress) TableName() string {
	return "quranprogress"
}
