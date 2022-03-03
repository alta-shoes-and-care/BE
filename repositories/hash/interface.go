package hash

type Hash interface {
	HashPassword(password string) (string, error)
	CheckPassword(password string, hash string) bool
}
