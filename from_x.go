package ripo

import (
	"time"
)

type FromX interface {
	GetString(req Request, key string) (*string, error)
	GetStringList(req Request, key string) ([]string, error)
	GetInt(req Request, key string) (*int, error)
	GetFloat(req Request, key string) (*float64, error)
	GetBool(req Request, key string) (*bool, error)
	GetTime(req Request, key string) (*time.Time, error)
}
