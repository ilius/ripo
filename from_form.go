package ripo

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var FromForm FromX = &fromForm{}

type fromForm struct{}

func (f *fromForm) GetString(req ExtendedRequest, key string) (*string, error) {
	value := req.GetFormValue(key)
	if value != "" {
		return &value, nil
	}
	return nil, nil
}

func (f *fromForm) GetStringList(req ExtendedRequest, key string) ([]string, error) {
	valueStr := req.GetFormValue(key)
	if valueStr != "" {
		valueSlice := strings.Split(valueStr, ",")
		return valueSlice, nil
	}
	return nil, nil
}

func (f *fromForm) GetInt(req ExtendedRequest, key string) (*int, error) {
	valueStr := req.GetFormValue(key)
	if valueStr != "" {
		value, err := strconv.ParseInt(valueStr, 10, 64)
		if err != nil {
			return nil, NewError(
				InvalidArgument,
				fmt.Sprintf("invalid '%v', must be integer", key),
				err,
			).Add("valueStr", valueStr)
		}
		valueInt := int(value)
		return &valueInt, nil
	}
	return nil, nil
}

func (f *fromForm) GetFloat(req ExtendedRequest, key string) (*float64, error) {
	valueStr := req.GetFormValue(key)
	if valueStr != "" {
		value, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			return nil, NewError(
				InvalidArgument,
				fmt.Sprintf("invalid '%v', must be float", key),
				err,
			).Add("valueStr", valueStr)
		}
		valueF := float64(value)
		return &valueF, nil
	}
	return nil, nil
}

func (f *fromForm) GetBool(req ExtendedRequest, key string) (*bool, error) {
	valueStr := req.GetFormValue(key)
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
		).Add("valueStr", valueStr)
	}
	return nil, nil
}

func (f *fromForm) GetTime(req ExtendedRequest, key string) (*time.Time, error) {
	valueStr := req.GetFormValue(key)
	if valueStr != "" {
		valueTm, err := time.Parse(time.RFC3339, valueStr)
		if err != nil {
			return nil, NewError(
				InvalidArgument,
				fmt.Sprintf("invalid '%v', must be RFC3339 time string", key),
				err,
			).Add("valueStr", valueStr)
		}
		return &valueTm, nil
	}
	return nil, nil
}

func (f *fromForm) GetObject(req ExtendedRequest, key string, _type reflect.Type) (interface{}, error) {
	return nil, nil
}
