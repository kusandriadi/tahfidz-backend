package repository

import (
	"tahfidz-backend/model"
	"tahfidz-backend/service"
	"time"
)

func FetchSubjectProgress() []model.SubjectProgress {
	db := service.ConnectToDatabase()
	var subjectProgress []model.SubjectProgress

	db.Where("markForDelete = ?", false).Find(&subjectProgress)

	return subjectProgress
}

func FetchSubjectProgressByUserIdAndSubjectId(userId int, subjectId int) []model.SubjectProgress {
	db := service.ConnectToDatabase()
	var subjectProgress []model.SubjectProgress

	db.Where("userId = ? AND subjectId = ? AND markForDelete = ?", userId, subjectId, false).Find(&subjectProgress)

	return subjectProgress
}

func FetchSubjectProgressBySubjectId(subjectId int) []model.SubjectProgress {
	db := service.ConnectToDatabase()
	var subjectProgress []model.SubjectProgress
	var subject model.Subject

	result := db.Where("id = ? AND markForDelete = ?", subjectId, false).Find(&subject)
	if result.RowsAffected == 0 {
		return subjectProgress
	}

	db.Where("subjectId = ? AND markForDelete = ? AND createdDate = ?", subjectId, false, time.Now()).Find(&subjectProgress)

	if len(subjectProgress) == 0 {
		var users = FetchUsers()
		var constructSubjectProgress []model.SubjectProgress
		for _, u := range users {
			constructSubjectProgress = append(constructSubjectProgress, model.SubjectProgress{
				CreatedDate:   time.Now(),
				MarkForDelete: false,
				UserId:        u.Id,
				SubjectId:     subjectId})
		}

		db.Create(constructSubjectProgress)
	}

	db.Where("subjectId = ? AND markForDelete = ? AND createdDate = ?", subjectId, false, time.Now()).Find(&subjectProgress)

	return subjectProgress
}

func FetchSubjectProgressByUserId(userId int) []model.SubjectProgress {
	db := service.ConnectToDatabase()
	var subjectProgress []model.SubjectProgress

	db.Where("userId = ? AND markForDelete = ?", userId, false).Find(&subjectProgress)

	return subjectProgress
}
