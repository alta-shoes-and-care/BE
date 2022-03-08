package auth

type ResponseLogin struct {
	Token   string `json:"token"`
	IsAdmin bool   `json:"is_admin"`
}

func ToResponseLogin(token string, isAdmin bool) ResponseLogin {
	return ResponseLogin{
		Token:   token,
		IsAdmin: isAdmin,
	}
}
