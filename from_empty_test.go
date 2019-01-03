package ripo

import (
	"reflect"
	"testing"

	"github.com/tylerb/is"
)

func TestFromEmpty(t *testing.T) {
	is := is.New(t)
	{
		value, err := FromEmpty.GetString(nil, "foo")
		is.NotErr(err)
		is.Equal("", *value)
	}
	{
		value, err := FromEmpty.GetStringList(nil, "foo")
		is.NotErr(err)
		is.Equal([]string{}, value)
	}
	{
		value, err := FromEmpty.GetInt(nil, "foo")
		is.NotErr(err)
		is.Equal(0, *value)
	}
	{
		value, err := FromEmpty.GetFloat(nil, "foo")
		is.NotErr(err)
		is.Equal(0.0, *value)
	}
	{
		value, err := FromEmpty.GetBool(nil, "foo")
		is.NotErr(err)
		is.Equal(false, *value)
	}
	{
		value, err := FromEmpty.GetTime(nil, "foo")
		is.NotErr(err)
		is.Equal(int64(-62135596800), value.Unix())
	}
	{
		type Person struct {
			Name      string  `json:"name"`
			BirthDate []int   `json:"birthDate"` // mapstructure does not support [3]int
			Age       float64 `json:"age"`
		}
		value, err := FromEmpty.GetObject(nil, "person", reflect.TypeOf(&Person{}))
		is.NotErr(err)
		is.Equal(&Person{}, value)
	}
	type Person struct {
		Name      string  `json:"name"`
		BirthDate []int   `json:"birthDate"` // mapstructure does not support [3]int
		Age       float64 `json:"age"`
	}
	{
		value, err := FromEmpty.GetObject(nil, "person", reflect.TypeOf(Person{}))
		is.NotErr(err)
		is.Equal(Person{}, value)
	}
	{
		value, err := FromEmpty.GetObject(nil, "list", reflect.SliceOf(reflect.TypeOf(Person{})))
		is.NotErr(err)
		if !reflect.DeepEqual([]Person(nil), value) {
			t.Fatalf("bad value = %#v", value)
		}
		// is.Equal([]Person(nil), value)
		// above line panics: comparing uncomparable type []ripo.Person
	}
}
