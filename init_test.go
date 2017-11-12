package restpc

import (
	"fmt"
	"net/http"
)

func init() {
	r, _ := http.NewRequest("GET", "", nil)
	request := &requestImp{
		r:           r,
		handlerName: "foo",
	}
	rpcErr := NewError(Unavailable, "", fmt.Errorf("boo")).Add("foo", "bar")
	errorDispatcher(request, rpcErr)
	SetErrorDispatcher(func(request Request, rpcErr RPCError) {})
}
