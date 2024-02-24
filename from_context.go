package ripo

import (
	"fmt"
	"reflect"
	"time"
)

var FromContext FromX = &fromContext{}

type fromContext struct{}

func (f *fromContext) GetString(req ExtendedRequest, key string) (*string, error) {
	ctx := req.Context()
	valueIn := ctx.Value(key)
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
				fmt.Errorf("ctx.Value(%#v) = %#v", key, valueIn),
			).Add("ctx", ctx)
		}
	}
	return nil, nil
}

func (f *fromContext) GetStringList(req ExtendedRequest, key string) ([]string, error) {
	ctx := req.Context()
	valueIn := ctx.Value(key)
	if valueIn != nil {
		switch value := valueIn.(type) {
		case []string:
			valueSlice := append([]string(nil), value...) // to copy
			return valueSlice, nil
		default:
			return nil, NewError(
				InvalidArgument,
				fmt.Sprintf("invalid '%v', must be array of strings", key),
				fmt.Errorf("ctx.Value(%#v) = %#v", key, valueIn),
			).Add("ctx", ctx)
		}
	}
	return nil, nil
}

func (f *fromContext) GetInt(req ExtendedRequest, key string) (*int, error) {
	ctx := req.Context()
	valueIn := ctx.Value(key)
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
				fmt.Errorf("ctx.Value(%#v) = %#v", key, valueIn),
			).Add("ctx", ctx)
		}
	}
	return nil, nil
}

func (f *fromContext) GetFloat(req ExtendedRequest, key string) (*float64, error) {
	ctx := req.Context()
	valueIn := ctx.Value(key)
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
				fmt.Errorf("ctx.Value(%#v) = %#v", key, valueIn),
			).Add("ctx", ctx)
		}
	}
	return nil, nil
}

func (f *fromContext) GetBool(req ExtendedRequest, key string) (*bool, error) {
	ctx := req.Context()
	valueIn := ctx.Value(key)
	if valueIn != nil {
		switch value := valueIn.(type) {
		case bool:
			valueBool := value // to copy
			return &valueBool, nil
		default:
			return nil, NewError(
				InvalidArgument,
				fmt.Sprintf("invalid '%v', must be bool", key),
				fmt.Errorf("ctx.Value(%#v) = %#v", key, valueIn),
			).Add("ctx", ctx)
		}
	}
	return nil, nil
}

func (f *fromContext) GetTime(req ExtendedRequest, key string) (*time.Time, error) {
	ctx := req.Context()
	valueIn := ctx.Value(key)
	if valueIn != nil {
		switch value := valueIn.(type) {
		case time.Time:
			valueTime := value // to copy
			return &valueTime, nil
		case string:
			valueTm, err := time.Parse(time.RFC3339, value)
			if err != nil {
				return nil, NewError(
					InvalidArgument,
					fmt.Sprintf("invalid '%v', must be RFC3339 time string", key),
					err,
				).Add("value", value)
			}
			return &valueTm, nil
		default:
			return nil, NewError(
				InvalidArgument,
				fmt.Sprintf("invalid '%v', must be RFC3339 time string", key),
				fmt.Errorf("ctx.Value(%#v) = %#v", key, valueIn),
			).Add("ctx", ctx)
		}
	}
	return nil, nil
}

func (f *fromContext) GetObject(req ExtendedRequest, key string, _type reflect.Type) (any, error) {
	// TODO
	return nil, nil
}
