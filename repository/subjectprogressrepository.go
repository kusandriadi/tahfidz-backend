package repository

import (
	"tahfidz-backend/model"
	"tahfidz-backend/service"
)

func FetchSubjectProgress() []model.SubjectProgress {
	db := service.ConnectToDatabase()
	var subjectProgress []model.SubjectProgress

	db.Where("markForDelete = ?", false).Find(&subjectProgress)

	return subjectProgress
}

func FetchSubjectProgressByUserIdAndSubjectId(userId int, subjectId int) model.SubjectProgress {
	db := service.ConnectToDatabase()
	var subjectProgress model.SubjectProgress

	db.Where("userId = ? AND subjectId = ? AND markForDelete = ?", userId, subjectId, false).Find(&subjectProgress)

	return subjectProgress
}

func FetchSubjectProgressBySubjectId(subjectId int) model.SubjectProgress {
	db := service.ConnectToDatabase()
	var subjectProgress model.SubjectProgress

	db.Where("subjectId = ? AND markForDelete = ?", subjectId, false).Find(&subjectProgress)

	return subjectProgress
}

func FetchSubjectProgressByUserId(userId int) []model.SubjectProgress {
	db := service.ConnectToDatabase()
	var subjectProgress []model.SubjectProgress

	db.Where("userId = ? AND markForDelete = ?", userId, false).Find(&subjectProgress)

	return subjectProgress
}
