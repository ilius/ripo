package ripo

// SmallT is a minimal subset of testing.TB (implemented by testing.T) that we use
type SmallT interface {
	Helper()
	Fatalf(format string, args ...interface{})
	// Errorf(format string, args ...interface{})
	// Logf(format string, args ...interface{})
}

// for non-rpc errors, pass code=Unknown
func AssertError(t SmallT, err error, code Code, msg string) bool {
	t.Helper()
	if err == nil {
		t.Fatalf("got err==nil, expected code=%v msg=%#v", code, msg)
		return false
	}
	rpcErr, isRPC := err.(RPCError)
	if isRPC {
		if code != rpcErr.Code() {
			t.Fatalf("got code=%v in err==%#v, expected code=%v", rpcErr.Code(), err.Error(), code)
			return false
		}
	} else {
		if code != Unknown {
			t.Fatalf("got non-rpc err==%#v, expected code=%v", err.Error(), code)
			return false
		}
	}
	if err.Error() != msg {
		t.Fatalf("got err.Error()==%#v, expected %#v", err.Error(), msg)
		return false
	}
	return true
}
