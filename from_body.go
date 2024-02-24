package ripo

import (
	"fmt"
	"reflect"
	"time"

	"github.com/mitchellh/mapstructure"
)

var FromBody FromX = &fromBody{}

type fromBody struct{}

func (f *fromBody) GetString(req ExtendedRequest, key string) (*string, error) {
	data, err := req.BodyMap()
	if err != nil {
		return nil, err
	}
	valueIn := data[key]
	if valueIn != nil {
		switch value := valueIn.(type) {
		case string:
			if value != "" {
				valueStr := value // to copy
				return &valueStr, nil
			}
		case []byte:
			if len(value) > 0 {
				valueStr := string(value)
				return &valueStr, nil
			}
		default:
			return nil, NewError(
				InvalidArgument,
				fmt.Sprintf("invalid '%v', must be string", key),
				nil,
			).Add("value", value)
		}
	}
	return nil, nil
}

func (f *fromBody) GetStringList(req ExtendedRequest, key string) ([]string, error) {
	data, err := req.BodyMap()
	if err != nil {
		return nil, err
	}
	valueIn := data[key]
	if valueIn != nil {
		switch value := valueIn.(type) {
		case []any:
			valueSlice := make([]string, len(value))
			for index, item := range value {
				itemStr, ok := item.(string)
				if !ok {
					return nil, NewError(
						InvalidArgument,
						fmt.Sprintf("invalid '%v', must be array of strings", key),
						nil,
					).Add("value", value)
				}
				valueSlice[index] = itemStr
			}
			return valueSlice, nil
		case []string:
			valueSlice := append([]string(nil), value...) // to copy
			return valueSlice, nil
		default:
			return nil, NewError(
				InvalidArgument,
				fmt.Sprintf("invalid '%v', must be array of strings", key),
				nil,
			).Add("value", value)
		}
	}
	return nil, nil
}

func (f *fromBody) GetInt(req ExtendedRequest, key string) (*int, error) {
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
			).Add("value", value)
		}
	}
	return nil, nil
}

func (f *fromBody) GetFloat(req ExtendedRequest, key string) (*float64, error) {
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
			).Add("value", value)
		}
	}
	return nil, nil
}

func (f *fromBody) GetBool(req ExtendedRequest, key string) (*bool, error) {
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
			).Add("value", value)
		}
	}
	return nil, nil
}

func (f *fromBody) GetTime(req ExtendedRequest, key string) (*time.Time, error) {
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
				).Add("value", value)
			}
			return &valueTm, nil
		default:
			return nil, NewError(
				InvalidArgument,
				fmt.Sprintf("invalid '%v', must be RFC3339 time string", key),
				nil,
			).Add("value", value)
		}
	}
	return nil, nil
}

func (f *fromBody) GetObject(req ExtendedRequest, key string, _type reflect.Type) (any, error) {
	data, err := req.BodyMap()
	if err != nil {
		return nil, err
	}
	mValueIn := data[key]
	if mValueIn != nil {
		mValueType := reflect.TypeOf(mValueIn)
		if mValueType == _type {
			// FIXME: not covered
			return &mValueIn, nil
		}
		switch mValueType.Kind() {
		case reflect.Map, reflect.Slice:
			valuePtrValue := reflect.New(_type)
			valuePtrIn := valuePtrValue.Interface()
			err := mapstructure.Decode(mValueIn, valuePtrIn)
			if err != nil {
				// FIXME: not covered
				return nil, NewError(
					InvalidArgument,
					fmt.Sprintf("invalid '%v', must be a compatible object", key),
					err,
				).Add("_type", _type)
			}
			valueIn := valuePtrValue.Elem().Interface()
			return valueIn, nil
		}
		return nil, NewError(
			InvalidArgument,
			fmt.Sprintf("invalid '%v', must be a compatible object", key),
			nil,
		)
	}
	return nil, nil
}
