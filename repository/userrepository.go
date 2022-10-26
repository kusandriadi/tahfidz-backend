package repository

import (
	"tahfidz-backend/model"
	"tahfidz-backend/service"
)

func FetchUsers() []model.User {
	db := service.ConnectToDatabase()
	var users []model.User

	db.Select("id", "createdDate", "markForDelete", "name", "address", "username", "guardian", "userPhone",
		"guardianPhone", "birthDate", "birthPlace", "city", "role", "lastEducation", "grade").Where("markForDelete = ?", false).Find(&users)

	return constructStructs(users)
}

func FetchUserByUsername(username string, login bool) model.User {
	db := service.ConnectToDatabase()
	var user model.User

	if login {
		db.Where("username = ? AND markForDelete = ?", username, false).Find(&user)
	} else {
		db.Select("id", "createdDate", "markForDelete", "name", "address", "username", "guardian", "userPhone",
			"guardianPhone", "birthDate", "birthPlace", "city", "role", "lastEducation", "grade").Where("username = ? AND markForDelete = ?", username, false).Find(&user)
	}
	return constructUser(user)
}

func FetchUserById(id int) model.User {
	db := service.ConnectToDatabase()
	var user model.User

	db.Select("id", "createdDate", "markForDelete", "name", "address", "username", "guardian", "userPhone",
		"guardianPhone", "birthDate", "birthPlace", "city", "role", "lastEducation", "grade").Find(&user, id)

	return constructUser(user)
}

func FetchUserByName(name string) []model.User {
	db := service.ConnectToDatabase()
	var users []model.User

	db.Select("id", "createdDate", "markForDelete", "name", "address", "username", "guardian", "userPhone",
		"guardianPhone", "birthDate", "birthPlace", "city", "role", "lastEducation", "grade").Where("markForDelete = ? AND name LIKE ?", false, "%"+name+"%").Find(&users)

	return constructStructs(users)
}

func FetchUserByNameAndRole(name string, role string) []model.User {
	db := service.ConnectToDatabase()
	var users []model.User

	db.Select("id", "createdDate", "markForDelete", "name", "address", "username", "guardian", "userPhone",
		"guardianPhone", "birthDate", "birthPlace", "city", "role", "lastEducation", "grade").Where("markForDelete = ? AND role = ? AND name LIKE ?", false, role, "%"+name+"%").Find(&users)

	return constructStructs(users)
}

func FetchUserByRole(role string) []model.User {
	db := service.ConnectToDatabase()
	var users []model.User

	db.Select("id", "createdDate", "markForDelete", "name", "address", "username", "guardian", "userPhone",
		"guardianPhone", "birthDate", "birthPlace", "city", "role", "lastEducation", "grade").Where("role = ? AND markForDelete = ?", role, false).Find(&users)

	return constructStructs(users)
}

func CountUser() []model.UserCount {
	db := service.ConnectToDatabase()
	var userCount []model.UserCount

	db.Raw("Select Role, COUNT(Role) as total FROM user WHERE markForDelete = false GROUP BY role").Scan(&userCount)

	return userCount
}

func constructStructs(users []model.User) []model.User {
	for index, u := range users {
		users[index] = constructUser(u)
	}

	return users
}

func constructUser(user model.User) model.User {
	if user.BirthDate != nil {
		user.UserBirthDate = user.BirthDate.Format("02-01-2006")
		user.BirthDate = nil
	}

	return user
}
