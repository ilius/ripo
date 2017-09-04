package restpc

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandler_1(t *testing.T) {
	handlerFunc := TranslateHandler(func(req Request) (res *Response, err error) {
		return nil, fmt.Errorf("go away")
	})
	r, err := http.NewRequest("GET", "", nil)
	if err != nil {
		panic(err)
	}
	w := httptest.NewRecorder()
	handlerFunc(w, r)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	body := w.Body.String()
	assert.Equal(t, "{\"code\":\"Unknown\",\"error\":\"Unknown\"}\n", body)
}
