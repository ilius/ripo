package restpc

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_NewError_Twice(t *testing.T) {
	err := NewError(InvalidArgument, "something is missing", nil)
	assert.Equal(t, "something is missing", err.Message())
	assert.Equal(t, "something is missing", err.Error())
	assert.Equal(t, InvalidArgument, err.Code())

	err = NewError(Unavailable, "not sure what", err)
	assert.Equal(t, "something is missing", err.Message())
	assert.Equal(t, "something is missing", err.Error())
	assert.Equal(t, InvalidArgument, err.Code())
}

func unimplementedHandler(req Request) (*Response, error) {
	err := fmt.Errorf("just an unexposed message")
	return nil, NewError(Unimplemented, "we didn't implement this", err).Add("name", "June The Girl")
}

func TestErrorFull(t *testing.T) {
	handlerName := "unimplementedHandler"

	origErrorDispatcher := errorDispatcher
	defer func() {
		errorDispatcher = origErrorDispatcher
	}()
	var rpcErr RPCError
	var request Request
	errorDispatcher = func(requestArg Request, rpcErrArg RPCError) {
		rpcErr = rpcErrArg
		request = requestArg
	}
	handlerFunc := TranslateHandler(unimplementedHandler)
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "", nil)
	handlerFunc(w, r)
	assert.Equal(t, http.StatusNotImplemented, w.Code)
	time.Sleep(500 * time.Millisecond)
	if rpcErr == nil {
		panic("rpcErr == nil")
	}
	if !strings.HasSuffix(request.HandlerName(), "."+handlerName) {
		t.Errorf("request.HandlerName()=%v", request.HandlerName())
		return
	}
	handlerNameFull := request.HandlerName()
	assert.Equal(t, Unimplemented, rpcErr.Code())
	assert.Equal(t, "we didn't implement this", rpcErr.Error())
	assert.Equal(t, "we didn't implement this", rpcErr.Message())
	assert.Equal(t, "just an unexposed message", rpcErr.Private().Error())
	assert.Equal(t, "June The Girl", rpcErr.Details()["name"])
	{
		tb := rpcErr.Traceback("")
		assert.Equal(t, 6, len(tb.Callers()))
		assert.Equal(t, 6, len(tb.Records()))
	}
	{
		tb := rpcErr.Traceback(handlerNameFull)
		assert.Equal(t, 6, len(tb.Callers()))
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
		assert.Equal(t, handlerNameFull, record.Function())
		assert.Equal(t, handlerName, record.FunctionLocal())
		assert.Equal(t, 28, record.Line())
		mapRecords := tb.MapRecords()
		if len(mapRecords) != 1 {
			t.Errorf("len(mapRecords) = %v", len(mapRecords))
			return
		}
		mapRecord := mapRecords[0]
		assert.Equal(t, record.File(), mapRecord["file"])
		assert.Equal(t, record.Function(), mapRecord["function"])
		assert.Equal(t, record.FunctionLocal(), mapRecord["functionLocal"])
		assert.Equal(t, record.Line(), mapRecord["line"])
	}
	{
		tb := rpcErr.Traceback("")
		assert.Equal(t, 6, len(tb.Callers()))
		assert.Equal(t, 6, len(tb.Records()))
	}
}

func TestFunctionLocalNoPanic(t *testing.T) {
	record := &tracebackRecordImp{}
	assert.Equal(t, "", record.FunctionLocal())
}

func Test_tracebackImp_empty(t *testing.T) {
	tb := &tracebackImp{}
	assert.Equal(t, 0, len(tb.Callers()))
	assert.Equal(t, 0, len(tb.Records()))
}
