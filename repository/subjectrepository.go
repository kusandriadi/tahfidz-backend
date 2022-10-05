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

func FetchSubjectByType(subjectType string) []model.Subject {
	db := service.ConnectToDatabase()
	var subjects []model.Subject

	db.Where("type = ? AND markForDelete = ?", subjectType, false).Find(&subjects)

	return subjects
}

func FetchSubjectById(id int) model.Subject {
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

func CountSubject() []model.SubjectCount {
	db := service.ConnectToDatabase()
	var subjectCount []model.SubjectCount

	db.Raw("Select Type, COUNT(Type) as total FROM subject WHERE markForDelete = false GROUP BY Type").Scan(&subjectCount)

	return subjectCount
}
