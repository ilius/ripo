package ripo

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
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
	body := strings.TrimSpace(w.Body.String())
	assert.Equal(t, "{\"code\":\"Internal\",\"error\":\"Internal\"}", body)
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
	body := strings.TrimSpace(w.Body.String())
	assert.Equal(t, "{\"code\":\"Unknown\",\"error\":\"Unknown\"}", body)
}

func TestHandler_PostNilBody(t *testing.T) {
	handlerFunc := TranslateHandler(func(req Request) (res *Response, err error) {
		name, err := req.GetString("name")
		if err != nil {
			return nil, err
		}
		return &Response{
			Data: map[string]string{
				"name": *name,
			},
		}, nil
	})
	r, err := http.NewRequest("POST", "", nil)
	if err != nil {
		panic(err)
	}
	w := httptest.NewRecorder()
	handlerFunc(w, r)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	body := strings.TrimSpace(w.Body.String())
	assert.Equal(t, "error in parsing form", body)
}

func TestHandler_PostSimple(t *testing.T) {
	handlerFunc := TranslateHandler(func(req Request) (res *Response, err error) {
		name, err := req.GetString("name")
		if err != nil {
			return nil, err
		}
		return &Response{
			Data: map[string]string{
				"msg": "hello " + *name,
			},
		}, nil
	})
	r, err := http.NewRequest("POST", "", strings.NewReader(`{"name": "John"}`))
	if err != nil {
		panic(err)
	}
	w := httptest.NewRecorder()
	handlerFunc(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
	body := strings.TrimSpace(w.Body.String())
	assert.Equal(t, "{\"msg\":\"hello John\"}", body)
}

func TestHandler_MockBody(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockBody := NewMockReadCloser(ctrl)
	r, err := http.NewRequest("POST", "http://127.0.0.1/test", mockBody)
	assert.NoError(t, err)
	handlerFunc := TranslateHandler(func(req Request) (res *Response, err error) {
		name, err := req.GetString("name")
		if err != nil {
			return nil, err
		}
		return &Response{
			Data: map[string]string{
				"msg": "hello " + *name,
			},
		}, nil
	})
	w := httptest.NewRecorder()

	mockBody.EXPECT().Read(gomock.Any()).Return(0, fmt.Errorf("no data for you"))
	mockBody.EXPECT().Close()

	handlerFunc(w, r)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	body := strings.TrimSpace(w.Body.String())
	assert.Equal(t, "{\"code\":\"Unknown\",\"error\":\"Unknown\"}", body)
}

func TestHandler_ResNil(t *testing.T) {
	handlerFunc := TranslateHandler(func(req Request) (res *Response, err error) {
		return nil, nil
	})
	r, err := http.NewRequest("GET", "", nil)
	if err != nil {
		panic(err)
	}
	w := httptest.NewRecorder()
	handlerFunc(w, r)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	body := strings.TrimSpace(w.Body.String())
	assert.Equal(t, "{\"code\":\"Internal\",\"error\":\"Internal\"}", body)
}

func TestHandler_ResData_JsonBytes(t *testing.T) {
	handlerFunc := TranslateHandler(func(req Request) (res *Response, err error) {
		return &Response{
			Data: []byte(`{"refNo": "1234"}`),
		}, nil
	})
	r, err := http.NewRequest("GET", "", nil)
	if err != nil {
		panic(err)
	}
	w := httptest.NewRecorder()
	handlerFunc(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "text/plain; charset=utf-8", w.Header().Get("Content-Type")) // not json
	body := strings.TrimSpace(w.Body.String())
	assert.Equal(t, "{\"refNo\": \"1234\"}", body)
}

func TestHandler_ResData_JsonString(t *testing.T) {
	handlerFunc := TranslateHandler(func(req Request) (res *Response, err error) {
		return &Response{
			Data: `{"refNo": "1234"}`,
		}, nil
	})
	r, err := http.NewRequest("GET", "", nil)
	if err != nil {
		panic(err)
	}
	w := httptest.NewRecorder()
	handlerFunc(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "text/plain; charset=utf-8", w.Header().Get("Content-Type")) // not json
	body := strings.TrimSpace(w.Body.String())
	assert.Equal(t, "{\"refNo\": \"1234\"}", body)
}

func TestHandler_ResData_Struct(t *testing.T) {
	handlerFunc := TranslateHandler(func(req Request) (res *Response, err error) {
		return &Response{
			Data: struct {
				RefNo string `json:"refNo"`
			}{
				RefNo: "1234",
			},
		}, nil
	})
	r, err := http.NewRequest("GET", "", nil)
	if err != nil {
		panic(err)
	}
	w := httptest.NewRecorder()
	handlerFunc(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json; charset=UTF-8", w.Header().Get("Content-Type")) // not json
	body := strings.TrimSpace(w.Body.String())
	assert.Equal(t, "{\"refNo\":\"1234\"}", body)
}

func TestHandler_ResData_Nil(t *testing.T) {
	handlerFunc := TranslateHandler(func(req Request) (res *Response, err error) {
		return &Response{
			Data: nil,
		}, nil
	})
	r, err := http.NewRequest("GET", "", nil)
	if err != nil {
		panic(err)
	}
	w := httptest.NewRecorder()
	handlerFunc(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json; charset=UTF-8", w.Header().Get("Content-Type"))
	body := strings.TrimSpace(w.Body.String())
	assert.Equal(t, "{}", body)
}

func TestHandler_ResData_HugeNum(t *testing.T) {
	handlerFunc := TranslateHandler(func(req Request) (res *Response, err error) {
		return &Response{
			Data: math.Pow10(1000), // not json-marshallable
		}, nil
	})
	r, err := http.NewRequest("GET", "", nil)
	if err != nil {
		panic(err)
	}
	w := httptest.NewRecorder()
	handlerFunc(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "text/plain; charset=utf-8", w.Header().Get("Content-Type")) // not json
	body := strings.TrimSpace(w.Body.String())
	assert.Equal(t, "", body)
}

func TestHandler_ResRedirectPath_DefaultCode(t *testing.T) {
	handlerFunc := TranslateHandler(func(req Request) (res *Response, err error) {
		return &Response{
			RedirectPath: "login",
		}, nil
	})
	r, err := http.NewRequest("GET", "", nil)
	if err != nil {
		panic(err)
	}
	w := httptest.NewRecorder()
	handlerFunc(w, r)
	assert.Equal(t, http.StatusSeeOther, w.Code)
	assert.Equal(t, "", w.Header().Get("Content-Type"))
	body := strings.TrimSpace(w.Body.String())
	assert.Equal(t, "<a href=\"/login\">See Other</a>.", body)
}

func TestHandler_ResRedirectPath_MovedPermanently(t *testing.T) {
	handlerFunc := TranslateHandler(func(req Request) (res *Response, err error) {
		return &Response{
			RedirectPath:       "login",
			RedirectStatusCode: http.StatusMovedPermanently,
		}, nil
	})
	r, err := http.NewRequest("GET", "", nil)
	if err != nil {
		panic(err)
	}
	w := httptest.NewRecorder()
	handlerFunc(w, r)
	assert.Equal(t, http.StatusMovedPermanently, w.Code)
	body := strings.TrimSpace(w.Body.String())
	assert.Equal(t, "<a href=\"/login\">Moved Permanently</a>.", body)
}

func TestHandler_ResHeader(t *testing.T) {
	handlerFunc := TranslateHandler(func(req Request) (res *Response, err error) {
		return &Response{
			Header: http.Header{
				"Content-Language": []string{"en"},
			},
		}, nil
	})
	r, err := http.NewRequest("GET", "", nil)
	if err != nil {
		panic(err)
	}
	w := httptest.NewRecorder()
	handlerFunc(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "en", w.Header().Get("Content-Language"))
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
		birthDateIn, err := req.GetObject("birthDate", reflect.TypeOf([]int{}))
		if err != nil {
			return nil, err
		}
		birthDate := birthDateIn.([]int)
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
				"birthDate":  birthDate,
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
			"birthDate": [1987, 1, 1],
			"maxCount": 20
		}`))
		r.RemoteAddr = "127.0.0.1:1234"
		if err != nil {
			panic(err)
		}
		w := httptest.NewRecorder()
		handlerFunc(w, r)
		assert.Equal(t, http.StatusOK, w.Code)
		resBody := strings.TrimSpace(w.Body.String())
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
		resBody := strings.TrimSpace(w.Body.String())
		assert.Equal(t, "{\"code\":\"MissingArgument\",\"error\":\"missing 'firstName'\"}", resBody)
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
		resBody := strings.TrimSpace(w.Body.String())
		assert.Equal(t, "{\"code\":\"MissingArgument\",\"error\":\"missing 'lastName'\"}", resBody)
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
		resBody := strings.TrimSpace(w.Body.String())
		assert.Equal(t, "{\"code\":\"MissingArgument\",\"error\":\"missing 'age'\"}", resBody)
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
		resBody := strings.TrimSpace(w.Body.String())
		assert.Equal(t, "{\"code\":\"InvalidArgument\",\"error\":\"invalid 'age', must be float\"}", resBody)
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
		resBody := strings.TrimSpace(w.Body.String())
		assert.Equal(t, "{\"code\":\"MissingArgument\",\"error\":\"missing 'subscribed'\"}", resBody)
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
		resBody := strings.TrimSpace(w.Body.String())
		assert.Equal(t, "{\"code\":\"InvalidArgument\",\"error\":\"invalid 'subscribed', must be true or false\"}", resBody)
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
		resBody := strings.TrimSpace(w.Body.String())
		assert.Equal(t, "{\"code\":\"MissingArgument\",\"error\":\"missing 'interests'\"}", resBody)
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
		resBody := strings.TrimSpace(w.Body.String())
		assert.Equal(t, "{\"code\":\"InvalidArgument\",\"error\":\"invalid 'interests', must be array of strings\"}", resBody)
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
		resBody := strings.TrimSpace(w.Body.String())
		assert.Equal(t, "{\"code\":\"MissingArgument\",\"error\":\"missing 'count'\"}", resBody)
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
		resBody := strings.TrimSpace(w.Body.String())
		assert.Equal(t, "{\"code\":\"InvalidArgument\",\"error\":\"invalid 'count', must be integer\"}", resBody)
	}
	{
		r, err := http.NewRequest("POST", myUrlStr, strings.NewReader(`{
			"firstName": "John",
			"lastName": "Smith",
			"age": 30,
			"subscribed": true,
			"interests": ["Tech", "Sports"],
			"birthDate": [1987, 1, 1],
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
		resBody := strings.TrimSpace(w.Body.String())
		assert.Equal(t, "{\"code\":\"InvalidArgument\",\"error\":\"invalid 'maxCount', must be integer\"}", resBody)
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
		assert.Equal(t, http.StatusBadRequest, w.Code)
		resBody := strings.TrimSpace(w.Body.String())
		assert.Equal(t, "{\"code\":\"MissingArgument\",\"error\":\"missing 'birthDate'\"}", resBody)
	}
	{
		r, err := http.NewRequest("POST", myUrlStr, strings.NewReader(`{
			"firstName": "John",
			"lastName": "Smith",
			"age": 30,
			"subscribed": true,
			"interests": ["Tech", "Sports"],
			"count": 10,
			"birthDate": "[1987, 1, 1]",
			"unsubTime": "2017-12-20T17:30:00Z"
		}`))
		if err != nil {
			panic(err)
		}
		r.RemoteAddr = "127.0.0.1:1234"
		w := httptest.NewRecorder()
		handlerFunc(w, r)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		resBody := strings.TrimSpace(w.Body.String())
		assert.Equal(t, "{\"code\":\"InvalidArgument\",\"error\":\"invalid 'birthDate', must be a compatible object\"}", resBody)
	}
	{
		r, err := http.NewRequest("POST", myUrlStr, strings.NewReader(`{
			"firstName": "John",
			"lastName": "Smith",
			"age": 30,
			"subscribed": true,
			"interests": ["Tech", "Sports"],
			"count": 10,
			"birthDate": [1987, 1, 1],
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
		resBody := strings.TrimSpace(w.Body.String())
		assert.Equal(t, "{\"code\":\"InvalidArgument\",\"error\":\"invalid 'firstName', must be string\"}", resBody)
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
		resBody := strings.TrimSpace(w.Body.String())
		assert.Equal(t, "{\"code\":\"MissingArgument\",\"error\":\"missing 'unsubTime'\"}", resBody)
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
		resBody := strings.TrimSpace(w.Body.String())
		assert.Equal(t, "{\"code\":\"InvalidArgument\",\"error\":\"invalid 'unsubTime', must be RFC3339 time string\"}", resBody)
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
