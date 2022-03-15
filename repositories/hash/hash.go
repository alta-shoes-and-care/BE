package hash

import "golang.org/x/crypto/bcrypt"

func CheckPasswordHash(hashed, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err == nil
}
