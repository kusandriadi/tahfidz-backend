package enum

func UserRoleEnum() *userRole {
	return &userRole{
		TEACHER: "TEACHER",
		STUDENT: "STUDENT",
		ADMIN:   "ADMIN",
	}
}

type userRole struct {
	TEACHER string
	STUDENT string
	ADMIN   string
}
