package ripo

import (
	"context"
	"testing"
	"time"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestFromContext_GetString(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReq := NewMockRequest(ctrl)
	var req Request = mockReq
	{
		mockReq.EXPECT().Context().Return(context.Background())
		value, err := FromContext.GetString(req, "name")
		assert.NoError(t, err)
		assert.Nil(t, value)
	}
	{
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "name", 123))
		value, err := FromContext.GetString(req, "name")
		assert.EqualError(t, err, "invalid 'name', must be string")
		assert.Nil(t, value)
	}
	{
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "name", "John"))
		value, err := FromContext.GetString(req, "name")
		assert.NoError(t, err)
		assert.Equal(t, "John", *value)
	}
	{
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "name", []byte("John")))
		value, err := FromContext.GetString(req, "name")
		assert.Equal(t, "John", *value)
		assert.NoError(t, err)
	}
}

func TestFromContext_GetStringList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReq := NewMockRequest(ctrl)
	var req Request = mockReq
	{
		mockReq.EXPECT().Context().Return(context.Background())
		value, err := FromContext.GetStringList(req, "names")
		assert.NoError(t, err)
		assert.Nil(t, value)
	}
	{
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "names", 123))
		value, err := FromContext.GetStringList(req, "names")
		assert.EqualError(t, err, "invalid 'names', must be array of strings")
		assert.Nil(t, value)
	}
	{
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "names", "John"))
		value, err := FromContext.GetStringList(req, "names")
		assert.EqualError(t, err, "invalid 'names', must be array of strings")
		assert.Nil(t, value)
	}
	{
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "names", []string{"John", "Smith"}))
		value, err := FromContext.GetStringList(req, "names")
		assert.NoError(t, err)
		assert.Equal(t, []string{"John", "Smith"}, value)
	}
}

func TestFromContext_GetInt(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReq := NewMockRequest(ctrl)
	var req Request = mockReq
	{
		mockReq.EXPECT().Context().Return(context.Background())
		value, err := FromContext.GetInt(req, "count")
		assert.NoError(t, err)
		assert.Nil(t, value)
	}
	{
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "count", "John"))
		value, err := FromContext.GetInt(req, "count")
		assert.EqualError(t, err, "invalid 'count', must be integer")
		assert.Nil(t, value)
	}
	{
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "count", 123))
		value, err := FromContext.GetInt(req, "count")
		assert.NoError(t, err)
		assert.Equal(t, 123, *value)
	}
	{
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "count", 12.234))
		value, err := FromContext.GetInt(req, "count")
		assert.NoError(t, err)
		assert.Equal(t, 12, *value)
	}
	{
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "count", int32(1234)))
		value, err := FromContext.GetInt(req, "count")
		assert.NoError(t, err)
		assert.Equal(t, 1234, *value)
	}
	{
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "count", int64(12345)))
		value, err := FromContext.GetInt(req, "count")
		assert.NoError(t, err)
		assert.Equal(t, 12345, *value)
	}
}

func TestFromContext_GetFloat(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReq := NewMockRequest(ctrl)
	var req Request = mockReq
	{
		mockReq.EXPECT().Context().Return(context.Background())
		value, err := FromContext.GetFloat(req, "height")
		assert.NoError(t, err)
		assert.Nil(t, value)
	}
	{
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "height", "John"))
		value, err := FromContext.GetFloat(req, "height")
		assert.EqualError(t, err, "invalid 'height', must be float")
		assert.Nil(t, value)
	}
	{
		v := float32(12.234)
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "height", v))
		value, err := FromContext.GetFloat(req, "height")
		assert.NoError(t, err)
		assert.Equal(t, float64(v), *value)
	}
	{
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "height", 123.45))
		value, err := FromContext.GetFloat(req, "height")
		assert.NoError(t, err)
		assert.Equal(t, 123.45, *value)
	}
	{
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "height", 123))
		value, err := FromContext.GetFloat(req, "height")
		assert.NoError(t, err)
		assert.Equal(t, 123.0, *value)
	}
	{
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "height", int32(1234)))
		value, err := FromContext.GetFloat(req, "height")
		assert.NoError(t, err)
		assert.Equal(t, 1234.0, *value)
	}
	{
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "height", int64(12345)))
		value, err := FromContext.GetFloat(req, "height")
		assert.NoError(t, err)
		assert.Equal(t, 12345.0, *value)
	}
}

func TestFromContext_GetBool(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReq := NewMockRequest(ctrl)
	var req Request = mockReq
	{
		mockReq.EXPECT().Context().Return(context.Background())
		value, err := FromContext.GetBool(req, "agree")
		assert.NoError(t, err)
		assert.Nil(t, value)
	}
	{
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "agree", "yes"))
		value, err := FromContext.GetBool(req, "agree")
		assert.EqualError(t, err, "invalid 'agree', must be bool")
		assert.Nil(t, value)
	}
	{
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "agree", true))
		value, err := FromContext.GetBool(req, "agree")
		assert.NoError(t, err)
		assert.Equal(t, true, *value)
	}
	{
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "agree", false))
		value, err := FromContext.GetBool(req, "agree")
		assert.NoError(t, err)
		assert.Equal(t, false, *value)
	}
}

func TestFromContext_GetTime(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReq := NewMockRequest(ctrl)
	var req Request = mockReq
	{
		mockReq.EXPECT().Context().Return(context.Background())
		value, err := FromContext.GetTime(req, "since")
		assert.Nil(t, value)
		assert.NoError(t, err)
	}
	{
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "since", "abcd"))
		value, err := FromContext.GetTime(req, "since")
		assert.Nil(t, value)
		assert.EqualError(t, err, "invalid 'since', must be RFC3339 time string")
	}
	{
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "since", 3465))
		value, err := FromContext.GetTime(req, "since")
		assert.Nil(t, value)
		assert.EqualError(t, err, "invalid 'since', must be RFC3339 time string")
	}
	{
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "since", "2017-12-20T17:30:00Z"))
		value, err := FromContext.GetTime(req, "since")
		assert.NoError(t, err)
		assert.Equal(t, time.Date(2017, time.Month(12), 20, 17, 30, 0, 0, time.UTC), *value)
	}
	{
		tm := time.Date(2017, time.Month(12), 20, 17, 30, 0, 0, time.UTC)
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "since", tm))
		value, err := FromContext.GetTime(req, "since")
		assert.NoError(t, err)
		assert.Equal(t, tm, *value)
	}
}
