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
	Error() string    // shown to user
	Code() Code       // shown to user
	GrpcCode() uint32 // shown to user
	Message() string  // shown to user (if set), can be empty

	Cause() error                           // not shown to user
	Unwrap() error                          // not shown to user
	Traceback(handlerName string) Traceback // not shown to user
	Details() map[string]interface{}        // not shown to user

	Add(key string, value interface{}) RPCError
}

type rpcErrorImp struct {
	publicMsg string // shown to user
	code      Code   // shown to user

	cause     error                  // not shown to user
	traceback *tracebackImp          // not shown to user
	details   map[string]interface{} // not shown to user
}

func (e *rpcErrorImp) Error() string {
	if e.publicMsg != "" {
		return e.publicMsg
	}
	return e.code.String()
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

func (e *rpcErrorImp) Cause() error {
	return e.cause
}

func (e *rpcErrorImp) Unwrap() error {
	return e.cause
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
