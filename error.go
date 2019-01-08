package ripo

import (
	"runtime"
)

// code and publicMsg are exposed to client by api
// while causeErr is not exposed to client by api
func NewError(code Code, publicMsg string, causeErr error) RPCError {
	if causeErr != nil {
		rpcErr, isRpcErr := causeErr.(RPCError)
		if isRpcErr {
			return rpcErr
		}
	}
	pc := make([]uintptr, 10)
	n := runtime.Callers(2, pc)
	return &rpcErrorImp{
		code:      code,
		cause:     causeErr,
		publicMsg: publicMsg,
		traceback: &tracebackImp{callers: pc[:n]},
		details:   map[string]interface{}{},
	}
}

type RPCError interface {
	Error() string // shown to user
	Private() error
	Cause() error
	Code() Code
	GrpcCode() uint32
	Message() string
	Traceback(handlerName string) Traceback
	Details() map[string]interface{}
	Add(key string, value interface{}) RPCError
}

type rpcErrorImp struct {
	publicMsg string // shown to user
	cause     error
	code      Code
	traceback *tracebackImp
	details   map[string]interface{}
}

func (e *rpcErrorImp) Error() string {
	if e.publicMsg != "" {
		return e.publicMsg
	}
	return e.code.String()
}

func (e *rpcErrorImp) Private() error {
	return e.cause
}

func (e *rpcErrorImp) Cause() error {
	return e.cause
}

func (e *rpcErrorImp) Code() Code {
	return e.code
}

func (e *rpcErrorImp) GrpcCode() uint32 {
	switch e.code {
	case MissingArgument:
		return uint32(InvalidArgument)
	case ResourceLocked:
		return uint32(Aborted)
	}
	return uint32(e.code)
}

func (e *rpcErrorImp) Message() string {
	return e.publicMsg
}

func (e *rpcErrorImp) Traceback(handlerName string) Traceback {
	return e.traceback.SetHandlerName(handlerName)
}

func (e *rpcErrorImp) Details() map[string]interface{} {
	return e.details
}

func (e *rpcErrorImp) Add(key string, value interface{}) RPCError {
	_, hasKey := e.details[key]
	if !hasKey {
		e.details[key] = value
	}
	return e
}
