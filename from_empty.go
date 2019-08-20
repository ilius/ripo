package ripo

import (
	"reflect"
	"time"
)

var FromEmpty FromX = &fromEmpty{}

type fromEmpty struct{}

func (f *fromEmpty) GetString(req ExtendedRequest, key string) (*string, error) {
	v := ""
	return &v, nil
}

func (f *fromEmpty) GetStringList(req ExtendedRequest, key string) ([]string, error) {
	return []string{}, nil
}

func (f *fromEmpty) GetInt(req ExtendedRequest, key string) (*int, error) {
	v := 0
	return &v, nil
}

func (f *fromEmpty) GetFloat(req ExtendedRequest, key string) (*float64, error) {
	v := 0.0
	return &v, nil
}

func (f *fromEmpty) GetBool(req ExtendedRequest, key string) (*bool, error) {
	v := false
	return &v, nil
}

func (f *fromEmpty) GetTime(req ExtendedRequest, key string) (*time.Time, error) {
	var v time.Time
	return &v, nil
}

func (f *fromEmpty) GetObject(req ExtendedRequest, key string, _type reflect.Type) (interface{}, error) {
	givePointer := false
	if _type.Kind() == reflect.Ptr {
		_type = _type.Elem()
		givePointer = true
	}
	valueValue := reflect.New(_type)
	if !givePointer {
		valueValue = valueValue.Elem()
	}
	valueIn := valueValue.Interface()
	return valueIn, nil
}
