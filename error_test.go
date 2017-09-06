package restpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewError_Twice(t *testing.T) {
	err := NewError(InvalidArgument, "something is missing", nil)
	assert.Equal(t, "something is missing", err.Message())
	assert.Equal(t, "something is missing", err.Error())
	assert.Equal(t, InvalidArgument, err.Code())

	err = NewError(Unavailable, "not sure what", err)
	assert.Equal(t, "something is missing", err.Message())
	assert.Equal(t, "something is missing", err.Error())
	assert.Equal(t, InvalidArgument, err.Code())
}
