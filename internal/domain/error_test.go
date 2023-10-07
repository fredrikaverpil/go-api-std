package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	var err error
	var ierr *Error

	err = InternalError("oops")
	ierr = err.(*Error)
	assert.Equal(t, ierr.Code, ErrInternal)
	assert.Equal(t, ierr.Message, "oops")

	err = NotFoundError("foo not found")
	ierr = err.(*Error)
	assert.Equal(t, ierr.Code, ErrNotFound)
	assert.Equal(t, ierr.Message, "foo not found")

	err = InvalidArgumentError("unrecognized argument")
	ierr = err.(*Error)
	assert.Equal(t, ierr.Code, ErrInvalidArgument)
	assert.Equal(t, ierr.Message, "unrecognized argument")
}
