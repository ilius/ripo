package restpc

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
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
func TestHandler_Full_Happy(t *testing.T) {
	myUrlStr := "http://127.0.0.1/test/full"
	handlerFunc := TranslateHandler(func(req Request) (*Response, error) {
		if req.URL().String() != myUrlStr {
			return nil, NewError(NotFound, "bad URL", nil)
		}
		remoteIP, err := req.RemoteIP()
		if err != nil {
			return nil, NewError(Internal, "error in req.RemoteIP", err)
		}
		if remoteIP != "127.0.0.1" {
			return nil, NewError(PermissionDenied, "bad RemoteIP", nil)
		}
		if req.Host() != "127.0.0.1" {
			return nil, NewError(PermissionDenied, "bad Host", nil)
		}
		firstName, err := req.GetString("firstName")
		if err != nil {
			return nil, err
		}
		lastName, err := req.GetString("lastName")
		if err != nil {
			return nil, err
		}
		medianName, err := req.GetString("medianName", FromBody, FromEmpty)
		if err != nil {
			return nil, err
		}
		age, err := req.GetFloat("age")
		if err != nil {
			return nil, err
		}
		subscribed, err := req.GetBool("subscribed")
		if err != nil {
			return nil, err
		}
		interests, err := req.GetStringList("interests")
		if err != nil {
			return nil, err
		}
		count, err := req.GetInt("count")
		if err != nil {
			return nil, err
		}
		maxCount, err := req.GetIntDefault("maxCount", 100)
		if err != nil {
			return nil, err
		}
		unsubTime, err := req.GetTime("unsubTime", FromBody, FromEmpty)
		if err != nil {
			return nil, err
		}
		return &Response{
			Data: map[string]interface{}{
				"firstName":  *firstName,
				"lastName":   *lastName,
				"medianName": *medianName,
				"age":        *age,
				"subscribed": *subscribed,
				"interests":  interests,
				"count":      *count,
				"maxCount":   maxCount,
				"unsubTime":  unsubTime,
			},
		}, nil
	})
	{

		r, err := http.NewRequest("POST", myUrlStr, strings.NewReader(`{
			"firstName": "John",
			"lastName": "Smith",
			"age": 30,
			"subscribed": true,
			"interests": ["Tech", "Sports"],
			"count": 10,
			"maxCount": 20
		}`))
		r.RemoteAddr = "127.0.0.1:1234"
		if err != nil {
			panic(err)
		}
		w := httptest.NewRecorder()
		handlerFunc(w, r)
		assert.Equal(t, http.StatusOK, w.Code)
		resBody := w.Body.String()
		resMap := map[string]interface{}{}
		err = json.Unmarshal([]byte(resBody), &resMap)
		if err != nil {
			panic(err)
		}
		assert.Equal(t, "John", resMap["firstName"])
		assert.Equal(t, "Smith", resMap["lastName"])
		assert.Equal(t, "", resMap["medianName"])
		assert.EqualValues(t, 30, resMap["age"])
		assert.Equal(t, true, resMap["subscribed"])
		assert.Equal(t, []interface{}{"Tech", "Sports"}, resMap["interests"])
		assert.EqualValues(t, 10, resMap["count"])
		assert.EqualValues(t, 20, resMap["maxCount"])
	}
	{
		r, err := http.NewRequest("POST", myUrlStr, strings.NewReader(`{}`))
		if err != nil {
			panic(err)
		}
		r.RemoteAddr = "127.0.0.1:1234"
		w := httptest.NewRecorder()
		handlerFunc(w, r)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		resBody := w.Body.String()
		assert.Equal(t, "{\"code\":\"MissingArgument\",\"error\":\"missing 'firstName'\"}\n", resBody)
	}
	{
		r, err := http.NewRequest("POST", myUrlStr, strings.NewReader(`{
			"firstName": "John"
		}`))
		if err != nil {
			panic(err)
		}
		r.RemoteAddr = "127.0.0.1:1234"
		w := httptest.NewRecorder()
		handlerFunc(w, r)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		resBody := w.Body.String()
		assert.Equal(t, "{\"code\":\"MissingArgument\",\"error\":\"missing 'lastName'\"}\n", resBody)
	}
	{
		r, err := http.NewRequest("POST", myUrlStr, strings.NewReader(`{
			"firstName": "John",
			"lastName": "Smith"
		}`))
		if err != nil {
			panic(err)
		}
		r.RemoteAddr = "127.0.0.1:1234"
		w := httptest.NewRecorder()
		handlerFunc(w, r)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		resBody := w.Body.String()
		assert.Equal(t, "{\"code\":\"MissingArgument\",\"error\":\"missing 'age'\"}\n", resBody)
	}
	{
		r, err := http.NewRequest("POST", myUrlStr, strings.NewReader(`{
			"firstName": "John",
			"lastName": "Smith",
			"age": "forty"
		}`))
		if err != nil {
			panic(err)
		}
		r.RemoteAddr = "127.0.0.1:1234"
		w := httptest.NewRecorder()
		handlerFunc(w, r)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		resBody := w.Body.String()
		assert.Equal(t, "{\"code\":\"InvalidArgument\",\"error\":\"invalid 'age', must be float\"}\n", resBody)
	}
	{
		r, err := http.NewRequest("POST", myUrlStr, strings.NewReader(`{
			"firstName": "John",
			"lastName": "Smith",
			"age": 30
		}`))
		if err != nil {
			panic(err)
		}
		r.RemoteAddr = "127.0.0.1:1234"
		w := httptest.NewRecorder()
		handlerFunc(w, r)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		resBody := w.Body.String()
		assert.Equal(t, "{\"code\":\"MissingArgument\",\"error\":\"missing 'subscribed'\"}\n", resBody)
	}
	{
		r, err := http.NewRequest("POST", myUrlStr, strings.NewReader(`{
			"firstName": "John",
			"lastName": "Smith",
			"age": 30,
			"subscribed": "no"
		}`))
		if err != nil {
			panic(err)
		}
		r.RemoteAddr = "127.0.0.1:1234"
		w := httptest.NewRecorder()
		handlerFunc(w, r)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		resBody := w.Body.String()
		assert.Equal(t, "{\"code\":\"InvalidArgument\",\"error\":\"invalid 'subscribed', must be true or false\"}\n", resBody)
	}
	{
		r, err := http.NewRequest("POST", myUrlStr, strings.NewReader(`{
			"firstName": "John",
			"lastName": "Smith",
			"age": 30,
			"subscribed": true
		}`))
		if err != nil {
			panic(err)
		}
		r.RemoteAddr = "127.0.0.1:1234"
		w := httptest.NewRecorder()
		handlerFunc(w, r)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		resBody := w.Body.String()
		assert.Equal(t, "{\"code\":\"MissingArgument\",\"error\":\"missing 'interests'\"}\n", resBody)
	}
	{
		r, err := http.NewRequest("POST", myUrlStr, strings.NewReader(`{
			"firstName": "John",
			"lastName": "Smith",
			"age": 30,
			"subscribed": true,
			"interests": "Tech"
		}`))
		if err != nil {
			panic(err)
		}
		r.RemoteAddr = "127.0.0.1:1234"
		w := httptest.NewRecorder()
		handlerFunc(w, r)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		resBody := w.Body.String()
		assert.Equal(t, "{\"code\":\"InvalidArgument\",\"error\":\"invalid 'interests', must be array of strings\"}\n", resBody)
	}
	{
		r, err := http.NewRequest("POST", myUrlStr, strings.NewReader(`{
			"firstName": "John",
			"lastName": "Smith",
			"age": 30,
			"subscribed": true,
			"interests": ["Tech", "Sports"]
		}`))
		if err != nil {
			panic(err)
		}
		r.RemoteAddr = "127.0.0.1:1234"
		w := httptest.NewRecorder()
		handlerFunc(w, r)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		resBody := w.Body.String()
		assert.Equal(t, "{\"code\":\"MissingArgument\",\"error\":\"missing 'count'\"}\n", resBody)
	}
	{
		r, err := http.NewRequest("POST", myUrlStr, strings.NewReader(`{
			"firstName": "John",
			"lastName": "Smith",
			"age": 30,
			"subscribed": true,
			"interests": ["Tech", "Sports"],
			"count": "ten"
		}`))
		if err != nil {
			panic(err)
		}
		r.RemoteAddr = "127.0.0.1:1234"
		w := httptest.NewRecorder()
		handlerFunc(w, r)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		resBody := w.Body.String()
		assert.Equal(t, "{\"code\":\"InvalidArgument\",\"error\":\"invalid 'count', must be integer\"}\n", resBody)
	}
	{
		r, err := http.NewRequest("POST", myUrlStr, strings.NewReader(`{
			"firstName": "John",
			"lastName": "Smith",
			"age": 30,
			"subscribed": true,
			"interests": ["Tech", "Sports"],
			"count": 10
		}`))
		if err != nil {
			panic(err)
		}
		r.RemoteAddr = "127.0.0.1:1234"
		w := httptest.NewRecorder()
		handlerFunc(w, r)
		assert.Equal(t, http.StatusOK, w.Code)
	}
	{
		r, err := http.NewRequest("POST", myUrlStr, strings.NewReader(`{
			"firstName": "John",
			"lastName": "Smith",
			"age": 30,
			"subscribed": true,
			"interests": ["Tech", "Sports"],
			"count": 10,
			"maxCount": "123"
		}`))
		if err != nil {
			panic(err)
		}
		r.RemoteAddr = "127.0.0.1:1234"
		w := httptest.NewRecorder()
		handlerFunc(w, r)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		resBody := w.Body.String()
		assert.Equal(t, "{\"code\":\"InvalidArgument\",\"error\":\"invalid 'maxCount', must be integer\"}\n", resBody)
	}
	{
		r, err := http.NewRequest("POST", myUrlStr, strings.NewReader(`{
			"firstName": "John",
			"lastName": "Smith",
			"age": 30,
			"subscribed": true,
			"interests": ["Tech", "Sports"],
			"count": 10,
			"unsubTime": "2017-12-20T17:30:00Z"
		}`))
		if err != nil {
			panic(err)
		}
		r.RemoteAddr = "127.0.0.1:1234"
		w := httptest.NewRecorder()
		handlerFunc(w, r)
		assert.Equal(t, http.StatusOK, w.Code)
		t.Log(w.Body.String())
	}
}

func TestHandler_2(t *testing.T) {
	myUrlStr := "http://127.0.0.1/test/full"
	handlerFunc := TranslateHandler(func(req Request) (*Response, error) {
		firstName, err := req.GetString("firstName")
		if err != nil {
			return nil, err
		}
		lastName, err := req.GetString("lastName")
		if err != nil {
			return nil, err
		}
		unsubTime, err := req.GetTime("unsubTime")
		if err != nil {
			return nil, err
		}
		return &Response{
			Data: map[string]interface{}{
				"firstName": *firstName,
				"lastName":  *lastName,
				"unsubTime": unsubTime,
			},
		}, nil
	})
	{
		r, err := http.NewRequest("POST", myUrlStr, strings.NewReader(`{
			"firstName": 123
		}`))
		if err != nil {
			panic(err)
		}
		r.RemoteAddr = "127.0.0.1:1234"
		w := httptest.NewRecorder()
		handlerFunc(w, r)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		resBody := w.Body.String()
		assert.Equal(t, "{\"code\":\"InvalidArgument\",\"error\":\"invalid 'firstName', must be string\"}\n", resBody)
	}
	{
		r, err := http.NewRequest("POST", myUrlStr, strings.NewReader(`{
			"firstName": "John",
			"lastName": "Smith"
		}`))
		if err != nil {
			panic(err)
		}
		r.RemoteAddr = "127.0.0.1:1234"
		w := httptest.NewRecorder()
		handlerFunc(w, r)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		resBody := w.Body.String()
		assert.Equal(t, "{\"code\":\"MissingArgument\",\"error\":\"missing 'unsubTime'\"}\n", resBody)
	}
	{
		r, err := http.NewRequest("POST", myUrlStr, strings.NewReader(`{
			"firstName": "John",
			"lastName": "Smith",
			"unsubTime": "bad time"
		}`))
		if err != nil {
			panic(err)
		}
		r.RemoteAddr = "127.0.0.1:1234"
		w := httptest.NewRecorder()
		handlerFunc(w, r)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		resBody := w.Body.String()
		assert.Equal(t, "{\"code\":\"InvalidArgument\",\"error\":\"invalid 'unsubTime', must be RFC3339 time string\"}\n", resBody)
	}
	{
		r, err := http.NewRequest("POST", myUrlStr, strings.NewReader(`{
			"firstName": "John",
			"lastName": "Smith",
			"unsubTime": "2017-12-20T17:30:00Z"
		}`))
		if err != nil {
			panic(err)
		}
		r.RemoteAddr = "127.0.0.1:1234"
		w := httptest.NewRecorder()
		handlerFunc(w, r)
		assert.Equal(t, http.StatusOK, w.Code)
		t.Log(w.Body.String())
	}
}
