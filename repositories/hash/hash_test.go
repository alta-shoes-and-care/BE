package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	t.Run("succeed to hash password", func(t *testing.T) {
		res, err := HashPassword("abc")
		assert.Nil(t, err)
		assert.NotNil(t, res)
	})
}

func TestCheckPasswordHash(t *testing.T) {
	t.Run("succeed to check hashed password", func(t *testing.T) {
		hashed, _ := HashPassword("abc")
		res := CheckPasswordHash("abc", hashed)
		assert.True(t, res)
	})

	t.Run("failed to check hashed password", func(t *testing.T) {
		hashed, _ := HashPassword("abc")
		res := CheckPasswordHash("abc", hashed)
		assert.True(t, res)
	})
}
