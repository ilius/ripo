package restpc

import (
	"time"
)

var FromEmpty = &fromEmpty{}

type fromEmpty struct{}

func (f *fromEmpty) GetString(req Request, key string) (*string, error) {
	v := ""
	return &v, nil
}
func (f *fromEmpty) GetStringList(req Request, key string) ([]string, error) {
	return []string{}, nil
}
func (f *fromEmpty) GetInt(req Request, key string) (*int, error) {
	v := 0
	return &v, nil
}
func (f *fromEmpty) GetFloat(req Request, key string) (*float64, error) {
	v := 0.0
	return &v, nil
}
func (f *fromEmpty) GetBool(req Request, key string) (*bool, error) {
	v := false
	return &v, nil
}
func (f *fromEmpty) GetTime(req Request, key string) (*time.Time, error) {
	var v time.Time
	return &v, nil
}
