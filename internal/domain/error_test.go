package domain

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestError(t *testing.T) {
	var err error
	var ierr *Error

	err = InternalError("oops")
	ierr, ok := err.(*Error)
	assert.Assert(t, ok)
	assert.Equal(t, ierr.Code, ErrInternal)
	assert.Equal(t, ierr.Message, "oops")

	err = NotFoundError("foo not found")
	ierr, ok = err.(*Error)
	assert.Assert(t, ok)
	assert.Equal(t, ierr.Code, ErrNotFound)
	assert.Equal(t, ierr.Message, "foo not found")

	err = InvalidArgumentError("unrecognized argument")
	ierr, ok = err.(*Error)
	assert.Assert(t, ok)
	assert.Equal(t, ierr.Code, ErrInvalidArgument)
	assert.Equal(t, ierr.Message, "unrecognized argument")
}
