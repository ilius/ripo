package restpc

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
)

func TestFromBody_GetString(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReq := NewMockRequest(ctrl)
	var req Request = mockReq
	{
		mockReq.EXPECT().BodyMap().Return(nil, fmt.Errorf("unknown error"))
		value, err := FromBody.GetString(req, "name")
		assert.Nil(t, value)
		assert.EqualError(t, err, "unknown error")
	}
	{
		mockReq.EXPECT().BodyMap().Return(nil, nil)
		value, err := FromBody.GetString(req, "name")
		assert.Nil(t, value)
		assert.Nil(t, err)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"name": 123,
		}, nil)
		value, err := FromBody.GetString(req, "name")
		assert.Nil(t, value)
		assert.EqualError(t, err, "invalid 'name', must be string")
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"name": "John Smith",
		}, nil)
		value, err := FromBody.GetString(req, "name")
		assert.Equal(t, "John Smith", *value)
		assert.Nil(t, err)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"name": []byte("John Smith"),
		}, nil)
		value, err := FromBody.GetString(req, "name")
		assert.Equal(t, "John Smith", *value)
		assert.Nil(t, err)
	}
}
