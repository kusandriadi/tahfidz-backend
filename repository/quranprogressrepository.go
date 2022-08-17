package repository

import (
	"tahfidz-backend/model"
	"tahfidz-backend/service"
	"time"
)

func FetchQuranProgress() []model.QuranProgress {
	db := service.ConnectToDatabase()
	var quranProgress []model.QuranProgress

	db.Where("markForDelete = ?", false).Find(&quranProgress)

	return quranProgress
}

func FetchQuranProgressByUserId(userId int) []model.QuranProgress {
	db := service.ConnectToDatabase()
	var quranProgress []model.QuranProgress

	db.Where("userId = ? AND markForDelete = ?", userId, false).Find(&quranProgress)

	return quranProgress
}

func FetchQuranProgressByMethod(method string) []model.QuranProgress {
	db := service.ConnectToDatabase()
	var quranProgress []model.QuranProgress

	db.Where("method = ? AND markForDelete = ? AND createdDate = ?", method, false, time.Now()).Find(&quranProgress)

	if len(quranProgress) == 0 {
		var users = FetchUsers()
		var constructQuranProgress []model.QuranProgress
		for _, u := range users {
			constructQuranProgress = append(constructQuranProgress, model.QuranProgress{
				CreatedDate:   time.Now(),
				MarkForDelete: false,
				UserId:        u.Id,
				Method:        method})
		}

		db.Create(constructQuranProgress)
	}

	db.Where("method = ? AND markForDelete = ? AND createdDate = ?", method, false, time.Now()).Find(&quranProgress)

	return quranProgress
}

func FetchQuranProgressByUserIdAndMethod(userId int, method string) []model.QuranProgress {
	db := service.ConnectToDatabase()
	var quranProgress []model.QuranProgress

	db.Where("userId = ? AND method = ? AND markForDelete = ?", userId, method, false).Find(&quranProgress)

	return quranProgress
}

func CountQuranProgressMethod(userId int) []model.QuranProgressMethodCount {
	db := service.ConnectToDatabase()
	var quranProgressMethod []model.QuranProgressMethodCount

	db.Raw("Select Method, COUNT(Method) as total FROM quranprogress GROUP BY Method").Scan(&quranProgressMethod)

	return quranProgressMethod
}

func CurrentQuranProgress(userId int) model.CurrentQuranProgress {
	db := service.ConnectToDatabase()
	var currentQuranProgress model.CurrentQuranProgress
	var quranProgress model.QuranProgress

	result := db.Where("userId = ? AND markForDelete = ?", userId, false).Find(&quranProgress)
	db.Where("userId = ? AND markForDelete = ?", userId, false).Last(&currentQuranProgress)

	currentQuranProgress.TotalSurat = result.RowsAffected

	return currentQuranProgress
}
