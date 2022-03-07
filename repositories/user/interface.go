package user

import U "final-project/entities/user"

type User interface {
	Create(newUser U.Users) (U.Users, error)
	Get(userID uint) (U.Users, error)
	Update(userUpdate U.Users) (U.Users, error)
	Delete(userID uint) error
}
