package restpc

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBadCodeString(t *testing.T) {
	{
		i := len(_Code_index) - 1
		c := Code(i)
		assert.Equal(t, fmt.Sprintf("Code(%d)", i), c.String())
	}
}

func TestHTTPStatusFromCodeOK(t *testing.T) {
	assert.Equal(t, http.StatusOK, HTTPStatusFromCode(OK))
}
