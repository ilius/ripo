package restpc

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func panicerHandler(req Request) (res *Response, err error) {
	panic("we screwed up")
	panic("this line never happens")
}

func TestHandler_Panic(t *testing.T) {
	handlerFunc := TranslateHandler(panicerHandler)
	r, err := http.NewRequest("GET", "", nil)
	if err != nil {
		panic(err)
	}
	w := httptest.NewRecorder()
	handlerFunc(w, r)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	body := w.Body.String()
	assert.Equal(t, "{\"code\":\"Internal\",\"error\":\"Internal\"}\n", body)
}

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
