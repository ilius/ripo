package restpc

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

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
