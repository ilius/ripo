package restpc

import (
	"context"
	"testing"

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
		assert.Nil(t, value)
		assert.NoError(t, err)
	}
	{
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "name", 123))
		value, err := FromContext.GetString(req, "name")
		assert.Nil(t, value)
		assert.EqualError(t, err, "invalid 'name', must be string")
	}
	{
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "name", "John"))
		value, err := FromContext.GetString(req, "name")
		assert.Equal(t, "John", *value)
		assert.NoError(t, err)
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
		assert.Nil(t, value)
		assert.NoError(t, err)
	}
	{
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "names", 123))
		value, err := FromContext.GetStringList(req, "names")
		assert.Nil(t, value)
		assert.EqualError(t, err, "invalid 'names', must be array of strings")
	}
	{
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "names", "John"))
		value, err := FromContext.GetStringList(req, "names")
		assert.Nil(t, value)
		assert.EqualError(t, err, "invalid 'names', must be array of strings")
	}
	{
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "names", []string{"John", "Smith"}))
		value, err := FromContext.GetStringList(req, "names")
		assert.Equal(t, []string{"John", "Smith"}, value)
		assert.NoError(t, err)
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
		assert.Nil(t, value)
		assert.NoError(t, err)
	}
	{
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "count", "John"))
		value, err := FromContext.GetInt(req, "count")
		assert.Nil(t, value)
		assert.EqualError(t, err, "invalid 'count', must be integer")
	}
	{
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "count", 123))
		value, err := FromContext.GetInt(req, "count")
		assert.Equal(t, 123, *value)
		assert.NoError(t, err)
	}
	{
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "count", 12.234))
		value, err := FromContext.GetInt(req, "count")
		assert.Equal(t, 12, *value)
		assert.NoError(t, err)
	}
	{
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "count", int32(1234)))
		value, err := FromContext.GetInt(req, "count")
		assert.Equal(t, 1234, *value)
		assert.NoError(t, err)
	}
	{
		mockReq.EXPECT().Context().Return(context.WithValue(context.Background(), "count", int64(12345)))
		value, err := FromContext.GetInt(req, "count")
		assert.Equal(t, 12345, *value)
		assert.NoError(t, err)
	}
}
