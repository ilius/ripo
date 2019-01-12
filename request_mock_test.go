package ripo

import (
	"io"
	"net/url"
	"reflect"
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
		mockReq.EXPECT().Header("test").Return("bar")
		mockReq.Header("test")
	}
	{
		mockReq.EXPECT().HeaderKeys().Return([]string{"bar"})
		mockReq.HeaderKeys()
	}
	{
		mockReq.EXPECT().Cookie("test")
		mockReq.Cookie("test")
	}
	{
		mockReq.EXPECT().CookieNames()
		mockReq.CookieNames()
	}
	{
		mockReq.EXPECT().Context()
		mockReq.Context()
	}
	{
		mockReq.EXPECT().GetString("test", FromForm).Return(nil, nil)
		mockReq.GetString("test", FromForm)
	}
	{
		mockReq.EXPECT().GetStringDefault("test", "foo", FromForm).Return("", nil)
		mockReq.GetStringDefault("test", "foo", FromForm)
	}
	{
		mockReq.EXPECT().GetStringList("test", FromBody).Return(nil, nil)
		mockReq.GetStringList("test", FromBody)
	}
	{
		mockReq.EXPECT().GetInt("test", FromBody).Return(nil, nil)
		mockReq.GetInt("test", FromBody)
	}
	{
		mockReq.EXPECT().GetIntDefault("test", 0, FromContext).Return(0, nil)
		mockReq.GetIntDefault("test", 0, FromContext)
	}
	{
		mockReq.EXPECT().GetFloat("test", FromBody).Return(nil, nil)
		mockReq.GetFloat("test", FromBody)
	}
	{
		mockReq.EXPECT().GetFloatDefault("test", 0.0, FromBody).Return(0.0, nil)
		mockReq.GetFloatDefault("test", 0.0, FromBody)
	}
	{
		mockReq.EXPECT().GetBool("test", FromForm).Return(nil, nil)
		mockReq.GetBool("test", FromForm)
	}
	{
		mockReq.EXPECT().GetTime("test", FromBody).Return(nil, nil)
		mockReq.GetTime("test", FromBody)
	}
	{
		mockReq.EXPECT().FullMap().Return(nil)
		mockReq.FullMap()
	}
	{
		type Person struct {
			Name string `json:"name"`
		}
		PersonType := reflect.TypeOf(Person{})
		mockReq.EXPECT().GetObject("person", PersonType, FromBody).Return(nil, nil)
		mockReq.GetObject("person", PersonType, FromBody)
	}
}

func Test_ExtendedRequestMock(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockReq := NewMockExtendedRequest(ctrl)
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
		mockReq.EXPECT().Header("test").Return("bar")
		mockReq.Header("test")
	}
	{
		mockReq.EXPECT().HeaderKeys().Return([]string{"bar"})
		mockReq.HeaderKeys()
	}
	{
		mockReq.EXPECT().Cookie("test")
		mockReq.Cookie("test")
	}
	{
		mockReq.EXPECT().CookieNames()
		mockReq.CookieNames()
	}
	{
		mockReq.EXPECT().Context()
		mockReq.Context()
	}
	{
		mockReq.EXPECT().GetString("test", FromForm).Return(nil, nil)
		mockReq.GetString("test", FromForm)
	}
	{
		mockReq.EXPECT().GetStringList("test", FromBody).Return(nil, nil)
		mockReq.GetStringList("test", FromBody)
	}
	{
		mockReq.EXPECT().GetInt("test", FromBody).Return(nil, nil)
		mockReq.GetInt("test", FromBody)
	}
	{
		mockReq.EXPECT().GetIntDefault("test", 0, FromContext).Return(0, nil)
		mockReq.GetIntDefault("test", 0, FromContext)
	}
	{
		mockReq.EXPECT().GetFloat("test", FromBody).Return(nil, nil)
		mockReq.GetFloat("test", FromBody)
	}
	{
		mockReq.EXPECT().GetBool("test", FromForm).Return(nil, nil)
		mockReq.GetBool("test", FromForm)
	}
	{
		mockReq.EXPECT().GetTime("test", FromBody).Return(nil, nil)
		mockReq.GetTime("test", FromBody)
	}
	{
		mockReq.EXPECT().FullMap().Return(nil)
		mockReq.FullMap()
	}
	{
		type Person struct {
			Name string `json:"name"`
		}
		PersonType := reflect.TypeOf(Person{})
		mockReq.EXPECT().GetObject("person", PersonType, FromBody).Return(nil, nil)
		mockReq.GetObject("person", PersonType, FromBody)
	}
}
