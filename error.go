package ripo

import (
	"runtime"
)

func NewError(code Code, publicMsg string, privateErr error) RPCError {
	if privateErr != nil {
		rpcErr, isRpcErr := privateErr.(RPCError)
		if isRpcErr {
			return rpcErr
		}
	}
	pc := make([]uintptr, 10)
	n := runtime.Callers(2, pc)
	return &rpcErrorImp{
		code:      code,
		private:   privateErr,
		publicMsg: publicMsg,
		traceback: &tracebackImp{callers: pc[:n]},
		details:   map[string]interface{}{},
	}
}

type RPCError interface {
	Error() string // shown to user
	Private() error
	Code() Code
	Message() string
	Traceback(handlerName string) Traceback
	Details() map[string]interface{}
	Add(key string, value interface{}) RPCError
}

type rpcErrorImp struct {
	publicMsg string // shown to user
	private   error
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
	return e.private
}

func (e *rpcErrorImp) Code() Code {
	return e.code
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
