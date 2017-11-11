package restpc

import (
	"fmt"
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
