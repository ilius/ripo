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

	"github.com/stretchr/testify/assert"
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
	{
		err := callSetDefaultParamSources()
		assert.EqualError(t, err, "Panic: SetDefaultParamSources: no arguments given")
	}
	assert.NoError(t, callSetDefaultParamSources(FromBody))
	assert.NoError(t, callSetDefaultParamSources(FromBody, FromContext))
	assert.NoError(t, callSetDefaultParamSources(FromBody, FromContext, FromEmpty, FromEmpty))
	assert.NoError(t, callSetDefaultParamSources(FromEmpty))
}

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

func Test_requestImp_FullMap(t *testing.T) {
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
	assert.Equal(t, expectedFullMap, fullMap)
	t.Log(fullMap)
}

func Test_requestImp_Body_Json(t *testing.T) {
	bodyStr := `{
		"firstName": "John",
		"lastName": "Smith"
	}`
	r, err := http.NewRequest("POST", "http://127.0.0.1/test", strings.NewReader(bodyStr))
	assert.NoError(t, err)
	req := &requestImp{
		r:           r,
		handlerName: "Test",
	}
	{
		bodyMap, err := req.BodyMap()
		assert.NoError(t, err)
		assert.Equal(t, map[string]interface{}{
			"firstName": "John",
			"lastName":  "Smith",
		}, bodyMap)
	}
	{
		body, err := req.Body()
		assert.NoError(t, err)
		assert.Equal(t, bodyStr, string(body))
	}
}

func Test_requestImp_Body_NonJson(t *testing.T) {
	bodyStr := `hello world`
	r, err := http.NewRequest("POST", "http://127.0.0.1/test", strings.NewReader(bodyStr))
	assert.NoError(t, err)
	req := &requestImp{
		r:           r,
		handlerName: "Test",
	}
	{
		bodyMap, err := req.BodyMap()
		assert.EqualError(t, err, "request body is not a valid json")
		assert.Nil(t, bodyMap)
	}
	{
		body, err := req.Body()
		assert.NoError(t, err)
		assert.Equal(t, bodyStr, string(body))
	}
}

func Test_requestImp_BodyTo_OK(t *testing.T) {
	bodyStr := `{
		"firstName": "John",
		"lastName": "Smith"
	}`
	bodyStruct := struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	}{}
	r, err := http.NewRequest("POST", "http://127.0.0.1/test", strings.NewReader(bodyStr))
	assert.NoError(t, err)
	req := &requestImp{
		r:           r,
		handlerName: "Test",
	}
	{
		err := req.BodyTo(&bodyStruct)
		assert.NoError(t, err)
		assert.Equal(t, "John", bodyStruct.FirstName)
		assert.Equal(t, "Smith", bodyStruct.LastName)
	}
}

func Test_requestImp_BodyTo_Bad(t *testing.T) {
	bodyStr := `{
		"firstName": "John",
		"lastName": "Smith",
	}`
	bodyStruct := struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	}{}
	r, err := http.NewRequest("POST", "http://127.0.0.1/test", strings.NewReader(bodyStr))
	assert.NoError(t, err)
	req := &requestImp{
		r:           r,
		handlerName: "Test",
	}
	{
		err := req.BodyTo(&bodyStruct)
		assert.EqualError(t, err, "request body is not a valid json")
		assert.Equal(t, "", bodyStruct.FirstName)
		assert.Equal(t, "", bodyStruct.LastName)
	}
}

func Test_requestImp_Header(t *testing.T) {
	r, err := http.NewRequest("POST", "http://127.0.0.1/test", nil)
	assert.NoError(t, err)
	r.Header.Add("Authorization", "bearer foobar")
	req := &requestImp{
		r:           r,
		handlerName: "Test",
	}
	assert.Equal(t, "bearer foobar", req.Header("Authorization"))
	assert.Equal(t, "", req.Header("foo"))
}

func Test_requestImp_Context(t *testing.T) {
	r, err := http.NewRequest("POST", "http://127.0.0.1/test", nil)
	assert.NoError(t, err)
	req := &requestImp{
		r:           r,
		handlerName: "Test",
	}
	ctx := req.Context()
	assert.Equal(t, context.Background(), ctx)
}

func Test_requestImp_MockBody(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockBody := NewMockReadCloser(ctrl)
	r, err := http.NewRequest("POST", "http://127.0.0.1/test", mockBody)
	assert.NoError(t, err)
	req := &requestImp{
		r:           r,
		handlerName: "Test",
	}
	{
		mockBody.EXPECT().Read(gomock.Any()).Return(0, fmt.Errorf("no data for you"))
		mockBody.EXPECT().Close()
		body, err := req.Body()
		assert.EqualError(t, err, "no data for you")
		assert.Nil(t, body)
	}
	{
		body, err := req.Body()
		assert.EqualError(t, err, "no data for you")
		assert.Nil(t, body)
	}
	{
		bodyMap, err := req.BodyMap()
		assert.EqualError(t, err, "no data for you")
		assert.Nil(t, bodyMap)
	}
	{
		bodyMap := map[string]string{}
		err := req.BodyTo(&bodyMap)
		assert.EqualError(t, err, "no data for you")
		assert.Equal(t, 0, len(bodyMap))
	}
}

func Test_requestImp_MockBody2(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockBody := NewMockReadCloser(ctrl)
	r, err := http.NewRequest("POST", "http://127.0.0.1/test", mockBody)
	assert.NoError(t, err)
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
		assert.NoError(t, err)
		assert.Equal(t, []byte("a"), body)
	}
	{
		bodyMap, err := req.BodyMap()
		assert.EqualError(t, err, "request body is not a valid json")
		assert.Nil(t, bodyMap)
	}
	{
		bodyMap, err := req.BodyMap()
		assert.EqualError(t, err, "request body is not a valid json")
		assert.Nil(t, bodyMap)
	}
	{
		bodyMap := map[string]string{}
		err := req.BodyTo(&bodyMap)
		assert.EqualError(t, err, "request body is not a valid json")
		assert.Equal(t, 0, len(bodyMap))
	}
}
