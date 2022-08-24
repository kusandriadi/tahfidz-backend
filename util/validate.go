package util

import (
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"strconv"
	"strings"
	"tahfidz-backend/model"
	"tahfidz-backend/model/enum"
	"tahfidz-backend/repository"
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
	regexLetter := regexp.MustCompile(`^[a-zA-Z ]+$`).MatchString
	var isLetter = true
	if user != nil {
		if len(user.Name) > 0 {
			isLetter = regexLetter(user.Name)
		}
	}

	if !isLetter {
		return false, "Nama user harus karakter a-z A-Z."
	}

	if strings.EqualFold("TEACHER", user.Role) {
		if len(user.Name) > 0 && len(user.Username) > 0 &&
			len(user.Password) > 0 && len(user.UserPhone) > 0 && len(user.LastEducation) > 0 {
			passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
			if err != nil {
				return false, "Gagal membuat/mengubah user."
			}
			user.Password = string(passwordHash)
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

func ValidateSubject(subject *model.Subject) (bool, string) {
	regexLetter := regexp.MustCompile(`^[a-zA-Z ]+$`).MatchString
	var isLetter = true
	if subject != nil {
		if len(subject.Name) > 0 {
			isLetter = regexLetter(subject.Name)
		}
	}

	if !isLetter {
		return false, "Nama Subject harus karakter a-z A-Z."
	}

	if strings.EqualFold("KAJIAN", subject.Type) && strings.EqualFold("PELAJARAN", subject.Type) {
		if len(subject.Name) > 0 && len(subject.Book) > 0 &&
			len(subject.Author) > 0 {
			return true, ""
		} else {
			return false, "nama, buku dan pengarang harus dimasukan."
		}
	}

	return false, "Subject type " + subject.Type + " tidak ditemukan."

}

func quranMethodPass(method string) bool {
	if strings.EqualFold(enum.QuranMethodEnum().SABAQ, method) ||
		strings.EqualFold(enum.QuranMethodEnum().MANZIL, method) ||
		strings.EqualFold(enum.QuranMethodEnum().SABAQI, method) {
		return true
	}

	return false
}

func ValidateQuranProgress(quranProgress *model.QuranProgress) (bool, string) {
	if len(quranProgress.Juz) > 0 && len(quranProgress.Ayat) > 0 &&
		len(quranProgress.Surat) > 0 && &quranProgress.UserId != nil &&
		len(quranProgress.Method) > 0 && quranMethodPass(quranProgress.Method) {
		user := repository.FetchUserById(quranProgress.UserId)
		if &user != nil {
			return true, ""
		}

		return false, "Data user tidak ditemukan."
	}

	return false, "Data yang diberikan tidak lengkap."
}

func ValidateSubjectProgress(subjectProgress *model.SubjectProgress) (bool, string) {
	if &subjectProgress.SubjectId != nil &&
		&subjectProgress.UserId != nil {
		user := repository.FetchUserById(subjectProgress.UserId)
		subject := repository.FetchSubjectById(subjectProgress.SubjectId)
		if &user != nil && &subject != nil {
			subjectProgressExisted := repository.FetchSubjectProgressByUserIdAndSubjectId(subjectProgress.UserId,
				subjectProgress.SubjectId)
			if &subjectProgressExisted != nil {
				return false, "Data subject progress sudah ada."
			}
			return true, ""
		}

		return false, "Data user dan subject tidak ditemukan."
	}

	return false, "Data yang diberikan tidak lengkap."
}
