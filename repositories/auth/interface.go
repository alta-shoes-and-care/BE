package auth

import U "final-project/entities/user"

type Auth interface {
	Login(email, password string) (U.Users, error)
}
