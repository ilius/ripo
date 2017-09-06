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

func convert(key string, valueIn interface{}, typ reflect.Type) (valueOut interface{}, err error) {
	if reflect.TypeOf(valueIn) == typ {
		valueOut = valueIn
		return
	}
	defer func() {
		r := recover()
		if r != nil {
			err = NewError(
				InvalidArgument,
				fmt.Sprintf(
					"invalid '%v', must be %v",
					key,
					typ,
				),
				fmt.Errorf("panic in convert: %v", r),
			)
		}
	}()
	v := reflect.ValueOf(valueIn)
	v2 := v.Convert(typ)
	valueOut = v2.Interface()
	return
}
