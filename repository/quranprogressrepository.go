package repository

import (
	"gorm.io/gorm"
	"strconv"
	"tahfidz-backend/model"
	"tahfidz-backend/model/enum"
	"tahfidz-backend/service"
	"time"
)

func FetchQuranProgress() []model.QuranProgress {
	db := service.ConnectToDatabase()
	var quranProgress []model.QuranProgress

	db.Where("markForDelete = ?", false).Find(&quranProgress)

	return quranProgress
}

func FetchQuranProgressById(id int) model.QuranProgress {
	db := service.ConnectToDatabase()
	var quranProgress model.QuranProgress

	db.Where("id = ? AND markForDelete = ?", id, false).Find(&quranProgress)

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

	db.Where("method = ? AND markForDelete = ? AND date(createdDate) = ?", method, false, time.Now().Format("2006-01-02")).
		Find(&quranProgress)

	if len(quranProgress) == 0 {
		var users = FetchUserByRole(enum.UserRoleEnum().STUDENT)
		now := time.Now()

		createSabaq(users, now, db)
		createSabaqi(users, now, db)
		createManzil(users, now, db)
	}

	db.Select("quranprogress.id, quranprogress.createdDate, quranprogress.markForDelete, quranprogress.userId, "+
		"user.name, quranprogress.surat, quranprogress.ayat, quranprogress.juz, quranprogress.method").
		Joins("join user on user.id = quranprogress.userId").
		Where("quranprogress.method = ? AND quranprogress.markForDelete = ? AND date(quranprogress.createdDate) = ?", method, false, time.Now().Format("2006-01-02")).
		Find(&quranProgress)

	return quranProgress
}

func createSabaq(users []model.User, now time.Time, db *gorm.DB) {
	var constructQuranProgress []model.QuranProgress

	for _, u := range users {
		constructQuranProgress = append(constructQuranProgress, model.QuranProgress{
			CreatedDate:   &now,
			MarkForDelete: false,
			UserId:        u.Id,
			Method:        enum.QuranMethodEnum().SABAQ})
	}

	db.Create(constructQuranProgress)
}

func createSabaqi(users []model.User, now time.Time, db *gorm.DB) {
	var constructQuranProgress []model.QuranProgress

	for _, u := range users {
		constructQuranProgress = append(constructQuranProgress, model.QuranProgress{
			CreatedDate:   &now,
			MarkForDelete: false,
			UserId:        u.Id,
			Method:        enum.QuranMethodEnum().SABAQI})
	}

	db.Create(constructQuranProgress)
}

func createManzil(users []model.User, now time.Time, db *gorm.DB) {
	var constructQuranProgress []model.QuranProgress

	for _, u := range users {
		constructQuranProgress = append(constructQuranProgress, model.QuranProgress{
			CreatedDate:   &now,
			MarkForDelete: false,
			UserId:        u.Id,
			Method:        enum.QuranMethodEnum().MANZIL})
	}

	db.Create(constructQuranProgress)
}

func FetchQuranProgressByUserIdAndMethodAndCreatedDate(userId int, method string, now time.Time) []model.QuranProgress {
	db := service.ConnectToDatabase()
	var quranProgress []model.QuranProgress

	db.Where("userId = ? AND method = ? AND markForDelete = ? AND date(createdDate) = ?", userId, method, false, time.Now().Format("2006-01-02")).
		Find(&quranProgress)

	return quranProgress
}

func CountQuranProgressMethod(userId int) []model.QuranProgressMethodCount {
	db := service.ConnectToDatabase()
	var quranProgressMethod []model.QuranProgressMethodCount

	db.Raw("Select Method, COUNT(Method) as total FROM quranprogress WHERE userId = " + strconv.Itoa(userId) + " GROUP BY Method").Scan(&quranProgressMethod)

	return quranProgressMethod
}

func CurrentQuranProgress(userId int) model.CurrentQuranProgress {
	db := service.ConnectToDatabase()
	var currentQuranProgress model.CurrentQuranProgress
	var quranProgress []model.QuranProgress

	db.Raw("SELECT id, createdDate, surat, ayat FROM quranprogress WHERE userId = " + strconv.Itoa(userId) + " AND markForDelete = false AND surat != '' " +
		"ORDER BY id desc").Scan(&quranProgress)

	if len(quranProgress) > 0 {
		currentQuranProgress.TotalSurat = len(quranProgress)
		currentQuranProgress.Surat = quranProgress[0].Surat
		currentQuranProgress.Ayat = quranProgress[0].Ayat
		currentQuranProgress.UserId = userId
	}

	return currentQuranProgress
}

func GetAllQuranProgress() []model.AllUserQuranProgress {
	db := service.ConnectToDatabase()
	var allUserQuranProgress []model.AllUserQuranProgress

	db.Raw("SELECT quranprogress.userId as UserId, user.name as Name, COUNT(quranprogress.userId) as Total FROM quranprogress " +
		"JOIN user ON user.id = quranprogress.userId " +
		"WHERE quranprogress.markForDelete = false AND quranprogress.surat != '' " +
		"GROUP BY quranprogress.userId").Scan(&allUserQuranProgress)

	return allUserQuranProgress
}
