package restpc

func init() {
	errorDispatcher = func(request Request, rpcErr RPCError) {}
}
