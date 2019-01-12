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
	"github.com/tylerb/is"
)

func panicerHandler(req Request) (res *Response, err error) {
	panic("we screwed up")
	panic("this line never happens")
}

func TestHandler_Panic(t *testing.T) {
	is := is.New(t)
	handlerFunc := TranslateHandler(panicerHandler)
	r, err := http.NewRequest("GET", "", nil)
	if err != nil {
		panic(err)
	}
	w := httptest.NewRecorder()
	handlerFunc(w, r)
	is.Equal(http.StatusInternalServerError, w.Code)
	body := strings.TrimSpace(w.Body.String())
	is.Equal("{\"code\":\"Internal\",\"error\":\"Internal\"}", body)
}

func TestHandler_1(t *testing.T) {
	is := is.New(t)
	handlerFunc := TranslateHandler(func(req Request) (res *Response, err error) {
		return nil, fmt.Errorf("go away")
	})
	r, err := http.NewRequest("GET", "", nil)
	if err != nil {
		panic(err)
	}
	w := httptest.NewRecorder()
	handlerFunc(w, r)
	is.Equal(http.StatusInternalServerError, w.Code)
	body := strings.TrimSpace(w.Body.String())
	is.Equal("{\"code\":\"Unknown\",\"error\":\"Unknown\"}", body)
}

func TestHandler_PostNilBody(t *testing.T) {
	is := is.New(t)
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
	is.Equal(http.StatusBadRequest, w.Code)
	body := strings.TrimSpace(w.Body.String())
	is.Equal("error in parsing form", body)
}

func TestHandler_PostSimple(t *testing.T) {
	is := is.New(t)
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
	is.Equal(http.StatusOK, w.Code)
	body := strings.TrimSpace(w.Body.String())
	is.Equal("{\"msg\":\"hello John\"}", body)
}

func TestHandler_MockBody(t *testing.T) {
	is := is.New(t)
	ctrl := gomock.NewController(t)
	mockBody := NewMockReadCloser(ctrl)
	r, err := http.NewRequest("POST", "http://127.0.0.1/test", mockBody)
	is.NotErr(err)
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
	is.Equal(http.StatusInternalServerError, w.Code)
	body := strings.TrimSpace(w.Body.String())
	is.Equal("{\"code\":\"Unknown\",\"error\":\"Unknown\"}", body)
}

func TestHandler_ResNil(t *testing.T) {
	is := is.New(t)
	handlerFunc := TranslateHandler(func(req Request) (res *Response, err error) {
		return nil, nil
	})
	r, err := http.NewRequest("GET", "", nil)
	if err != nil {
		panic(err)
	}
	w := httptest.NewRecorder()
	handlerFunc(w, r)
	is.Equal(http.StatusInternalServerError, w.Code)
	body := strings.TrimSpace(w.Body.String())
	is.Equal("{\"code\":\"Internal\",\"error\":\"Internal\"}", body)
}

func TestHandler_ResData_JsonBytes(t *testing.T) {
	is := is.New(t)
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
	is.Equal(http.StatusOK, w.Code)
	is.Equal("text/plain; charset=utf-8", w.Header().Get("Content-Type")) // not json
	body := strings.TrimSpace(w.Body.String())
	is.Equal("{\"refNo\": \"1234\"}", body)
}

func TestHandler_ResData_JsonString(t *testing.T) {
	is := is.New(t)
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
	is.Equal(http.StatusOK, w.Code)
	is.Equal("text/plain; charset=utf-8", w.Header().Get("Content-Type")) // not json
	body := strings.TrimSpace(w.Body.String())
	is.Equal("{\"refNo\": \"1234\"}", body)
}

func TestHandler_ResData_Struct(t *testing.T) {
	is := is.New(t)
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
	is.Equal(http.StatusOK, w.Code)
	is.Equal("application/json; charset=UTF-8", w.Header().Get("Content-Type")) // not json
	body := strings.TrimSpace(w.Body.String())
	is.Equal("{\"refNo\":\"1234\"}", body)
}

func TestHandler_ResData_Nil(t *testing.T) {
	is := is.New(t)
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
	is.Equal(http.StatusOK, w.Code)
	is.Equal("application/json; charset=UTF-8", w.Header().Get("Content-Type"))
	body := strings.TrimSpace(w.Body.String())
	is.Equal("{}", body)
}

func TestHandler_ResData_HugeNum(t *testing.T) {
	is := is.New(t)
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
	is.Equal(http.StatusOK, w.Code)
	is.Equal("text/plain; charset=utf-8", w.Header().Get("Content-Type")) // not json
	body := strings.TrimSpace(w.Body.String())
	is.Equal("", body)
}

func TestHandler_ResRedirectPath_DefaultCode(t *testing.T) {
	is := is.New(t)
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
	is.Equal(http.StatusSeeOther, w.Code)
	is.Equal("text/html; charset=utf-8", w.Header().Get("Content-Type"))
	body := strings.TrimSpace(w.Body.String())
	is.Equal("<a href=\"/login\">See Other</a>.", body)
}

func TestHandler_ResRedirectPath_MovedPermanently(t *testing.T) {
	is := is.New(t)
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
	is.Equal(http.StatusMovedPermanently, w.Code)
	body := strings.TrimSpace(w.Body.String())
	is.Equal("<a href=\"/login\">Moved Permanently</a>.", body)
}

func TestHandler_ResHeader(t *testing.T) {
	is := is.New(t)
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
	is.Equal(http.StatusOK, w.Code)
	is.Equal("en", w.Header().Get("Content-Language"))
}

func TestHandler_CodeMapping(t *testing.T) {
	is := is.New(t)
	{
		handlerFunc := TranslateHandler(func(req Request) (res *Response, err error) {
			return nil, NewError(Code(len(_Code_index)-1), "", nil) // added by Saeed Rasooli
		})
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "", nil)
		handlerFunc(w, r)
		is.Equal(http.StatusInternalServerError, w.Code)
	}
	{
		handlerFunc := TranslateHandler(func(req Request) (res *Response, err error) {
			return nil, NewError(Canceled, "", nil)
		})
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "", nil)
		handlerFunc(w, r)
		is.Equal(http.StatusRequestTimeout, w.Code)
	}
	{
		handlerFunc := TranslateHandler(func(req Request) (res *Response, err error) {
			return nil, NewError(InvalidArgument, "", nil)
		})
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "", nil)
		handlerFunc(w, r)
		is.Equal(http.StatusBadRequest, w.Code)
	}
	{
		handlerFunc := TranslateHandler(func(req Request) (res *Response, err error) {
			return nil, NewError(DeadlineExceeded, "", nil)
		})
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "", nil)
		handlerFunc(w, r)
		is.Equal(http.StatusRequestTimeout, w.Code)
	}
	{
		handlerFunc := TranslateHandler(func(req Request) (res *Response, err error) {
			return nil, NewError(NotFound, "", nil)
		})
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "", nil)
		handlerFunc(w, r)
		is.Equal(http.StatusNotFound, w.Code)
	}
	{
		handlerFunc := TranslateHandler(func(req Request) (res *Response, err error) {
			return nil, NewError(AlreadyExists, "", nil)
		})
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "", nil)
		handlerFunc(w, r)
		is.Equal(http.StatusConflict, w.Code)
	}
	{
		handlerFunc := TranslateHandler(func(req Request) (res *Response, err error) {
			return nil, NewError(PermissionDenied, "", nil)
		})
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "", nil)
		handlerFunc(w, r)
		is.Equal(http.StatusForbidden, w.Code)
	}
	{
		handlerFunc := TranslateHandler(func(req Request) (res *Response, err error) {
			return nil, NewError(Unauthenticated, "", nil)
		})
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "", nil)
		handlerFunc(w, r)
		is.Equal(http.StatusUnauthorized, w.Code)
	}
	{
		handlerFunc := TranslateHandler(func(req Request) (res *Response, err error) {
			return nil, NewError(ResourceExhausted, "", nil)
		})
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "", nil)
		handlerFunc(w, r)
		is.Equal(http.StatusForbidden, w.Code)
	}
	{
		handlerFunc := TranslateHandler(func(req Request) (res *Response, err error) {
			return nil, NewError(FailedPrecondition, "", nil)
		})
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "", nil)
		handlerFunc(w, r)
		is.Equal(http.StatusPreconditionFailed, w.Code)
	}
	{
		handlerFunc := TranslateHandler(func(req Request) (res *Response, err error) {
			return nil, NewError(Aborted, "", nil)
		})
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "", nil)
		handlerFunc(w, r)
		is.Equal(http.StatusConflict, w.Code)
	}
	{
		handlerFunc := TranslateHandler(func(req Request) (res *Response, err error) {
			return nil, NewError(OutOfRange, "", nil)
		})
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "", nil)
		handlerFunc(w, r)
		is.Equal(http.StatusBadRequest, w.Code)
	}
	{
		handlerFunc := TranslateHandler(func(req Request) (res *Response, err error) {
			return nil, NewError(Unimplemented, "", nil)
		})
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "", nil)
		handlerFunc(w, r)
		is.Equal(http.StatusNotImplemented, w.Code)
	}
	{
		handlerFunc := TranslateHandler(func(req Request) (res *Response, err error) {
			return nil, NewError(Unavailable, "", nil)
		})
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "", nil)
		handlerFunc(w, r)
		is.Equal(http.StatusServiceUnavailable, w.Code)
	}
	{
		handlerFunc := TranslateHandler(func(req Request) (res *Response, err error) {
			return nil, NewError(DataLoss, "", nil)
		})
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "", nil)
		handlerFunc(w, r)
		is.Equal(http.StatusInternalServerError, w.Code)
	}
	{
		handlerFunc := TranslateHandler(func(req Request) (res *Response, err error) {
			return nil, NewError(MissingArgument, "", nil) // added by Saeed Rasooli
		})
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "", nil)
		handlerFunc(w, r)
		is.Equal(http.StatusBadRequest, w.Code)
	}
}
func TestHandler_Full_Happy(t *testing.T) {
	is := is.New(t)
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
		is.Equal(http.StatusOK, w.Code)
		resBody := strings.TrimSpace(w.Body.String())
		resMap := map[string]interface{}{}
		err = json.Unmarshal([]byte(resBody), &resMap)
		if err != nil {
			panic(err)
		}
		is.Equal("John", resMap["firstName"])
		is.Equal("Smith", resMap["lastName"])
		is.Equal("", resMap["medianName"])
		is.Equal(30, resMap["age"])
		is.Equal(true, resMap["subscribed"])
		is.Equal([]interface{}{"Tech", "Sports"}, resMap["interests"])
		is.Equal(10, resMap["count"])
		is.Equal(20, resMap["maxCount"])
	}
	{
		r, err := http.NewRequest("POST", myUrlStr, strings.NewReader(`{}`))
		if err != nil {
			panic(err)
		}
		r.RemoteAddr = "127.0.0.1:1234"
		w := httptest.NewRecorder()
		handlerFunc(w, r)
		is.Equal(http.StatusBadRequest, w.Code)
		resBody := strings.TrimSpace(w.Body.String())
		is.Equal("{\"code\":\"MissingArgument\",\"error\":\"missing 'firstName'\"}", resBody)
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
		is.Equal(http.StatusBadRequest, w.Code)
		resBody := strings.TrimSpace(w.Body.String())
		is.Equal("{\"code\":\"MissingArgument\",\"error\":\"missing 'lastName'\"}", resBody)
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
		is.Equal(http.StatusBadRequest, w.Code)
		resBody := strings.TrimSpace(w.Body.String())
		is.Equal("{\"code\":\"MissingArgument\",\"error\":\"missing 'age'\"}", resBody)
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
		is.Equal(http.StatusBadRequest, w.Code)
		resBody := strings.TrimSpace(w.Body.String())
		is.Equal("{\"code\":\"InvalidArgument\",\"error\":\"invalid 'age', must be float\"}", resBody)
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
		is.Equal(http.StatusBadRequest, w.Code)
		resBody := strings.TrimSpace(w.Body.String())
		is.Equal("{\"code\":\"MissingArgument\",\"error\":\"missing 'subscribed'\"}", resBody)
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
		is.Equal(http.StatusBadRequest, w.Code)
		resBody := strings.TrimSpace(w.Body.String())
		is.Equal("{\"code\":\"InvalidArgument\",\"error\":\"invalid 'subscribed', must be true or false\"}", resBody)
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
		is.Equal(http.StatusBadRequest, w.Code)
		resBody := strings.TrimSpace(w.Body.String())
		is.Equal("{\"code\":\"MissingArgument\",\"error\":\"missing 'interests'\"}", resBody)
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
		is.Equal(http.StatusBadRequest, w.Code)
		resBody := strings.TrimSpace(w.Body.String())
		is.Equal("{\"code\":\"InvalidArgument\",\"error\":\"invalid 'interests', must be array of strings\"}", resBody)
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
		is.Equal(http.StatusBadRequest, w.Code)
		resBody := strings.TrimSpace(w.Body.String())
		is.Equal("{\"code\":\"MissingArgument\",\"error\":\"missing 'count'\"}", resBody)
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
		is.Equal(http.StatusBadRequest, w.Code)
		resBody := strings.TrimSpace(w.Body.String())
		is.Equal("{\"code\":\"InvalidArgument\",\"error\":\"invalid 'count', must be integer\"}", resBody)
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
		is.Equal(http.StatusOK, w.Code)
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
		is.Equal(http.StatusBadRequest, w.Code)
		resBody := strings.TrimSpace(w.Body.String())
		is.Equal("{\"code\":\"InvalidArgument\",\"error\":\"invalid 'maxCount', must be integer\"}", resBody)
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
		is.Equal(http.StatusBadRequest, w.Code)
		resBody := strings.TrimSpace(w.Body.String())
		is.Equal("{\"code\":\"MissingArgument\",\"error\":\"missing 'birthDate'\"}", resBody)
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
		is.Equal(http.StatusBadRequest, w.Code)
		resBody := strings.TrimSpace(w.Body.String())
		is.Equal("{\"code\":\"InvalidArgument\",\"error\":\"invalid 'birthDate', must be a compatible object\"}", resBody)
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
		is.Equal(http.StatusOK, w.Code)
		t.Log(w.Body.String())
	}
}

func TestHandler_2(t *testing.T) {
	is := is.New(t)
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
		alias, err := req.GetStringDefault("alias", "")
		if err != nil {
			return nil, err
		}
		return &Response{
			Data: map[string]interface{}{
				"firstName": *firstName,
				"lastName":  *lastName,
				"alias":     alias,
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
		is.Equal(http.StatusBadRequest, w.Code)
		resBody := strings.TrimSpace(w.Body.String())
		is.Equal("{\"code\":\"InvalidArgument\",\"error\":\"invalid 'firstName', must be string\"}", resBody)
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
		is.Equal(http.StatusBadRequest, w.Code)
		resBody := strings.TrimSpace(w.Body.String())
		is.Equal("{\"code\":\"MissingArgument\",\"error\":\"missing 'unsubTime'\"}", resBody)
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
		is.Equal(http.StatusBadRequest, w.Code)
		resBody := strings.TrimSpace(w.Body.String())
		is.Equal("{\"code\":\"InvalidArgument\",\"error\":\"invalid 'unsubTime', must be RFC3339 time string\"}", resBody)
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
		is.Equal(http.StatusOK, w.Code)
		is.Equal(w.Body.String(), `{"alias":"","firstName":"John","lastName":"Smith","unsubTime":"2017-12-20T17:30:00Z"}`)
	}
	{
		r, err := http.NewRequest("POST", myUrlStr, strings.NewReader(`{
			"firstName": "John",
			"lastName": "Smith",
			"alias": 123,
			"unsubTime": "2017-12-20T17:30:00Z"
		}`))
		if err != nil {
			panic(err)
		}
		r.RemoteAddr = "127.0.0.1:1234"
		w := httptest.NewRecorder()
		handlerFunc(w, r)
		is.Equal(http.StatusBadRequest, w.Code)
		resBody := strings.TrimSpace(w.Body.String())
		is.Equal(resBody, `{"code":"InvalidArgument","error":"invalid 'alias', must be string"}`)
	}
	{
		r, err := http.NewRequest("POST", myUrlStr, strings.NewReader(`{
			"firstName": "John",
			"lastName": "Smith",
			"alias": "jsm",
			"unsubTime": "2017-12-20T17:30:00Z"
		}`))
		if err != nil {
			panic(err)
		}
		r.RemoteAddr = "127.0.0.1:1234"
		w := httptest.NewRecorder()
		handlerFunc(w, r)
		is.Equal(http.StatusOK, w.Code)
		resBody := strings.TrimSpace(w.Body.String())
		is.Equal(resBody, `{"alias":"jsm","firstName":"John","lastName":"Smith","unsubTime":"2017-12-20T17:30:00Z"}`)
	}
}

func TestHandler_3(t *testing.T) {
	is := is.New(t)
	myUrlStr := "http://127.0.0.1/test/full"
	handlerFunc := TranslateHandler(func(req Request) (*Response, error) {
		start, err := req.GetFloatDefault("start", 0)
		if err != nil {
			return nil, err
		}
		end, err := req.GetFloatDefault("end", 10)
		if err != nil {
			return nil, err
		}
		step, err := req.GetFloatDefault("step", 1)
		if err != nil {
			return nil, err
		}
		nums := []float64{}
		for x := start; x < end; x += step {
			nums = append(nums, x)
		}
		return &Response{
			Data: nums,
		}, nil
	})
	{
		r, err := http.NewRequest("POST", myUrlStr, strings.NewReader(`{}`))
		if err != nil {
			panic(err)
		}
		r.RemoteAddr = "127.0.0.1:1234"
		w := httptest.NewRecorder()
		handlerFunc(w, r)
		is.Equal(http.StatusOK, w.Code)
		resBody := strings.TrimSpace(w.Body.String())
		is.Equal(`[0,1,2,3,4,5,6,7,8,9]`, resBody)
	}
	{
		r, err := http.NewRequest("POST", myUrlStr, strings.NewReader(`{}`))
		if err != nil {
			panic(err)
		}
		r.RemoteAddr = "127.0.0.1:1234"
		w := httptest.NewRecorder()
		handlerFunc(w, r)
		is.Equal(http.StatusOK, w.Code)
		resBody := strings.TrimSpace(w.Body.String())
		is.Equal(`[0,1,2,3,4,5,6,7,8,9]`, resBody)
	}
	{
		r, err := http.NewRequest("POST", myUrlStr, strings.NewReader(`{
			"start": 6
		}`))
		if err != nil {
			panic(err)
		}
		r.RemoteAddr = "127.0.0.1:1234"
		w := httptest.NewRecorder()
		handlerFunc(w, r)
		is.Equal(http.StatusOK, w.Code)
		resBody := strings.TrimSpace(w.Body.String())
		is.Equal(`[6,7,8,9]`, resBody)
	}
	{
		r, err := http.NewRequest("POST", myUrlStr, strings.NewReader(`{
			"start": 15,
			"end": 19
		}`))
		if err != nil {
			panic(err)
		}
		r.RemoteAddr = "127.0.0.1:1234"
		w := httptest.NewRecorder()
		handlerFunc(w, r)
		is.Equal(http.StatusOK, w.Code)
		resBody := strings.TrimSpace(w.Body.String())
		is.Equal(`[15,16,17,18]`, resBody)
	}
	{
		r, err := http.NewRequest("POST", myUrlStr, strings.NewReader(`{
			"start": 15,
			"end": 19,
			"step": 2
		}`))
		if err != nil {
			panic(err)
		}
		r.RemoteAddr = "127.0.0.1:1234"
		w := httptest.NewRecorder()
		handlerFunc(w, r)
		is.Equal(http.StatusOK, w.Code)
		resBody := strings.TrimSpace(w.Body.String())
		is.Equal(`[15,17]`, resBody)
	}
	{
		r, err := http.NewRequest("POST", myUrlStr, strings.NewReader(`{
			"start": 15,
			"end": 19,
			"step": "2"
		}`))
		if err != nil {
			panic(err)
		}
		r.RemoteAddr = "127.0.0.1:1234"
		w := httptest.NewRecorder()
		handlerFunc(w, r)
		is.Equal(http.StatusBadRequest, w.Code)
		resBody := strings.TrimSpace(w.Body.String())
		is.Equal(`{"code":"InvalidArgument","error":"invalid 'step', must be float"}`, resBody)
	}
}
