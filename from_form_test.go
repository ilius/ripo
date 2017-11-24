package ripo

import (
	"testing"
	"time"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestFromForm_GetString(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReq := NewMockRequest(ctrl)
	var req Request = mockReq
	{
		mockReq.EXPECT().GetFormValue("name").Return("")
		value, err := FromForm.GetString(req, "name")
		assert.NoError(t, err)
		assert.Nil(t, value)
	}
	{
		mockReq.EXPECT().GetFormValue("name").Return("John Smith")
		value, err := FromForm.GetString(req, "name")
		assert.NoError(t, err)
		assert.Equal(t, "John Smith", *value)
	}
}
func TestFromForm_GetStringList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReq := NewMockRequest(ctrl)
	var req Request = mockReq
	{
		mockReq.EXPECT().GetFormValue("names").Return("")
		value, err := FromForm.GetStringList(req, "names")
		assert.NoError(t, err)
		assert.Nil(t, value)
	}
	{
		mockReq.EXPECT().GetFormValue("names").Return("John Smith,John Doe")
		value, err := FromForm.GetStringList(req, "names")
		assert.NoError(t, err)
		assert.Equal(t, []string{"John Smith", "John Doe"}, value)
	}
}

func TestFromForm_GetInt(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReq := NewMockRequest(ctrl)
	var req Request = mockReq
	{
		mockReq.EXPECT().GetFormValue("count").Return("")
		value, err := FromForm.GetInt(req, "count")
		assert.NoError(t, err)
		assert.Nil(t, value)
	}
	{
		mockReq.EXPECT().GetFormValue("count").Return("abc")
		value, err := FromForm.GetInt(req, "count")
		assert.EqualError(t, err, "invalid 'count', must be integer")
		assert.Nil(t, value)
	}
	{
		mockReq.EXPECT().GetFormValue("count").Return("1.23")
		value, err := FromForm.GetInt(req, "count")
		assert.EqualError(t, err, "invalid 'count', must be integer")
		assert.Nil(t, value)
	}
	{
		mockReq.EXPECT().GetFormValue("count").Return("5001")
		value, err := FromForm.GetInt(req, "count")
		assert.NoError(t, err)
		assert.Equal(t, 5001, *value)
	}
}

func TestFromForm_GetFloat(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReq := NewMockRequest(ctrl)
	var req Request = mockReq
	{
		mockReq.EXPECT().GetFormValue("weight").Return("")
		value, err := FromForm.GetFloat(req, "weight")
		assert.Nil(t, value)
		assert.NoError(t, err)
	}
	{
		mockReq.EXPECT().GetFormValue("weight").Return("abc")
		value, err := FromForm.GetFloat(req, "weight")
		assert.Nil(t, value)
		assert.EqualError(t, err, "invalid 'weight', must be float")
	}
	{
		mockReq.EXPECT().GetFormValue("weight").Return("1231")
		value, err := FromForm.GetFloat(req, "weight")
		assert.Equal(t, 1231.0, *value)
		assert.NoError(t, err)
	}
	{
		mockReq.EXPECT().GetFormValue("weight").Return("104.15")
		value, err := FromForm.GetFloat(req, "weight")
		assert.Equal(t, 104.15, *value)
		assert.NoError(t, err)
	}
}

func TestFromForm_GetBool(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReq := NewMockRequest(ctrl)
	var req Request = mockReq
	{
		mockReq.EXPECT().GetFormValue("agree").Return("")
		value, err := FromForm.GetBool(req, "agree")
		assert.Nil(t, value)
		assert.NoError(t, err)
	}
	{
		mockReq.EXPECT().GetFormValue("agree").Return("abcd")
		value, err := FromForm.GetBool(req, "agree")
		assert.Nil(t, value)
		assert.EqualError(t, err, "invalid 'agree', must be true or false")
	}
	{
		mockReq.EXPECT().GetFormValue("agree").Return("3465")
		value, err := FromForm.GetBool(req, "agree")
		assert.Nil(t, value)
		assert.EqualError(t, err, "invalid 'agree', must be true or false")
	}
	{
		mockReq.EXPECT().GetFormValue("agree").Return("true")
		value, err := FromForm.GetBool(req, "agree")
		assert.Equal(t, true, *value)
		assert.NoError(t, err)
	}
	{
		mockReq.EXPECT().GetFormValue("agree").Return("false")
		value, err := FromForm.GetBool(req, "agree")
		assert.Equal(t, false, *value)
		assert.NoError(t, err)
	}
}

func TestFromForm_GetTime(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReq := NewMockRequest(ctrl)
	var req Request = mockReq
	{
		mockReq.EXPECT().GetFormValue("since").Return("")
		value, err := FromForm.GetTime(req, "since")
		assert.Nil(t, value)
		assert.NoError(t, err)
	}
	{
		mockReq.EXPECT().GetFormValue("since").Return("abcd")
		value, err := FromForm.GetTime(req, "since")
		assert.Nil(t, value)
		assert.EqualError(t, err, "invalid 'since', must be RFC3339 time string")
	}
	{
		mockReq.EXPECT().GetFormValue("since").Return("3465")
		value, err := FromForm.GetTime(req, "since")
		assert.Nil(t, value)
		assert.EqualError(t, err, "invalid 'since', must be RFC3339 time string")
	}
	{
		mockReq.EXPECT().GetFormValue("since").Return("2017-12-20T17:30:00Z")
		value, err := FromForm.GetTime(req, "since")
		assert.NoError(t, err)
		assert.Equal(t, time.Date(2017, time.Month(12), 20, 17, 30, 0, 0, time.UTC), *value)
	}
}
