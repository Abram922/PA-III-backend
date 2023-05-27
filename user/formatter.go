package user

type UserFormatter struct {
	ID    int    `json:"id"`
	Nama  string `json:"nama"`
	Email string `json:"email"`
	Token string `json:"token"`
	Role  int       `json:"role"`
}

func FormatUser(user User, token string, role int) UserFormatter {
	formatter := UserFormatter{
		ID:    user.ID,
		Nama:  user.Nama,
		Email: user.Email,
		Token: token,
		Role:  role,
	}

	return formatter
}
