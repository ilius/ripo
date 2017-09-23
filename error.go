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
	details := mapFromKeyValuePairs(detailsKVPairs...)
	pc := make([]uintptr, 10)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	traceback := []map[string]interface{}{}
	processFrame := func(frame runtime.Frame) bool {
		if frame.Func == nil {
			return true
		}
		traceback = append(traceback, map[string]interface{}{
			"file":     frame.File,
			"function": frame.Function,
			"line":     frame.Line,
		})
		return true
	}
	for {
		frame, more := frames.Next()
		if !processFrame(frame) || !more {
			break
		}
	}
	details["traceback"] = traceback
	return &rpcErrorImp{
		code:      code,
		private:   privateErr,
		publicMsg: publicMsg,
		details:   details,
	}
}

type RPCError interface {
	Error() string // shown to user
	Private() error
	Code() Code
	Message() string
	Details() map[string]interface{}
}

type rpcErrorImp struct {
	publicMsg string // shown to user
	private   error
	code      Code
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

func (e *rpcErrorImp) Details() map[string]interface{} {
	return e.details
}
