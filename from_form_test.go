package ripo

import (
	"reflect"
	"testing"
	"time"

	gomock "github.com/golang/mock/gomock"
	"github.com/ilius/is/v2"
)

func TestFromForm_GetString(t *testing.T) {
	is := is.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReq := NewMockExtendedRequest(ctrl)
	var req ExtendedRequest = mockReq
	{
		mockReq.EXPECT().GetFormValue("name").Return("")
		value, err := FromForm.GetString(req, "name")
		is.NotErr(err)
		is.Nil(value)
	}
	{
		mockReq.EXPECT().GetFormValue("name").Return("John Smith")
		value, err := FromForm.GetString(req, "name")
		is.NotErr(err)
		is.Equal("John Smith", *value)
	}
}

func TestFromForm_GetStringList(t *testing.T) {
	is := is.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReq := NewMockExtendedRequest(ctrl)
	var req ExtendedRequest = mockReq
	{
		mockReq.EXPECT().GetFormValue("names").Return("")
		value, err := FromForm.GetStringList(req, "names")
		is.NotErr(err)
		is.Nil(value)
	}
	{
		mockReq.EXPECT().GetFormValue("names").Return("John Smith,John Doe")
		value, err := FromForm.GetStringList(req, "names")
		is.NotErr(err)
		is.Equal([]string{"John Smith", "John Doe"}, value)
	}
}

func TestFromForm_GetInt(t *testing.T) {
	is := is.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReq := NewMockExtendedRequest(ctrl)
	var req ExtendedRequest = mockReq
	{
		mockReq.EXPECT().GetFormValue("count").Return("")
		value, err := FromForm.GetInt(req, "count")
		is.NotErr(err)
		is.Nil(value)
	}
	{
		mockReq.EXPECT().GetFormValue("count").Return("abc")
		value, err := FromForm.GetInt(req, "count")
		AssertError(t, err, InvalidArgument, "invalid 'count', must be integer")
		is.Nil(value)
	}
	{
		mockReq.EXPECT().GetFormValue("count").Return("1.23")
		value, err := FromForm.GetInt(req, "count")
		AssertError(t, err, InvalidArgument, "invalid 'count', must be integer")
		is.Nil(value)
	}
	{
		mockReq.EXPECT().GetFormValue("count").Return("5001")
		value, err := FromForm.GetInt(req, "count")
		is.NotErr(err)
		is.Equal(5001, *value)
	}
}

func TestFromForm_GetFloat(t *testing.T) {
	is := is.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReq := NewMockExtendedRequest(ctrl)
	var req ExtendedRequest = mockReq
	{
		mockReq.EXPECT().GetFormValue("weight").Return("")
		value, err := FromForm.GetFloat(req, "weight")
		is.Nil(value)
		is.NotErr(err)
	}
	{
		mockReq.EXPECT().GetFormValue("weight").Return("abc")
		value, err := FromForm.GetFloat(req, "weight")
		is.Nil(value)
		AssertError(t, err, InvalidArgument, "invalid 'weight', must be float")
	}
	{
		mockReq.EXPECT().GetFormValue("weight").Return("1231")
		value, err := FromForm.GetFloat(req, "weight")
		is.Equal(1231.0, *value)
		is.NotErr(err)
	}
	{
		mockReq.EXPECT().GetFormValue("weight").Return("104.15")
		value, err := FromForm.GetFloat(req, "weight")
		is.Equal(104.15, *value)
		is.NotErr(err)
	}
}

func TestFromForm_GetBool(t *testing.T) {
	is := is.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReq := NewMockExtendedRequest(ctrl)
	var req ExtendedRequest = mockReq
	{
		mockReq.EXPECT().GetFormValue("agree").Return("")
		value, err := FromForm.GetBool(req, "agree")
		is.Nil(value)
		is.NotErr(err)
	}
	{
		mockReq.EXPECT().GetFormValue("agree").Return("abcd")
		value, err := FromForm.GetBool(req, "agree")
		is.Nil(value)
		AssertError(t, err, InvalidArgument, "invalid 'agree', must be true or false")
	}
	{
		mockReq.EXPECT().GetFormValue("agree").Return("3465")
		value, err := FromForm.GetBool(req, "agree")
		is.Nil(value)
		AssertError(t, err, InvalidArgument, "invalid 'agree', must be true or false")
	}
	{
		mockReq.EXPECT().GetFormValue("agree").Return("true")
		value, err := FromForm.GetBool(req, "agree")
		is.Equal(true, *value)
		is.NotErr(err)
	}
	{
		mockReq.EXPECT().GetFormValue("agree").Return("false")
		value, err := FromForm.GetBool(req, "agree")
		is.Equal(false, *value)
		is.NotErr(err)
	}
}

func TestFromForm_GetTime(t *testing.T) {
	is := is.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReq := NewMockExtendedRequest(ctrl)
	var req ExtendedRequest = mockReq
	{
		mockReq.EXPECT().GetFormValue("since").Return("")
		value, err := FromForm.GetTime(req, "since")
		is.Nil(value)
		is.NotErr(err)
	}
	{
		mockReq.EXPECT().GetFormValue("since").Return("abcd")
		value, err := FromForm.GetTime(req, "since")
		is.Nil(value)
		AssertError(t, err, InvalidArgument, "invalid 'since', must be RFC3339 time string")
	}
	{
		mockReq.EXPECT().GetFormValue("since").Return("3465")
		value, err := FromForm.GetTime(req, "since")
		is.Nil(value)
		AssertError(t, err, InvalidArgument, "invalid 'since', must be RFC3339 time string")
	}
	{
		mockReq.EXPECT().GetFormValue("since").Return("2017-12-20T17:30:00Z")
		value, err := FromForm.GetTime(req, "since")
		is.NotErr(err)
		is.Equal(time.Date(2017, time.Month(12), 20, 17, 30, 0, 0, time.UTC), *value)
	}
}

func TestFromForm_GetObject(t *testing.T) {
	is := is.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReq := NewMockExtendedRequest(ctrl)
	var req ExtendedRequest = mockReq
	type Person struct {
		Name string `json:"name"`
	}
	{
		value, err := FromForm.GetObject(req, "since", reflect.TypeOf(Person{}))
		is.Nil(value)
		is.NotErr(err)
	}
}
