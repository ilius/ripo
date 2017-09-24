package restpc

import (
	"fmt"
	"reflect"
	"runtime"
)

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func mapFromKeyValuePairs(kvPairs ...interface{}) map[string]interface{} {
	if len(kvPairs)%2 != 0 {
		panic(fmt.Sprintf(
			"mapFromKeyValuePairs: must give even number of args, give %v",
			len(kvPairs),
		))
	}
	m := map[string]interface{}{}
	for i := 0; i < len(kvPairs)/2; i++ {
		m[fmt.Sprintf("%v", kvPairs[i])] = kvPairs[i+1]
	}
	return m
}
