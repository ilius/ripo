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
	details := mapFromKeyValuePairs(detailsKVPairs...)
	return &rpcErrorImp{
		code:      code,
		private:   privateErr,
		publicMsg: publicMsg,
		traceback: &tracebackImp{callers: pc[:n]},
		details:   details,
	}
}

type RPCError interface {
	Error() string // shown to user
	Private() error
	Code() Code
	Message() string
	Traceback() Traceback
	Details() map[string]interface{}
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

func (e *rpcErrorImp) Traceback() Traceback {
	return e.traceback
}

func (e *rpcErrorImp) Details() map[string]interface{} {
	return e.details
}
