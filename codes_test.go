package ripo

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/tylerb/is"
)

func TestBadCodeString(t *testing.T) {
	is := is.New(t)
	{
		i := len(_Code_index) - 1
		c := Code(i)
		is.Equal(fmt.Sprintf("Code(%d)", i), c.String())
	}
}

func TestHTTPStatusFromCodeOK(t *testing.T) {
	is := is.New(t)
	is.Equal(http.StatusOK, HTTPStatusFromCode(OK))
}
