package ripo

import (
	"reflect"
	"time"
)

type FromX interface {
	GetString(req ExtendedRequest, key string) (*string, error)
	GetStringList(req ExtendedRequest, key string) ([]string, error)
	GetInt(req ExtendedRequest, key string) (*int, error)
	GetFloat(req ExtendedRequest, key string) (*float64, error)
	GetBool(req ExtendedRequest, key string) (*bool, error)
	GetTime(req ExtendedRequest, key string) (*time.Time, error)
	GetObject(req ExtendedRequest, key string, _type reflect.Type) (interface{}, error)
}
