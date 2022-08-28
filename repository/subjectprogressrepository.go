package repository

import (
	"tahfidz-backend/model"
	"tahfidz-backend/model/enum"
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

func FetchSubjectProgressByUserIdAndSubjectIdAndCreatedDate(userId int, subjectId int, createdDate time.Time) model.SubjectProgress {
	db := service.ConnectToDatabase()
	var subjectProgress model.SubjectProgress

	db.Where("userId = ? AND subjectId = ? AND markForDelete = ? AND date(createdDate) = ?",
		userId, subjectId, false, createdDate.Format("2006-01-02")).Find(&subjectProgress)

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

	db.Where("subjectId = ? AND markForDelete = ? AND date(createdDate) = ?", subjectId, false, time.Now().Format("2006-01-02")).Find(&subjectProgress)

	if len(subjectProgress) == 0 {
		var users = FetchUserByRole(enum.UserRoleEnum().STUDENT)
		var constructSubjectProgress []model.SubjectProgress
		now := time.Now()
		for _, u := range users {
			constructSubjectProgress = append(constructSubjectProgress, model.SubjectProgress{
				CreatedDate:   &now,
				MarkForDelete: false,
				UserId:        u.Id,
				SubjectId:     subjectId})
		}

		db.Create(constructSubjectProgress)
	}
	db.Select("subjectprogress.id, subjectprogress.createdDate, subjectprogress.markForDelete, subjectprogress.userId, "+
		"user.name, subjectprogress.presence, subjectprogress.subjectId").
		Joins("join user on user.id = subjectprogress.userId").
		Where("subjectprogress.subjectId = ? AND subjectprogress.markForDelete = ? AND date(subjectprogress.createdDate) = ?", subjectId, false, time.Now().Format("2006-01-02")).
		Find(&subjectProgress)

	return subjectProgress
}

func FetchSubjectProgressByUserId(userId int) []model.SubjectProgress {
	db := service.ConnectToDatabase()
	var subjectProgress []model.SubjectProgress

	db.Where("userId = ? AND markForDelete = ?", userId, false).Find(&subjectProgress)

	return subjectProgress
}
