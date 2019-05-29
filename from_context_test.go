package ripo

import (
	"context"
	"reflect"
	"testing"
	"time"

	gomock "github.com/golang/mock/gomock"
	"github.com/ilius/is"
)

func TestFromContext_GetString(t *testing.T) {
	is := is.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReq := NewMockExtendedRequest(ctrl)
	var req ExtendedRequest = mockReq
	{
		mockReq.EXPECT().Context().Return(context.Background())
		value, err := FromContext.GetString(req, "name")
		is.NotErr(err)
		is.Nil(value)
	}
	{
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "name", 123))
		value, err := FromContext.GetString(req, "name")
		AssertError(t, err, InvalidArgument, "invalid 'name', must be string")
		is.Nil(value)
	}
	{
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "name", "John"))
		value, err := FromContext.GetString(req, "name")
		is.NotErr(err)
		is.Equal("John", *value)
	}
	{
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "name", []byte("John")))
		value, err := FromContext.GetString(req, "name")
		is.Equal("John", *value)
		is.NotErr(err)
	}
}

func TestFromContext_GetStringList(t *testing.T) {
	is := is.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReq := NewMockExtendedRequest(ctrl)
	var req ExtendedRequest = mockReq
	{
		mockReq.EXPECT().Context().Return(context.Background())
		value, err := FromContext.GetStringList(req, "names")
		is.NotErr(err)
		is.Nil(value)
	}
	{
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "names", 123))
		value, err := FromContext.GetStringList(req, "names")
		AssertError(t, err, InvalidArgument, "invalid 'names', must be array of strings")
		is.Nil(value)
	}
	{
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "names", "John"))
		value, err := FromContext.GetStringList(req, "names")
		AssertError(t, err, InvalidArgument, "invalid 'names', must be array of strings")
		is.Nil(value)
	}
	{
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "names", []string{"John", "Smith"}))
		value, err := FromContext.GetStringList(req, "names")
		is.NotErr(err)
		is.Equal([]string{"John", "Smith"}, value)
	}
}

func TestFromContext_GetInt(t *testing.T) {
	is := is.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReq := NewMockExtendedRequest(ctrl)
	var req ExtendedRequest = mockReq
	{
		mockReq.EXPECT().Context().Return(context.Background())
		value, err := FromContext.GetInt(req, "count")
		is.NotErr(err)
		is.Nil(value)
	}
	{
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "count", "John"))
		value, err := FromContext.GetInt(req, "count")
		AssertError(t, err, InvalidArgument, "invalid 'count', must be integer")
		is.Nil(value)
	}
	{
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "count", 123))
		value, err := FromContext.GetInt(req, "count")
		is.NotErr(err)
		is.Equal(123, *value)
	}
	{
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "count", 12.234))
		value, err := FromContext.GetInt(req, "count")
		is.NotErr(err)
		is.Equal(12, *value)
	}
	{
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "count", int32(1234)))
		value, err := FromContext.GetInt(req, "count")
		is.NotErr(err)
		is.Equal(1234, *value)
	}
	{
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "count", int64(12345)))
		value, err := FromContext.GetInt(req, "count")
		is.NotErr(err)
		is.Equal(12345, *value)
	}
}

func TestFromContext_GetFloat(t *testing.T) {
	is := is.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReq := NewMockExtendedRequest(ctrl)
	var req ExtendedRequest = mockReq
	{
		mockReq.EXPECT().Context().Return(context.Background())
		value, err := FromContext.GetFloat(req, "height")
		is.NotErr(err)
		is.Nil(value)
	}
	{
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "height", "John"))
		value, err := FromContext.GetFloat(req, "height")
		AssertError(t, err, InvalidArgument, "invalid 'height', must be float")
		is.Nil(value)
	}
	{
		v := float32(12.234)
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "height", v))
		value, err := FromContext.GetFloat(req, "height")
		is.NotErr(err)
		is.Equal(float64(v), *value)
	}
	{
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "height", 123.45))
		value, err := FromContext.GetFloat(req, "height")
		is.NotErr(err)
		is.Equal(123.45, *value)
	}
	{
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "height", 123))
		value, err := FromContext.GetFloat(req, "height")
		is.NotErr(err)
		is.Equal(123.0, *value)
	}
	{
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "height", int32(1234)))
		value, err := FromContext.GetFloat(req, "height")
		is.NotErr(err)
		is.Equal(1234.0, *value)
	}
	{
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "height", int64(12345)))
		value, err := FromContext.GetFloat(req, "height")
		is.NotErr(err)
		is.Equal(12345.0, *value)
	}
}

func TestFromContext_GetBool(t *testing.T) {
	is := is.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReq := NewMockExtendedRequest(ctrl)
	var req ExtendedRequest = mockReq
	{
		mockReq.EXPECT().Context().Return(context.Background())
		value, err := FromContext.GetBool(req, "agree")
		is.NotErr(err)
		is.Nil(value)
	}
	{
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "agree", "yes"))
		value, err := FromContext.GetBool(req, "agree")
		AssertError(t, err, InvalidArgument, "invalid 'agree', must be bool")
		is.Nil(value)
	}
	{
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "agree", true))
		value, err := FromContext.GetBool(req, "agree")
		is.NotErr(err)
		is.Equal(true, *value)
	}
	{
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "agree", false))
		value, err := FromContext.GetBool(req, "agree")
		is.NotErr(err)
		is.Equal(false, *value)
	}
}

func TestFromContext_GetTime(t *testing.T) {
	is := is.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReq := NewMockExtendedRequest(ctrl)
	var req ExtendedRequest = mockReq
	{
		mockReq.EXPECT().Context().Return(context.Background())
		value, err := FromContext.GetTime(req, "since")
		is.Nil(value)
		is.NotErr(err)
	}
	{
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "since", "abcd"))
		value, err := FromContext.GetTime(req, "since")
		is.Nil(value)
		AssertError(t, err, InvalidArgument, "invalid 'since', must be RFC3339 time string")
	}
	{
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "since", 3465))
		value, err := FromContext.GetTime(req, "since")
		is.Nil(value)
		AssertError(t, err, InvalidArgument, "invalid 'since', must be RFC3339 time string")
	}
	{
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "since", "2017-12-20T17:30:00Z"))
		value, err := FromContext.GetTime(req, "since")
		is.NotErr(err)
		is.Equal(time.Date(2017, time.Month(12), 20, 17, 30, 0, 0, time.UTC), *value)
	}
	{
		tm := time.Date(2017, time.Month(12), 20, 17, 30, 0, 0, time.UTC)
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "since", tm))
		value, err := FromContext.GetTime(req, "since")
		is.NotErr(err)
		is.Equal(tm, *value)
	}
}

func TestFromContext_GetObject(t *testing.T) {
	is := is.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReq := NewMockExtendedRequest(ctrl)
	var req ExtendedRequest = mockReq
	type Person struct {
		Name string `json:"name"`
	}
	{
		value, err := FromContext.GetObject(req, "since", reflect.TypeOf(Person{}))
		is.Nil(value)
		is.NotErr(err)
	}
}
