package util

import (
	"strconv"
	"strings"
	"tahfidz-backend/model"
)

func IsNumber(parameter string) bool {
	_, err := strconv.Atoi(parameter)

	if err == nil {
		return true
	}

	return false
}

func Role(role string) bool {
	if strings.EqualFold("TEACHER", role) || strings.EqualFold("ADMIN", role) ||
		strings.EqualFold("STUDENT", role) {
		return true
	}

	return false
}

func ValidateUser(user *model.User) (bool, string) {
	if strings.EqualFold("TEACHER", user.Role) {
		if len(user.Name) > 0 && len(user.Username) > 0 &&
			len(user.Password) > 0 && len(user.UserPhone) > 0 && len(user.LastEducation) > 0 {
			//user.Password =
			return true, ""
		} else {
			return false, "nama, username, password, nomor handphone dan pendidikan terakhir harus di isi."
		}
	}

	if strings.EqualFold("STUDENT", user.Role) {
		if len(user.Name) > 0 && len(user.UserPhone) > 0 && len(user.Guardian) > 0 && len(user.GuardianPhone) > 0 &&
			&user.BirthDate != nil && len(user.City) > 0 {
			return true, ""
		} else {
			return false, "nama, nomor handphone, wali, nomor handphone wali, tanggal lahir dan kota asal harus di isi."
		}
	}

	if strings.EqualFold("ADMIN", user.Role) {
		if len(user.Name) > 0 && len(user.Username) > 0 && len(user.Password) > 0 {
			return true, ""
		} else {
			return false, "nama, username dan password harus di isi."
		}
	}

	return false, "Role " + user.Role + " tidak ditemukan."
}

func ValidateSubject() {

}
