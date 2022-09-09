package repository

import (
	"tahfidz-backend/model"
	"tahfidz-backend/service"
)

func FetchUsers() []model.User {
	db := service.ConnectToDatabase()
	var users []model.User

	db.Select("id", "createdDate", "markForDelete", "name", "username", "guardian", "userPhone",
		"guardianPhone", "birthDate", "city", "role", "lastEducation").Where("markForDelete = ?", false).Find(&users)

	return users
}

func FetchUserByUsername(username string, login bool) model.User {
	db := service.ConnectToDatabase()
	var user model.User

	if login {
		db.Where("username = ? AND markForDelete = ?", username, false).Find(&user)
	} else {
		db.Select("id", "createdDate", "markForDelete", "name", "username", "guardian", "userPhone",
			"guardianPhone", "birthDate", "city", "role", "lastEducation").Where("username = ? AND markForDelete = ?", username, false).Find(&user)
	}
	return user
}

func FetchUserById(id int) model.User {
	db := service.ConnectToDatabase()
	var user model.User

	db.Select("id", "createdDate", "markForDelete", "name", "username", "guardian", "userPhone",
		"guardianPhone", "birthDate", "city", "role", "lastEducation").Find(&user, id)

	return user
}

func FetchUserByName(name string) []model.User {
	db := service.ConnectToDatabase()
	var users []model.User

	db.Select("id", "createdDate", "markForDelete", "name", "username", "guardian", "userPhone",
		"guardianPhone", "birthDate", "city", "role", "lastEducation").Where("markForDelete = ? AND name LIKE ?", false, "%"+name+"%").Find(&users)

	return users
}

func FetchUserByNameAndRole(name string, role string) []model.User {
	db := service.ConnectToDatabase()
	var users []model.User

	db.Select("id", "createdDate", "markForDelete", "name", "username", "guardian", "userPhone",
		"guardianPhone", "birthDate", "city", "role", "lastEducation").Where("markForDelete = ? AND role = ? AND name LIKE ?", false, role, "%"+name+"%").Find(&users)

	return users
}

func FetchUserByRole(role string) []model.User {
	db := service.ConnectToDatabase()
	var users []model.User

	db.Select("id", "createdDate", "markForDelete", "name", "username", "guardian", "userPhone",
		"guardianPhone", "birthDate", "city", "role", "lastEducation").Where("role = ? AND markForDelete = ?", role, false).Find(&users)

	return users
}

func CountUser() []model.UserCount {
	db := service.ConnectToDatabase()
	var userCount []model.UserCount

	db.Raw("Select Role, COUNT(Role) as total FROM user GROUP BY role").Scan(&userCount)

	return userCount
}
