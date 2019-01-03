package ripo

import (
	"fmt"
	"log"
	"reflect"
	"testing"
	"time"

	"github.com/tylerb/is"

	"github.com/golang/mock/gomock"
)

func TestFromBody_GetString(t *testing.T) {
	is := is.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReq := NewMockExtendedRequest(ctrl)
	var req ExtendedRequest = mockReq
	{
		mockReq.EXPECT().BodyMap().Return(nil, fmt.Errorf("unknown error"))
		value, err := FromBody.GetString(req, "name")
		AssertError(t, err, Unknown, "unknown error")
		is.Nil(value)
	}
	{
		mockReq.EXPECT().BodyMap().Return(nil, nil)
		value, err := FromBody.GetString(req, "name")
		is.NotErr(err)
		is.Nil(value)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"name": 123,
		}, nil)
		value, err := FromBody.GetString(req, "name")
		AssertError(t, err, InvalidArgument, "invalid 'name', must be string")
		is.Nil(value)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"name": "John Smith",
		}, nil)
		value, err := FromBody.GetString(req, "name")
		is.NotErr(err)
		is.Equal("John Smith", *value)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"name": []byte("John Smith"),
		}, nil)
		value, err := FromBody.GetString(req, "name")
		is.NotErr(err)
		is.Equal("John Smith", *value)
	}
}
func TestFromBody_GetStringList(t *testing.T) {
	is := is.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReq := NewMockExtendedRequest(ctrl)
	var req ExtendedRequest = mockReq
	{
		mockReq.EXPECT().BodyMap().Return(nil, fmt.Errorf("unknown error"))
		value, err := FromBody.GetStringList(req, "names")
		AssertError(t, err, Unknown, "unknown error")
		is.Nil(value)
	}
	{
		mockReq.EXPECT().BodyMap().Return(nil, nil)
		value, err := FromBody.GetStringList(req, "names")
		is.NotErr(err)
		is.Nil(value)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"names": 123,
		}, nil)
		value, err := FromBody.GetStringList(req, "names")
		AssertError(t, err, InvalidArgument, "invalid 'names', must be array of strings")
		is.Nil(value)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"names": "John Smith",
		}, nil)
		value, err := FromBody.GetStringList(req, "names")
		AssertError(t, err, InvalidArgument, "invalid 'names', must be array of strings")
		is.Nil(value)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"names": []interface{}{"John Smith", 1234},
		}, nil)
		value, err := FromBody.GetStringList(req, "names")
		AssertError(t, err, InvalidArgument, "invalid 'names', must be array of strings")
		is.Nil(value)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"names": []string{"John Smith", "John Doe"},
		}, nil)
		value, err := FromBody.GetStringList(req, "names")
		is.NotErr(err)
		is.Equal([]string{"John Smith", "John Doe"}, value)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"names": []interface{}{"John Smith", "John Doe"},
		}, nil)
		value, err := FromBody.GetStringList(req, "names")
		is.NotErr(err)
		is.Equal([]string{"John Smith", "John Doe"}, value)
	}

}

func TestFromBody_GetInt(t *testing.T) {
	is := is.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReq := NewMockExtendedRequest(ctrl)
	var req ExtendedRequest = mockReq
	{
		mockReq.EXPECT().BodyMap().Return(nil, fmt.Errorf("unknown error"))
		value, err := FromBody.GetInt(req, "count")
		AssertError(t, err, Unknown, "unknown error")
		is.Nil(value)
	}
	{
		mockReq.EXPECT().BodyMap().Return(nil, nil)
		value, err := FromBody.GetInt(req, "count")
		is.NotErr(err)
		is.Nil(value)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"count": "abc",
		}, nil)
		value, err := FromBody.GetInt(req, "count")
		AssertError(t, err, InvalidArgument, "invalid 'count', must be integer")
		is.Nil(value)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"count": "345",
		}, nil)
		value, err := FromBody.GetInt(req, "count")
		AssertError(t, err, InvalidArgument, "invalid 'count', must be integer")
		is.Nil(value)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"count": 5001,
		}, nil)
		value, err := FromBody.GetInt(req, "count")
		is.NotErr(err)
		is.Equal(5001, *value)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"count": int32(5003),
		}, nil)
		value, err := FromBody.GetInt(req, "count")
		is.NotErr(err)
		is.Equal(5003, *value)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"count": int64(6123),
		}, nil)
		value, err := FromBody.GetInt(req, "count")
		is.NotErr(err)
		is.Equal(6123, *value)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"count": 14.15,
		}, nil)
		value, err := FromBody.GetInt(req, "count")
		is.NotErr(err)
		is.Equal(14, *value)
	}
}

func TestFromBody_GetFloat(t *testing.T) {
	is := is.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReq := NewMockExtendedRequest(ctrl)
	var req ExtendedRequest = mockReq
	{
		mockReq.EXPECT().BodyMap().Return(nil, fmt.Errorf("unknown error"))
		value, err := FromBody.GetFloat(req, "weight")
		is.Nil(value)
		AssertError(t, err, Unknown, "unknown error")
	}
	{
		mockReq.EXPECT().BodyMap().Return(nil, nil)
		value, err := FromBody.GetFloat(req, "weight")
		is.Nil(value)
		is.NotErr(err)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"weight": "abc",
		}, nil)
		value, err := FromBody.GetFloat(req, "weight")
		is.Nil(value)
		AssertError(t, err, InvalidArgument, "invalid 'weight', must be float")
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"weight": "345",
		}, nil)
		value, err := FromBody.GetFloat(req, "weight")
		is.Nil(value)
		AssertError(t, err, InvalidArgument, "invalid 'weight', must be float")
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"weight": 1231,
		}, nil)
		value, err := FromBody.GetFloat(req, "weight")
		is.Equal(1231.0, *value)
		is.NotErr(err)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"weight": int32(2345),
		}, nil)
		value, err := FromBody.GetFloat(req, "weight")
		is.Equal(2345.0, *value)
		is.NotErr(err)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"weight": int64(7123),
		}, nil)
		value, err := FromBody.GetFloat(req, "weight")
		is.Equal(7123.0, *value)
		is.NotErr(err)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"weight": 104.15,
		}, nil)
		value, err := FromBody.GetFloat(req, "weight")
		is.Equal(104.15, *value)
		is.NotErr(err)
	}
	{
		weight := float32(104.15)
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"weight": weight,
		}, nil)
		value, err := FromBody.GetFloat(req, "weight")
		is.Equal(float64(weight), *value)
		is.NotErr(err)
	}
}

func TestFromBody_GetBool(t *testing.T) {
	is := is.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReq := NewMockExtendedRequest(ctrl)
	var req ExtendedRequest = mockReq
	{
		mockReq.EXPECT().BodyMap().Return(nil, fmt.Errorf("unknown error"))
		value, err := FromBody.GetBool(req, "agree")
		is.Nil(value)
		AssertError(t, err, Unknown, "unknown error")
	}
	{
		mockReq.EXPECT().BodyMap().Return(nil, nil)
		value, err := FromBody.GetBool(req, "agree")
		is.Nil(value)
		is.NotErr(err)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"agree": "abcd",
		}, nil)
		value, err := FromBody.GetBool(req, "agree")
		is.Nil(value)
		AssertError(t, err, InvalidArgument, "invalid 'agree', must be true or false")
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"agree": "3465",
		}, nil)
		value, err := FromBody.GetBool(req, "agree")
		is.Nil(value)
		AssertError(t, err, InvalidArgument, "invalid 'agree', must be true or false")
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"agree": 1231,
		}, nil)
		value, err := FromBody.GetBool(req, "agree")
		is.Nil(value)
		AssertError(t, err, InvalidArgument, "invalid 'agree', must be true or false")
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"agree": true,
		}, nil)
		value, err := FromBody.GetBool(req, "agree")
		is.Equal(true, *value)
		is.NotErr(err)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"agree": false,
		}, nil)
		value, err := FromBody.GetBool(req, "agree")
		is.Equal(false, *value)
		is.NotErr(err)
	}
}

func TestFromBody_GetTime(t *testing.T) {
	is := is.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReq := NewMockExtendedRequest(ctrl)
	var req ExtendedRequest = mockReq
	{
		mockReq.EXPECT().BodyMap().Return(nil, fmt.Errorf("unknown error"))
		value, err := FromBody.GetTime(req, "since")
		is.Nil(value)
		AssertError(t, err, Unknown, "unknown error")
	}
	{
		mockReq.EXPECT().BodyMap().Return(nil, nil)
		value, err := FromBody.GetTime(req, "since")
		is.Nil(value)
		is.NotErr(err)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"since": "abcd",
		}, nil)
		value, err := FromBody.GetTime(req, "since")
		is.Nil(value)
		AssertError(t, err, InvalidArgument, "invalid 'since', must be RFC3339 time string")
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"since": 3465,
		}, nil)
		value, err := FromBody.GetTime(req, "since")
		is.Nil(value)
		AssertError(t, err, InvalidArgument, "invalid 'since', must be RFC3339 time string")
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"since": "2017-12-20T17:30:00Z",
		}, nil)
		value, err := FromBody.GetTime(req, "since")
		is.NotErr(err)
		is.Equal(time.Date(2017, time.Month(12), 20, 17, 30, 0, 0, time.UTC), *value)
	}
}

func TestFromBody_GetObject(t *testing.T) {
	is := is.New(t)
	type Person struct {
		Name      string  `json:"name"`
		BirthDate []int   `json:"birthDate"` // mapstructure does not support [3]int
		Age       float64 `json:"age"`
	}
	PersonType := reflect.TypeOf(Person{})
	PersonTypePtr := reflect.TypeOf(&Person{})
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReq := NewMockExtendedRequest(ctrl)
	var req ExtendedRequest = mockReq
	{
		mockReq.EXPECT().BodyMap().Return(nil, fmt.Errorf("unknown error"))
		value, err := FromBody.GetObject(req, "info", PersonTypePtr)
		AssertError(t, err, Unknown, "unknown error")
		is.Nil(value)
	}
	{
		mockReq.EXPECT().BodyMap().Return(nil, nil)
		value, err := FromBody.GetObject(req, "info", PersonTypePtr)
		is.NotErr(err)
		is.Nil(value)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"info": 123,
		}, nil)
		value, err := FromBody.GetObject(req, "info", PersonTypePtr)
		AssertError(t, err, InvalidArgument, "invalid 'info', must be a compatible object")
		is.Nil(value)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"info": map[string]interface{}{},
		}, nil)
		value, err := FromBody.GetObject(req, "info", PersonTypePtr)
		is.NotErr(err)
		if err != nil {
			log.Println("Private:", err.(RPCError).Private())
		}
		is.Equal(&Person{}, value)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"info": map[string]interface{}{},
		}, nil)
		value, err := FromBody.GetObject(req, "info", PersonType)
		is.NotErr(err)
		if err != nil {
			log.Println("Private:", err.(RPCError).Private())
		}
		is.Equal(Person{}, value)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"info": map[string]interface{}{
				"name": "John Smith",
			},
		}, nil)
		value, err := FromBody.GetObject(req, "info", PersonTypePtr)
		is.NotErr(err)
		if err != nil {
			log.Println("Private:", err.(RPCError).Private())
		}
		is.Equal(&Person{
			Name: "John Smith",
		}, value)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"info": map[string]interface{}{
				"name":      "John Smith",
				"birthDate": []int{1987, 12, 30},
			},
		}, nil)
		value, err := FromBody.GetObject(req, "info", PersonTypePtr)
		is.NotErr(err)
		if err != nil {
			log.Println("Private:", err.(RPCError).Private())
		}
		is.Equal(&Person{
			Name:      "John Smith",
			BirthDate: []int{1987, 12, 30},
		}, value)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"info": map[string]interface{}{
				"name":      "John Smith",
				"birthDate": []int{1987, 12, 30},
				"age":       30.8,
			},
		}, nil)
		value, err := FromBody.GetObject(req, "info", PersonTypePtr)
		is.NotErr(err)
		if err != nil {
			log.Println("Private:", err.(RPCError).Private())
		}
		is.Equal(&Person{
			Name:      "John Smith",
			BirthDate: []int{1987, 12, 30},
			Age:       30.8,
		}, value)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"info": map[string]interface{}{
				"name":      "John Smith",
				"birthDate": []int{1987, 12, 30},
				"age":       30.8,
			},
		}, nil)
		value, err := FromBody.GetObject(req, "info", PersonType)
		is.NotErr(err)
		if err != nil {
			log.Println("Private:", err.(RPCError).Private())
		}
		is.Equal(Person{
			Name:      "John Smith",
			BirthDate: []int{1987, 12, 30},
			Age:       30.8,
		}, value)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"guestList": []interface{}{
				map[string]interface{}{
					"name":      "John Smith",
					"birthDate": []int{1987, 12, 30},
					"age":       30.8,
				},
				map[string]interface{}{
					"name":      "Jane Smith",
					"birthDate": []int{1991, 6, 30},
					"age":       27.3,
				},
			},
		}, nil)
		value, err := FromBody.GetObject(req, "guestList", reflect.SliceOf(PersonType))
		is.NotErr(err)
		if err != nil {
			log.Println("Private:", err.(RPCError).Private())
		}
		is.Equal([]Person{
			{
				Name:      "John Smith",
				BirthDate: []int{1987, 12, 30},
				Age:       30.8,
			},
			{
				Name:      "Jane Smith",
				BirthDate: []int{1991, 6, 30},
				Age:       27.3,
			},
		}, value)
	}
}
