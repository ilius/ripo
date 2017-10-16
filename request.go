package restpc

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"time"
)

type Request interface {
	RemoteIP() (string, error)
	URL() *url.URL
	Host() string
	HandlerName() string

	Body() ([]byte, error)
	BodyMap() (map[string]interface{}, error)
	BodyTo(model interface{}) error

	GetHeader(string) string
	GetFormValue(key string) string
	Context() context.Context

	GetString(key string, sources ...FromX) (*string, error)
	GetStringList(key string, sources ...FromX) ([]string, error)
	GetInt(key string, sources ...FromX) (*int, error)
	GetIntDefault(key string, defaultValue int, sources ...FromX) (int, error)
	GetFloat(key string, sources ...FromX) (*float64, error)
	GetBool(key string, sources ...FromX) (*bool, error)
	GetTime(key string, sources ...FromX) (*time.Time, error)

	FullMap() map[string]interface{}
}

type FromX interface {
	GetString(req Request, key string) (*string, error)
	GetStringList(req Request, key string) ([]string, error)
	GetInt(req Request, key string) (*int, error)
	GetFloat(req Request, key string) (*float64, error)
	GetBool(req Request, key string) (*bool, error)
	GetTime(req Request, key string) (*time.Time, error)
}

var defaultParamSources = []FromX{
	FromBody,
	FromForm,
	// FromContext, // I don't have any use case for it, enable if you want
	// FromEmpty, // makes every param optional, enable if you want
}

// SetDefaultParamSources: set default parameter sources for req.Get* methods
// Typical arguments (that are implemented by the library): FromBody, FromForm, FromContext, FromEmpty
// Adding `FromEmpty` at the end, makes the parameter optional, meaning Get* methods return empty value
// with no error if the parameter is missing (or empty) in all these parameter sources
// You can also write your own implementation of `FromX` interface, and pass it here
func SetDefaultParamSources(sources ...FromX) {
	if len(sources) == 0 {
		panic("SetDefaultParamSources: no arguments given")
	}
	defaultParamSources = sources
}

type requestImp struct {
	r           *http.Request // must be set initially
	handlerName string        // must be set initially
	body        []byte
	bodyErr     error
	bodyMap     map[string]interface{}
	bodyMapErr  error
}

func (req *requestImp) RemoteIP() (string, error) {
	remoteIp, _, err := net.SplitHostPort(req.r.RemoteAddr)
	if err != nil {
		return "", NewError(
			Internal, "", err,
			"r.RemoteAddr", req.r.RemoteAddr,
		)
	}
	return remoteIp, nil
}

func (req *requestImp) URL() *url.URL {
	u := *req.r.URL
	return &u
}

func (req *requestImp) Host() string {
	return req.r.Host
}

func (req *requestImp) HandlerName() string {
	return req.handlerName
}

func (req *requestImp) Body() ([]byte, error) {
	if req.body != nil {
		return req.body, nil
	}
	if req.bodyErr != nil {
		return nil, req.bodyErr
	}
	if req.r.Body == nil {
		return nil, nil
	}
	body, err := ioutil.ReadAll(req.r.Body)
	if err != nil {
		req.bodyErr = err
		log.Println(err)
	}
	req.body = body
	req.r.Body.Close()
	req.r.Body = nil
	return body, nil
}

func (req *requestImp) BodyMap() (map[string]interface{}, error) {
	if req.bodyMap != nil {
		return req.bodyMap, nil
	}
	if req.bodyMapErr != nil {
		return nil, req.bodyMapErr
	}
	data := map[string]interface{}{}
	body, err := req.Body()
	if err != nil {
		req.bodyMapErr = err
		return nil, err
	}
	if len(body) > 0 {
		err = json.Unmarshal(body, &data)
		if err != nil {
			req.bodyMapErr = err
			log.Println(err)
			// return nil, err // FIXME
		}
	}
	req.bodyMap = data
	return data, nil
}

func (req *requestImp) BodyTo(model interface{}) error {
	body, err := req.Body()
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, model)
	if err != nil {
		return NewError(InvalidArgument, "request body is not a valid json", err)
	}
	return nil
}

func (req *requestImp) GetHeader(key string) string {
	return req.r.Header.Get(key)
}

func (req *requestImp) GetFormValue(key string) string {
	return req.r.FormValue(key)
}

func (req *requestImp) Context() context.Context {
	return req.r.Context()
}

func (req *requestImp) GetString(key string, sources ...FromX) (*string, error) {
	if len(sources) == 0 {
		sources = defaultParamSources
	}
	for _, source := range sources {
		value, err := source.GetString(req, key)
		if err != nil {
			return nil, err
		}
		if value != nil {
			return value, nil
		}
	}
	return nil, NewError(
		MissingArgument,
		fmt.Sprintf("missing '%v'", key),
		nil,
	)
}

func (req *requestImp) GetStringList(key string, sources ...FromX) ([]string, error) {
	if len(sources) == 0 {
		sources = defaultParamSources
	}
	for _, source := range sources {
		value, err := source.GetStringList(req, key)
		if err != nil {
			return nil, err
		}
		if value != nil {
			return value, nil
		}
	}
	return nil, NewError(
		MissingArgument,
		fmt.Sprintf("missing '%v'", key),
		nil,
	)
}

func (req *requestImp) GetInt(key string, sources ...FromX) (*int, error) {
	if len(sources) == 0 {
		sources = defaultParamSources
	}
	for _, source := range sources {
		value, err := source.GetInt(req, key)
		if err != nil {
			return nil, err
		}
		if value != nil {
			return value, nil
		}
	}
	return nil, NewError(
		MissingArgument,
		fmt.Sprintf("missing '%v'", key),
		nil,
	)
}

func (req *requestImp) GetIntDefault(key string, defaultValue int, sources ...FromX) (int, error) {
	if len(sources) == 0 {
		sources = []FromX{
			FromBody,
			FromForm,
		}
	}
	for _, source := range sources {
		value, err := source.GetInt(req, key)
		if err != nil {
			return defaultValue, err
		}
		if value != nil {
			return *value, nil
		}
	}
	return defaultValue, nil
}

func (req *requestImp) GetFloat(key string, sources ...FromX) (*float64, error) {
	if len(sources) == 0 {
		sources = defaultParamSources
	}
	for _, source := range sources {
		value, err := source.GetFloat(req, key)
		if err != nil {
			return nil, err
		}
		if value != nil {
			return value, nil
		}
	}
	return nil, NewError(
		MissingArgument,
		fmt.Sprintf("missing '%v'", key),
		nil,
	)
}

func (req *requestImp) GetBool(key string, sources ...FromX) (*bool, error) {
	if len(sources) == 0 {
		sources = defaultParamSources
	}
	for _, source := range sources {
		value, err := source.GetBool(req, key)
		if err != nil {
			return nil, err
		}
		if value != nil {
			return value, nil
		}
	}
	return nil, NewError(
		MissingArgument,
		fmt.Sprintf("missing '%v'", key),
		nil,
	)
}

func (req *requestImp) GetTime(key string, sources ...FromX) (*time.Time, error) {
	if len(sources) == 0 {
		sources = defaultParamSources
	}
	for _, source := range sources {
		value, err := source.GetTime(req, key)
		if err != nil {
			return nil, err
		}
		if value != nil {
			return value, nil
		}
	}
	return nil, NewError(
		MissingArgument,
		fmt.Sprintf("missing '%v'", key),
		nil,
	)
}

func (req *requestImp) HeaderCopy() http.Header {
	header := http.Header{}
	for key, values := range req.r.Header {
		header[key] = values
	}
	return header
}

func (req *requestImp) HeaderStrippedAuth() http.Header {
	header := req.HeaderCopy()
	authHader, ok := header["Authorization"]
	if ok {
		authHaderNew := make([]string, len(authHader))
		for i := 0; i < len(authHader); i++ {
			authHaderNew[i] = "[REMOVED]"
		}
		header["Authorization"] = authHaderNew
	}
	return header
}

func (req *requestImp) FullMap() map[string]interface{} {
	bodyMap, _ := req.BodyMap()
	urlStr := req.URL().String()
	remoteIP, _ := req.RemoteIP()
	return map[string]interface{}{
		"bodyMap":  bodyMap,
		"url":      urlStr,
		"form":     req.r.Form,
		"header":   req.HeaderStrippedAuth(),
		"remoteIP": remoteIP,
	}
}
