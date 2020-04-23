package ripo

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ilius/is/v2"
)

func callSetDefaultParamSources(sources ...FromX) (err error) {
	defer func() {
		r := recover()
		if r != nil {
			err = fmt.Errorf("Panic: %v", r)
		}
	}()
	SetDefaultParamSources(sources...)
	return
}

func TestSetDefaultParamSources(t *testing.T) {
	defer func() {
		r := recover()
		if r != nil {
			t.Errorf("Panic: %v", r)
		}
	}()
	is := is.New(t)
	{
		err := callSetDefaultParamSources()
		AssertError(t, err, Unknown, "Panic: SetDefaultParamSources: no arguments given")
	}
	is.NotErr(callSetDefaultParamSources(FromBody))
	is.NotErr(callSetDefaultParamSources(FromBody, FromContext))
	is.NotErr(callSetDefaultParamSources(FromBody, FromContext, FromEmpty, FromEmpty))
	is.NotErr(callSetDefaultParamSources(FromEmpty))
}

func Test_requestImp_URL(t *testing.T) {
	is := is.New(t)
	r, err := http.NewRequest("GET", "http://127.0.0.1/test1", nil)
	if err != nil {
		panic(err)
	}
	req := &requestImp{
		r:           r,
		handlerName: "Test",
	}
	is.Equal("/test1", req.URL().Path)
	u := req.URL()
	u.Path = "/test2"
	is.Equal("/test1", req.URL().Path)
	is.Equal("/test2", u.Path)
}

func Test_requestImp_FullMap(t *testing.T) {
	is := is.New(t)
	r, err := http.NewRequest("POST", "http://127.0.0.1/test1?name=John", nil)
	if err != nil {
		panic(err)
	}
	r.Header.Add("Authorization", "bearer foobar")
	req := &requestImp{
		r:           r,
		handlerName: "Test",
	}
	fullMap := req.FullMap()
	expectedFullMap := map[string]interface{}{
		"url":     "http://127.0.0.1/test1?name=John",
		"bodyMap": map[string]interface{}{},
		"header": http.Header{
			"Authorization": []string{"[REMOVED]"},
		},
		"remoteIP": "",
		"form": url.Values{
			"name": []string{"John"},
		},
	}
	is.Equal(expectedFullMap, fullMap)
	t.Log(fullMap)
}

func Test_requestImp_Body_Json(t *testing.T) {
	is := is.New(t)
	bodyStr := `{
		"firstName": "John",
		"lastName": "Smith"
	}`
	r, err := http.NewRequest("POST", "http://127.0.0.1/test", strings.NewReader(bodyStr))
	is.NotErr(err)
	req := &requestImp{
		r:           r,
		handlerName: "Test",
	}
	{
		bodyMap, err := req.BodyMap()
		is.NotErr(err)
		is.Equal(map[string]interface{}{
			"firstName": "John",
			"lastName":  "Smith",
		}, bodyMap)
	}
	{
		body, err := req.Body()
		is.NotErr(err)
		is.Equal(bodyStr, string(body))
	}
}

func Test_requestImp_Body_NonJson(t *testing.T) {
	is := is.New(t)
	bodyStr := `hello world`
	r, err := http.NewRequest("POST", "http://127.0.0.1/test", strings.NewReader(bodyStr))
	is.NotErr(err)
	req := &requestImp{
		r:           r,
		handlerName: "Test",
	}
	{
		bodyMap, err := req.BodyMap()
		AssertError(t, err, InvalidArgument, "request body is not a valid json")
		is.Nil(bodyMap)
	}
	{
		body, err := req.Body()
		is.NotErr(err)
		is.Equal(bodyStr, string(body))
	}
}

func Test_requestImp_BodyTo_OK(t *testing.T) {
	is := is.New(t)
	bodyStr := `{
		"firstName": "John",
		"lastName": "Smith"
	}`
	bodyStruct := struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	}{}
	r, err := http.NewRequest("POST", "http://127.0.0.1/test", strings.NewReader(bodyStr))
	is.NotErr(err)
	req := &requestImp{
		r:           r,
		handlerName: "Test",
	}
	{
		err := req.BodyTo(&bodyStruct)
		is.NotErr(err)
		is.Equal("John", bodyStruct.FirstName)
		is.Equal("Smith", bodyStruct.LastName)
	}
}

func Test_requestImp_BodyTo_Bad(t *testing.T) {
	is := is.New(t)
	bodyStr := `{
		"firstName": "John",
		"lastName": "Smith",
	}`
	bodyStruct := struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	}{}
	r, err := http.NewRequest("POST", "http://127.0.0.1/test", strings.NewReader(bodyStr))
	is.NotErr(err)
	req := &requestImp{
		r:           r,
		handlerName: "Test",
	}
	{
		err := req.BodyTo(&bodyStruct)
		AssertError(t, err, InvalidArgument, "request body is not a valid json")
		is.Equal("", bodyStruct.FirstName)
		is.Equal("", bodyStruct.LastName)
	}
}

func Test_requestImp_Header(t *testing.T) {
	is := is.New(t)
	r, err := http.NewRequest("POST", "http://127.0.0.1/test", nil)
	is.NotErr(err)
	r.Header.Add("Authorization", "bearer foobar")
	req := &requestImp{
		r:           r,
		handlerName: "Test",
	}
	is.Equal("bearer foobar", req.Header("Authorization"))
	is.Equal("", req.Header("foo"))
}

func Test_requestImp_HeaderKeys(t *testing.T) {
	is := is.New(t)
	r, err := http.NewRequest("POST", "http://127.0.0.1/test", nil)
	is.NotErr(err)
	r.Header.Add("Authorization", "bearer foobar")
	req := &requestImp{
		r:           r,
		handlerName: "Test",
	}
	is.Equal([]string{"Authorization"}, req.HeaderKeys())
}

func Test_requestImp_Cookie(t *testing.T) {
	is := is.New(t)
	r, err := http.NewRequest("POST", "http://127.0.0.1/test", nil)
	is.NotErr(err)
	expectedCookie := &http.Cookie{
		Name: "token",
	}
	r.AddCookie(expectedCookie)
	req := &requestImp{
		r:           r,
		handlerName: "Test",
	}
	actualCookie, err := req.Cookie("token")
	if err != nil {
		panic(err)
	}
	is.Equal(expectedCookie, actualCookie)
	is.Equal([]string{"token"}, req.CookieNames())
}

func Test_requestImp_Context(t *testing.T) {
	is := is.New(t)
	r, err := http.NewRequest("POST", "http://127.0.0.1/test", nil)
	is.NotErr(err)
	req := &requestImp{
		r:           r,
		handlerName: "Test",
	}
	ctx := req.Context()
	is.Equal(context.Background(), ctx)
}

func Test_requestImp_MockBody(t *testing.T) {
	is := is.New(t)
	ctrl := gomock.NewController(t)
	mockBody := NewMockReadCloser(ctrl)
	r, err := http.NewRequest("POST", "http://127.0.0.1/test", mockBody)
	is.NotErr(err)
	req := &requestImp{
		r:           r,
		handlerName: "Test",
	}
	{
		mockBody.EXPECT().Read(gomock.Any()).Return(0, fmt.Errorf("no data for you"))
		mockBody.EXPECT().Close()
		body, err := req.Body()
		AssertError(t, err, Unknown, "no data for you")
		is.Nil(body)
	}
	{
		body, err := req.Body()
		AssertError(t, err, Unknown, "no data for you")
		is.Nil(body)
	}
	{
		bodyMap, err := req.BodyMap()
		AssertError(t, err, Unknown, "no data for you")
		is.Nil(bodyMap)
	}
	{
		bodyMap := map[string]string{}
		err := req.BodyTo(&bodyMap)
		AssertError(t, err, Unknown, "no data for you")
		is.Equal(0, len(bodyMap))
	}
}

func Test_requestImp_MockBody2(t *testing.T) {
	is := is.New(t)
	ctrl := gomock.NewController(t)
	mockBody := NewMockReadCloser(ctrl)
	r, err := http.NewRequest("POST", "http://127.0.0.1/test", mockBody)
	is.NotErr(err)
	req := &requestImp{
		r:           r,
		handlerName: "Test",
	}
	{
		mockBody.EXPECT().Read(gomock.Any()).Do(func(b []byte) {
			b[0] = 'a'
		}).Return(1, nil).Return(1, io.EOF)
		mockBody.EXPECT().Close()
		body, err := req.Body()
		is.NotErr(err)
		is.Equal([]byte("a"), body)
	}
	{
		bodyMap, err := req.BodyMap()
		AssertError(t, err, InvalidArgument, "request body is not a valid json")
		is.Nil(bodyMap)
	}
	{
		bodyMap, err := req.BodyMap()
		AssertError(t, err, InvalidArgument, "request body is not a valid json")
		is.Nil(bodyMap)
	}
	{
		bodyMap := map[string]string{}
		err := req.BodyTo(&bodyMap)
		AssertError(t, err, InvalidArgument, "request body is not a valid json")
		is.Equal(0, len(bodyMap))
	}
}
