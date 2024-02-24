// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ilius/ripo (interfaces: Request,ExtendedRequest)

package ripo

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	http "net/http"
	url "net/url"
	reflect "reflect"
	time "time"
)

// MockRequest is a mock of Request interface
type MockRequest struct {
	ctrl     *gomock.Controller
	recorder *MockRequestMockRecorder
}

// MockRequestMockRecorder is the mock recorder for MockRequest
type MockRequestMockRecorder struct {
	mock *MockRequest
}

// NewMockRequest creates a new mock instance
func NewMockRequest(ctrl *gomock.Controller) *MockRequest {
	mock := &MockRequest{ctrl: ctrl}
	mock.recorder = &MockRequestMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRequest) EXPECT() *MockRequestMockRecorder {
	return m.recorder
}

// Body mocks base method
func (m *MockRequest) Body() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Body")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Body indicates an expected call of Body
func (mr *MockRequestMockRecorder) Body() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Body", reflect.TypeOf((*MockRequest)(nil).Body))
}

// BodyTo mocks base method
func (m *MockRequest) BodyTo(arg0 any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BodyTo", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// BodyTo indicates an expected call of BodyTo
func (mr *MockRequestMockRecorder) BodyTo(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BodyTo", reflect.TypeOf((*MockRequest)(nil).BodyTo), arg0)
}

// Context mocks base method
func (m *MockRequest) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockRequestMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockRequest)(nil).Context))
}

// Cookie mocks base method
func (m *MockRequest) Cookie(arg0 string) (*http.Cookie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cookie", arg0)
	ret0, _ := ret[0].(*http.Cookie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cookie indicates an expected call of Cookie
func (mr *MockRequestMockRecorder) Cookie(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cookie", reflect.TypeOf((*MockRequest)(nil).Cookie), arg0)
}

// CookieNames mocks base method
func (m *MockRequest) CookieNames() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CookieNames")
	ret0, _ := ret[0].([]string)
	return ret0
}

// CookieNames indicates an expected call of CookieNames
func (mr *MockRequestMockRecorder) CookieNames() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CookieNames", reflect.TypeOf((*MockRequest)(nil).CookieNames))
}

// FullMap mocks base method
func (m *MockRequest) FullMap() map[string]any {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FullMap")
	ret0, _ := ret[0].(map[string]any)
	return ret0
}

// FullMap indicates an expected call of FullMap
func (mr *MockRequestMockRecorder) FullMap() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FullMap", reflect.TypeOf((*MockRequest)(nil).FullMap))
}

// GetBool mocks base method
func (m *MockRequest) GetBool(arg0 string, arg1 ...FromX) (*bool, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBool", varargs...)
	ret0, _ := ret[0].(*bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBool indicates an expected call of GetBool
func (mr *MockRequestMockRecorder) GetBool(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBool", reflect.TypeOf((*MockRequest)(nil).GetBool), varargs...)
}

// GetFloat mocks base method
func (m *MockRequest) GetFloat(arg0 string, arg1 ...FromX) (*float64, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetFloat", varargs...)
	ret0, _ := ret[0].(*float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFloat indicates an expected call of GetFloat
func (mr *MockRequestMockRecorder) GetFloat(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFloat", reflect.TypeOf((*MockRequest)(nil).GetFloat), varargs...)
}

// GetFloatDefault mocks base method
func (m *MockRequest) GetFloatDefault(arg0 string, arg1 float64, arg2 ...FromX) (float64, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetFloatDefault", varargs...)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFloatDefault indicates an expected call of GetFloatDefault
func (mr *MockRequestMockRecorder) GetFloatDefault(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFloatDefault", reflect.TypeOf((*MockRequest)(nil).GetFloatDefault), varargs...)
}

// GetInt mocks base method
func (m *MockRequest) GetInt(arg0 string, arg1 ...FromX) (*int, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetInt", varargs...)
	ret0, _ := ret[0].(*int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInt indicates an expected call of GetInt
func (mr *MockRequestMockRecorder) GetInt(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInt", reflect.TypeOf((*MockRequest)(nil).GetInt), varargs...)
}

// GetIntDefault mocks base method
func (m *MockRequest) GetIntDefault(arg0 string, arg1 int, arg2 ...FromX) (int, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetIntDefault", varargs...)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIntDefault indicates an expected call of GetIntDefault
func (mr *MockRequestMockRecorder) GetIntDefault(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIntDefault", reflect.TypeOf((*MockRequest)(nil).GetIntDefault), varargs...)
}

// GetObject mocks base method
func (m *MockRequest) GetObject(arg0 string, arg1 reflect.Type, arg2 ...FromX) (any, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetObject", varargs...)
	ret0, _ := ret[0].(any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObject indicates an expected call of GetObject
func (mr *MockRequestMockRecorder) GetObject(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObject", reflect.TypeOf((*MockRequest)(nil).GetObject), varargs...)
}

// GetString mocks base method
func (m *MockRequest) GetString(arg0 string, arg1 ...FromX) (*string, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetString", varargs...)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetString indicates an expected call of GetString
func (mr *MockRequestMockRecorder) GetString(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetString", reflect.TypeOf((*MockRequest)(nil).GetString), varargs...)
}

// GetStringDefault mocks base method
func (m *MockRequest) GetStringDefault(arg0, arg1 string, arg2 ...FromX) (string, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetStringDefault", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStringDefault indicates an expected call of GetStringDefault
func (mr *MockRequestMockRecorder) GetStringDefault(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStringDefault", reflect.TypeOf((*MockRequest)(nil).GetStringDefault), varargs...)
}

// GetStringList mocks base method
func (m *MockRequest) GetStringList(arg0 string, arg1 ...FromX) ([]string, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetStringList", varargs...)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStringList indicates an expected call of GetStringList
func (mr *MockRequestMockRecorder) GetStringList(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStringList", reflect.TypeOf((*MockRequest)(nil).GetStringList), varargs...)
}

// GetTime mocks base method
func (m *MockRequest) GetTime(arg0 string, arg1 ...FromX) (*time.Time, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTime", varargs...)
	ret0, _ := ret[0].(*time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTime indicates an expected call of GetTime
func (mr *MockRequestMockRecorder) GetTime(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTime", reflect.TypeOf((*MockRequest)(nil).GetTime), varargs...)
}

// HandlerName mocks base method
func (m *MockRequest) HandlerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandlerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// HandlerName indicates an expected call of HandlerName
func (mr *MockRequestMockRecorder) HandlerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandlerName", reflect.TypeOf((*MockRequest)(nil).HandlerName))
}

// Header mocks base method
func (m *MockRequest) Header(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// Header indicates an expected call of Header
func (mr *MockRequestMockRecorder) Header(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockRequest)(nil).Header), arg0)
}

// HeaderKeys mocks base method
func (m *MockRequest) HeaderKeys() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HeaderKeys")
	ret0, _ := ret[0].([]string)
	return ret0
}

// HeaderKeys indicates an expected call of HeaderKeys
func (mr *MockRequestMockRecorder) HeaderKeys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HeaderKeys", reflect.TypeOf((*MockRequest)(nil).HeaderKeys))
}

// Host mocks base method
func (m *MockRequest) Host() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Host")
	ret0, _ := ret[0].(string)
	return ret0
}

// Host indicates an expected call of Host
func (mr *MockRequestMockRecorder) Host() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Host", reflect.TypeOf((*MockRequest)(nil).Host))
}

// RemoteIP mocks base method
func (m *MockRequest) RemoteIP() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoteIP")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoteIP indicates an expected call of RemoteIP
func (mr *MockRequestMockRecorder) RemoteIP() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteIP", reflect.TypeOf((*MockRequest)(nil).RemoteIP))
}

// URL mocks base method
func (m *MockRequest) URL() *url.URL {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "URL")
	ret0, _ := ret[0].(*url.URL)
	return ret0
}

// URL indicates an expected call of URL
func (mr *MockRequestMockRecorder) URL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "URL", reflect.TypeOf((*MockRequest)(nil).URL))
}

// MockExtendedRequest is a mock of ExtendedRequest interface
type MockExtendedRequest struct {
	ctrl     *gomock.Controller
	recorder *MockExtendedRequestMockRecorder
}

// MockExtendedRequestMockRecorder is the mock recorder for MockExtendedRequest
type MockExtendedRequestMockRecorder struct {
	mock *MockExtendedRequest
}

// NewMockExtendedRequest creates a new mock instance
func NewMockExtendedRequest(ctrl *gomock.Controller) *MockExtendedRequest {
	mock := &MockExtendedRequest{ctrl: ctrl}
	mock.recorder = &MockExtendedRequestMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockExtendedRequest) EXPECT() *MockExtendedRequestMockRecorder {
	return m.recorder
}

// Body mocks base method
func (m *MockExtendedRequest) Body() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Body")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Body indicates an expected call of Body
func (mr *MockExtendedRequestMockRecorder) Body() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Body", reflect.TypeOf((*MockExtendedRequest)(nil).Body))
}

// BodyMap mocks base method
func (m *MockExtendedRequest) BodyMap() (map[string]any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BodyMap")
	ret0, _ := ret[0].(map[string]any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BodyMap indicates an expected call of BodyMap
func (mr *MockExtendedRequestMockRecorder) BodyMap() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BodyMap", reflect.TypeOf((*MockExtendedRequest)(nil).BodyMap))
}

// BodyTo mocks base method
func (m *MockExtendedRequest) BodyTo(arg0 any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BodyTo", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// BodyTo indicates an expected call of BodyTo
func (mr *MockExtendedRequestMockRecorder) BodyTo(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BodyTo", reflect.TypeOf((*MockExtendedRequest)(nil).BodyTo), arg0)
}

// Context mocks base method
func (m *MockExtendedRequest) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockExtendedRequestMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockExtendedRequest)(nil).Context))
}

// Cookie mocks base method
func (m *MockExtendedRequest) Cookie(arg0 string) (*http.Cookie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cookie", arg0)
	ret0, _ := ret[0].(*http.Cookie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cookie indicates an expected call of Cookie
func (mr *MockExtendedRequestMockRecorder) Cookie(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cookie", reflect.TypeOf((*MockExtendedRequest)(nil).Cookie), arg0)
}

// CookieNames mocks base method
func (m *MockExtendedRequest) CookieNames() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CookieNames")
	ret0, _ := ret[0].([]string)
	return ret0
}

// CookieNames indicates an expected call of CookieNames
func (mr *MockExtendedRequestMockRecorder) CookieNames() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CookieNames", reflect.TypeOf((*MockExtendedRequest)(nil).CookieNames))
}

// FullMap mocks base method
func (m *MockExtendedRequest) FullMap() map[string]any {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FullMap")
	ret0, _ := ret[0].(map[string]any)
	return ret0
}

// FullMap indicates an expected call of FullMap
func (mr *MockExtendedRequestMockRecorder) FullMap() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FullMap", reflect.TypeOf((*MockExtendedRequest)(nil).FullMap))
}

// GetBool mocks base method
func (m *MockExtendedRequest) GetBool(arg0 string, arg1 ...FromX) (*bool, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBool", varargs...)
	ret0, _ := ret[0].(*bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBool indicates an expected call of GetBool
func (mr *MockExtendedRequestMockRecorder) GetBool(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBool", reflect.TypeOf((*MockExtendedRequest)(nil).GetBool), varargs...)
}

// GetFloat mocks base method
func (m *MockExtendedRequest) GetFloat(arg0 string, arg1 ...FromX) (*float64, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetFloat", varargs...)
	ret0, _ := ret[0].(*float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFloat indicates an expected call of GetFloat
func (mr *MockExtendedRequestMockRecorder) GetFloat(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFloat", reflect.TypeOf((*MockExtendedRequest)(nil).GetFloat), varargs...)
}

// GetFloatDefault mocks base method
func (m *MockExtendedRequest) GetFloatDefault(arg0 string, arg1 float64, arg2 ...FromX) (float64, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetFloatDefault", varargs...)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFloatDefault indicates an expected call of GetFloatDefault
func (mr *MockExtendedRequestMockRecorder) GetFloatDefault(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFloatDefault", reflect.TypeOf((*MockExtendedRequest)(nil).GetFloatDefault), varargs...)
}

// GetFormValue mocks base method
func (m *MockExtendedRequest) GetFormValue(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFormValue", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetFormValue indicates an expected call of GetFormValue
func (mr *MockExtendedRequestMockRecorder) GetFormValue(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFormValue", reflect.TypeOf((*MockExtendedRequest)(nil).GetFormValue), arg0)
}

// GetInt mocks base method
func (m *MockExtendedRequest) GetInt(arg0 string, arg1 ...FromX) (*int, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetInt", varargs...)
	ret0, _ := ret[0].(*int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInt indicates an expected call of GetInt
func (mr *MockExtendedRequestMockRecorder) GetInt(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInt", reflect.TypeOf((*MockExtendedRequest)(nil).GetInt), varargs...)
}

// GetIntDefault mocks base method
func (m *MockExtendedRequest) GetIntDefault(arg0 string, arg1 int, arg2 ...FromX) (int, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetIntDefault", varargs...)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIntDefault indicates an expected call of GetIntDefault
func (mr *MockExtendedRequestMockRecorder) GetIntDefault(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIntDefault", reflect.TypeOf((*MockExtendedRequest)(nil).GetIntDefault), varargs...)
}

// GetObject mocks base method
func (m *MockExtendedRequest) GetObject(arg0 string, arg1 reflect.Type, arg2 ...FromX) (any, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetObject", varargs...)
	ret0, _ := ret[0].(any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObject indicates an expected call of GetObject
func (mr *MockExtendedRequestMockRecorder) GetObject(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObject", reflect.TypeOf((*MockExtendedRequest)(nil).GetObject), varargs...)
}

// GetString mocks base method
func (m *MockExtendedRequest) GetString(arg0 string, arg1 ...FromX) (*string, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetString", varargs...)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetString indicates an expected call of GetString
func (mr *MockExtendedRequestMockRecorder) GetString(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetString", reflect.TypeOf((*MockExtendedRequest)(nil).GetString), varargs...)
}

// GetStringDefault mocks base method
func (m *MockExtendedRequest) GetStringDefault(arg0, arg1 string, arg2 ...FromX) (string, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetStringDefault", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStringDefault indicates an expected call of GetStringDefault
func (mr *MockExtendedRequestMockRecorder) GetStringDefault(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStringDefault", reflect.TypeOf((*MockExtendedRequest)(nil).GetStringDefault), varargs...)
}

// GetStringList mocks base method
func (m *MockExtendedRequest) GetStringList(arg0 string, arg1 ...FromX) ([]string, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetStringList", varargs...)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStringList indicates an expected call of GetStringList
func (mr *MockExtendedRequestMockRecorder) GetStringList(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStringList", reflect.TypeOf((*MockExtendedRequest)(nil).GetStringList), varargs...)
}

// GetTime mocks base method
func (m *MockExtendedRequest) GetTime(arg0 string, arg1 ...FromX) (*time.Time, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTime", varargs...)
	ret0, _ := ret[0].(*time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTime indicates an expected call of GetTime
func (mr *MockExtendedRequestMockRecorder) GetTime(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTime", reflect.TypeOf((*MockExtendedRequest)(nil).GetTime), varargs...)
}

// HandlerName mocks base method
func (m *MockExtendedRequest) HandlerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandlerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// HandlerName indicates an expected call of HandlerName
func (mr *MockExtendedRequestMockRecorder) HandlerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandlerName", reflect.TypeOf((*MockExtendedRequest)(nil).HandlerName))
}

// Header mocks base method
func (m *MockExtendedRequest) Header(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// Header indicates an expected call of Header
func (mr *MockExtendedRequestMockRecorder) Header(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockExtendedRequest)(nil).Header), arg0)
}

// HeaderKeys mocks base method
func (m *MockExtendedRequest) HeaderKeys() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HeaderKeys")
	ret0, _ := ret[0].([]string)
	return ret0
}

// HeaderKeys indicates an expected call of HeaderKeys
func (mr *MockExtendedRequestMockRecorder) HeaderKeys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HeaderKeys", reflect.TypeOf((*MockExtendedRequest)(nil).HeaderKeys))
}

// Host mocks base method
func (m *MockExtendedRequest) Host() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Host")
	ret0, _ := ret[0].(string)
	return ret0
}

// Host indicates an expected call of Host
func (mr *MockExtendedRequestMockRecorder) Host() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Host", reflect.TypeOf((*MockExtendedRequest)(nil).Host))
}

// RemoteIP mocks base method
func (m *MockExtendedRequest) RemoteIP() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoteIP")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoteIP indicates an expected call of RemoteIP
func (mr *MockExtendedRequestMockRecorder) RemoteIP() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteIP", reflect.TypeOf((*MockExtendedRequest)(nil).RemoteIP))
}

// URL mocks base method
func (m *MockExtendedRequest) URL() *url.URL {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "URL")
	ret0, _ := ret[0].(*url.URL)
	return ret0
}

// URL indicates an expected call of URL
func (mr *MockExtendedRequestMockRecorder) URL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "URL", reflect.TypeOf((*MockExtendedRequest)(nil).URL))
}
