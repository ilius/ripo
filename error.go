package restpc

import (
	"runtime"
)

func NewError(code Code, publicMsg string, privateErr error, detailsKVPairs ...interface{}) RPCError {
	if privateErr != nil {
		rpcErr, isRpcErr := privateErr.(RPCError)
		if isRpcErr {
			return rpcErr
		}
	}
	pc := make([]uintptr, 10)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	traceback := []TracebackRecord{}
	processFrame := func(frame runtime.Frame) bool {
		if frame.Func == nil {
			return true
		}
		traceback = append(traceback, &tracebackRecordImp{
			file:     frame.File,
			function: frame.Function,
			line:     frame.Line,
		})
		_, isHandler := handlers[frame.Function]
		if isHandler {
			return false
		}
		return true
	}
	for {
		frame, more := frames.Next()
		if !processFrame(frame) || !more {
			break
		}
	}
	details := mapFromKeyValuePairs(detailsKVPairs...)
	return &rpcErrorImp{
		code:      code,
		private:   privateErr,
		publicMsg: publicMsg,
		traceback: traceback,
		details:   details,
	}
}

type TracebackRecord interface {
	File() string
	Function() string
	Line() int
}

type RPCError interface {
	Error() string // shown to user
	Private() error
	Code() Code
	Message() string
	Traceback() []TracebackRecord
	Details() map[string]interface{}
}

type tracebackRecordImp struct {
	file     string
	function string
	line     int
}

func (tr *tracebackRecordImp) File() string {
	return tr.file
}
func (tr *tracebackRecordImp) Function() string {
	return tr.function
}
func (tr *tracebackRecordImp) Line() int {
	return tr.line
}

type rpcErrorImp struct {
	publicMsg string // shown to user
	private   error
	code      Code
	traceback []TracebackRecord
	details   map[string]interface{}
}

func (e *rpcErrorImp) Error() string {
	if e.publicMsg != "" {
		return e.publicMsg
	}
	return e.code.String()
}

func (e *rpcErrorImp) Private() error {
	return e.private
}

func (e *rpcErrorImp) Code() Code {
	return e.code
}

func (e *rpcErrorImp) Message() string {
	return e.publicMsg
}

func (e *rpcErrorImp) Traceback() []TracebackRecord {
	return e.traceback
}

func (e *rpcErrorImp) Details() map[string]interface{} {
	return e.details
}
