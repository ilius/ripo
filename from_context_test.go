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
