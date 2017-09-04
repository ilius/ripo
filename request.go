package myrpc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Request interface {
	RemoteIP() (string, error)
	URL() *url.URL
	Host() string
	Body() ([]byte, error)
	GetString(key string, flags ...ParamFlag) (*string, error)
	GetInt(key string, flags ...ParamFlag) (*int, error)
	GetFloat(key string, flags ...ParamFlag) (*float64, error)
	GetTime(key string, flags ...ParamFlag) (*time.Time, error)
	JSONData() (map[string]interface{}, error)
	GetHeader(string) string
}

type requestImp struct {
	r        *http.Request
	body     []byte
	jsonData map[string]interface{}
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
	return req.r.URL
}

func (req *requestImp) Host() string {
	return req.r.Host
}

func (req *requestImp) Body() ([]byte, error) {
	if req.body != nil {
		return req.body, nil
	}
	if req.r.Body == nil {
		return nil, nil
	}
	body, err := ioutil.ReadAll(req.r.Body)
	if err != nil {
		log.Println(err)
	}
	req.body = body
	req.r.Body.Close()
	req.r.Body = nil
	return body, nil
}

func (req *requestImp) GetString(key string, flags ...ParamFlag) (*string, error) {
	flag := mergeParamFlags(flags...)
	if flag.FromJSON() {
		data, err := req.JSONData()
		if err != nil {
			return nil, err
		}
		valueIn := data[key]
		if valueIn != nil {
			switch value := valueIn.(type) {
			case string:
				valueStr := value // to copy
				return &valueStr, nil
			case []byte:
				valueStr := string(value)
				return &valueStr, nil
			default:
				return nil, NewError(
					InvalidArgument,
					fmt.Sprintf("invalid '%v', must be string", key),
					nil,
				)
			}
		}
	}
	if flag.FromForm() {
		value := req.r.FormValue(key)
		if value != "" {
			return &value, nil
		}
	}
	if flag.Mandatory() {
		return nil, NewError(
			InvalidArgument,
			fmt.Sprintf("missing '%v'", key),
			nil,
		)
	}
	return nil, nil
}

func (req *requestImp) GetInt(key string, flags ...ParamFlag) (*int, error) {
	flag := mergeParamFlags(flags...)
	if flag.FromJSON() {
		data, err := req.JSONData()
		if err != nil {
			return nil, err
		}
		valueIn := data[key]
		if valueIn != nil {
			switch value := valueIn.(type) {
			case int:
				valueInt := value // to copy
				return &valueInt, nil
			case int32:
				valueInt := int(value)
				return &valueInt, nil
			case int64:
				valueInt := int(value)
				return &valueInt, nil
			default:
				return nil, NewError(
					InvalidArgument,
					fmt.Sprintf("invalid '%v', must be integer", key),
					nil,
				)
			}
		}
	}
	if flag.FromForm() {
		valueStr := req.r.FormValue(key)
		if valueStr != "" {
			value, err := strconv.ParseInt(valueStr, 10, 64)
			if err != nil {
				return nil, NewError(
					InvalidArgument,
					fmt.Sprintf("invalid '%v', must be integer", key),
					err,
				)
			}
			valueInt := int(value)
			return &valueInt, nil
		}
	}
	if flag.Mandatory() {
		return nil, NewError(
			InvalidArgument,
			fmt.Sprintf("missing '%v'", key),
			nil,
		)
	}
	return nil, nil
}

func (req *requestImp) GetFloat(key string, flags ...ParamFlag) (*float64, error) {
	flag := mergeParamFlags(flags...)
	if flag.FromJSON() {
		data, err := req.JSONData()
		if err != nil {
			return nil, err
		}
		valueIn := data[key]
		if valueIn != nil {
			switch value := valueIn.(type) {
			case float64:
				valueF := value // to copy
				return &valueF, nil
			case float32:
				valueF := float64(value)
				return &valueF, nil
			case int:
				valueF := float64(value)
				return &valueF, nil
			case int64:
				valueF := float64(value)
				return &valueF, nil
			case int32:
				valueF := float64(value)
				return &valueF, nil
			default:
				return nil, NewError(
					InvalidArgument,
					fmt.Sprintf("invalid '%v', must be float", key),
					nil,
				)
			}
		}
	}
	if flag.FromForm() {
		valueStr := req.r.FormValue(key)
		if valueStr != "" {
			value, err := strconv.ParseFloat(valueStr, 64)
			if err != nil {
				return nil, NewError(
					InvalidArgument,
					fmt.Sprintf("invalid '%v', must be float", key),
					err,
				)
			}
			valueF := float64(value)
			return &valueF, nil
		}
	}
	if flag.Mandatory() {
		return nil, NewError(
			InvalidArgument,
			fmt.Sprintf("missing '%v'", key),
			nil,
		)
	}
	return nil, nil
}

func (req *requestImp) GetTime(key string, flags ...ParamFlag) (*time.Time, error) {
	flag := mergeParamFlags(flags...)
	if flag.FromJSON() {
		data, err := req.JSONData()
		if err != nil {
			return nil, err
		}
		valueIn := data[key]
		if valueIn != nil {
			switch value := valueIn.(type) {
			case string:
				valueTm, err := time.Parse(time.RFC3339, value)
				if err != nil {
					return nil, NewError(
						InvalidArgument,
						fmt.Sprintf("invalid '%v', must be RFC3339 time string", key),
						err,
					)
				}
				return &valueTm, nil
			default:
				return nil, NewError(
					InvalidArgument,
					fmt.Sprintf("invalid '%v', must be RFC3339 time string", key),
					nil,
				)
			}
		}
	}
	if flag.FromForm() {
		valueStr := req.r.FormValue(key)
		if valueStr != "" {
			valueTm, err := time.Parse(time.RFC3339, valueStr)
			if err != nil {
				return nil, NewError(
					InvalidArgument,
					fmt.Sprintf("invalid '%v', must be RFC3339 time string", key),
					err,
				)
			}
			return &valueTm, nil
		}
	}
	if flag.Mandatory() {
		return nil, NewError(
			InvalidArgument,
			fmt.Sprintf("missing '%v'", key),
			nil,
		)
	}
	return nil, nil
}

func (req *requestImp) JSONData() (map[string]interface{}, error) {
	if req.jsonData != nil {
		return req.jsonData, nil
	}
	data := map[string]interface{}{}
	body, err := req.Body()
	if err != nil {
		return nil, err
	}
	if len(body) > 0 {
		json.Unmarshal(body, &data)
	}
	req.jsonData = data
	return data, nil
}

func (req *requestImp) GetHeader(key string) string {
	return req.r.Header.Get(key)
}
