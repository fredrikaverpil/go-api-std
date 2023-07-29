package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	var err error
	var ierr *CustomError

	err = InternalError("oops")
	ierr = err.(*CustomError)
	assert.Equal(t, ierr.Code, ErrInternal)
	assert.Equal(t, ierr.Message, "oops")

	err = NotFoundError("foo not found")
	ierr = err.(*CustomError)
	assert.Equal(t, ierr.Code, ErrNotFound)
	assert.Equal(t, ierr.Message, "foo not found")

	err = InvalidInputError("unrecognized argument")
	ierr = err.(*CustomError)
	assert.Equal(t, ierr.Code, ErrInvalidInput)
	assert.Equal(t, ierr.Message, "unrecognized argument")
}
