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

func TestHandler_CodeMapping(t *testing.T) {
	{
		handlerFunc := TranslateHandler(func(req Request) (res *Response, err error) {
			return nil, NewError(Code(len(_Code_index)-1), "", nil) // added by Saeed Rasooli
		})
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "", nil)
		handlerFunc(w, r)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	}
	{
		handlerFunc := TranslateHandler(func(req Request) (res *Response, err error) {
			return nil, NewError(Canceled, "", nil)
		})
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "", nil)
		handlerFunc(w, r)
		assert.Equal(t, http.StatusRequestTimeout, w.Code)
	}
	{
		handlerFunc := TranslateHandler(func(req Request) (res *Response, err error) {
			return nil, NewError(InvalidArgument, "", nil)
		})
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "", nil)
		handlerFunc(w, r)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	}
	{
		handlerFunc := TranslateHandler(func(req Request) (res *Response, err error) {
			return nil, NewError(DeadlineExceeded, "", nil)
		})
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "", nil)
		handlerFunc(w, r)
		assert.Equal(t, http.StatusRequestTimeout, w.Code)
	}
	{
		handlerFunc := TranslateHandler(func(req Request) (res *Response, err error) {
			return nil, NewError(NotFound, "", nil)
		})
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "", nil)
		handlerFunc(w, r)
		assert.Equal(t, http.StatusNotFound, w.Code)
	}
	{
		handlerFunc := TranslateHandler(func(req Request) (res *Response, err error) {
			return nil, NewError(AlreadyExists, "", nil)
		})
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "", nil)
		handlerFunc(w, r)
		assert.Equal(t, http.StatusConflict, w.Code)
	}
	{
		handlerFunc := TranslateHandler(func(req Request) (res *Response, err error) {
			return nil, NewError(PermissionDenied, "", nil)
		})
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "", nil)
		handlerFunc(w, r)
		assert.Equal(t, http.StatusForbidden, w.Code)
	}
	{
		handlerFunc := TranslateHandler(func(req Request) (res *Response, err error) {
			return nil, NewError(Unauthenticated, "", nil)
		})
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "", nil)
		handlerFunc(w, r)
		assert.Equal(t, http.StatusUnauthorized, w.Code)
	}
	{
		handlerFunc := TranslateHandler(func(req Request) (res *Response, err error) {
			return nil, NewError(ResourceExhausted, "", nil)
		})
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "", nil)
		handlerFunc(w, r)
		assert.Equal(t, http.StatusForbidden, w.Code)
	}
	{
		handlerFunc := TranslateHandler(func(req Request) (res *Response, err error) {
			return nil, NewError(FailedPrecondition, "", nil)
		})
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "", nil)
		handlerFunc(w, r)
		assert.Equal(t, http.StatusPreconditionFailed, w.Code)
	}
	{
		handlerFunc := TranslateHandler(func(req Request) (res *Response, err error) {
			return nil, NewError(Aborted, "", nil)
		})
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "", nil)
		handlerFunc(w, r)
		assert.Equal(t, http.StatusConflict, w.Code)
	}
	{
		handlerFunc := TranslateHandler(func(req Request) (res *Response, err error) {
			return nil, NewError(OutOfRange, "", nil)
		})
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "", nil)
		handlerFunc(w, r)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	}
	{
		handlerFunc := TranslateHandler(func(req Request) (res *Response, err error) {
			return nil, NewError(Unimplemented, "", nil)
		})
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "", nil)
		handlerFunc(w, r)
		assert.Equal(t, http.StatusNotImplemented, w.Code)
	}
	{
		handlerFunc := TranslateHandler(func(req Request) (res *Response, err error) {
			return nil, NewError(Unavailable, "", nil)
		})
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "", nil)
		handlerFunc(w, r)
		assert.Equal(t, http.StatusServiceUnavailable, w.Code)
	}
	{
		handlerFunc := TranslateHandler(func(req Request) (res *Response, err error) {
			return nil, NewError(DataLoss, "", nil)
		})
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "", nil)
		handlerFunc(w, r)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	}
	{
		handlerFunc := TranslateHandler(func(req Request) (res *Response, err error) {
			return nil, NewError(MissingArgument, "", nil) // added by Saeed Rasooli
		})
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "", nil)
		handlerFunc(w, r)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	}
}
