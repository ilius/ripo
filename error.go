package restpc

func NewError(code Code, publicMsg string, privateErr error, detailsKVPairs ...interface{}) RPCError {
	return &rpcErrorImp{
		code:      code,
		private:   privateErr,
		publicMsg: publicMsg,
		details:   mapFromKeyValuePairs(detailsKVPairs...),
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
