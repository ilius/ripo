package ripo

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/ilius/is"
)

func Test_NewError_Twice(t *testing.T) {
	is := is.New(t)

	err := NewError(InvalidArgument, "something is missing", nil)
	is.Equal("something is missing", err.Message())
	is.Equal("something is missing", err.Error())
	is.Equal(InvalidArgument, err.Code())

	err = NewError(Unavailable, "not sure what", err)
	is.Equal("something is missing", err.Message())
	is.Equal("something is missing", err.Error())
	is.Equal(InvalidArgument, err.Code())
}

func unimplementedHandler(req Request) (*Response, error) {
	err := fmt.Errorf("just an unexposed message")
	return nil, NewError(Unimplemented, "we didn't implement this", err).Add("name", "June The Girl")
}

func TestError_GrpcCodeExtra(t *testing.T) {
	is := is.New(t)
	is.Equal(NewError(MissingArgument, "", nil).GrpcCode(), uint32(InvalidArgument))
	is.Equal(NewError(ResourceLocked, "", nil).GrpcCode(), uint32(Aborted))
}

func TestErrorFull(t *testing.T) {
	is := is.New(t)

	handlerName := "unimplementedHandler"

	origErrorDispatcher := errorDispatcher
	defer func() {
		errorDispatcher = origErrorDispatcher
	}()
	var rpcErr RPCError
	var request Request
	errorDispatcher = func(requestArg ExtendedRequest, rpcErrArg RPCError) {
		rpcErr = rpcErrArg
		request = requestArg
	}
	handlerFunc := TranslateHandler(unimplementedHandler)
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "", nil)
	handlerFunc(w, r)
	is.Equal(http.StatusNotImplemented, w.Code)
	time.Sleep(500 * time.Millisecond)
	if rpcErr == nil {
		panic("rpcErr == nil")
	}
	if !strings.HasSuffix(request.HandlerName(), "."+handlerName) {
		t.Errorf("request.HandlerName()=%v", request.HandlerName())
		return
	}
	handlerNameFull := request.HandlerName()
	is.Equal(Unimplemented, rpcErr.Code())
	is.Equal(uint32(Unimplemented), rpcErr.GrpcCode())
	is.Equal("we didn't implement this", rpcErr.Error())
	is.Equal("we didn't implement this", rpcErr.Message())
	is.Equal("just an unexposed message", rpcErr.Cause().Error())
	is.Equal("just an unexposed message", rpcErr.Cause().Error())
	is.Equal("June The Girl", rpcErr.Details()["name"])
	{
		tb := rpcErr.Traceback("")
		is.Equal(6, len(tb.(*tracebackImp).Callers()))
		is.Equal(6, len(tb.Records()))
	}
	{
		tb := rpcErr.Traceback(handlerNameFull)
		is.Equal(6, len(tb.(*tracebackImp).Callers()))
		records := tb.Records()
		if len(records) != 1 {
			for _, record := range records {
				t.Log(record)
			}
			t.Errorf("len(records) = %v", len(records))
			return
		}
		record := records[0]
		if !strings.HasSuffix(record.File(), "/error_test.go") {
			t.Errorf("record.File()=%v", record.File())
		}
		is.Equal(handlerNameFull, record.Function())
		is.Equal(handlerName, record.FunctionLocal())
		is.Equal(30, record.Line())
		mapRecords := tb.MapRecords()
		if len(mapRecords) != 1 {
			t.Errorf("len(mapRecords) = %v", len(mapRecords))
			return
		}
		mapRecord := mapRecords[0]
		is.Equal(record.File(), mapRecord["file"])
		is.Equal(record.Function(), mapRecord["function"])
		is.Equal(record.FunctionLocal(), mapRecord["functionLocal"])
		is.Equal(record.Line(), mapRecord["line"])
	}
	{
		tb := rpcErr.Traceback("")
		is.Equal(6, len(tb.(*tracebackImp).Callers()))
		is.Equal(6, len(tb.Records()))
	}
}

func TestFunctionLocalNoPanic(t *testing.T) {
	is := is.New(t)
	record := &tracebackRecordImp{}
	is.Equal("", record.FunctionLocal())
}

func Test_tracebackImp_empty(t *testing.T) {
	is := is.New(t)
	tb := &tracebackImp{}
	is.Equal(0, len(tb.Callers()))
	is.Equal(0, len(tb.Records()))
}
