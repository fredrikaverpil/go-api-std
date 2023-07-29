package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	var err *internalError

	err = InternalError("oops")
	assert.Equal(t, err.Code, ErrInternal)
	assert.Equal(t, err.Message, "oops")

	err = NotFoundError("foo not found")
	assert.Equal(t, err.Code, ErrNotFound)
	assert.Equal(t, err.Message, "foo not found")

	err = InvalidInputError("unrecognized argument")
	assert.Equal(t, err.Code, ErrInvalidInput)
	assert.Equal(t, err.Message, "unrecognized argument")
}
