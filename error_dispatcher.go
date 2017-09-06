package restpc

import (
	"fmt"
	"log"
	"strings"
)

var errorDispatcher = func(rpcErr RPCError) {
	parts := []string{
		fmt.Sprintf("Code=%v", rpcErr.Code()),
		fmt.Sprintf("Message=%#v", rpcErr.Message()),
	}
	if rpcErr.Private() != nil {
		parts = append(parts, fmt.Sprintf("Original=%#v", rpcErr.Private().Error()))
	}
	if len(rpcErr.Details()) > 0 {
		parts = append(parts, fmt.Sprintf("Details=%#v", rpcErr.Details()))
	}
	log.Printf("RPCError: %v", strings.Join(parts, ", "))
}

func SetErrorDispatcher(dispatcher func(rpcErr RPCError)) {
	errorDispatcher = dispatcher
}
