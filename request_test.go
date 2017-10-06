package restpc

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_requestImp_URL(t *testing.T) {
	r, err := http.NewRequest("GET", "http://127.0.0.1/test1", nil)
	if err != nil {
		panic(err)
	}
	req := &requestImp{
		r:           r,
		handlerName: "Test",
	}
	assert.Equal(t, "/test1", req.URL().Path)
	u := req.URL()
	u.Path = "/test2"
	assert.Equal(t, "/test1", req.URL().Path)
	assert.Equal(t, "/test2", u.Path)
}
