package restpc

import (
	"io"
	"net/url"
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_RequestMock(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockReq := NewMockRequest(ctrl)
	{
		mockReq.EXPECT().RemoteIP().Return("192.168.100.100", nil)
		mockReq.RemoteIP()
	}
	{
		testUrl, err := url.Parse("http://127.0.0.1/test")
		if err != nil {
			panic(err)
		}
		mockReq.EXPECT().URL().Return(testUrl)
		mockReq.URL()
	}
	{
		mockReq.EXPECT().Host().Return("localhost")
		mockReq.Host()
	}
	{
		mockReq.EXPECT().HandlerName().Return("TestHandler")
		mockReq.HandlerName()
	}
	{
		mockReq.EXPECT().Body().Return(nil, io.EOF)
		mockReq.Body()
	}
	{
		mockReq.EXPECT().BodyTo(gomock.Any()).Return(nil)
		m := map[string]interface{}{}
		mockReq.BodyTo(&m)
	}
	{
		mockReq.EXPECT().GetHeader("test").Return("bar")
		mockReq.GetHeader("test")
	}
	{
		mockReq.EXPECT().GetString("test").Return(nil, nil)
		mockReq.GetString("test")
	}
	{
		mockReq.EXPECT().GetStringList("test").Return(nil, nil)
		mockReq.GetStringList("test")
	}
	{
		mockReq.EXPECT().GetInt("test").Return(nil, nil)
		mockReq.GetInt("test")
	}
	{
		mockReq.EXPECT().GetIntDefault("test", 0).Return(0, nil)
		mockReq.GetIntDefault("test", 0)
	}
	{
		mockReq.EXPECT().GetFloat("test").Return(nil, nil)
		mockReq.GetFloat("test")
	}
	{
		mockReq.EXPECT().GetBool("test").Return(nil, nil)
		mockReq.GetBool("test")
	}
	{
		mockReq.EXPECT().GetTime("test").Return(nil, nil)
		mockReq.GetTime("test")
	}
	{
		mockReq.EXPECT().FullMap().Return(nil)
		mockReq.FullMap()
	}

}
