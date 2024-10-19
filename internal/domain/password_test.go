package domain

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
	"gotest.tools/v3/assert"
)

func TestPassword(t *testing.T) {
	password := "mysecretpassword"
	hashedPassword, err := HashPassword(password)
	assert.NilError(t, err)
	assert.Assert(t, hashedPassword != "")

	err = CheckPassword(password, hashedPassword)
	assert.NilError(t, err)

	wrongPassword := "wrongpassword"
	err = CheckPassword(wrongPassword, hashedPassword)
	assert.ErrorIs(t, err, bcrypt.ErrMismatchedHashAndPassword)
}
