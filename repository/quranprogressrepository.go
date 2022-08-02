package repository

import (
	"tahfidz-backend/model"
	"tahfidz-backend/service"
)

func FetchQuranProgress() []model.QuranProgress {
	db := service.ConnectToDatabase()
	var quranProgress []model.QuranProgress

	db.Where("markForDelete = ?", false).Find(&quranProgress)

	return quranProgress
}

func FetchQuranProgressByUserId(userId int) model.QuranProgress {
	db := service.ConnectToDatabase()
	var quranProgress model.QuranProgress

	db.Where("userId = ? AND markForDelete = ?", userId, false).Find(&quranProgress)

	return quranProgress
}

func FetchQuranProgressByUserIdAndMethod(userId int, method string) model.QuranProgress {
	db := service.ConnectToDatabase()
	var quranProgress model.QuranProgress

	db.Where("userId = ? AND method = ? AND markForDelete = ?", userId, method, false).Find(&quranProgress)

	return quranProgress
}
