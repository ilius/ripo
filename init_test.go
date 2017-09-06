package restpc

func init() {
	errorDispatcher = func(rpcErr RPCError) {}
}
