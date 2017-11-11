package restpc

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
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
	return nil, NewError(
		Unimplemented,
		"we didn't implement this",
		fmt.Errorf("just an unexposed message"),
	).Add("name", "June The Girl")
}

func TestErrorFull(t *testing.T) {
	pkgPath := "github.com/ilius/restpc"
	goPath := os.Getenv("GOPATH")
	expectedHandlerName := pkgPath + ".unimplementedHandler"
	expectedTracebackRecord := &tracebackRecordImp{
		file:     goPath + "/src/" + pkgPath + "/error_test.go",
		function: expectedHandlerName,
		line:     31,
	}

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
	assert.Equal(t, expectedHandlerName, request.HandlerName())
	assert.Equal(t, Unimplemented, rpcErr.Code())
	assert.Equal(t, "we didn't implement this", rpcErr.Error())
	assert.Equal(t, "we didn't implement this", rpcErr.Message())
	assert.Equal(t, "just an unexposed message", rpcErr.Private().Error())
	assert.Equal(t, "June The Girl", rpcErr.Details()["name"])
	{
		tb := rpcErr.Traceback()
		records := tb.Records(expectedHandlerName)
		if len(records) != 1 {
			panic(fmt.Errorf("len(records) = %v", len(records)))
		}
		record := records[0]
		assert.Equal(t, expectedTracebackRecord, record)
		assert.Equal(t, expectedTracebackRecord.File(), record.File())
		assert.Equal(t, expectedTracebackRecord.Function(), record.Function())
		assert.Equal(t, expectedTracebackRecord.FunctionLocal(), record.FunctionLocal())
		assert.Equal(t, expectedTracebackRecord.Line(), record.Line())
		mapRecords := tb.MapRecords(expectedHandlerName)
		if len(mapRecords) != 1 {
			panic(fmt.Errorf("len(mapRecords) = %v", len(mapRecords)))
		}
		mapRecord := mapRecords[0]
		assert.Equal(t, expectedTracebackRecord.File(), mapRecord["file"])
		assert.Equal(t, expectedTracebackRecord.Function(), mapRecord["function"])
		assert.Equal(t, expectedTracebackRecord.FunctionLocal(), mapRecord["functionLocal"])
		assert.Equal(t, expectedTracebackRecord.Line(), mapRecord["line"])
	}
}
