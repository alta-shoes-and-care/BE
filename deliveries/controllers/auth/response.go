package auth

import U "final-project/entities/user"

type ResponseLogin struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Token   string `json:"token"`
	IsAdmin bool   `json:"is_admin"`
}

func ToResponseLogin(user U.Users, token string, isAdmin bool) ResponseLogin {
	return ResponseLogin{
		ID:      user.ID,
		Name:    user.Name,
		Token:   token,
		IsAdmin: isAdmin,
	}
}
