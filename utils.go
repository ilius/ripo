package ripo

import (
	"reflect"
	"runtime"
)

func getFunctionName(i any) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
