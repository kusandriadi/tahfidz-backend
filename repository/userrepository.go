package repository

import (
	"tahfidz-backend/model"
	"tahfidz-backend/service"
)

func FetchUsers() []model.User {
	db := service.ConnectToDatabase()
	var users []model.User

	db.Where("markForDelete = ?", false).Find(&users)

	return users
}

func FetchUserByUsername(username string) model.User {
	db := service.ConnectToDatabase()
	var user model.User

	db.Where("username = ? AND markForDelete = ?", username, false).Find(&user)

	return user
}

func FetchUserById(id int) model.User {
	db := service.ConnectToDatabase()
	var user model.User

	db.Find(&user, id)

	return user
}

func FetchUserByName(name string) []model.User {
	db := service.ConnectToDatabase()
	var users []model.User

	db.Where("markForDelete = ? AND name LIKE ?", false, "%"+name+"%").Find(&users)

	return users
}

func FetchUserByRole(role string) []model.User {
	db := service.ConnectToDatabase()
	var users []model.User

	db.Where("role = ? AND markForDelete = ?", role, false).Find(&users)

	return users
}

func CountUser() model.UserCount {
	db := service.ConnectToDatabase()
	var userCount model.UserCount

	db.Select("DISTINCT ON (role) role", "total").Find(&userCount)

	return userCount
}
