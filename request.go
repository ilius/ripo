package restpc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Request interface {
	RemoteIP() (string, error)
	URL() *url.URL
	Host() string

	Body() ([]byte, error)
	BodyMap() (map[string]interface{}, error)
	BodyTo(model interface{}) error

	GetHeader(string) string

	GetString(key string, flags ...ParamFlag) (*string, error)
	GetStringList(key string, flags ...ParamFlag) ([]string, error)
	GetInt(key string, flags ...ParamFlag) (*int, error)
	GetFloat(key string, flags ...ParamFlag) (*float64, error)
	GetBool(key string, flags ...ParamFlag) (*bool, error)
	GetTime(key string, flags ...ParamFlag) (*time.Time, error)

	FullMap() map[string]interface{}
}

type requestImp struct {
	r          *http.Request
	body       []byte
	bodyErr    error
	bodyMap    map[string]interface{}
	bodyMapErr error
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

func (req *requestImp) GetString(key string, flags ...ParamFlag) (*string, error) {
	flag := mergeParamFlags(flags...)
	if flag.FromBody() {
		data, err := req.BodyMap()
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
					"value", value,
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

func (req *requestImp) GetStringList(key string, flags ...ParamFlag) ([]string, error) {
	flag := mergeParamFlags(flags...)
	if flag.FromBody() {
		data, err := req.BodyMap()
		if err != nil {
			return nil, err
		}
		valueIn := data[key]
		if valueIn != nil {
			switch value := valueIn.(type) {
			case []string:
				valueSlice := append([]string(nil), value...) // to copy
				return valueSlice, nil
			default:
				return nil, NewError(
					InvalidArgument,
					fmt.Sprintf("invalid '%v', must be array of strings", key),
					nil,
					"value", value,
				)
			}
		}
	}
	if flag.FromForm() {
		valueStr := req.r.FormValue(key)
		if valueStr != "" {
			valueSlice := strings.Split(valueStr, ",")
			return valueSlice, nil
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
	if flag.FromBody() {
		data, err := req.BodyMap()
		if err != nil {
			return nil, err
		}
		valueIn := data[key]
		if valueIn != nil {
			switch value := valueIn.(type) {
			case float64:
				valueInt := int(value)
				return &valueInt, nil
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
					"value", value,
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
					"valueStr", valueStr,
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
	if flag.FromBody() {
		data, err := req.BodyMap()
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
					"value", value,
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
					"valueStr", valueStr,
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

func (req *requestImp) GetBool(key string, flags ...ParamFlag) (*bool, error) {
	flag := mergeParamFlags(flags...)
	if flag.FromBody() {
		data, err := req.BodyMap()
		if err != nil {
			return nil, err
		}
		valueIn := data[key]
		if valueIn != nil {
			switch value := valueIn.(type) {
			case bool:
				valueBool := value // to copy
				return &valueBool, nil
			default:
				return nil, NewError(
					InvalidArgument,
					fmt.Sprintf("invalid '%v', must be true or false", key),
					nil,
					"value", value,
				)
			}
		}
	}
	if flag.FromForm() {
		valueStr := req.r.FormValue(key)
		if valueStr != "" {
			valueStr = strings.ToLower(valueStr)
			switch valueStr {
			case "true":
				valueBool := true
				return &valueBool, nil
			case "false":
				valueBool := false
				return &valueBool, nil
			}
			return nil, NewError(
				InvalidArgument,
				fmt.Sprintf("invalid '%v', must be true or false", key),
				nil,
				"valueStr", valueStr,
			)
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
	if flag.FromBody() {
		data, err := req.BodyMap()
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
						"value", value,
					)
				}
				return &valueTm, nil
			default:
				return nil, NewError(
					InvalidArgument,
					fmt.Sprintf("invalid '%v', must be RFC3339 time string", key),
					nil,
					"value", value,
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
					"valueStr", valueStr,
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
