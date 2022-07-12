package repository

import (
	"tahfidz-backend/model"
	"tahfidz-backend/service"
)

func FetchSubjects() []model.Subject {
	db := service.ConnectToDatabase()
	var subjects []model.Subject

	db.Where("markForDelete = ?", false).Find(&subjects)

	return subjects
}

func FetchSubjectByType(subjectType string) model.Subject {
	db := service.ConnectToDatabase()
	var subject model.Subject

	db.Where("type = ? AND markForDelete = ?", subjectType, false).Find(&subject)

	return subject
}

func FetchSubjectById(id string) model.Subject {
	db := service.ConnectToDatabase()
	var subject model.Subject

	db.Find(&subject, id)

	return subject
}

func FetchSubjectByName(name string) []model.Subject {
	db := service.ConnectToDatabase()
	var subjects []model.Subject

	db.Where("markForDelete = ? AND name LIKE ?", false, "%"+name+"%").Find(&subjects)

	return subjects
}

func CountSubject() model.SubjectCount {
	db := service.ConnectToDatabase()
	var subjectCount model.SubjectCount

	db.Select("DISTINCT ON (role) role", "total").Find(&subjectCount)

	return subjectCount
}
