package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := "mysecretpassword"
	hashedPassword, err := HashPassword(password)
	assert.NoError(t, err)
	assert.NotEmpty(t, hashedPassword)

	err = CheckPassword(password, hashedPassword)
	assert.NoError(t, err)

	wrongPassword := "wrongpassword"
	err = CheckPassword(wrongPassword, hashedPassword)
	assert.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}
