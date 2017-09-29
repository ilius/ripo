package restpc

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var FromForm FromX

func init() {
	FromForm = &fromForm{}
}

type fromForm struct{}

func (f *fromForm) GetString(req Request, key string) (*string, error) {
	value := req.HTTP().FormValue(key)
	if value != "" {
		return &value, nil
	}
	return nil, nil
}
func (f *fromForm) GetStringList(req Request, key string) ([]string, error) {
	valueStr := req.HTTP().FormValue(key)
	if valueStr != "" {
		valueSlice := strings.Split(valueStr, ",")
		return valueSlice, nil
	}
	return nil, nil
}
func (f *fromForm) GetInt(req Request, key string) (*int, error) {
	valueStr := req.HTTP().FormValue(key)
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
	return nil, nil
}
func (f *fromForm) GetFloat(req Request, key string) (*float64, error) {
	valueStr := req.HTTP().FormValue(key)
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
	return nil, nil
}
func (f *fromForm) GetBool(req Request, key string) (*bool, error) {
	valueStr := req.HTTP().FormValue(key)
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
	return nil, nil
}
func (f *fromForm) GetTime(req Request, key string) (*time.Time, error) {
	valueStr := req.HTTP().FormValue(key)
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
	return nil, nil
}
